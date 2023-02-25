package model

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
)

func GetAnalyzesCitilab(document *goquery.Document, key string) ListAnalyses {
	result := make(ListAnalyses, 0)
	re := regexp.MustCompile("[0-9]+")

	// Find the review items
	document.Find(".col-md-14 .row").Each(func(i int, selection *goquery.Selection) {
		// For each item found, get the title
		title := selection.Find("h2").Text()
		price := selection.Find(".price-block .new-price").Text()

		// если не пустые
		if price != "" && title != "" {
			totalPrice, err := strconv.Atoi(re.FindString(price))
			if err == nil {
				result = append(result, Analysis{
					Name:  title,
					Price: totalPrice,
				})
			}
		}
	})
	return result
}
