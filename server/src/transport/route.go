package transport

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(app *gin.Engine) {
	//groupApiV1 := app.Group(API_V1)
	app.GET(API_V1+"/", GetIndexHtml)
	app.GET(API_V1+"/get_labs", GetLabs)
	app.GET(API_V1+"/get_names_labs", GetLabsNames)
	app.GET(API_V1+"/analysis", GetListAnalyses)
	app.GET(API_V1+"/get_default_city", GetDefaultCity)
	app.GET(API_V1+"/get_city_info", GetCityInfo)
	app.GET(API_V1+"/get_list_of_cities", GetListCities)
	app.GET(API_V1+"/get_popular", GetPopular)
}
