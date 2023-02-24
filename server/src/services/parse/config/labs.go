package config

type Laboratory struct {
	// Название
	Name string `json:"name"`
	// url офф сайта
	Url string `json:"url"`
	// параметры запроса
	Params []string `json:"params"`
}
