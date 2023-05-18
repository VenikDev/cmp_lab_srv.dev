package tests

import (
	"comparisonLaboratories/src/global"
	"testing"
)

func TestGetName(t *testing.T) {
	lab := global.Laboratory{Name: "test name", Url: "http://example.com", ParamForFind: "foo=bar"}
	name := lab.GetName()
	if name != "test name" {
		t.Errorf("Expected GetName() to return 'test name', but got '%s'", name)
	}
}

func TestGetUrl(t *testing.T) {
	lab := global.Laboratory{Name: "test name", Url: "http://example.com", ParamForFind: "foo=bar"}
	url := lab.GetUrl()
	if url != "http://example.com" {
		t.Errorf("Expected GetUrl() to return 'http://example.com', but got '%s'", url)
	}
}

func TestGetParamForFind(t *testing.T) {
	lab := global.Laboratory{Name: "test name", Url: "http://example.com", ParamForFind: "foo=bar"}
	param := lab.GetParamForFind()
	if param != "foo=bar" {
		t.Errorf("Expected GetParamForFind() to return 'foo=bar', but got '%s'", param)
	}
}
