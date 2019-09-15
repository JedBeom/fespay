package main

import (
	"os"

	"github.com/JedBeom/fespay/models"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	_ "github.com/joho/godotenv/autoload"
)

var db *pg.DB

func connectDB() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	db = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
	})

	orm.RegisterTable(&models.OrderToProduct{})
}

func insertStudentsIfNotExists(db *pg.DB) {
	ok, err := db.Model(&models.Student{}).Where("class = 3").Exists()
	if err != nil {
		panic(err)
	}

	if ok {
		return
	}

	loc := os.Getenv("STUDENT_CSV")
	file, err := os.Open(loc)
	if err != nil {
		panic(err)
	}

	if err := models.CopyStudentsCSV(db, file); err != nil {
		panic(err)
	}
}
