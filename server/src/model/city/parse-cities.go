package city

import (
	"comparisonLaboratories/src/model/paths"
	"github.com/goccy/go-json"
	"io"
	"os"
)

func ParseCities() ([]City, error) {
	pathToConfig := paths.PathToConfig()

	jsonFile, err := os.Open(pathToConfig + paths.JSON_RUSSIAN_CITIES)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var arrCities []City
	err = json.Unmarshal(byteValue, &arrCities)
	if err != nil {
		return nil, err
	}

	return arrCities, nil
}
