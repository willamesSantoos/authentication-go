package repository

import (
	"log"

	"github.com/willamesSantoos/authentication/internal/db"
	"github.com/willamesSantoos/authentication/internal/models"
)

func SelectAll() (users []models.User, err error) {
	conn := db.GetInstace()

	rows, err := conn.Query(`SELECT * FROM users`)

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.Id, &user.Email, &user.Password)

		if err != nil {
			log.Fatalln(err)
		}

		users = append(users, user)
	}

	return users, err
}
