package version2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(context *gin.Context) {
	result := make(chan gin.H)

	go func(context *gin.Context) {
		result <- gin.H{
			"message":  "pong",
			"fullpath": context.Request.URL.Path,
		}
	}(context.Copy())

	context.JSON(http.StatusOK, <-result)
}

func GetListAnalyses(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}
