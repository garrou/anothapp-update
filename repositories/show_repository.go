package repositories

import (
	"anothapp_update/database"
	"anothapp_update/models"
	"database/sql"
	"fmt"
)

func GetShows() []models.Show {
	query := "SELECT id, title, poster, kinds, duration, seasons, country FROM shows"
	rows, err := database.Db.Query(query)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return toShows(rows)
}

func UpdateShows(toUpdate []models.Show, toDelete []models.Show) {
	if len(toUpdate)+len(toDelete) == 0 {
		return
	}
	updateShows(toUpdate)
	deleteShows(toDelete)
}

func updateShows(shows []models.Show) {
	if len(shows) == 0 {
		return
	}
	query := ""

	for _, s := range shows {
		query += fmt.Sprintf(
			"UPDATE shows SET kinds = '%s', poster = '%s', duration = %d, seasons = %d, country = '%s' WHERE id = %d;\n",
			s.Kinds, s.Poster, s.Duration, s.Seasons, s.Country, s.Id)
	}
	if _, err := database.Db.Query(query); err != nil {
		panic(err)
	}
}

func deleteShows(shows []models.Show) {
	if len(shows) == 0 {
		return
	}
	query := ""

	for _, s := range shows {
		query += fmt.Sprintf("DELETE shows WHERE id = %d;\n", s.Id)
	}
	if _, err := database.Db.Query(query); err != nil {
		panic(err)
	}
}

func toShows(rows *sql.Rows) []models.Show {
	var id, duration, seasons int
	var title, poster string
	var kinds, country interface{}
	var shows []models.Show

	for rows.Next() {

		err := rows.Scan(&id, &title, &poster, &kinds, &duration, &seasons, &country)

		if err != nil {
			panic(err)
		}

		shows = append(shows, models.Show{
			Id:       id,
			Title:    title,
			Poster:   poster,
			Kinds:    fmt.Sprintf("%s", kinds),
			Duration: duration,
			Seasons:  seasons,
			Country:  fmt.Sprintf("%s", country),
		})
	}
	return shows
}
