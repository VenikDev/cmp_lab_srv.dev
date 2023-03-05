package core

import (
	"comparisonLaboratories/src/clog"
	"comparisonLaboratories/src/herr"
	"comparisonLaboratories/src/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// CreateURLFrom
// The code defines a function named "CreateURLFrom," which takes two parameters: a string named "key" and a "model.Laboratory" named "lab."
// Within this function, the code calls the "fmt.Sprintf" function and sets it up with three placeholders: "%s?%s=%s".
// The three placeholders represent the components that will be used to build a URL.
// The first "%s" placeholder represents the URL of the lab, which is obtained by calling the method
// "GetUrl()" on the "lab" object.
// The second "%s" placeholder represents the parameter that is used to find the laboratory details,
// which is obtained by calling the method "GetParamForFind()" on the "lab" object.
// Finally, the third "%s" placeholder represents the value of "key," which is passed in as a parameter.
// The function returns the resulting string, which is the URL created from the components: lab URL, find parameter,
// and key.
func CreateURLFrom(key string, lab model.Laboratory) string {
	return fmt.Sprintf("%s?%s=%s", lab.GetUrl(), lab.GetParamForFind(), key)
}

// GetHtmlFrom
// This is a Go function that takes a URL string as input and returns a *goquery.Document pointer.
// It first sends an HTTP GET request to the URL using the "http" package
// and stores the response and any error encountered in the process. "herr.
// HandlerError" is likely a customized function to handle error and it is called with the error and an empty string
// as arguments.
// It then checks if the response status code is not 200 (OK), using "response.StatusCode != 200",
// and raises an error if that's the case, using "clog.Logger.Error".
// Finally, it creates a new *goquery.Document from the response body using "goquery.
// NewDocumentFromReader" and returns it. The HTTP response body is closed before returning using "defer response.
// Body.Close()".
func GetHtmlFrom(url string) *goquery.Document {
	response, err := http.Get(url)
	herr.HandlerError(err, "")

	defer response.Body.Close()
	if response.StatusCode != 200 {
		clog.Logger.Error("status code error: %d %s", response.StatusCode, response.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	herr.HandlerError(err, "")

	return doc
}
