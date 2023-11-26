package helpers

import (
	"anothapp_update/models"
	"database/sql"
	"fmt"
)

func RowsToSeasons(rows *sql.Rows) []models.Season {

	var number, episode, showId int
	var image interface{}
	var seasons []models.Season

	for rows.Next() {

		err := rows.Scan(&number, &episode, &image, &showId)

		if err != nil {
			panic(err)
		}
		if image == nil {
			image = ""
		}
		seasons = append(seasons, models.Season{
			Number:   number,
			Episodes: episode,
			Image:    fmt.Sprintf("%s", image),
			ShowId:   showId,
		})
	}
	return seasons
}