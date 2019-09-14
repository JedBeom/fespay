package models

import (
	"github.com/go-pg/pg"
)

func (s *Seller) FillAll(db *pg.DB) (err error) {
	err = db.Model(s).WherePK().
		Relation("Student").Relation("Booth.Products").
		Select()
	return
}
