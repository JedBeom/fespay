package models

import "github.com/go-pg/pg"

func BoothByID(db *pg.DB, id string) (b Booth, err error) {
	b.ID = id
	err = db.Model(&b).WherePK().Select()
	return
}

func BoothByIDForUpdate(tx *pg.Tx, id string) (b Booth, err error) {
	err = tx.Model(&b).For("UPDATE").Where("id = ?", id).Select()
	return
}
