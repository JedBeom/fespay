package models

import "github.com/go-pg/pg"

func BoothByID(db *pg.DB, id string) (b Booth, err error) {
	b.ID = id
	err = db.Model(&b).WherePK().Select()
	return
}
