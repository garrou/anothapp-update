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

func CompareShows(shows []models.Show) ([]models.Show, []models.Show) {
	var toUpdate []models.Show
	var toDelete []models.Show
	var wg sync.WaitGroup
	var apiKey = os.Getenv("BETASERIES_KEY")

	for _, show := range shows {
		wg.Add(1)
		go func(show models.Show) {
			defer wg.Done()
			body, err := helpers.HttpGet(fmt.Sprintf("%s/display?id=%d", BaseUrl, show.Id), apiKey)

			if err != nil {
				helpers.SendTelegramMessage(fmt.Sprintf("Error while fetching show %d, reason %v", show.Id, err))
				panic(err)
			}
			current := models.ShowInfo{}

			if showErr := json.Unmarshal(body, &current); showErr != nil {
				helpers.SendTelegramMessage(fmt.Sprintf("Error during deserialize show %d, reason %v", show.Id, showErr))
				panic(showErr)
			}
			if current.Show.Id == 0 {
				toDelete = append(toDelete, show)
				return
			}
			kinds := helpers.MapToString(current.Show.Kinds)
			duration, _ := strconv.Atoi(current.Show.Duration)
			seasons := len(current.Show.Seasons)
			image := current.GetImage()

			if show.Poster != image ||
				show.Seasons != seasons ||
				show.Country != current.Show.Country ||
				(duration != 0 && show.Duration != duration) {
				toUpdate = append(toUpdate, models.Show{
					Id:       current.Show.Id,
					Title:    current.Show.Title,
					Kinds:    kinds,
					Poster:   image,
					Duration: duration,
					Seasons:  seasons,
					Country:  current.Show.Country,
				})
			}
		}(show)
	}
	wg.Wait()
	return toUpdate, toDelete
}

func CompareSeasons(seasons []models.Season) ([]models.Season, []models.Season) {
	var toUpdate []models.Season
	var toDelete []models.Season
	var current models.SeasonInfos
	var apiKey = os.Getenv("BETASERIES_KEY")
	var previous int

	for _, season := range seasons {
		if season.ShowId != previous {
			body, getErr := helpers.HttpGet(fmt.Sprintf("%s/seasons?id=%d", BaseUrl, season.ShowId), apiKey)

			if getErr != nil {
				helpers.SendTelegramMessage(fmt.Sprintf("Error with season %d, reason %v", season.ShowId, getErr))
				panic(getErr)
			}
			current.Seasons = nil

			if err := json.Unmarshal(body, &current); err != nil {
				helpers.SendTelegramMessage(fmt.Sprintf("Error during deserialize seasons of show %d, reason %v", season.ShowId, err))
				panic(err)
			}
		}
		if season.Number > len(current.Seasons) {
			toDelete = append(toDelete, season)
			continue
		}
		currSeason := current.Seasons[season.Number-1]

		if season.Number == currSeason.Number && (season.Episodes != currSeason.Episodes || (currSeason.Image != "" && season.Image != currSeason.Image)) {
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
