package version2

import (
	"cmp_lab/src/clog"
	"cmp_lab/src/global"
	"cmp_lab/src/model"
	"cmp_lab/src/model/responce"
	"cmp_lab/src/services/process"
	"cmp_lab/src/structs/opt"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

func Ping(ctx *gin.Context) {
	result := make(chan gin.H)

	go func(context *gin.Context) {
		clog.Info("[v2/router/ping]", "ping", context.Request.Host)
		result <- gin.H{
			"message":  "pong",
			"fullpath": context.Request.URL.Path,
		}
	}(ctx.Copy())

	ctx.JSON(http.StatusOK, <-result)
}

func GetListAnalyses(ctx *gin.Context) {
	result := make(chan responce.Response[model.LaboratoryAnalyzes])
	key := ctx.Query("key")
	go func(key string) {
		if key == "" {
			clog.Error("[v2/get_analysis]", "key", key)
			result <- responce.Response[model.LaboratoryAnalyzes]{
				Result: opt.Option[model.LaboratoryAnalyzes]{},
				Error: responce.Error{
					Code:    404,
					Message: "Key not found",
				},
			}
		}
	}(key)

	city := ctx.Query("city")
	go func(city string) {
		if city == "" {
			clog.Error("[v2/get_analysis]", "city", city)
			result <- responce.Response[model.LaboratoryAnalyzes]{
				Result: opt.Option[model.LaboratoryAnalyzes]{},
				Error: responce.Error{
					Code:    404,
					Message: "city not found",
				},
			}
		}
	}(city)

	parameters := model.Bundle{
		"city": city,
		"key":  key,
	}

	go func(parameters model.Bundle) {
		clog.Info("[router/parse]", "parsing", "start")

		//analysis := make(chan)

		for _, lab := range global.Laboratories {
			go func(lab global.Laboratory, parameters model.Bundle) {
				// formation an url
				url := fmt.Sprintf("%s?%s=%s", lab.Url, lab.ParamForFind, parameters["key"])
				clog.Info("[client/request]", "url to lab", url)

				req := fasthttp.AcquireRequest()
				defer fasthttp.ReleaseRequest(req)

				req.SetRequestURI(url)
				req.Header.SetMethod(fasthttp.MethodGet)

				// Set a custom timeout
				timeout := 5 * time.Second
				client := &fasthttp.Client{
					ReadTimeout:  timeout,
					WriteTimeout: timeout,
				}

				resp := fasthttp.AcquireResponse()
				defer fasthttp.ReleaseResponse(resp)

				err := client.DoTimeout(req, resp, timeout)
				if err != nil {
					clog.Error("[client/request]", "request to the lab ended with an error", err)
					return
				}

				parameters["lab"] = lab
				// Process response
				process.Body(resp.Body(), parameters, func(err error) {
					clog.Error("[client/request]", fmt.Sprintf("parsing on %s. Error", lab.Name), err)
				})

				clog.Info("[client/request]", fmt.Sprintf("status code from %s", lab.Name), resp.StatusCode())
			}(lab, parameters)
		}

		clog.Info("[router/parse]", "parsing", "end")
	}(parameters)

	badResponse := <-result
	switch badResponse.Error.Code {
	case 404:
		ctx.AbortWithStatusJSON(badResponse.Error.Code, badResponse.Error)
	default:
		ctx.JSON(200, gin.H{
			"msg": "success",
		})
	}

}
