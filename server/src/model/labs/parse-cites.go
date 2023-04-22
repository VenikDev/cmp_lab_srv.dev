package labs

import (
	"bufio"
	"comparisonLaboratories/src/clog"
	"comparisonLaboratories/src/model/city"
	"comparisonLaboratories/src/model/paths"
	"fmt"
	"github.com/goccy/go-json"
	"os"
)

var Cities []city.City // assuming City is the type of the elements in the JSON array

func InitCities() {
	err := parseCities()
	if err != nil {
		clog.Logger.Fatal("InitCities", "failed to parse cities", err)
	}
}

func parseCities() error {
	pathToConfig := paths.PathToConfig()

	file, err := os.Open(pathToConfig + paths.JSON_USE_RUSSIAN_CITIES)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&Cities); err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}
	return nil
}