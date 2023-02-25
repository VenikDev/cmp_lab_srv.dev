package model

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
	"strings"
)

const (
	PROFILE = "ПРОФИЛЬ"
)

func findProfile(str string) bool {
	return strings.Contains(str, PROFILE)
}

func GetAnalyzesCitilab(document *goquery.Document) ListAnalyses {
	result := make(ListAnalyses, 0)
	re := regexp.MustCompile("[0-9]+")

	// Find the review items
	document.Find(".col-md-14 .row").Each(func(i int, selection *goquery.Selection) {
		// For each item found, get the title
		title := selection.Find("h2").Text()
		if !findProfile(title) || title != "" {
			price := selection.Find(".price-block .new-price").Text()

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
