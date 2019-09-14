package models

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

func (s Seller) NewSession(db *pg.DB) (sess Session, err error) {
	sess.SellerID = s.ID
	sess.ID = uuid.New().String()
	err = db.Insert(&sess)
	return
}

func SessionByUUID(db *pg.DB, uuid string) (sess Session, err error) {
	err = db.Model(&sess).Where("id = ?", uuid).Select()
	return
}

func SessionAndSellerByUUID(db *pg.DB, uuid string) (sess Session, seller Seller, err error) {
	sess, err = SessionByUUID(db, uuid)
	if err != nil || sess.ID == "" {
		return
	}

	seller.ID = sess.SellerID

	err = db.Model(&seller).WherePK().Relation("Booth").Select()
	return
}

func (sess *Session) Delete(db *pg.DB) (err error) {
	_, err = db.Model(sess).WherePK().Delete()
	return
}
