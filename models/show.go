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
			Banner interface{} `json:"banner"`
			Box    interface{} `json:"box"`
		} `json:"images"`
	} `json:"show"`
}

func (s ShowInfo) GetImage() string {
	images := s.Show.Images

	if images.Poster != nil {
		return fmt.Sprintf("%s", images.Poster)
	}
	if images.Show != nil {
		return fmt.Sprintf("%s", images.Show)
	}
	if images.Banner != nil {
		return fmt.Sprintf("%s", images.Banner)
	}
	if images.Box != nil {
		return fmt.Sprintf("%s", images.Box)
	}
	return ""
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
