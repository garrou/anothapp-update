package repositories

import (
	"anothapp_update/database"
	"anothapp_update/models"
	"database/sql"
	"fmt"
)

func GetSeasons() []models.Season {
	query := "SELECT number, episodes, duration, image, show_id FROM seasons ORDER BY show_id, number"
	rows, err := database.Db.Query(query)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return toSeasons(rows)
}

func UpdateSeasons(toUpdate []models.Season, toDelete []models.Season) {
	if len(toUpdate)+len(toDelete) == 0 {
		return
	}
	deleteSeasons(toDelete)
	updateSeasons(toUpdate)
}

func toSeasons(rows *sql.Rows) []models.Season {
	var number, episodes, duration, showId int
	var image string
	var seasons []models.Season

	for rows.Next() {

		err := rows.Scan(&number, &episodes, &duration, &image, &showId)

		if err != nil {
			panic(err)
		}
		seasons = append(seasons, models.Season{
			Number:   number,
			Episodes: episodes,
			Duration: duration,
			Image:    image,
			ShowId:   showId,
		})
	}
	return seasons
}

func deleteSeasons(seasons []models.Season) {
	if len(seasons) == 0 {
		return
	}
	query := ""

	for _, s := range seasons {
		query += fmt.Sprintf("DELETE FROM seasons WHERE show_id = %d AND number = %d;\n", s.ShowId, s.Number)
	}
	if _, err := database.Db.Query(query); err != nil {
		panic(err)
	}
}

func updateSeasons(seasons []models.Season) {
	if len(seasons) == 0 {
		return
	}
	query := ""

	for _, s := range seasons {
		if s.Image == "" {
			query += fmt.Sprintf("UPDATE seasons SET episodes = %d WHERE show_id = %d AND number = %d;\n", s.Episodes, s.ShowId, s.Number)
		} else {
			query += fmt.Sprintf("UPDATE seasons SET image = '%s', episodes = %d WHERE show_id = %d AND number = %d;\n", s.Image, s.Episodes, s.ShowId, s.Number)
		}
	}
	if _, err := database.Db.Query(query); err != nil {
		panic(err)
	}
}
