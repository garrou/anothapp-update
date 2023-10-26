package main

import (
	"anothapp_update/database"
	"anothapp_update/helpers"
	"anothapp_update/repositories"
)

func main() {

	database.Open()
	defer database.Close()

	seasonRows := repositories.GetSeasonsWithNoPicture()
	defer seasonRows.Close()

	seasons := helpers.CompareSeasons(seasonRows)
}
