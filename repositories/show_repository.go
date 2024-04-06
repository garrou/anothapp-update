package repositories

import (
	"anothapp_update/database"
	"anothapp_update/models"
	"database/sql"
	"fmt"
)

func GetShows() []models.Show {
	query := "SELECT id, title, poster, kinds, duration FROM shows"
	rows, err := database.Db.Query(query)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return toShows(rows)
}

func UpdateShows(shows []models.Show) {
	if len(shows) == 0 {
		return
	}
	query := ""

	for _, s := range shows {
		query += fmt.Sprintf("UPDATE shows SET kinds = '%s', poster = '%s', duration = %d WHERE id = %d;\n", s.Kinds, s.Poster, s.Duration, s.Id)
	}
	if _, err := database.Db.Query(query); err != nil {
		panic(err)
	}
}

func toShows(rows *sql.Rows) []models.Show {
	var id, duration int
	var title, poster string
	var kinds interface{}
	var shows []models.Show

	for rows.Next() {

		err := rows.Scan(&id, &title, &poster, &kinds, &duration)

		if err != nil {
			panic(err)
		}

		shows = append(shows, models.Show{
			Id:       id,
			Title:    title,
			Poster:   poster,
			Kinds:    fmt.Sprintf("%s", kinds),
			Duration: duration,
		})
	}
	return shows
}
