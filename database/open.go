package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	PSQLHost     = "localhost"
	PSQLPort     = 5431
	PSQLUser     = "cfabrica46"
	PSQLPassword = "01234"
	PSQLDBName   = "engine_test"
)

var PsqlInfoDefault = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", PSQLHost, PSQLPort, PSQLUser, PSQLPassword, PSQLDBName)
var PostgresDriver = "postgres"

func init() {
	err := Open(PostgresDriver, PsqlInfoDefault, "./database.sql")
	if err != nil {
		log.Fatal(err)
	}
}

func Open(driver, psqlInfo, migrationPath string) (err error) {

	if db != nil {
		Close()
	}

	if migrationPath == "" {
		db, err = sql.Open(driver, psqlInfo)
		if err != nil {
			return
		}
		err = db.Ping()
		if err != nil {
			db.Close()
			return
		}
	} else {
		var dataSQL []byte
		dataSQL, err = ioutil.ReadFile(migrationPath)
		if err != nil {
			return
		}

		db, err = sql.Open(driver, psqlInfo)
		if err != nil {
			return
		}

		err = db.Ping()
		if err != nil {
			db.Close()
			return
		}

		_, err = db.Exec(string(dataSQL))
		if err != nil {
			db.Close()
			return
		}
	}
	return
}

func Close() {
	db.Close()
}
