package core

import (
	"comparisonLaboratories/src/model"
	"comparisonLaboratories/src/services/parse"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

// This code defines two variables:
//  1. `Server`: It is an instance of a `gin` framework's default router engine.
//     It is used to handle incoming HTTP requests.
//  2. `Laboratories`: It is a slice of `model.Laboratory` type. It is used to store a list of laboratory objects.
var (
	Server       = gin.Default()
	Laboratories []model.Laboratory
)

// CORSMiddleware
// This code defines a middleware function called "CORSMiddleware" that will be used
// in Gin web framework to handle Cross-Origin Resource Sharing (CORS) requests.
// The middleware sets the following headers:
// - Access-Control-Allow-Origin: allows requests from any origin(*).
// - Access-Control-Allow-Credentials: allows cookies to be sent in cross-origin requests.
// - Access-Control-Allow-Headers: specifies the allowed headers for a request.
// - Access-Control-Allow-Methods: specifies the allowed HTTP methods for a request.
// The middleware checks if the method of incoming request is OPTIONS, if so,
// returns a 204 status code and aborts the middleware chain. Otherwise,
// it proceeds to the next middleware and/or to the final handler.
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

// InitConfig
// The code defines a function called "InitConfig" which when called,
// assigns the value returned by the function "parse.ParseLabs()" to the variable "Laboratories".
// It's possible that there's a commented-out line that assigns a value to a different variable called "KeyValuesDict".
func InitConfig() {
	//KeyValuesDict = parse.ParseKeyValues()
	Laboratories = parse.ParseLabs()
}

// InitServer
// The code defines a function called `InitServer` that takes in a `*gin.Engine` as an argument.
// The function then sets up middleware for the `app` using `gin.Logger()` and a custom CORS middleware,
// and also sets up a route for serving static files located at `..
// client/dist/assets` and loads HTML templates located at `../client/dist/*.html`.
// The function also attempts to set a trusted proxy at IP `192.168.1.2`,
// which is used to correctly handle client IPs when the application is behind a reverse proxy.
// If an error occurs while setting the trusted proxies, the function simply returns without doing anything further.
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

// InitEnv
// This is a function called `InitEnv` that loads environment variables from a file named `.env`. Specifically,
// it uses the `godotenv` package to load the variables, and if an error occurs during the process,
// the function panics and displays the message "Error loading .env file".
func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
