package parse

import (
	"comparisonLaboratories/src/model"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

const (
	CONFIG_FOR_SEARCH = "\\config_for_search"
	JSON_KEY_WORD     = "\\key_words.json"
	JSON_LABORATORY   = "\\laboratory.json"
)

func getWorkDir() (wordDir string) {
	wordDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return wordDir
}

func pathToConfig() string {
	return getWorkDir() + CONFIG_FOR_SEARCH
}

func ParseKeyValues() (arrKeyValues []model.KeyValue) {
	pathToConfig := pathToConfig()

	jsonFile, err := os.Open(pathToConfig + JSON_KEY_WORD)
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our KeyValue array
	arrKeyValues = make([]model.KeyValue, 2)

	err = json.Unmarshal(byteValue, &arrKeyValues)
	if err != nil {
		panic(err)
	}

	return arrKeyValues
}

func ParseLabs() (arrLabs []model.Laboratory) {
	pathToConfig := pathToConfig()

	jsonFile, err := os.Open(pathToConfig + JSON_LABORATORY)
	if err != nil {
		panic(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	// we initialize our KeyValue array
	arrLabs = make([]model.Laboratory, 2)
	err = json.Unmarshal(byteValue, &arrLabs)
	if err != nil {
		panic(err)
	}

	return arrLabs
}
