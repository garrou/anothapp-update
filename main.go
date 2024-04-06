package main

import (
	"anothapp_update/database"
	"anothapp_update/repositories"
	"anothapp_update/services"
	"fmt"
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
	// seasons := repositories.GetSeasons()
	// users := repositories.GetUsers()

	showsToUp := services.CompareShows(shows)
	// seasonsToUp, seasonsToDel := services.CompareSeasons(seasons)

	for _, show := range showsToUp {
		fmt.Println(show)
	}
}
