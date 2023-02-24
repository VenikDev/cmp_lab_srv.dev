package parse

import (
	"comparisonLaboratories/src/services/parse/config"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func Parse(key string, lab config.Laboratory) map[string]string {
	request := lab.Url + key
	log.Printf("request = %s", request)
	// получаем страницу
	response, err := http.Get(request)
	if err != nil {
		log.Fatal(err)
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
	result := make(map[string]string)
	// Find the review items
	doc.Find(".search-item").Each(func(i int, selection *goquery.Selection) {
		// For each item found, get the title
		title := selection.Find(".row").Find("h3").Text()
		price := selection.Find(".price-block .new-price").Text()

		if len(price) != 0 {
			result[title] = price
		}
	})

	return result
}
