package models

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

func (u User) NewSession(db *pg.DB, ua string) (s Session, err error) {
	s.UserID = u.ID
	s.ID = uuid.New().String()
	s.UserAgent = ua
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

	user, err = UserByID(db, id, true)
	return
}

func (sess *Session) Delete(db *pg.DB) (err error) {
	_, err = db.Model(sess).WherePK().Delete()
	return
}
