package models

import "fmt"

type SeasonInfos struct {
	Seasons []Season `json:"seasons"`
}

type Season struct {
	Number   int    `json:"number"`
	Episodes int    `json:"episodes"`
	Image    string `json:"image"`
	ShowId   int
}

func (s Season) String() string {
	return fmt.Sprintf("[Serie %d - season %d]", s.ShowId, s.Number)
}
