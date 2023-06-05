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

type LabSearchResults struct {
	Name string
	Data ListAnalyses
}

func GetAnalyzes(labName string, document *goquery.Document, params Bundle) ListAnalyses {
	switch labName {
	case CITILAB:
		return GetAnalyzesCitilab(document, params)
	case INVITRO:
		return GetAnalyzesInvitro(document, params)
	case GEMOTEST:
		return GetAnalyzesGemotest(document, params)
	default:
		return nil
	}
}
