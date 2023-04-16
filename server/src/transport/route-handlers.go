package transport

import (
	"comparisonLaboratories/src/algorithm"
	"comparisonLaboratories/src/clog"
	"comparisonLaboratories/src/global"
	"comparisonLaboratories/src/model"
	"comparisonLaboratories/src/model/favorite"
	"comparisonLaboratories/src/redis"
	"comparisonLaboratories/src/services"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"strings"
)

func GetIndexHtml(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}

func GetListAnalyses(context *gin.Context) {
	key := context.Query("key")
	city := context.Query("city")

	if city != "" {
		clog.Logger.Info("InitRouters", "city", city)
		query := city + ":" + key
		jsonData, err := redis.GetAnalysisByCity(query)

		if err != nil && key != "" {
			strings.Trim(key, " ")
			clog.Logger.Info("InitRouters", "key word", key)
			// add to redis for statistics
			err := redis.AddToPopular(key)
			if err != nil {
				clog.Logger.Error("GetListAnalyses", "Couldn't save", key)
			}
			result, err := services.GetLaboratoryAnalyses(key)
			if err == nil {
				err := redis.AddAnalysisByCity(query, result)
				if err != nil {
					clog.Logger.Info("InitRouters", "No added data to redis")
				}
				context.IndentedJSON(http.StatusOK, result)
				return
			}
		} else {
			_ = redis.AddToPopular(key)

			var analysis model.LabAndListAnalyses
			_ = json.Unmarshal([]byte(jsonData), &analysis)

			context.IndentedJSON(http.StatusOK, analysis)
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

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	result := algorithm.QuickSort(
		allFavorite,
		func(left favorite.Favorite, right favorite.Favorite) bool {
			return left.Count > right.Count
		},
	)
	if len(result) > 5 {
		result = result[:5]
	}
	context.IndentedJSON(http.StatusOK, result)
}
