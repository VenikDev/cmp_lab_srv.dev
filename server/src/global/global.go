package global

type Laboratory struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	ParamForFind string `json:"param_for_find"`
}

func NewLaboratory(name, url, paramForFind string) *Laboratory {
	return &Laboratory{
		Name:         name,
		Url:          url,
		ParamForFind: paramForFind,
	}
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
