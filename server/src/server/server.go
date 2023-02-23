package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	Server = gin.Default()
)

func SetupServer(app *gin.Engine) {
	app.Use(gin.Logger())
	app.StaticFS("/assets", http.Dir("../client/dist/assets"))
	app.LoadHTMLGlob("../client/dist/*.html")

	err := app.SetTrustedProxies([]string{"192.168.1.2"})
	if err != nil {
		return
	}
}
