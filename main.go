package main

import (
	"anothapp_update/database"
	"anothapp_update/helpers"
	"anothapp_update/repositories"
	"anothapp_update/services"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func main() {
	mode := getEnvFile(os.Args[1:])
	errEnv := godotenv.Load(mode)

	if errEnv != nil {
		panic(errEnv)
	}
	database.Open()
	defer database.Close()

	filename := fmt.Sprintf("logs/%s.txt", time.Now().Format("20060102"))
	helpers.Write(filename, fmt.Sprintf("File %s\n\n", mode))
	updateShows(filename)
	updateSeasons(filename)
}

func getEnvFile(args []string) string {
	if len(args) == 1 && args[0] == "prod" {
		return "prod.env"
	}
	return ".env"
}

func updateShows(filename string) {
	shows := repositories.GetShows()
	showsToUp := services.CompareShows(shows)
	repositories.UpdateShows(showsToUp)

	helpers.Write(filename, helpers.Format("Series updated", showsToUp))
}

func updateSeasons(filename string) {
	seasons := repositories.GetSeasons()
	seasonsToUp, seasonsToDel := services.CompareSeasons(seasons)
	repositories.UpdateSeasons(seasonsToUp, seasonsToDel)

	msg := fmt.Sprint(helpers.Format("Seasons updated", seasonsToUp), helpers.Format("Seasons deleted", seasonsToDel))
	helpers.Write(filename, msg)
}
