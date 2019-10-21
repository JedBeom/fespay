package models

import "github.com/go-pg/pg"

func UserByID(db *pg.DB, id string) (u User, err error) {
	err = db.Model(&u).Where("id = ?", id).Select()
	return
}

func UserByCardCode(db *pg.DB, code string) (u User, err error) {
	err = db.Model(&u).Where("card_code = ?", code).Relation("Wallet").Select()
	return
}
