package models

type Laboratory struct {
	// Название
	Name string `json:"name"`
	// url офф сайта
	Url string `json:"url"`
	// параметры запроса
	ParamForFind string `json:"param_for_find"`
}
