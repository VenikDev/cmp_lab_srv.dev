package transport

import (
	"cmp_lab/src/algorithm"
	"cmp_lab/src/clog"
	"cmp_lab/src/global"
	"cmp_lab/src/model"
	"cmp_lab/src/model/city"
	"cmp_lab/src/model/favorite"
	"cmp_lab/src/model/labs"
	"cmp_lab/src/redis"
	"cmp_lab/src/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetIndexHtml(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}

// GetListAnalyses
// This is a Go function that handles an HTTP GET request to retrieve laboratory analyses for a given city and key.
// The function first checks if the city parameter is provided,
// if not it returns an error response using the gin library. If the city is provided,
// the function constructs a query from the city and key parameters to look up cached data in a Redis database.
// If the cached data is found,
// it adds the key to a Redis set of popular keys and returns the cached jsonData as a JSON response. Otherwise,
// it attempts to fetch the data from an external service using the services.GetLaboratoryAnalyses(key) function,
// caches the result in Redis,
// and adds the key to the Redis set of popular keys before returning the result as a JSON response.
// The function uses logging via the clog.Logger object to track execution steps and errors.
// The function expects a model.LabAndListAnalyses struct to represent the fetched analysis data and uses the json
// package to parse the jsonData response. Overall,
// the function provides a simple RESTful API endpoint for retrieving laboratory analyses that is backed by an
// external service and cache.
func GetListAnalyses(context *gin.Context) {
	key := strings.ToLower(context.Query("key"))
	cityForSearch := strings.ToLower(context.Query("city"))

	if key == "" {
		clog.Logger.Error("[router/get_list_ana]", "message", "Key is not provided")
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Key is not provided"})
		return
	}

	if cityForSearch == "" {
		clog.Logger.Error("[router/get_list_ana]", "message", "City is not provided")
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "City is not provided"})
		return
	}

	addToPopular := func(key string) {
		if err := redis.AddToPopular(key); err != nil {
			clog.Logger.Error("[router/get_list_ana]", "Couldn't save", key)
		}
	}

	clog.Logger.Info("[router/get_list_ana]", "cityForSearch", cityForSearch)
	query := cityForSearch + ":" + key

	jsonData, err := redis.GetAnalysisByCity(query)
	// if there isn't a redis, then
	if err != nil {
		result, err := services.GetLaboratoryAnalyses(key)
		if err != nil {
			headers := gin.H{
				"message": "Failed to fetch analyses from service",
				"error":   err.Error(),
			}

			context.AbortWithStatusJSON(http.StatusInternalServerError, headers)
			return
		}

		addToPopular(key)

		if err := redis.AddAnalysisByCity(query, result); err != nil {
			clog.Logger.Error("[router/get_list_ana]", "No added data to redis")
		}

		context.JSON(http.StatusOK, result)
	} else {
		addToPopular(key)

		var analysis model.LabAndListAnalyses
		if err := json.Unmarshal([]byte(jsonData), &analysis); err != nil {
			clog.Logger.Error("[router/get_list_ana]", "message", "Failed to unmarshal data")

			headers := gin.H{
				"message": "Failed to unmarshal data",
				"error":   err.Error(),
			}
			context.AbortWithStatusJSON(http.StatusInternalServerError, headers)
			return
		}

		context.JSON(http.StatusOK, analysis)
	}
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
	idx := algorithm.LinearSearch(labs.Cities, func(city city.City) bool {
		return city.Name == "Нижний Тагил"
	})

	defaultCity := labs.Cities[idx]
	clog.Logger.Info("[router/default_city]", defaultCity.Name)

	context.IndentedJSON(http.StatusOK, labs.Cities[idx])
}

// GetListCities
// TODO change on regis in future
func GetListCities(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, labs.Cities)
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

// GetCityInfo
// handles an HTTP GET request for city information.
// The function first retrieves the value of the city query parameter from the HTTP request using the Query method of
// the gin.Context object.
// If the city parameter is empty, the function calls a callback function callbackNotFound()
// and returns an error response with status code 404 and an error message ({ "error": "<city> not found" }).
// If the city parameter is not empty, the function searches for the corresponding city in an array called labs.
// Cities using linear search algorithm. If the city is not found, the function again calls callbackNotFound()
// and returns an error response.
// Finally, if the city is found,
// the function returns the city information as a JSON response with status code 200 using the IndentedJSON method of
// the gin.Context object.
func GetCityInfo(context *gin.Context) {
	cityForFound := context.Query("city")
	callbackNotFound := func() {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": cityForFound + " not found"})
	}

	// check
	if cityForFound == "" {
		callbackNotFound()
		return
	}

	// find
	idx := algorithm.LinearSearch(labs.Cities, func(city city.City) bool {
		return city.Name == cityForFound
	})

	if idx == -1 {
		callbackNotFound()
		return
	}

	context.IndentedJSON(http.StatusOK, labs.Cities[idx])
}
