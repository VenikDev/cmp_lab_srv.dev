package process

import (
	"bytes"
	"cmp_lab/src/clog"
	"cmp_lab/src/global"
	"cmp_lab/src/model"
	"cmp_lab/src/structs/opt"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/liyue201/gostl/ds/vector"
	"gopkg.in/errgo.v2/errors"
	"regexp"
	"strconv"
	"strings"
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

	doc.Find(".analize_list").Find(".analize-item").Each(func(i int, selection *goquery.Selection) {
		attrText, isExist := selection.Find("a").Attr("href")
		if isExist {
			clog.Info("[parsing/gemotest]", "url", attrText)
			result.PushBack(attrText)
		}
	})

	if !result.Empty() {
		return result.Data(), nil
	} else {
		return nil, errors.New("items not found")
	}
}

func parsingInvitro(doc *goquery.Document) ([]string, error) {
	result := vector.New[string](vector.WithCapacity(10))

	doc.Find("div.iwg_margin").Each(func(i int, selection *goquery.Selection) {
		attrText, isExist := selection.Find("a.search__result-title").Attr("href")
		if isExist {
			clog.Info("[parsing/invitro]", "url", attrText)
			result.PushBack(attrText)
		}
	})

	if !result.Empty() {
		return result.Data(), nil
	} else {
		return nil, errors.New("items not found")
	}
}

func parsingCitilab(doc *goquery.Document) ([]string, error) {
	result := vector.New[string](vector.WithCapacity(10))

	container := doc.Find(".search-category")

	container.Find(".search-item").Each(func(i int, selection *goquery.Selection) {
		attrText, isExist := selection.Find("a").Attr("href")
		if isExist {
			clog.Info("[parsing/citilab]", "url", attrText)
			result.PushBack(attrText)
		}
	})

	if !result.Empty() {
		return result.Data(), nil
	} else {
		return nil, errors.New("items not found")
	}
}

func GetAllUrlFrom(body []byte, params model.Bundle) opt.Option[[]string] {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return opt.None[[]string]()
	}

	nameLab := params["lab"].(global.Laboratory).Name
	urls, err := getAllUrlAnalysis(nameLab, doc)
	if err != nil {
		return opt.None[[]string]()
	}

	return opt.Some(urls)
}

func GetDataAbout(body []byte, params model.Bundle) *model.Analysis {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil
	}

	nameLab := params["lab"].(global.Laboratory).Name
	switch nameLab {
	case model.CITILAB:
		{
			return parsingAnalysisCitilab(doc)
		}
	case model.GEMOTEST:
		{
			return parsingAnalysisGemotest(doc)
		}
	case model.INVITRO:
		{
			return parsingAnalysisInvitro(doc)
		}
	}
	return nil
}

func parsingAnalysisInvitro(doc *goquery.Document) *model.Analysis {
	re := regexp.MustCompile("[0-9]+")

	rawPrice := doc.Find(".info-block__price").Text()
	if rawPrice == "" {
		return nil
	}

	price, err := strconv.Atoi(re.FindString(rawPrice))
	if err != nil {
		clog.Error("[parsing/invitro]", "price parsing error", err)
		return nil
	}

	title := doc.Find(".title-block .title-block--img").Text()
	if title == "" {
		clog.Error("[parsing/invitro]", "title is empty", err)
		return nil
	}
	title = strings.ReplaceAll(title, "\t", "")
	title = strings.ReplaceAll(title, "\n", "")

	return &model.Analysis{
		Name:        title,
		Price:       price,
		Description: "",
		OriginalURL: "",
	}
}

func parsingAnalysisGemotest(doc *goquery.Document) *model.Analysis {
	re := regexp.MustCompile("[0-9]+")

	price, err := strconv.Atoi(re.FindString(doc.Find(".add-to-cart__price").Text()))
	if err != nil {
		clog.Error("[parsing/gemotest]", "price parsing error", err)
		return nil
	}
	title := doc.Find(".content-wrapper__title").Text()

	return &model.Analysis{
		Name:        title,
		Price:       price,
		Description: "",
		OriginalURL: "",
	}
}

func parsingAnalysisCitilab(doc *goquery.Document) *model.Analysis {
	re := regexp.MustCompile("[0-9]+")

	rawPrice := doc.Find(".new-price").Text()
	if rawPrice == "" {
		return nil
	}

	price, err := strconv.Atoi(re.FindString(rawPrice))
	if err != nil {
		clog.Error("[parsing/citilab]", "price parsing error", err)
		return nil
	}
	title := doc.Find(".mini-information").Find(".dop-name").Text()
	if title == "" {
		clog.Error("[parsing/citilab]", "title is empty", err)
		return nil
	}
	title = strings.ReplaceAll(title, "\t", "")
	title = strings.ReplaceAll(title, "\n", "")

	return &model.Analysis{
		Name:        title,
		Price:       price,
		Description: "",
		OriginalURL: "",
	}
}
