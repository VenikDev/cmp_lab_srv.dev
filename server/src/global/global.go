package global

type Laboratory struct {
	// Название
	Name string `json:"name"`
	// url офф сайта
	Url string `json:"url"`
	// параметры запроса
	ParamForFind string `json:"param_for_find"`
}

func (lab *Laboratory) GetName() string {
	return lab.Name
}

func (lab *Laboratory) GetUrl() string {
	return lab.Url
}

func (lab *Laboratory) GetParamForFind() string {
	return lab.ParamForFind
}

var (
	Laboratories []Laboratory
)
