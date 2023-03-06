package parse

import (
	"comparisonLaboratories/src/global"
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

// getWorkDir
// This is a function in Go programming language called `getWorkDir()`. It returns a string value `wordDir` indicating the current working directory.
// The function uses the `os.Getwd()` function to get the current working directory and assigns it to the variable
// `wordDir`. It also checks for the error returned from the `os.Getwd()` function and if it is not `nil`,
// the function uses `panic()` to halt the program with an error message.
// Finally, the function returns the value of `wordDir`.
func getWorkDir() (wordDir string) {
	wordDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return wordDir
}

// pathToConfig
// This code defines a function called "pathToConfig" that returns a string representing the path to a configuration
// file. The function calls another function called "getWorkDir(
// )" to get the current working directory and concatenates it with the string constant "CONFIG_FOR_SEARCH".
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

// ParseLabs
// This code defines a function named "ParseLabs" that returns an array of "Laboratory" structs.
// The function opens a file containing JSON data,
// reads its contents and unmarshals it into an array of two "Laboratory" structs, which is returned.
// If there is an error opening the file or unmarshaling the JSON data, the function will panic.
func ParseLabs() (arrLabs []global.Laboratory) {
	pathToConfig := pathToConfig()

	jsonFile, err := os.Open(pathToConfig + JSON_LABORATORY)
	if err != nil {
		panic(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	// we initialize our KeyValue array
	arrLabs = make([]global.Laboratory, 2)
	err = json.Unmarshal(byteValue, &arrLabs)
	if err != nil {
		panic(err)
	}

	return arrLabs
}
