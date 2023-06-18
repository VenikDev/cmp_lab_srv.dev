package core

import (
	"cmp_lab/src/core/middleware"
	"cmp_lab/src/global"
	"cmp_lab/src/herr"
	"cmp_lab/src/model/labs"
	"cmp_lab/src/model/paths"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"strings"
)

var (
	Server = gin.Default()
)

func InitConfig() {
	global.Laboratories = labs.ParseLabs()
}

func InitServer(app *gin.Engine) {
	var pathToStaticFiles strings.Builder
	pathToStaticFiles.WriteString(paths.GetWorkDir())
	pathToStaticFiles.WriteString(`/static/`)

	app.Use(middleware.Logger())
	app.Use(middleware.CORSMiddleware())

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
