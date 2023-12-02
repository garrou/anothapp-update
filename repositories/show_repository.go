package repositories

import (
	"anothapp_update/database"
	"anothapp_update/helpers"
	"anothapp_update/models"
	"database/sql"
	"fmt"
)

func GetShows() *sql.Rows {
	query := "SELECT id, kinds FROM shows"
	rows, err := database.Db.Query(query)

	if err != nil {
		panic(err)
	}
	return rows
}

func UpdateShows(shows []models.Show) {

	updated := len(shows)

	if updated == 0 {
		return
	}
	query := ""

	for _, s := range shows {
		query += fmt.Sprintf("UPDATE shows SET kinds = '%s' WHERE id = %d;\n", s.Kinds, s.Id)
	}
	if _, err := database.Db.Query(query); err != nil {
		panic(err)
	}
	helpers.SendTelegramMessage(fmt.Sprintf("%d updated show(s)", updated))
}
