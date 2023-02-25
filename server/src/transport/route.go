package transport

import (
	"comparisonLaboratories/src/core"
	"comparisonLaboratories/src/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func InitRouters(app *gin.Engine) {
	app.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	app.GET(API_V1+"/get_labs", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, core.Laboratories)
	})

	app.GET(API_V1+"/analysis", func(context *gin.Context) {
		key := context.Query("key")
		strings.Trim(key, " ")
		log.Printf("key = %s", key)

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
