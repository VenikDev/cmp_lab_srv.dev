package city

type City struct {
	Coords     map[string]string `json:"coords"`
	District   string            `json:"district"`
	Name       string            `json:"name"`
	Population int               `json:"population"`
	Subject    string            `json:"subject"`
}
