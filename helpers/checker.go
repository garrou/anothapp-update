package helpers

import (
	"anothapp_update/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
)

func RowsToSeasons(rows *sql.Rows) []models.Season {

	var number, episode, showId int
	var image interface{}
	var seasons []models.Season

	for rows.Next() {

		err := rows.Scan(&number, &episode, &image, &showId)

		if err != nil {
			panic(err.Error())
		}
		seasons = append(seasons, models.Season{
			Number:   number,
			Episodes: episode,
			Image:    fmt.Sprintf("%v", image),
			ShowId:   showId,
		})
	}
	return seasons
}

func CompareSeasons(seasons []models.Season) ([]models.Season, []models.Season) {

	var previous int
	var current models.SeasonInfos
	var toUpdate []models.Season
	var toDelete []models.Season
	apiKey := os.Getenv("BETASERIES_KEY")

	for _, season := range seasons {

		if previous != season.ShowId {
			body := HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/seasons?id=%d&key=%s", season.ShowId, apiKey))
			current.Seasons = nil

			if err := json.Unmarshal(body, &current); err != nil {
				panic(err.Error())
			}
		}
		if season.Number > len(current.Seasons) {
			toDelete = append(toDelete, season)
			continue
		}
		s := current.Seasons[season.Number-1]

		if season.Number == s.Number && (season.Episodes != s.Episodes || season.Image != s.Image) {
			toUpdate = append(toUpdate, models.Season{
				ShowId:   season.ShowId,
				Number:   s.Number,
				Episodes: s.Episodes,
				Image:    s.Image,
			})
		}
		previous = season.ShowId
	}
	return toUpdate, toDelete
}
