package core

import (
	"comparisonLaboratories/src/services/parse/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

var (
	Server        = gin.Default()
	KeyValuesDict []config.KeyValue
	Laboratories  []config.Laboratory
)

func InitConfig() {
	KeyValuesDict = config.ParseKeyValues()
	Laboratories = config.ParseLabs()
}

func InitServer(app *gin.Engine) {
	app.Use(gin.Logger())
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
