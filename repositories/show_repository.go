package repositories

import (
	"anothapp_update/database"
	"anothapp_update/helpers"
	"anothapp_update/models"
	"database/sql"
	"fmt"
)

func GetShows() []models.Show {
	query := "SELECT id, poster, kinds, duration FROM shows LIMIT 5"
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

func toShows(rows *sql.Rows) []models.Show {

	var id, duration int
	var poster string
	var kinds interface{}
	var shows []models.Show

	for rows.Next() {

		err := rows.Scan(&id, &poster, &kinds, &duration)

		if err != nil {
			panic(err)
		}

		if kinds == nil {
			kinds = ""
		}
		shows = append(shows, models.Show{
			Id:       id,
			Poster:   poster,
			Kinds:    fmt.Sprintf("%s", kinds),
			Duration: duration,
		})
	}
	return shows
}
