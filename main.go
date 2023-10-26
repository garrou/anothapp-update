package main

import (
	"anothapp_update/database"
	"anothapp_update/repositories"
	"fmt"
)

func main() {

	database.Open()
	defer database.Close()

	seasonRows := repositories.GetSeasonsWithNoPicture()
	defer seasonRows.Close()

	var number, episode, epDuration, showId int
	var image string

	for seasonRows.Next() {

		err := seasonRows.Scan(&number, &episode, &epDuration, &image, &showId)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(number, episode, epDuration, showId, image)
	}
}
