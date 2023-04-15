package transport

import (
	"comparisonLaboratories/src/algorithm"
	"comparisonLaboratories/src/clog"
	"comparisonLaboratories/src/global"
	"comparisonLaboratories/src/model/favorite"
	"comparisonLaboratories/src/redis"
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
	city := context.Query("city")

	strings.Trim(key, " ")
	clog.Logger.Info("InitRouters", "key word", key)
	clog.Logger.Info("InitRouters", "city", city)

	if key != "" && city != "" {
		// add to redis for statistics
		err := redis.AddToPopular(key)
		if err != nil {
			clog.Logger.Error("GetListAnalyses", "Couldn't save", key)
		}
		result, err := services.GetLaboratoryAnalyses(key)
		if err == nil {
			context.IndentedJSON(http.StatusOK, result)
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, "Key or City isn't installed")
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

func GetPopular(context *gin.Context) {
	allFavorite, err := redis.GetFavorite()

	result := algorithm.QuickSort(
		allFavorite,
		func(left favorite.Favorite, right favorite.Favorite) bool {
			return left.Count > right.Count
		},
	)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
	}
	context.IndentedJSON(http.StatusOK, result[:5])
}
