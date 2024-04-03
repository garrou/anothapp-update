package models

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