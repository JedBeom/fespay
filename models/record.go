package models

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/teris-io/shortid"
)

func IsInvalidAmount(amount int) bool {
	if amount <= 0 || (amount%100) > 0 {
		return false
	}
	return true
}

func RecordByID(db *pg.DB, id string) (r Record, err error) {
	r.ID = id
	err = db.Model(&r).WherePK().Select()
	return
}

func (r *Record) Create(db *pg.DB) error {
	var err error
	if r.ID, err = shortid.Generate(); err != nil {
		return echo.ErrInternalServerError
	}

	return db.Insert(r)
}

// run in tx
func (r *Record) CreateTx(tx *pg.Tx) error {
	var err error
	if r.ID, err = shortid.Generate(); err != nil {
		return echo.ErrInternalServerError
	}

	return tx.Insert(r)
}

func (u *User) Records(db *pg.DB) ([]Record, error) {
	rs, err := RecordsByUserID(db, u.ID)
	return rs, err
}

func (b *Booth) Records(db *pg.DB, all bool) (rs []Record, err error) {
	q := db.Model(&rs).Where("booth_id = ?", b.ID)
	if !all {
		q.Where("canceled_at IS NULL")
	}
	err = q.Select()
	return
}

func RecordsByUserID(db *pg.DB, id string) (rs []Record, err error) {
	err = db.Model(&rs).Where("user_id = ?", id).Select()
	return
}

func RecordsByBoothID(db *pg.DB, id string) (rs []Record, err error) {
	err = db.Model(&rs).Where("booth_id = ?", id).Select()
	return
}

// Run in tx
func (r *Record) Pay(tx *pg.Tx) error {
	if r.Type != RecordOrder {
		return NewFieldError("it is not paying")
	}
	u, err := UserByIDForUpdate(tx, r.UserID)
	if err != nil {
		return err
	}

	if (u.Coin - r.Amount) < 0 {
		return NewFieldError("not enough coin")
	}

	b, err := BoothByIDForUpdate(tx, r.BoothID)
	if err != nil {
		return err
	}

	now := time.Now()
	u.Coin -= r.Amount
	u.UpdatedAt = now
	b.Coin += r.Amount
	b.UpdatedAt = now

	if err := tx.Update(u); err != nil {
		return err
	}

	if err := tx.Update(b); err != nil {
		return err
	}
	return nil
}

func (r *Record) PayAndCreate(db *pg.DB) error {
	return db.RunInTransaction(func(tx *pg.Tx) error {
		if err := r.Pay(tx); err != nil {
			return err
		}

		return r.CreateTx(tx)
	})
}

func (r *Record) Charge(tx *pg.Tx) error {
	u, err := UserByIDForUpdate(tx, r.UserID)
	if err != nil {
		return err
	}

	u.Coin += r.Amount
	u.UpdatedAt = time.Now()
	return tx.Update(&u)
}

func (r *Record) ChargeAndCreate(db *pg.DB) error {
	return db.RunInTransaction(func(tx *pg.Tx) error {
		if err := r.Charge(tx); err != nil {
			return err
		}

		return r.CreateTx(tx)
	})
}

func (r *Record) CancelOrder(db *pg.DB) error {
	now := time.Now()
	r.CanceledAt = now
	if r.UserID == "" { // 돈 빼고 할 것도 없음
		return db.Update(r)
	}

	return db.RunInTransaction(func(tx *pg.Tx) error {
		u, err := UserByIDForUpdate(tx, r.UserID)
		if err != nil {
			return err
		}

		b, err := BoothByIDForUpdate(tx, r.BoothID)
		if err != nil {
			return err
		}

		u.Coin += r.Amount
		b.Coin -= r.Amount
		if b.Coin < 0 {
			return NewFieldError("??? 부스에 돈이 없음")
		}

		u.UpdatedAt = now
		b.UpdatedAt = now
		if err := tx.Update(&u); err != nil {
			return err
		}
		if err := tx.Update(&b); err != nil {
			return err
		}
		return tx.Update(r)
	})
}

func (r *Record) CancelCharge(db *pg.DB) error {
	now := time.Now()
	r.CanceledAt = now
	return db.RunInTransaction(func(tx *pg.Tx) error {
		u, err := UserByIDForUpdate(tx, r.UserID)
		if err != nil {
			return err
		}

		u.Coin -= r.Amount
		if u.Coin < 0 {
			return NewFieldError("user not enough coin")
		}
		u.UpdatedAt = now
		if err := tx.Update(&u); err != nil {
			return err
		}
		return tx.Update(r)
	})
}
