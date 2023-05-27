package transport

import (
	"cmp_lab/src/transport/v1"
	"cmp_lab/src/transport/v2"
	"github.com/gin-gonic/gin"
)

func InitRouters(app *gin.Engine) {
	app.GET("/", version1.GetIndexHtml)

	routerV1 := app.Group(API_V1)
	{
		routerV1.GET("/get_labs", version1.GetLabs)
		routerV1.GET("/get_names_labs", version1.GetLabsNames)
		routerV1.GET("/analysis", version1.GetListAnalyses)
		routerV1.GET("/get_default_city", version1.GetDefaultCity)
		routerV1.GET("/get_city_info", version1.GetCityInfo)
		routerV1.GET("/get_list_of_cities", version1.GetListCities)
		routerV1.GET("/get_popular", version1.GetPopular)
	}

	v2 := app.Group(API_V2)
	{
		v2.GET("/ping", version2.Ping)
		v2.POST("/analysis", version2.GetListAnalyses)
	}
}
