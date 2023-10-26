package models

type SeasonInfos struct {
	Seasons []Season `json:"seasons"`
}

type Season struct {
	ShowId   int
	Number   int    `json:"number"`
	Image    string `json:"image"`
	Episodes int    `json:"episodes"`
}
