package model

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
)

func GetAnalyzesGemotest(document *goquery.Document) ListAnalyses {
	result := make(ListAnalyses, 0)
	re := regexp.MustCompile("[0-9]+")

	// Find the review items
	document.Find(".analize-item_narrow").Each(func(i int, selection *goquery.Selection) {
		// For each item found, get the title
		title := selection.Find(".analize-item__title").Find("a").Text()
		if title != "" {
			price := selection.Find(".add-to-cart__price").Text()

			// если не пустые
			if price != "" {
				totalPrice, err := strconv.Atoi(re.FindString(price))
				if err == nil {
					result = append(result, Analysis{
						Name:  title,
						Price: totalPrice,
					})
				}
			}
		}
	})
	return result
}
