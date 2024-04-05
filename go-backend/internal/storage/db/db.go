package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
)

var DB *sql.DB = nil

func Connect() {
	pgUrl := os.Getenv("POSTGRES_URL")
	var err error
	DB, err = sql.Open("pgx", pgUrl)
	if err != nil {
		panic(err)
	}

}
