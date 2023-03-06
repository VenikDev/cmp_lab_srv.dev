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
