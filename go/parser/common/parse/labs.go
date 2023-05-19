package parse

import (
	"github.com/goccy/go-json"
	"io"
	"os"
	"parser_labs/common/models"
	"parser_labs/common/paths"
)

// ParseLabs
// This code defines a function named "ParseLabs" that returns an array of "Laboratory" structs.
// The function opens a file containing JSON data,
// reads its contents and unmarshals it into an array of two "Laboratory" structs, which is returned.
// If there is an error opening the file or unmarshaling the JSON data, the function will panic.
func ParseLabs() []models.Laboratory {
	pathToConfig := paths.PathToConfig()

	jsonFile, err := os.Open(pathToConfig + paths.JSON_LABORATORY)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var arrLabs []models.Laboratory
	err = json.Unmarshal(byteValue, &arrLabs)
	if err != nil {
		panic(err)
	}

	return arrLabs
}
