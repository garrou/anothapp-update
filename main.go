package main

import (
	"anothapp_update/database"
	"anothapp_update/repositories"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic(errEnv)
	}
	database.Open()
	defer database.Close()

	shows := repositories.GetShows()
	seasons := repositories.GetSeasons()
	// users := repositories.GetUsers()

	showsToUp := services.CompareShows(shows)
	seasonsToUp, seasonsToDel := services.CompareSeasons(seasons)

	fmt.Println(shows, seasons)
}