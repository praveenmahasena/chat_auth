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
	connection *sql.DB
	err        error
)

func Connect() error {
	for e := range c {
		val := viper.GetString(e)
		if val == "" {
			return errors.New("Error during reading config file")
		}
		c[e] = val
	}

	dns := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", c["HOST"], c["PORT"], c["USER"], c["PASSWORD"], c["DBNAME"], c["SSLMODE"])
	fmt.Println(dns)
	connection, err = sql.Open("postgres", dns)

	if err != nil {
		return err
	}
	return connection.Ping()
}
