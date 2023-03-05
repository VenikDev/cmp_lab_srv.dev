package transport

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(app *gin.Engine) {
	app.GET("/", GetIndexHtml)
	app.GET(API_V1+"/get_labs", GetLabs)
	app.GET(API_V1+"/analysis", GetListAnalyses)
}
