package model

import (
	"cmp_lab/src/algorithm"
	"cmp_lab/src/clog"
	"cmp_lab/src/global"
	"cmp_lab/src/herr"
	"cmp_lab/src/model/city"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
	"strings"
)

func GetAnalyzesInvitro(document *goquery.Document, params Bundle) ListAnalyses {
	result := make(ListAnalyses, 0)
	re := regexp.MustCompile("[0-9]+")

	// Find the review items
	document.Find(".iwg_margin").Each(func(i int,
		selection *goquery.Selection) {
		// For each item found, get the title
		tagWithTitle := selection.Find("a")
		title := tagWithTitle.Find("h3").Text()
		if title != "" {
			linkToAnalyses, _ := tagWithTitle.Attr("href")
			price := selection.Find(".search__result-order-meta-price ").Text()

			// The first line of code selects the text content of the `<p>` element within an HTML element with class
			// "search__result" using the Go package "selection". The result is assigned to the variable "description".
			// The second line of code trims any leading and trailing spaces from the "description" string using the
			// Go package "strings".
			description := selection.Find(".search__result").Find("p").Text()
			strings.Trim(description, " ")
			// если не пустые
			if price != "" {
				totalPrice, err := strconv.Atoi(re.FindString(price))
				herr.HandlerError(err, "Not parse price")
				if err == nil {
					idx := algorithm.LinearSearch(global.Laboratories, func(lab global.Laboratory) bool {
						if lab.Name == INVITRO {
							return true
						}
						return false
					})

					clog.Info("[get_analysis/invitro]", "url", linkToAnalyses)

					re := regexp.MustCompile(`/\d+/\d+/$`)
					tailUrl := re.FindString(linkToAnalyses)
					cityEn := params["city"].(city.City).NameEn
					linkToAnalyses := fmt.Sprintf("/analizes/for-doctors/%s%s", cityEn, tailUrl)

					result = append(result, Analysis{
						Name:        title,
						Price:       totalPrice,
						Description: description,
						OriginalURL: global.Laboratories[idx].Url + linkToAnalyses,
					})
				}
			}
		}
	})
	return result
}
