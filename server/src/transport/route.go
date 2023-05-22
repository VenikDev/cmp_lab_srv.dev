package transport

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(app *gin.Engine) {
	app.GET("/", GetIndexHtml)

	v1 := app.Group(API_V1)
	{
		v1.GET("/get_labs", GetLabs)
		v1.GET("/get_names_labs", GetLabsNames)
		v1.GET("/analysis", GetListAnalyses)
		v1.GET("/get_default_city", GetDefaultCity)
		v1.GET("/get_city_info", GetCityInfo)
		v1.GET("/get_list_of_cities", GetListCities)
		v1.GET("/get_popular", GetPopular)
	}

}
