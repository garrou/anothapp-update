package models

type ShowInfo struct {
	Show struct {
		Id    int               `json:"id"`
		Kinds map[string]string `json:"genres"`
		Seasons int `json:"seasons"`
		Status  string `json:"status"`
		Duration int `json:"length"`
		Images struct {
			Poster string `json:"poster"`
		} `json:"images"`
	} `json:"show"`
}

type Show struct {
	Id    int
	Poster string
	Kinds string
	Duration int
}
