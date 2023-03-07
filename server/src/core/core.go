package core

import (
	"comparisonLaboratories/src/global"
	"comparisonLaboratories/src/services/parse"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

var (
	Server = gin.Default()
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func InitConfig() {
	//KeyValuesDict = parse.ParseKeyValues()
	global.Laboratories = parse.ParseLabs()
}

func InitServer(app *gin.Engine) {
	app.Use(gin.Logger())
	app.Use(CORSMiddleware())
	app.StaticFS("/assets", http.Dir("../client/dist/assets"))
	app.LoadHTMLGlob("../client/dist/*.html")

	err := app.SetTrustedProxies([]string{"192.168.1.2"})
	if err != nil {
		return
	}
}

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
