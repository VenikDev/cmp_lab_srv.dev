package controllers

import (
	"parser_labs/lab"
	"parser_labs/models/store"
	"strings"
)

func GetLink(params store.StrStore) (string, error) {
	builder := strings.Builder{}

	for _, laboratory := range lab.ListLab {
		if laboratory.Name == params["name"] {
			builder.WriteString(laboratory.Url)
			builder.WriteString("?")
			builder.WriteString(laboratory.ParamForFind)
			builder.WriteString("=")
		}
	}

	builder.WriteString(params["key"])

	return builder.String(), nil
}
