package model

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
)

func GetAnalyzesInvitro(document *goquery.Document) ListAnalyses {
	result := make(ListAnalyses, 0)
	re := regexp.MustCompile("[0-9]+")

	// Find the review items
	document.Find(".iwg_margin").Each(func(i int,
		selection *goquery.Selection) {
		// For each item found, get the title
		title := selection.Find("a").Find("h3").Text()
		if title != "" {
			price := selection.Find(".search__result-order-meta-price ").Text()

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
