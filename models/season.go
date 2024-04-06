package models

import "fmt"

type SeasonInfos struct {
	Seasons []Season `json:"seasons"`
}

type Season struct {
	Number   int
	Episodes int
	Duration int
	Image    string
	ShowId   int
}

func (s Season) String() string {
	return fmt.Sprintf("[Serie %d - season %d]", s.ShowId, s.Number)
}
