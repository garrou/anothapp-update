package repositories

import (
	"anothapp_update/database"
	"anothapp_update/helpers"
	"anothapp_update/models"
	"database/sql"
	"fmt"
)

func GetSeasons() *sql.Rows {
	queryStmt := "SELECT number, episode, image, show_id FROM seasons ORDER BY show_id, number"
	rows, err := database.Db.Query(queryStmt)

	if err != nil {
		panic(err)
	}
	return rows
}

func DeleteSeasons(seasons []models.Season) {

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
	helpers.SendTelegramMessage(fmt.Sprintf("%d deleted season(s)", len(seasons)))
}

func UpdateSeasons(seasons []models.Season) {

	if len(seasons) == 0 {
		return
	}
	query := ""

	for _, s := range seasons {
		if s.Image == "" {
			query += fmt.Sprintf("UPDATE seasons SET image = NULL, episode = %d WHERE show_id = %d AND number = %d;\n", s.Episodes, s.ShowId, s.Number)
		} else {
			query += fmt.Sprintf("UPDATE seasons SET image = '%s', episode = %d WHERE show_id = %d AND number = %d;\n", s.Image, s.Episodes, s.ShowId, s.Number)
		}
	}
	if _, err := database.Db.Query(query); err != nil {
		panic(err)
	}
	helpers.SendTelegramMessage(fmt.Sprintf("%d updated season(s)", len(seasons)))
}
