package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitSqlxDatabase() (db *sqlx.DB) {
	data_source_name := "user=golang-lesson password=golang-lesson dbname=golang-lesson sslmode=disable"
	postgres_db, err := sqlx.Connect("postgres", data_source_name)
	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database")
	}

	postgres_db.SetMaxOpenConns(20)
	postgres_db.SetMaxIdleConns(10)

	return postgres_db
}
