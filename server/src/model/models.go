package model

type ListAnalyses []Analysis
type LabAndListAnalyses map[string]ListAnalyses
type AnalysesResponse []LabAndListAnalyses

type Analysis struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
