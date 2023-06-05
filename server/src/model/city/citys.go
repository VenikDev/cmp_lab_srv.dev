package city

type Coords struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type City struct {
	Coords     Coords `json:"coords"`
	District   string `json:"district"`
	Name       string `json:"name"`
	Population uint64 `json:"population"`
	Subject    string `json:"subject"`
	NameEn     string `json:"name_en"`
}
