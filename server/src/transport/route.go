package transport

import (
	"comparisonLaboratories/src/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func SetupRouters(app *gin.Engine) {
	app.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	app.GET("/api/v1/album/:id", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))
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

		context.IndentedJSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	})

	app.GET("/api/v1/albums", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, model.Albums)
	})

	app.POST("/api/v1/analysis", func(context *gin.Context) {
		url := context.Query("url")

		result := parse.Parse(url)
		if result != nil {
			context.IndentedJSON(http.StatusOK, result)
		} else {
			context.IndentedJSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		}
	})
}
