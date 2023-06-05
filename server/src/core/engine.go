package core

import (
	"bytes"
	"cmp_lab/src/clog"
	"cmp_lab/src/global"
	"cmp_lab/src/model"
	"cmp_lab/src/model/city"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func CreateURLFrom(params model.Bundle, lab global.Laboratory) string {
	key := params["key"].(string)

	switch lab.Name {
	case model.CITILAB:
		{
			return fmt.Sprintf("%s%s?%s=%s", lab.GetUrl(), "/search", lab.GetParamForFind(), key)
		}
	case model.INVITRO:
		{
			return fmt.Sprintf("%s%s?%s=%s", lab.GetUrl(), "/search", lab.GetParamForFind(), key)
		}
	default:
		{
			return fmt.Sprintf("%s%s?%s=%s", lab.GetUrl(), "/search", lab.GetParamForFind(), key)
		}
	}
}

func GetHtmlFrom(url string, lab global.Laboratory, params model.Bundle) *goquery.Document {
	client := resty.New()
	req := client.
		R().
		EnableTrace()

	AddHeadersTo(req, lab, params)

	response, err := req.Get(url)

	if err != nil {
		clog.Error("[parse_html/get_html_from]", "error", err)
		return nil
	}

	if response.StatusCode() != http.StatusOK {
		clog.Error(
			"[parse_html/get_html_from]",
			"error",
			fmt.Sprintf("status code error: %d %s", response.StatusCode, response.Status))
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(response.Body()))
	if err != nil {
		clog.Error("[parse_html/get_html_from]", "error", err)
		return nil
	}

	return doc
}

func AddHeadersTo(response *resty.Request, lab global.Laboratory, params model.Bundle) {
	city := params["city"].(city.City).NameEn
	setCookie := "Set-Cookie"
	switch lab.Name {
	case model.GEMOTEST:
		response.SetHeader(
			setCookie,
			fmt.Sprintf("BITRIX_SM_CITY_CODE=%s", city))
	case model.INVITRO:
		response.SetHeader(
			setCookie,
			fmt.Sprintf("INVITRO_REGION_CODE=%s", city))

	}
}
