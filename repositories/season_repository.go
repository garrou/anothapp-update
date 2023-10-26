package repositories

import (
	"anothapp_update/database"
	"database/sql"
)

func GetSeasonsWithNoPicture() *sql.Rows {
	queryStmt := `SELECT * FROM seasons WHERE image IS NULL`
	rows, err := database.Db.Query(queryStmt)

	if err != nil {
		panic(err.Error())
	}
	return rows
}
