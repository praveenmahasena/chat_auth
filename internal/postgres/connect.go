package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	connection *sql.DB
	err        error
)

func Connect() error {
	dns := fmt.Sprint("host=localhost port=5432 user=postgres password=Sahan2015 dbname=chat_app sslmode=disable")
	connection, err = sql.Open("postgres", dns)

	if err != nil {
		return err
	}
	return connection.Ping()
}
