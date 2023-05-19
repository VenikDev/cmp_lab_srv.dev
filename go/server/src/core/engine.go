package core

import (
	"cmp_lab/src/clog"
	"cmp_lab/src/global"
	"cmp_lab/src/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func CreateURLFrom(key string, lab models.Laboratory) string {
	return fmt.Sprintf("%s%s?%s=%s", lab.Url, "/search", lab.ParamForFind, key)
}

func GetHtmlFrom(url string) *goquery.Document {
	response, err := http.Get(url)

	if err != nil {
		clog.Logger.Error("[parse_html/get_html_from]", "error", err)
		return nil
	}

	//AddHeadersTo(response, parse)

	defer response.Body.Close()
	if response.StatusCode != 200 {
		clog.Logger.Error(
			"[parse_html/get_html_from]",
			"error",
			fmt.Sprintf("status code error: %d %s", response.StatusCode, response.Status))
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		clog.Logger.Error("[parse_html/get_html_from]", "error", err)
		return nil
	}

	return doc
}

func AddHeadersTo(response *http.Response, lab models.Laboratory) {
	switch lab.Name {
	case model.GEMOTEST:
		response.Header.Add("BITRIX_SM_CITY_CODE", "")
	}
}
