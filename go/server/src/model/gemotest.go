package model

import (
	"cmp_lab/src/algorithm"
	"cmp_lab/src/common/models"
	"cmp_lab/src/global"
	"cmp_lab/src/herr"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
)

func GetAnalyzesGemotest(document *goquery.Document) ListAnalyses {
	result := make(ListAnalyses, 0)
	re := regexp.MustCompile("[0-9]+")

	// This code is using the Go programming language and the goquery library to scrape data from an HTML document.
	// It is finding all elements with the class "analize-item_narrow" and then running a function for each of them.
	// For each element, it is finding the text of an element with the class "analize-item__title" inside it,
	// which represents the title of a product.
	// It is also finding the text of an element with the class "add-to-cart__price",
	// which represents the price of that product.
	// If both the title and price are not empty,
	// it is using a regular expression to extract the numerical value from the price string,
	// converting it to an integer, and then appending the title and price to a result slice of Analysis structs.
	document.Find(".analize-item_narrow").Each(func(i int, selection *goquery.Selection) {
		// For each item found, get the title
		tagA := selection.Find(".analize-item__title").Find("a")
		title := tagA.Text()
		if title != "" {
			price := selection.Find(".add-to-cart__price").Text()

			// если не пустые
			if price != "" {
				linkToAnalyses, _ := tagA.Attr("href")

				totalPrice, err := strconv.Atoi(re.FindString(price))
				herr.HandlerError(err, "Not parse price")
				if err == nil {
					idx := algorithm.LinearSearch(global.Laboratories, func(lab models.Laboratory) bool {
						if lab.Name == GEMOTEST {
							return true
						}
						return false
					})
					result = append(result, Analysis{
						Name:        title,
						Price:       totalPrice,
						Description: "",
						OriginalURL: global.Laboratories[idx].Url + linkToAnalyses,
					})
				}
			}
		}
	})
	return result
}
