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

func BoothsSearchName(db *pg.DB, like string, limit, page int) (bs []Booth, err error) {
	page -= 1
	like = "%" + like + "%"
	err = db.Model(&bs).
		Where("name LIKE ?", like).
		Relation("Staffs").
		Limit(limit).Offset(page * limit).Order("updated_at DESC").Select()
	return
}

func Booths(db *pg.DB, column string, limit, page int) (bs []Booth, err error) {
	page -= 1
	err = db.Model(&bs).Order(column + " DESC").Relation("Staffs").
		Limit(limit).Offset(limit * page).Select()
	return
}
