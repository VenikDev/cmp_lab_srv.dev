package services

import (
	"comparisonLaboratories/src/core"
	"comparisonLaboratories/src/model"
	"errors"
)

// Получить список анализов для каждой лаборатории
func GetAnalysis(key string) (model.LabAndListAnalyses, error) {
	labsAndListTests := make(model.LabAndListAnalyses)
	fillMapAnalyses(labsAndListTests, key)

	// если нашли хоть бы для обной лаборатории
	if len(labsAndListTests) != 0 {
		return labsAndListTests, nil
	} else {
		return nil, errors.New("not found")
	}
}

// заполняем список с анализвми для каждоый лаборатории
func fillMapAnalyses(labsAndListTests model.LabAndListAnalyses, key string) {
	for _, lab := range core.Laboratories {
		listTests := core.GetListTests(key, lab)
		if len(listTests) != 0 {
			labsAndListTests[lab.GetName()] = listTests
		}
	}
}
