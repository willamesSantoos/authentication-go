package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/willamesSantoos/authentication/internal/configs"
)

var connection *sql.DB

func openConnection() *sql.DB {
	config := configs.GetDBConfig()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	conn, err := sql.Open("postgres", stringConnection)

	if err != nil {
		log.Fatal("error when trying to open connection to the database")
		panic(err)
	}

	return conn
}

func GetInstace() *sql.DB {
	if connection == nil {
		connection = openConnection()
	}
	
	return connection
}
