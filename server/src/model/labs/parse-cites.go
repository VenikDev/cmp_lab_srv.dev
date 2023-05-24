package labs

import (
	"bufio"
	"cmp_lab/src/clog"
	"cmp_lab/src/model/city"
	"cmp_lab/src/model/paths"
	"fmt"
	"github.com/goccy/go-json"
	"os"
)

var (
	Cities = InitCities()
)

func InitCities() []city.City {
	cities, err := parseCities()
	if err != nil {
		clog.Logger.Fatal("InitCities", "failed to parse cities", err)
	}

	return cities
}

func parseCities() ([]city.City, error) {
	pathToConfig := paths.PathToConfig()

	file, err := os.Open(pathToConfig + paths.JSON_USE_RUSSIAN_CITIES)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	var cities []city.City
	if err = decoder.Decode(&cities); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}
	return cities, nil
}
