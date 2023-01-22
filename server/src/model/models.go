package model

type Album struct {
	ID     uint64 `json:"ID"`
	Name   string `json:"name"`
	Rating uint8  `json:"rating"`
	Author string `json:"author"`
}

var Albums = [...]Album{
	{ID: 1, Name: "Rammstein", Rating: 5, Author: "Rammstein"},
	{ID: 2, Name: "Raize Raize", Rating: 5, Author: "Rammstein"},
	{ID: 3, Name: "Main tail", Rating: 5, Author: "Rammstein"},
}
