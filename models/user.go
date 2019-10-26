package models

import (
	"io"

	"github.com/go-pg/pg"
)

func UserByID(db *pg.DB, id string) (u User, err error) {
	err = db.Model(&u).Where("id = ?", id).Select()
	return
}

func UserByCardCode(db *pg.DB, code string) (u User, err error) {
	err = db.Model(&u).Where("card_code = ?", code).Relation("Wallet").Select()
	return
}

func CheckCardAvailable(db *pg.DB, cardCode string) (bool, error) {
	u := User{}
	err := db.Model(&u).Where("card_code = ?", cardCode).Select()
	if u.ID == "" || err == pg.ErrNoRows || u.LoginID == "" {
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

func CopyUsersCSV(db *pg.DB, file io.Reader) (err error) {
	_, err = db.CopyFrom(file,
		`COPY users(id, grade, class, number, name, card_code, type) FROM STDIN DELIMITER ',' CSV HEADER`)
	return
}
