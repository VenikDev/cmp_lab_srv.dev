package core

import (
	"comparisonLaboratories/src/clog"
	"comparisonLaboratories/src/global"
	"comparisonLaboratories/src/herr"
	"comparisonLaboratories/src/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func CreateURLFrom(key string, lab global.Laboratory) string {
	return fmt.Sprintf("%s%s?%s=%s", lab.GetUrl(), "/search", lab.GetParamForFind(), key)
}

func GetHtmlFrom(url string, lab global.Laboratory) *goquery.Document {
	response, err := http.Get(url)
	herr.HandlerError(err, "")

	AddHeadersTo(response, lab)

	defer response.Body.Close()
	if response.StatusCode != 200 {
		clog.Logger.Error("status code error: %d %s", response.StatusCode, response.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	herr.HandlerError(err, "")

	return doc
}

func AddHeadersTo(response *http.Response, lab global.Laboratory) {
	switch lab.Name {
	case model.GEMOTEST:
		response.Header.Add("BITRIX_SM_CITY_CODE", "")
	}
}
