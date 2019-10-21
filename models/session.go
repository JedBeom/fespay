package models

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

func (u User) NewSession(db *pg.DB) (s Session, err error) {
	s.UserID = u.ID
	s.ID = uuid.New().String()
	err = db.Insert(&s)
	return
}

func SessionByID(db *pg.DB, id string) (sess Session, err error) {
	err = db.Model(&sess).Where("id = ?", id).Select()
	return
}

func SessionAndUserByID(db *pg.DB, id string) (sess Session, user User, err error) {
	sess, err = SessionByID(db, id)
	if err != nil || sess.ID == "" {
		return
	}

	user.ID = sess.UserID

	err = db.Model(&user).WherePK().Relation("Booth").Select()
	return
}

func (sess *Session) Delete(db *pg.DB) (err error) {
	_, err = db.Model(sess).WherePK().Delete()
	return
}
