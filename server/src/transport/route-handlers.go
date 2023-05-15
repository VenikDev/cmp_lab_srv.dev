package transport

import (
	"comparisonLaboratories/src/algorithm"
	"comparisonLaboratories/src/clog"
	"comparisonLaboratories/src/global"
	"comparisonLaboratories/src/model"
	"comparisonLaboratories/src/model/city"
	"comparisonLaboratories/src/model/favorite"
	"comparisonLaboratories/src/model/labs"
	"comparisonLaboratories/src/redis"
	"comparisonLaboratories/src/services"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
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
	city := strings.ToLower(context.Query("city"))

	if city == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "City is not provided"})
		return
	}

	clog.Logger.Info("router/init_routers", "city", city)
	query := city + ":" + key

	jsonData, err := redis.GetAnalysisByCity(query)
	if err != nil {
		if key == "" {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found for the given city"})
			return
		}

		if err := redis.AddToPopular(key); err != nil {
			clog.Logger.Error("GetListAnalyses", "Couldn't save", key)
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to add data to redis"})
			return
		}

		result, err := services.GetLaboratoryAnalyses(key)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch analyses from service"})
			return
		}

		if err := redis.AddToPopular(key); err != nil {
			clog.Logger.Error("GetListAnalyses", "Couldn't save", key)
		}

		if err := redis.AddAnalysisByCity(query, result); err != nil {
			clog.Logger.Error("InitRouters", "No added data to redis")
		}

		context.JSON(http.StatusOK, result)
		return
	}

	if err := redis.AddToPopular(key); err != nil {
		clog.Logger.Error("GetListAnalyses", "Couldn't save", key)
	}

	var analysis model.LabAndListAnalyses
	if err := json.Unmarshal([]byte(jsonData), &analysis); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to unmarshal data"})
		return
	}

	context.JSON(http.StatusOK, analysis)
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
