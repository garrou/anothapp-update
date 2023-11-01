package main

import (
	"anothapp_update/database"
	"anothapp_update/helpers"
	"anothapp_update/repositories"
)

func main() {

	database.Open()
	defer database.Close()

	seasonRows := repositories.GetSeasons()
	defer seasonRows.Close()

	seasons := helpers.RowsToSeasons(seasonRows)
	toUpdate, toDelete := helpers.CompareSeasons(seasons)
	repositories.DeleteSeasons(toDelete)
	repositories.UpdateSeasons(toUpdate)
}
