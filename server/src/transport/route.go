package transport

import (
	"comparisonLaboratories/src/model"
	"comparisonLaboratories/src/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func SetupRouters(app *gin.Engine) {
	app.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	app.GET(API_V1+"/album/:id", func(context *gin.Context) {
		paramID := context.Param("id")

		if paramID != "" {
			id, err := strconv.Atoi(paramID)
			if err != nil {
				log.Fatalln("id album isn int")
				return
			}
			for _, a := range model.Albums {
				if a.ID == uint64(id) {
					context.IndentedJSON(http.StatusOK, a)
					return
				}
			}
		}

		context.IndentedJSON(
			http.StatusNotFound,
			http.StatusText(http.StatusNotFound))
	})

	app.GET(API_V1+"/albums", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, model.Albums)
	})

	app.POST(API_V1+"/analysis", func(context *gin.Context) {
		url := context.Query("url")

		if url != "" {
			result, err := services.GetAnalysis(url)
			if err == nil {
				context.IndentedJSON(http.StatusOK, result)
			} else {
				context.IndentedJSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
			}
		}
	})
}
