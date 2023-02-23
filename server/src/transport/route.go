package transport

import (
	"comparisonLaboratories/src/services"
	"comparisonLaboratories/src/services/parse/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouters(app *gin.Engine) {
	app.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	app.GET(API_V1+"/get_labs", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, config.ParseLabs())
	})

	app.GET(API_V1+"/analysis", func(context *gin.Context) {
		key := context.Query("key")

		if key != "" {
			result, err := services.GetAnalysis(key)
			if err == nil {
				context.IndentedJSON(http.StatusOK, result)
				return
			}
		}

		context.IndentedJSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	})
}
