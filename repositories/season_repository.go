package repositories

import (
	"anothapp_update/database"
	"anothapp_update/models"
	"database/sql"
)

func GetSeasonsWithNoPicture() *sql.Rows {
	queryStmt := `SELECT number, show_id FROM seasons WHERE image IS NULL`
	rows, err := database.Db.Query(queryStmt)

	if err != nil {
		panic(err.Error())
	}
	return rows
}

func UpdateSeasons(seasons []models.Season) {

}
