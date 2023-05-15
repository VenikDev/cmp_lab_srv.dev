package paths

import "os"

// GetWorkDir
// This is a function in Go programming language called `GetWorkDir()`. It returns a string value `wordDir` indicating the current working directory.
// The function uses the `os.Getwd()` function to get the current working directory and assigns it to the variable
// `wordDir`. It also checks for the error returned from the `os.Getwd()` function and if it is not `nil`,
// the function uses `panic()` to halt the program with an error message.
// Finally, the function returns the value of `wordDir`.
func GetWorkDir() (wordDir string) {
	wordDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return wordDir
}

// PathToConfig
// This code defines a function called "PathToConfig" that returns a string representing the path to a configuration
// file. The function calls another function called "GetWorkDir(
// )" to get the current working directory and concatenates it with the string constant "CONFIG_FOR_SEARCH".
func PathToConfig() string {
	return GetWorkDir() + CONFIG_FOR_SEARCH
}

// /usr/src/app\config_for_search\cities-to-use.json
