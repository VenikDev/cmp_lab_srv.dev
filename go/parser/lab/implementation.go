package lab

import (
	"parser_labs/common/models"
	"parser_labs/common/parse"
)

var (
	ListLab []models.Laboratory = parse.ParseLabs()
)
