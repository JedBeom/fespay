package models

import (
	"io"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-pg/pg"
)

func userByColumn(db *pg.DB, column, value string, fillAll bool) (u User, err error) {
	q := db.Model(&u).Where(column+" = ?", value)
	if err = q.Select(); err != nil {
		return
	}

	if !fillAll {
		return
	}

	if u.BoothID != "" {
		if b, err := BoothByID(db, u.BoothID, false); err == nil {
			u.Booth = &b
		} else {
			return u, err
		}
	}

	return
}

func UserByID(db *pg.DB, id string, fillAll bool) (u User, err error) {
	u, err = userByColumn(db, "id", id, fillAll)
	return
}

func UserByCardCode(db *pg.DB, code string) (u User, err error) {
	u, err = userByColumn(db, "card_code", code, true)
	return
}

func UserByLoginID(db *pg.DB, loginID string) (u User, err error) {
	u, err = userByColumn(db, "login_id", loginID, false)
	return
}

func CanCardRegistered(db *pg.DB, cardCode string) (bool, error) {
	u := User{}
	err := db.Model(&u).Where("card_code = ?", cardCode).Select()
	if u.ID == "" || err == pg.ErrNoRows || u.LoginID != "" {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *User) UpdateByCardCode(db *pg.DB, code string) error {
	_, err := db.Model(u).Where("card_code = ?", code).Update()
	return err
}

func Encrypt(pw string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pw), 16)
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}

func (u *User) Register(db *pg.DB) error {
	var err error
	u.Password, err = Encrypt(u.Password)
	if err != nil {
		return err
	}
	return db.Update(u)
}

func UserByIDForUpdate(tx *pg.Tx, id string) (u User, err error) {
	err = tx.Model(&u).Where("id = ?", id).For("UPDATE").Select()
	return
}

func CopyUsersCSV(db *pg.DB, file io.Reader) (err error) {
	_, err = db.CopyFrom(file,
		`COPY users(id, grade, class, number, name, card_code, type) FROM STDIN DELIMITER ',' CSV HEADER`)
	return
}
