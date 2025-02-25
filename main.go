package main

import (
	"anothapp_update/database"
	"anothapp_update/helpers"
	"anothapp_update/models"
	"anothapp_update/repositories"
	"anothapp_update/services"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	mode := getEnvFile(os.Args[1:])
	errEnv := godotenv.Load(mode)

	if errEnv != nil {
		panic(errEnv)
	}
	database.Open()
	defer database.Close()

	updatedShows, deletedShows := updateShows()
	updatedSeasons, deletedSeasons := updateSeasons()
	helpers.SendTelegramMessage(helpers.FormatMessage(updatedShows, deletedShows, updatedSeasons, deletedSeasons))
}

func getEnvFile(args []string) string {
	if len(args) == 1 && args[0] == "prod" {
		return "prod.env"
	}
	return ".env"
}

func updateShows() ([]models.Show, []models.Show) {
	shows := repositories.GetShows()
	showsToUp, showsToDelete := services.CompareShows(shows)
	repositories.UpdateShows(showsToUp, showsToDelete)
	return showsToUp, showsToDelete
}

func updateSeasons() ([]models.Season, []models.Season) {
	seasons := repositories.GetSeasons()
	seasonsToUp, seasonsToDel := services.CompareSeasons(seasons)
	repositories.UpdateSeasons(seasonsToUp, seasonsToDel)
	return seasonsToUp, seasonsToDel
}
