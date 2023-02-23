package parse

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func Parse(url string) map[string]string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
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
