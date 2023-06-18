package model

import (
	"cmp_lab/src/algorithm"
	"cmp_lab/src/clog"
	"cmp_lab/src/global"
	"cmp_lab/src/herr"
	"cmp_lab/src/model/city"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
	"strings"
)

const (
	PROFILE = "ПРОФИЛЬ"
)

// findProfile
func findProfile(str string) bool {
	return strings.Contains(str, PROFILE)
}

// GetAnalyzesCitilab
func GetAnalyzesCitilab(document *goquery.Document, params Bundle) ListAnalyses {
	listAnalysesResult := make(ListAnalyses, 0)
	re := regexp.MustCompile("[0-9]+")

	// The code uses the GoQuery library to parse a document and search for
	// elements that match the selector ".col-md-14 .row". For each matching element,
	// it then searches for a "h2" tag and extracts the title text.
	// If the title is not an existing profile and is not an empty string,
	// it proceeds to extract additional information such as a link to the analysis, a description, and the price,
	// which it converts to an integer using the "strconv.Atoi" function. If the price is not empty,
	// it creates an "Analysis" struct with the extracted information and adds it to the "listAnalysesResult" array.
	document.Find(".col-md-14 .row").Each(func(i int, selection *goquery.Selection) {
		// For each item found, get the title
		h2Tag := selection.Find("h2")

		title := h2Tag.Text()
		if !findProfile(title) || title != "" {
			linkToAnalyses, _ := h2Tag.Find("a").Attr("href")
			price := selection.Find(".price-block .new-price").Text()

			description := selection.Find(".description").Text()
			strings.Trim(description, " ")
			// если не пустые
			if price != "" {
				totalPrice, err := strconv.Atoi(re.FindString(price))
				herr.HandlerError(err, "Not parse price")
				if err == nil {
					idx := algorithm.LinearSearch(global.Laboratories, func(lab global.Laboratory) bool {
						if lab.Name == CITILAB {
							return true
						}
						return false
					})

					original := global.Laboratories[idx].Url + "/" + params["city"].(city.City).NameEn + linkToAnalyses
					clog.Info("[get_analysis/citilab]", "url", original)

					listAnalysesResult = append(listAnalysesResult, *NewAnalysis(title, totalPrice, description,
						original))
				}
			}
		}
	})
	return listAnalysesResult
}
