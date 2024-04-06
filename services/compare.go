package services

import (
	"anothapp_update/helpers"
	"anothapp_update/models"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
)

const BaseUrl = "https://api.betaseries.com/shows"

func CompareShows(shows []models.Show) []models.Show {
	var toUpdate []models.Show
	var wg sync.WaitGroup
	var apiKey = os.Getenv("BETASERIES_KEY")

	for _, show := range shows {
		wg.Add(1)
		go func(show models.Show) {
			defer wg.Done()
			body := helpers.HttpGet(fmt.Sprintf("%s/display?id=%d", BaseUrl, show.Id), apiKey)
			current := models.ShowInfo{}

			if showErr := json.Unmarshal(body, &current); showErr != nil {
				panic(showErr)
			}
			if current.Show.Id == 0 {
				panic("Error during deserialization")
			}
			kinds := helpers.MapToString(current.Show.Kinds)
			duration, _ := strconv.Atoi(current.Show.Duration)

			if show.Poster != current.Show.Images.Poster || show.Duration != duration {
				toUpdate = append(toUpdate, models.Show{
					Id:       current.Show.Id,
					Title:    current.Show.Title,
					Kinds:    kinds,
					Poster:   current.Show.Images.Poster,
					Duration: duration,
				})
			}
		}(show)
	}
	wg.Wait()
	return toUpdate
}

func CompareSeasons(seasons []models.Season) ([]models.Season, []models.Season) {
	var toUpdate []models.Season
	var toDelete []models.Season
	var wg sync.WaitGroup
	var apiKey = os.Getenv("BETASERIES_KEY")

	for _, season := range seasons {
		wg.Add(1)
		go func(show models.Season) {
			defer wg.Done()
			current := models.SeasonInfos{}
			length := len(current.Seasons)

			if length == 0 || (current.Seasons[0].ShowId != season.ShowId) {
				body := helpers.HttpGet(fmt.Sprintf("%s/seasons?id=%d", BaseUrl, season.ShowId), apiKey)
				current.Seasons = nil

				if err := json.Unmarshal(body, &current); err != nil {
					panic(err)
				}
			}
			if season.Number > len(current.Seasons) {
				toDelete = append(toDelete, season)
				return
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
		}(season)
	}
	wg.Wait()
	return toUpdate, toDelete
}
