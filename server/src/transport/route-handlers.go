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

// GetDefaultCity
// TODO for now default "Нижний Тагил"
func GetDefaultCity(context *gin.Context) {
	clog.Logger.Info("get default city")
	context.IndentedJSON(http.StatusOK, "Нижний Тагил")
}

// GetListCities
// TODO change on regis in future
func GetListCities(context *gin.Context) {
	clog.Logger.Info("get list of cities")
	context.IndentedJSON(http.StatusOK, global.Cities)

}
