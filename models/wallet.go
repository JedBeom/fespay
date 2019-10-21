package models

import (
	"errors"

	"github.com/go-pg/pg"
)

func (w *Wallet) Create(db *pg.DB) error {
	if w.OwnerID == "" {
		return FieldError{message: "OwnerID is blank"}
	}

	switch w.OwnerType {
	default:
		return FieldError{"invalid OwnerType"}
	case OwnerUser:
		u, err := UserByID(db, w.OwnerID)
		if err != nil || u.ID == "" {
			return err
		}

		if u.WalletID != "" {
			return errors.New("user has a wallet already")
		}

		if _, err := db.Model(&w).Insert(); err != nil {
			return err
		}

		u.WalletID = w.ID
		if _, err := db.Model(&u).Update(); err != nil {
			return err
		}
	case OwnerBooth:
		b, err := BoothByID(db, w.OwnerID)
		if err != nil || b.ID == "" {
			return err
		}

		if b.WalletID != "" {
			return errors.New("booth has a wallet already")
		}

		if _, err := db.Model(&w).Insert(); err != nil {
			return err
		}

		b.WalletID = w.ID
		if _, err := db.Model(&b).Update(); err != nil {
			return err
		}
	}

	return nil
}

func (w *Wallet) User(db *pg.DB) (u User, err error) {
	if w.OwnerType != OwnerUser {
		err = FieldError{"OwnerType is not OwnerUser"}
		return
	}

	err = db.Model(&u).Where("wallet_id = ?", w.ID).Select()
	u.Wallet = w
	return
}

func (w *Wallet) Booth(db *pg.DB) (b Booth, err error) {
	if w.OwnerType != OwnerBooth {
		err = FieldError{"OwnerType is not OwnerBooth"}
		return
	}

	err = db.Model(&b).Where("wallet_id = ?", w.ID).Select()
	b.Wallet = w
	return
}
