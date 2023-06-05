package services

import (
	"cmp_lab/src/clog"
	"cmp_lab/src/core"
	"cmp_lab/src/global"
	"cmp_lab/src/model"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"sync"
)

// resultDocument. Структура получения результатов запроса
type resultDocument struct {
	Name string
	Data *goquery.Document
}

// GetLaboratoryAnalyses
// Получить список анализов для каждой лаборатории
func GetLaboratoryAnalyses(params model.Bundle) (model.LabAndListAnalyses, error) {
	labsAndListTests := make(model.LabAndListAnalyses, 3)
	fillMapAnalyses(labsAndListTests, params)

	// если нашли хоть бы для обной лаборатории
	if len(labsAndListTests) != 0 {
		return labsAndListTests, nil
	} else {
		return nil, errors.New("not found")
	}
}

// SendRequests
// This code defines a function named "SendRequests" that sends requests to multiple APIs to get data for a given
// "key". It receives a channel of "resultDocument" type and "key" as an input parameter. It returns nothing (void).
// Inside the function, a loop is initiated to iterate over a slice of "Laboratories" belonging to the "global
// " variable. For each lab in the loop, a URL is created using the "CreateURLFrom" function from the "core" package.
// Then, an asynchronous goroutine is started using the "go" statement.
// Within this goroutine, a "resultDocument" type is created with the "Name" property set to the name of the current
// lab, and the "Data" property set to the HTML string returned by the "GetHtmlFrom" function from the "core" package
// using the URL. This "resultDocument" object then sent to the channel provided as an input parameter of the function.
// The code also includes a deferred close statement to close the channel when all requests have been sent.
func SendRequests(documentChannel chan resultDocument, params model.Bundle) {
	for _, lab := range global.Laboratories {
		url := core.CreateURLFrom(params, lab)
		clog.Info("[req/fill_map_analyses]", "Send request", url)

		go func(lab global.Laboratory, url string) {
			params := params
			data := core.GetHtmlFrom(url, lab, params)
			if data != nil {
				documentChannel <- resultDocument{
					Name: lab.Name,
					Data: data,
				}
			}
		}(lab, url)
	}
}

// fillMapAnalyses
func fillMapAnalyses(labsAndListTests model.LabAndListAnalyses, params model.Bundle) {
	var wg sync.WaitGroup
	sizeLabs := len(global.Laboratories)

	documentChannel := make(chan resultDocument, sizeLabs)

	params["key"] = strings.ReplaceAll(params["key"].(string), " ", "+")
	SendRequests(documentChannel, params)

	for idx := 0; idx < sizeLabs; idx++ {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()
			foundData := <-documentChannel

			clog.Info("[req/list_analysis]", "Received a list of analyzes from", foundData.Name)

			foundLaboratories := model.GetAnalyzes(foundData.Name, foundData.Data, params)
			labsAndListTests[idx] = model.LaboratoryAnalyzes{
				NameLab: foundData.Name,
				List:    foundLaboratories,
			}
		}(idx)
	}

	wg.Wait()
}

// GetNameLaboratories
// This function named "GetNameLaboratories" creates a new string slice of length 3, and then appends the names of all
// the laboratories from a global variable named "global.Laboratories" to that slice. Finally, the function returns
// the resulting string slice.
func GetNameLaboratories() []string {
	slice := make([]string, len(global.Laboratories))
	for idx, lab := range global.Laboratories {
		slice[idx] = lab.Name
	}

	return slice
}
