package models

import "github.com/go-pg/pg"

func BoothByID(db *pg.DB, id string, fillStaffs bool) (b Booth, err error) {
	b.ID = id
	q := db.Model(&b).WherePK()
	if fillStaffs {
		q.Relation("Staffs")
	}
	err = q.Select()
	return
}

func BoothByIDForUpdate(tx *pg.Tx, id string) (b Booth, err error) {
	err = tx.Model(&b).For("UPDATE").Where("id = ?", id).Select()
	return
}
