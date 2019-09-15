package models

import (
	"github.com/go-pg/pg"
)

func BoothBySeller(db *pg.DB, seller *Seller) (booth Booth, err error) {
	booth.ID = seller.BoothID
	err = db.Model(&booth).WherePK().Relation("Products").Relation("Sellers").Select()
	return
}
