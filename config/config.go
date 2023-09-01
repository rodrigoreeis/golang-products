package config

import (
	"database/sql"
	"log"
)

func ConnectionDB() *sql.DB {
	connStr := "user=process.env.user dbname=process.env.db password=process.env.dbpassword host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	return db
}
