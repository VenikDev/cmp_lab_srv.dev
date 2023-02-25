package model

type ListAnalyses []Analysis
type LabAndListAnalyses map[string]ListAnalyses

type Analysis struct {
	Name  string
	Price int
}
