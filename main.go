package main

import (
	"anothapp_update/database"
	"anothapp_update/helpers"
	"anothapp_update/repositories"
	"anothapp_update/services"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	arg, mode := getArgs(os.Args[1:])
	errEnv := godotenv.Load(getEnvFile(mode))

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

func getEnvFile(mode string) string {
	if mode == "prod" {
		return "production.env"
	}
	return ".env"
}

func getArgs(args []string) (string, string) {
	size := len(args)

	if size == 0 {
		log.Fatal("Needs one argument")
	}
	if size == 2 {
		return args[0], args[1]
	}
	return args[0], "dev"
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
