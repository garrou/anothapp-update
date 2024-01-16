package main

import (
	"anothapp_update/database"
	"anothapp_update/helpers"
	"anothapp_update/repositories"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	arg := getArg(os.Args[1:])
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic(errEnv)
	}
	database.Open()
	defer database.Close()

	switch arg {
	case "shows":
		updateShows()
	case "seasons":
		updateSeasons()
	case "all":
		updateShows()
		updateSeasons()
	default:
		log.Fatal("Invalid argument")
	}
}

func getArg(args []string) string {
	if len(args) != 1 {
		log.Fatal("Needs one argument")
	}
	return args[0]
}

func updateShows() {
	showRows := repositories.GetShows()
	defer showRows.Close()

	shows := helpers.RowsToShows(showRows)
	showsToUp := helpers.CompareShows(shows)
	repositories.UpdateShows(showsToUp)
}

func updateSeasons() {
	seasonRows := repositories.GetSeasons()
	defer seasonRows.Close()

	seasons := helpers.RowsToSeasons(seasonRows)
	seasonsToUp, seasonsToDel := helpers.CompareSeasons(seasons)
	repositories.UpdateSeasons(seasonsToUp, seasonsToDel)
}
