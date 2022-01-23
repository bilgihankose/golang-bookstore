package main

import (
	"database/sql"
	"fmt"
	"github.com/bilgihankose/golang-bookstore/handlers"
	. "github.com/bilgihankose/golang-bookstore/utils"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

var db *sql.DB

func init() {
	CheckError(gotenv.Load())

}

func main() {

	pgURL, err := pq.ParseURL(os.Getenv("POSTGRESQL_URL"))
	CheckError(err)
	fmt.Println("URL is", pgURL)

	db, err = sql.Open("postgres", pgURL)
	err = db.Ping() // Need to do this to check that the connection is valid
	if err != nil {
		log.Fatal(err)
	}

	handlers.Run()

}
