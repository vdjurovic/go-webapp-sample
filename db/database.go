package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", "dbname=goweb user=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetDBConnection() *sql.DB {
	return db
}
