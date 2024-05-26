package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	c = map[string]string{ //credentials
		"HOST":     "",
		"PORT":     "",
		"USER":     "",
		"PASSWORD": "",
		"DBNAME":   "",
		"SSLMODE":  "",
	}
)

func ConnectDSN(dsn string) (*sql.DB, error) {
	connection, connectionErr := sql.Open("postgres", dsn)

	if connectionErr != nil {
		return nil, connectionErr
	}

	if err := connection.Ping(); err != nil {
		return nil, err
	}

	return connection, nil
}

func AuthDSN() (string, error) {
	for e := range c {
		val := viper.GetString(e)
		if val == "" {
			return "", errors.New("Error during reading config file")
		}
		c[e] = val
	}

	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", c["HOST"], c["PORT"], c["USER"], c["PASSWORD"], c["DBNAME"], c["SSLMODE"]), nil
}

/*
func Connect() error {
	for e := range c {
		val := viper.GetString(e)
		if val == "" {
			return errors.New("Error during reading config file")
		}
		c[e] = val
	}

	dns := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", c["HOST"], c["PORT"], c["USER"], c["PASSWORD"], c["DBNAME"], c["SSLMODE"])
	connection, err = sql.Open("postgres", dns)

	if err != nil {
		return err
	}
	return connection.Ping()
}
*/
