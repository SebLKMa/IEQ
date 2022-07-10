package db

import (
	"database/sql"
	"fmt"
)

// TODO: hardcoded for now, to refactor to config
const (
	host     = "localhost"
	port     = 5432
	user     = "iequser"
	password = "iequser"
	dbname   = "ieqdb"
)

// connect opens and return the db object. Caller must close db when done with it.
func connect() (db *sql.DB) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}
