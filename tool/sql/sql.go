package sql

import (
	"fmt"
	"log"
	"os"
)

// ConnString is to handle parsing connection string to database
func ConnString() (uri string) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbssl := os.Getenv("DB_SSL")
	dbconn := os.Getenv("DB_CONN")

	if os.Getenv("DB_TYPE") == "MYSQL" {
		dbinstance := fmt.Sprintf("%s:%s", host, port)
		uri = fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbconn, dbinstance, dbname)
	} else if os.Getenv("DB_TYPE") == "POSTGRESQL" {
		if dbconn != "cloudsql" {
			uri = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, dbssl)
		} else {
			uri = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", host, user, password, dbname, dbssl)
		}
	} else {
		log.Fatalf(fmt.Sprintf("Cant handle connection type %s", os.Getenv("DB_TYPE")))
	}

	return
}
