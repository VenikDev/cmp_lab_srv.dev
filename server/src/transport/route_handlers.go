package transport

import (
	"comparisonLaboratories/src/clog"
	"comparisonLaboratories/src/global"
	"comparisonLaboratories/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetIndexHtml(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}

func GetListAnalyses(context *gin.Context) {
	key := context.Query("key")
	strings.Trim(key, " ")
	clog.Logger.Info("InitRouters", "key word", key)

	if key != "" {
		result, err := services.GetLaboratoryAnalyses(key)
		if err == nil {
			context.IndentedJSON(http.StatusOK, result)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

func GetLabs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, global.Laboratories)
}

func GetLabsNames(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, services.GetNameLaboratories())
}
