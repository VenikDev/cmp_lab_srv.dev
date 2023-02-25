package core

import (
	"comparisonLaboratories/src/model"
	"comparisonLaboratories/src/services/parse"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

var (
	Server        = gin.Default()
	KeyValuesDict []model.KeyValue
	Laboratories  []model.Laboratory
)

func InitConfig() {
	KeyValuesDict = parse.ParseKeyValues()
	Laboratories = parse.ParseLabs()
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
