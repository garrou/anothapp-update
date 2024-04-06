package main

import (
	"anothapp_update/database"
	"anothapp_update/helpers"
	"anothapp_update/repositories"
	"anothapp_update/services"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
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
	shows := repositories.GetShows()
	showsToUp := services.CompareShows(shows)
	isShowsUp := repositories.UpdateShows(showsToUp)

	if isShowsUp {
		helpers.SendTelegramMessage(helpers.Format("Series updated", showsToUp))
	} else {
		helpers.SendTelegramMessage("Series are up to date")
	}
}

func updateSeasons() {
	seasons := repositories.GetSeasons()
	seasonsToUp, seasonsToDel := services.CompareSeasons(seasons)
	isSeasonsUp := repositories.UpdateSeasons(seasonsToUp, seasonsToDel)

	if isSeasonsUp {
		msg := fmt.Sprintln(helpers.Format("Seasons updated", seasonsToUp), helpers.Format("Seasons deleted", seasonsToDel))
		helpers.SendTelegramMessage(msg)
	} else {
		helpers.SendTelegramMessage("Seasons are up to date")
	}
}
