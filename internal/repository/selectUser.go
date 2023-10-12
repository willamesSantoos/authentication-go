package repository

import (
	"github.com/willamesSantoos/authentication/internal/db"
	"github.com/willamesSantoos/authentication/internal/models"
)

func SelectUser(user models.User) (bool, error) {
	conn := db.GetInstace()

	rows, err := conn.Query(`SELECT * FROM users WHERE email = $1 AND password = $2`, user.Email, user.Password)

	for rows.Next() {
		return true, nil
	}

	return false, err
}
