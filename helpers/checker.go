package helpers

import (
	"anothapp_update/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
)

func CompareSeasons(rows *sql.Rows) []models.Season {
	var number, showId int
	var seasonInfos models.SeasonInfos
	var seasons []models.Season
	apiKey := os.Getenv("BETASERIES_KEY")

	for rows.Next() {

		err := rows.Scan(&number, &showId)
		if err != nil {
			panic(err.Error())
		}
		body := HttpGet(fmt.Sprintf("https://api.betaseries.com/shows/seasons?id=%d&key=%s", showId, apiKey))

		if err := json.Unmarshal(body, &seasonInfos); err != nil {
			panic(err.Error())
		}

		for _, s := range seasonInfos.Seasons {
			if s.Number == number {
				seasons = append(seasons, models.Season{
					ShowId:   showId,
					Number:   s.Number,
					Episodes: s.Episodes,
					Image:    s.Image,
				})
			}
		}
	}
	return seasons
}
