package repositories

import (
	"anothapp_update/database"
	"anothapp_update/models"
	"database/sql"
	"fmt"
)

func GetSeasonsWithNoPicture() *sql.Rows {
	queryStmt := "SELECT number, episode, image, show_id FROM seasons WHERE image IS NULL"
	rows, err := database.Db.Query(queryStmt)

	if err != nil {
		panic(err.Error())
	}
	return rows
}

func UpdateSeasons(seasons []models.Season) {

	query := ""

	for _, s := range seasons {
		query += fmt.Sprintf("UPDATE seasons SET image = '%s', episode = %d WHERE show_id = %d AND number = %d;\n", s.Image, s.Episodes, s.ShowId, s.Number)
	}
	if _, err := database.Db.Query(query); err != nil {
		panic(err)
	}
}
