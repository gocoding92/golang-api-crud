package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-echo-api/config"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	conn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME)
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic("Connecting Error")
	}

	err = db.Ping()

	if err != nil {
		panic("DSN Invalid")
	}

}

func CreateCon() *sql.DB {
	return db
}
