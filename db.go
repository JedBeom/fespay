package main

import (
	"github.com/go-pg/pg"
)

var db *pg.DB

func connectDB() {
	db = pg.Connect(&pg.Options{
		User:     "fespay",
		Password: "fespay",
		Database: "fespay",
	})
}
