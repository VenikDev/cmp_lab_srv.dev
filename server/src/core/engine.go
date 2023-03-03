package core

import (
	"comparisonLaboratories/src/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func CreateURLFrom(key string, lab model.Laboratory) string {
	return fmt.Sprintf("%s?%s=%s", lab.GetUrl(), lab.GetParamForFind(), key)
}

func GetHtmlFrom(url string) *goquery.Document {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Printf("status code error: %d %s", response.StatusCode, response.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(err)
	}

	return doc
}
