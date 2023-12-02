package models

type ShowInfo struct {
	Show struct {
		Id    int               `json:"id"`
		Kinds map[string]string `json:"genres"`
	} `json:"show"`
}

type Show struct {
	Id    int    `json:"id"`
	Kinds string `json:"genres"`
}
