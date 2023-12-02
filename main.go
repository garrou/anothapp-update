package main

import (
	"anothapp_update/database"
	"anothapp_update/helpers"
	"anothapp_update/repositories"

	"github.com/joho/godotenv"
)

func main() {

	errEnv := godotenv.Load()

	if errEnv != nil {
		panic(errEnv)
	}
	database.Open()
	defer database.Close()

	showRows := repositories.GetShows()
	defer showRows.Close()

	shows := helpers.RowsToShows(showRows)
	showsToUp := helpers.CompareShows(shows)
	repositories.UpdateShows(showsToUp)

	seasonRows := repositories.GetSeasons()
	defer seasonRows.Close()

	seasons := helpers.RowsToSeasons(seasonRows)
	seasonsToUp, seasonstoDel := helpers.CompareSeasons(seasons)
	repositories.UpdateSeasons(seasonsToUp, seasonstoDel)
}
