package services

import (
	"comparisonLaboratories/src/clog"
	"comparisonLaboratories/src/core"
	"comparisonLaboratories/src/global"
	"comparisonLaboratories/src/model"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"sync"
)

// resultDocument. Структура получения результатов запроса
type resultDocument struct {
	Name string
	Data *goquery.Document
}

// GetLaboratoryAnalyses
// Получить список анализов для каждой лаборатории
func GetLaboratoryAnalyses(key string) (model.LabAndListAnalyses, error) {
	labsAndListTests := make(model.LabAndListAnalyses)
	fillMapAnalyses(labsAndListTests, key)

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
func SendRequests(documentChannel chan resultDocument, key string) {
	for _, lab := range global.Laboratories {
		url := core.CreateURLFrom(key, lab)
		clog.Logger.Info("fillMapAnalyses: ", "Send request", url)

		go func(nameLab string, url string) {
			documentChannel <- resultDocument{
				Name: nameLab,
				Data: core.GetHtmlFrom(url),
			}

		}(lab.Name, url)
	}
}

// fillMapAnalyses
// Функция заполняет массив структур ключами и анализами с помощью указанного URL.
// Для каждой лаборатории из списка отправляется запрос, полученные данные помещаются в documentChannel,
// после чего на основании содержимого массива labsAndListTests создаются анализы. После того, как все запросы
// будут обработаны, выполнится метод Wait, который ожидает, пока не завершатся все задания из директивы Add.
func fillMapAnalyses(labsAndListTests model.LabAndListAnalyses, key string) {
	var wg sync.WaitGroup
	sizeLabs := len(global.Laboratories)

	documentChannel := make(chan resultDocument, sizeLabs)

	SendRequests(documentChannel, key)

	for idx := 0; idx < sizeLabs; idx++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			foundData := <-documentChannel

			clog.Logger.Info("fillMapAnalyses: ", "Received a list of analyzes from", foundData.Name)

			foundLaboratories := model.GetAnalyzes(foundData.Name, foundData.Data)
			labsAndListTests[foundData.Name] = foundLaboratories
		}()
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
