package repositories

import (
	"anothapp_update/database"
	"anothapp_update/models"
	"database/sql"
)

func GetUsers() []models.User {
	query := "SELECT id FROM users"
	rows, err := database.Db.Query(query)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return toUsers(rows)
}

func toUsers(rows *sql.Rows) []models.User {

	var id string
	var users []models.User

	for rows.Next() {

		err := rows.Scan(&id)

		if err != nil {
			panic(err)
		}
		users = append(users, models.User{
			Id: id,
		})
	}
	return users
}
