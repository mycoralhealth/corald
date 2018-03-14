package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mycoralhealth/corald/model"
	"github.com/mycoralhealth/corald/web"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPath := os.Getenv("CORALD_DB")
	log.Printf("Opening database %s", dbPath)

	dbCon, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dbCon.Close()

	dbCon.AutoMigrate(&model.User{})
	dbCon.Debug()

	log.Fatal(web.Run(dbCon))
}
