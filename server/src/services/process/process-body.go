package process

import (
	"bytes"
	"cmp_lab/src/clog"
	"cmp_lab/src/global"
	"cmp_lab/src/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/liyue201/gostl/ds/vector"
	"gopkg.in/errgo.v2/errors"
)

func getAllUrlAnalysis(nameLab string, doc *goquery.Document) ([]string, error) {
	headerLog := fmt.Sprintf("[process/%s]", nameLab)
	clog.Info(headerLog, "parsing", "start")

	var result []string
	var err error

	switch nameLab {
	case model.CITILAB:
		result, err = parsingCitilab(doc)
	case model.INVITRO:
		result, err = parsingInvitro(doc)
	case model.GEMOTEST:
		result, err = parsingGemotest(doc)
	}

	return result, err
}

func parsingGemotest(doc *goquery.Document) ([]string, error) {
	result := vector.New[string](vector.WithCapacity(10))

	doc.Find(".analize-item .analize-item_narrow").Each(func(i int, selection *goquery.Selection) {
		selection.Find(".analize-item__title").Each(func(i int, selection *goquery.Selection) {
			attrText, isExist := selection.Find("a").Attr("href")
			if isExist {
				clog.Info("[parsing/citilab]", "url", attrText)
				result.PushBack(attrText)
			}
		})
	})

	return result.Data(), errors.New("items not found")
}

func parsingInvitro(doc *goquery.Document) ([]string, error) {
	result := vector.New[string](vector.WithCapacity(10))

	doc.Find(".search__result .iwg_container").Each(func(i int, selection *goquery.Selection) {
		attrText, isExist := selection.Find("a").Attr("href")
		if isExist {
			clog.Info("[parsing/citilab]", "url", attrText)
			result.PushBack(attrText)
		}
	})

	return result.Data(), errors.New("items not found")
}

func parsingCitilab(doc *goquery.Document) ([]string, error) {
	result := vector.New[string](vector.WithCapacity(10))

	container := doc.Find(".col-md-14 .row")

	container.Each(func(i int, selection *goquery.Selection) {
		attrText, isExist := selection.Find("h2").Find("a").Attr("href")
		if isExist {
			clog.Info("[parsing/citilab]", "url", attrText)
			result.PushBack(attrText)
		}
	})

	return result.Data(), errors.New("items not found")
}

func Body(body []byte, params model.Bundle, onFailCallback func(err error)) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		onFailCallback(err)
		return
	}

	nameLab := params["lab"].(global.Laboratory).Name
	urls, err := getAllUrlAnalysis(nameLab, doc)
	if err != nil {
		onFailCallback(err)
		return
	}
	urls = urls
}
