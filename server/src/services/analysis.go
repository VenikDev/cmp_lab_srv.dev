package services

import (
	"comparisonLaboratories/src/core"
	"comparisonLaboratories/src/model"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"log"
	"sync"
)

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

// Функция заполняет массив структур ключами и анализами с помощью указанного URL.
// Для каждой лаборатории из списка отправляется запрос, полученные данные помещаются в documentChannel,
// после чего на основании содержимого массива labsAndListTests создаются анализы. После того, как все запросы
// будут обработаны, выполнится метод Wait, который ожидает, пока не завершатся все задания из директивы Add.
func fillMapAnalyses(labsAndListTests model.LabAndListAnalyses, key string) {
	var wg sync.WaitGroup

	sizeLabs := len(labsAndListTests)

	documentChannel := make(chan struct {
		Name string
		Data *goquery.Document
	}, sizeLabs)
	defer close(documentChannel)

	for _, lab := range core.Laboratories {
		wg.Add(1)

		url := core.CreateURLFrom(key, lab)
		log.Println("Send request on", url)

		go func(nameLab string, url string) {

			documentChannel <- struct {
				Name string
				Data *goquery.Document
			}{Name: nameLab, Data: core.GetHtmlFrom(url)}

			wg.Done()
		}(lab.Name, url)
	}

	for idx := 0; idx < sizeLabs; idx++ {
		foundData := <-documentChannel

		log.Println("Received a list of analyzes from", foundData.Name)

		foundLaboratories := model.GetAnalyzes(foundData.Name, foundData.Data)
		labsAndListTests[foundData.Name] = foundLaboratories
	}

	wg.Wait()
}
