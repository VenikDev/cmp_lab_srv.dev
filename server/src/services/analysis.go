package services

import (
	"comparisonLaboratories/src/services/parse"
	"errors"
)

func GetAnalysis(url string) (map[string]string, error) {
	result := parse.Parse(url)

	if len(result) != 0 {
		return result, nil
	} else {
		return nil, errors.New("not found")
	}
}
