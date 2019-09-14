package models

import (
	"github.com/go-pg/pg"
	"github.com/lithammer/shortuuid"
)

func InsertProducts(db *pg.DB, boothID int, products *[]Product) error {
	for i := range *(products) {
		(*products)[i].ID = shortuuid.New()
		(*products)[i].BoothID = boothID
	}

	return db.Insert(products)
}
