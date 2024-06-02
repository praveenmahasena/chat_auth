package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	connection *sql.DB
	err        error
)

const hostParam = "HOST"
const portParam = "PORT"
const userParam = "USER"
const passwordParam = "PASSWORD"
const dbNameParam = "DBNAME"
const sslModeParam = "SSLMODE"

func Connect() error {
	host := viper.GetString(hostParam)
	port := viper.GetString(portParam)
	user := viper.GetString(userParam)
	password := viper.GetString(passwordParam)
	dbName := viper.GetString(dbNameParam)
	sslmode := viper.GetString(sslModeParam)

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", host, port, user, password, dbName, sslmode)
	connection, err = sql.Open("postgres", dsn)

	if err != nil {
		return err
	}
	return connection.Ping()
}
