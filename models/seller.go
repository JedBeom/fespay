package models

import (
	"github.com/go-pg/pg"
)

func SellerByLoginID(db *pg.DB, id string) (s Seller, err error) {
	err = db.Model(&s).Where("login_id = ?", id).Relation("Booth").Select()
	return
}
