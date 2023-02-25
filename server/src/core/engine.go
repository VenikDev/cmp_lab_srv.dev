package core

import (
	"comparisonLaboratories/src/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func GetListTests(key string, lab model.Laboratory) model.ListAnalyses {
	request := fmt.Sprintf("%s?%s=%s", lab.GetUrl(), lab.GetParamForFind(), key)
	log.Printf("request = %s", request)

	// получаем страницу
	response, err := http.Get(request)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Printf("status code error: %d %s", response.StatusCode, response.Status)
		return nil
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	result := lab.GetAnalyzes(doc, key)

	return result
}
