package repositories

import (
	"anothapp_update/database"
	"anothapp_update/helpers"
	"anothapp_update/models"
	"database/sql"
	"fmt"
)

func GetShows() []models.Season {
	query := "SELECT id, poster, kinds, duration FROM shows"
	rows, err := database.Db.Query(query)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return toShows(rows)
}

func UpdateShows(shows []models.Show) {

	updated := len(shows)

	if updated == 0 {
		return
	}
	query := ""

	for _, s := range shows {
		query += fmt.Sprintf("UPDATE shows SET kinds = '%s', poster = '%s', duration = %d WHERE id = %d;\n", s.Kinds, s.Poster, s.Duration, s.Id)
	}
	if _, err := database.Db.Query(query); err != nil {
		panic(err)
	}
	helpers.SendTelegramMessage(fmt.Sprintf("%d updated show(s)", updated))
}

func toSeasons(rows *sql.Rows) []models.Season {

	var number, episodes, duration, showId int
	var image interface{}
	var seasons []models.Season

	for rows.Next() {

		err := rows.Scan(&number, &episodes, &duration, &image, &showId)

		if err != nil {
			panic(err)
		}
		if image == nil {
			image = ""
		}
		seasons = append(seasons, models.Season{
			Number:   number,
			Episodes: episode,
			Duration: duration,
			Image:    fmt.Sprintf("%s", image),
			ShowId:   showId,
		})
	}
	return seasons
}