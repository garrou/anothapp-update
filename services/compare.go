package services

import (
	"anothapp_update/helpers"
	"anothapp_update/models"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

const BaseUrl = "https://api.betaseries.com/shows"

var apiKey = os.Getenv("BETASERIES_KEY")

func CompareShows(shows []models.Show) []models.Show {

	var toUpdate []models.Show
	var wg sync.WaitGroup

	for _, show := range shows {
		wg.Add(1)
		go func(show models.Show) {
			defer wg.Done()
			body := helpers.HttpGet(fmt.Sprintf("%s/display?id=%d", BaseUrl, show.Id), apiKey)
			current := models.ShowInfo{}

			if showErr := json.Unmarshal(body, &current); showErr != nil {
				panic(showErr)
			}
			fmt.Println(current)
			kinds := helpers.MapToString(current.Show.Kinds)

			if kinds != show.Kinds || show.Poster != current.Show.Images.Poster || show.Duration != current.Show.Duration {
				toUpdate = append(toUpdate, models.Show{
					Id:       show.Id,
					Kinds:    kinds,
					Poster:   current.Show.Images.Poster,
					Duration: current.Show.Duration,
				})
			}
		}(show)
	}
	wg.Wait()
	return toUpdate
}

func CompareSeasons(seasons []models.Season) ([]models.Season, []models.Season) {

	var previous int
	var toUpdate []models.Season
	var toDelete []models.Season
	var current models.SeasonInfos

	for _, season := range seasons {

		if previous != season.ShowId {
			body := helpers.HttpGet(fmt.Sprintf("%s/seasons?id=%d", BaseUrl, season.ShowId), apiKey)
			current.Seasons = nil

			if err := json.Unmarshal(body, &current); err != nil {
				panic(err)
			}
		}
		if season.Number > len(current.Seasons) {
			toDelete = append(toDelete, season)
			continue
		}
		currSeason := current.Seasons[season.Number-1]

		if season.Number == currSeason.Number && (season.Episodes != currSeason.Episodes || season.Image != currSeason.Image) {
			toUpdate = append(toUpdate, models.Season{
				ShowId:   season.ShowId,
				Number:   currSeason.Number,
				Episodes: currSeason.Episodes,
				Image:    currSeason.Image,
			})
		}
		previous = season.ShowId
	}
	return toUpdate, toDelete
}
