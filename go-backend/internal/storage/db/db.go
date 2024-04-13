package db

import (
	"database/sql"
	"errors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
)

var database *sql.DB = nil

func Connect() {
	pgUrl := os.Getenv("POSTGRES_URL")
	var err error
	database, err = sql.Open("pgx", pgUrl)
	if err != nil {
		panic(err)
	}

}

func Close() {
	err := database.Close()
	if err != nil {
		panic(err)
	}
}

func GetInstance() (*sql.DB, error) {
	if database == nil {
		return nil, errors.New("error: Database can't be 'nil' try to 'Initialize'")
	}
	return database, nil
}
