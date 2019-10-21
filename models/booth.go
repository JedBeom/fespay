package models

import "github.com/go-pg/pg"

func BoothByID(db *pg.DB, id string) (b Booth, err error) {
	err = db.Model(&b).Where("id = ?", id).Select()
	return
}
