package transport

import (
	"github.com/gin-gonic/gin"
)

// InitRouters
// The code defines a function called "InitRouters" that takes a Gin engine pointer named "app" as an input parameter.
// The function uses this app to define three routes using the HTTP GET request method:
// - "/" : mapping to the function "GetIndexHtml"
// - "/api/v1/get_labs": mapping to the function "GetLabs"
// - "/api/v1/analysis": mapping to the function "GetListAnalyses"
func InitRouters(app *gin.Engine) {
	app.GET("/", GetIndexHtml)
	app.GET(API_V1+"/get_labs", GetLabs)
	app.GET(API_V1+"/get_names_labs", GetLabsNames)
	app.GET(API_V1+"/analysis", GetListAnalyses)
	app.GET(API_V1+"/get_default_city", GetDefaultCity)
	app.GET(API_V1+"/get_list_of_cities", GetListCities)
}
