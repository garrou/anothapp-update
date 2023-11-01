package models

type SeasonInfos struct {
	Seasons []Season `json:"seasons"`
}

type Season struct {
	Number   int    `json:"number"`
	Episodes int    `json:"episodes"`
	Image    string `json:"image"`
	ShowId   int
}
