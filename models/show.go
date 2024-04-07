package models

import "fmt"

type ShowInfo struct {
	Show struct {
		Id       int               `json:"id"`
		Title    string            `json:"title"`
		Kinds    map[string]string `json:"genres"`
		Seasons  string            `json:"seasons"`
		Status   string            `json:"status"`
		Duration string            `json:"length"`
		Images   struct {
			Poster interface{} `json:"poster"`
			Show   interface{} `json:"show"`
		} `json:"images"`
	} `json:"show"`
}

type Show struct {
	Id       int
	Title    string
	Poster   string
	Kinds    string
	Duration int
}

func (s Show) String() string {
	return fmt.Sprintf("[%d - %s]", s.Id, s.Title)
}
