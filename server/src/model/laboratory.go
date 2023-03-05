package model

import (
	"github.com/PuerkitoBio/goquery"
)

const (
	CITILAB  = "citilab"
	INVITRO  = "invitro"
	GEMOTEST = "gemotest"
)

type ILaboratory interface {
	GetAnalyzes(document *goquery.Document, key string) ListAnalyses
	GetName() string
	GetUrl() string
	GetParamForFind() string
}

type Laboratory struct {
	// Название
	Name string `json:"name"`
	// url офф сайта
	Url string `json:"url"`
	// параметры запроса
	ParamForFind string `json:"param_for_find"`
}

type LabSearchResults struct {
	Name string
	Data ListAnalyses
}

func GetAnalyzes(labName string, document *goquery.Document) ListAnalyses {
	switch labName {
	case CITILAB:
		return GetAnalyzesCitilab(document)
	case INVITRO:
		return GetAnalyzesInvitro(document)
	case GEMOTEST:
		return GetAnalyzesGemotest(document)
	default:
		return nil
	}
}

func (lab *Laboratory) GetName() string {
	return lab.Name
}

func (lab *Laboratory) GetUrl() string {
	return lab.Url
}

func (lab *Laboratory) GetParamForFind() string {
	return lab.ParamForFind
}
