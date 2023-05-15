package core

import (
	"comparisonLaboratories/src/global"
	"comparisonLaboratories/src/herr"
	"comparisonLaboratories/src/model/labs"
	"comparisonLaboratories/src/model/paths"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"strings"
)

var (
	Server = gin.Default()
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}

func InitConfig() {
	global.Laboratories = labs.ParseLabs()
}

func InitServer(app *gin.Engine) {
	var pathToStaticFiles strings.Builder
	pathToStaticFiles.WriteString(paths.GetWorkDir())
	pathToStaticFiles.WriteString(`/static/`)

	app.Use(gin.Logger())
	app.Use(CORSMiddleware())

	// static files
	app.StaticFS("/assets", http.Dir(pathToStaticFiles.String()+"assets"))
	app.LoadHTMLGlob(pathToStaticFiles.String() + "*.html")

	err := app.SetTrustedProxies([]string{"192.168.1.2"})
	if err != nil {
		return
	}
}

// InitEnv
// This Go function called InitEnv loads environment variables from a .env file using the godotenv package.
// If an error occurs during this process,
// it uses errorHandler (herr) package to handle the error and terminate the program.
func InitEnv() {
	err := godotenv.Load(paths.GetWorkDir() + `/.env`)
	herr.HandlerError(err, "Error loading .env file")
}
