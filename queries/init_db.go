package queries

import (
	"musematch/globals"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB() error {
	dbc, err := sqlx.Connect("sqlite3", globals.Env.DB_URL)
	if err != nil {
		return err
	}
	db = dbc
	return nil
}
