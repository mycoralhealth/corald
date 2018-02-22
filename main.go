package main

import (
	"log"
	"os"

	"github.com/mycoralhealth/mycoral-patient-server/web"

	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	dbPath := os.Getenv("MYCORAL_DB")
	log.Printf("Opening database %s", dbPath)

	dbCon, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dbCon.Close()

	log.Fatal(web.Run(dbCon))
}
