package helpers

import (
	"anothapp_update/models"
	"database/sql"
	"fmt"
)

func RowsToShows(rows *sql.Rows) []models.Show {

	var id int
	var kinds interface{}
	var shows []models.Show

	for rows.Next() {

		err := rows.Scan(&id, &kinds)

		if err != nil {
			panic(err)
		}

		if kinds == nil {
			kinds = ""
		}
		shows = append(shows, models.Show{
			Id:    id,
			Kinds: fmt.Sprintf("%s", kinds),
		})
	}
	return shows
}

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
