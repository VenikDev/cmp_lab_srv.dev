package model

import "github.com/google/uuid"

// ListAnalyses which is just a slice (array) of "Analysis" objects
type ListAnalyses []Analysis

// LabAndListAnalyses which is a map where the keys are strings and the values are ListAnalyses objects
type LabAndListAnalyses []LaboratoryAnalyzes

// AnalysesResponse which is a slice of LabAndListAnalyses objects.
type AnalysesResponse []LabAndListAnalyses

// Analysis
// The code defines a struct named "Analysis" with four fields: "Name" (of type string), "Price" (of type int),
// "Description" (of type string), and "OriginalURL" (of type string).
// Each field also has a tag named "json" with a corresponding value (i.e. "name", "price", "description",
// and "original_url"). This is used when encoding and decoding the struct to/from a JSON format,
// as the tag determines the JSON key that corresponds to each struct field.
type Analysis struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	OriginalURL string `json:"original_url"`
}

func NewAnalysis(name string, price int, description string, originalURL string) *Analysis {
	return &Analysis{
		ID:          uuid.New().String(),
		Name:        name,
		Price:       price,
		Description: description,
		OriginalURL: originalURL,
	}
}

type LaboratoryAnalyzes struct {
	NameLab string     `json:"name_lab"`
	List    []Analysis `json:"list"`
}
