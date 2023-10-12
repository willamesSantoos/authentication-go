package repository

import (
	"github.com/willamesSantoos/authentication/internal/db"
	"github.com/willamesSantoos/authentication/internal/models"
)

func Insert(user models.User) (id int64, err error) {
	conn := db.GetInstace()

	err = conn.QueryRow("INSERT INTO users (email, password) VALUES ($1, $2)  RETURNING id", user.Email, user.Password).Scan(&id)

	return id, err
}
