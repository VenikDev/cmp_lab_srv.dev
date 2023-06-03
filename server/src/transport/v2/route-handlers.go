package version2

import (
	"cmp_lab/src/algorithm"
	"cmp_lab/src/clog"
	"cmp_lab/src/global"
	"cmp_lab/src/model"
	"cmp_lab/src/model/responce"
	"cmp_lab/src/services/process"
	"cmp_lab/src/structs/opt"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/liyue201/gostl/ds/vector"
	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/iter"
	"net/http"
	"sync"
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

func GetAllAnalyses(ctx *gin.Context) {
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

	chAnalysis := make(chan model.Analysis, 10)
	var wg conc.WaitGroup

	for _, lab := range global.Laboratories {
		lab := lab
		parameters := parameters

		wg.Go(func() {
			parameters["lab"] = lab

			// formation an url
			url := GetUrl(parameters)
			clog.Info("[client/request]", "url to lab", url)

			resp, err := resty.New().R().Get(url)
			if err == nil {
				// Process response
				process.GetAllUrlFrom(resp.Body(), parameters).ProcessIfHas(func(urls *[]string) {
					iter.ForEach(*urls, func(urlToAnalysis *string) {
						fullUrl := fmt.Sprintf("%s%s", lab.Url, *urlToAnalysis)
						clog.Info("[request/get_data]", "send request to", fullUrl)

						resp, err := resty.New().R().Get(fullUrl)
						if err == nil {
							body := resp.Body()

							data := process.GetDataAbout(body, parameters)
							if data != nil {
								data.OriginalURL = fullUrl
								chAnalysis <- *data
							}
						}
					})
				})
				clog.Info("[client/request]", fmt.Sprintf("status code from %s", lab.Name), resp.StatusCode())
			}
		})
	}

	for analysis := range chAnalysis {
		clog.Info("[ch/analysis]", "analysis", analysis.Name)
		clog.Info("[ch/analysis/price]", "price", analysis.Price)
	}

	wg.Wait()

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

func GetUrl(parameters model.Bundle) string {
	lab := parameters["lab"].(global.Laboratory)
	key := parameters["key"].(string)

	return fmt.Sprintf("%s/search/?%s=%s", lab.Url, lab.ParamForFind, key)
}

func GetAnalysisByLab(ctx *gin.Context) {
	parLab := ctx.Param("lab")

	var lab global.Laboratory
	{
		idxLab := algorithm.LinearSearch(global.Laboratories, func(laboratory global.Laboratory) bool {
			if laboratory.Name == parLab {
				return true
			}

			return false
		})

		lab = global.Laboratories[idxLab]
	}

	key := ctx.Param("key")
	city := ctx.Param("city")

	parameters := model.Bundle{
		"lab":  lab,
		"city": city,
		"key":  key,
	}

	var listAnalysis struct {
		Result vector.Vector[model.Analysis]
		Mutex  sync.Mutex
	}

	switch parLab {
	case model.CITILAB:
		{
			for page := 1; ; page += 1 {
				pageUrl := fmt.Sprintf("%s/search/?%s=%s&s=page-%d", lab.Url, lab.ParamForFind, key, page)

				resp, err := resty.New().R().Get(pageUrl)
				if err != nil || resp.StatusCode() != http.StatusOK {
					break
				}

				// Process response
				urls := process.GetAllUrlFrom(resp.Body(), parameters)
				if urls.IsNone() {
					break
				}

				iter.ForEach(*urls.Value, func(urlToAnalysis *string) {
					fullUrl := fmt.Sprintf("%s%s", lab.Url, *urlToAnalysis)
					clog.Info("[request/get_data]", "send request to", fullUrl)

					resp, err := resty.New().R().Get(fullUrl)
					if err == nil {
						body := resp.Body()

						data := process.GetDataAbout(body, parameters)
						if data != nil {
							data.OriginalURL = fullUrl
							clog.Info("[req/citilab]", "name analysis", data.Name)

							listAnalysis.Mutex.Lock()
							listAnalysis.Result.PushBack(*data)
							listAnalysis.Mutex.Unlock()
						}
					}
				})
			}
		}
	case model.GEMOTEST:
		{

		}
	case model.INVITRO:
		{
			for page := 1; page <= 30; page += 1 {
				pageUrl := fmt.Sprintf("%s/search/?%s=%s&PAGEN_5=%d", lab.Url, lab.ParamForFind, key, page)

				resp, err := resty.New().R().Get(pageUrl)
				if err != nil || resp.StatusCode() != http.StatusOK {
					break
				}

				// Process response
				urls := process.GetAllUrlFrom(resp.Body(), parameters)
				if urls.IsNone() {
					break
				}

				iter.ForEach(*urls.Value, func(urlToAnalysis *string) {
					fullUrl := fmt.Sprintf("%s%s", lab.Url, *urlToAnalysis)
					clog.Info("[request/get_data]", "send request to", fullUrl)

					resp, err := resty.New().R().Get(fullUrl)
					if err == nil {
						body := resp.Body()

						data := process.GetDataAbout(body, parameters)
						if data != nil {
							data.OriginalURL = fullUrl
							clog.Info("[req/invitro]", "name analysis", data.Name)

							listAnalysis.Mutex.Lock()
							listAnalysis.Result.PushBack(*data)
							listAnalysis.Mutex.Unlock()
						}
					}
				})
			}
		}
	}

	ctx.JSON(http.StatusOK, listAnalysis.Result.Data())
}
