package models

import (
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/teris-io/shortid"
)

func IsInvalidAmount(amount int) bool {
	if amount == 0 {
		return false
	}

	if (amount % 100) > 0 {
		return false
	}

	return true
}

func OrdersRelatedWallet(db *pg.DB, id string) (o []Order, err error) {
	err = db.Model(&o).Where("from_id = ?", id).WhereOr("to_id = ?", id).Select()
	return
}

func OrderByID(db *pg.DB, id string) (o Order, err error) {
	o.ID = id
	err = db.Model(&o).WherePK().Relation("From").Select()
	if o.ToID != "" {
		w, err := WalletByID(db, o.ToID)
		if err != nil {
			return o, err
		}

		o.To = &w
	}
	return
}

func (o *Order) Create(db *pg.DB) error {
	var err error
	if o.ID, err = shortid.Generate(); err != nil {
		return echo.ErrInternalServerError
	}

	return db.Insert(o)
}

func (o *Order) WasRefunded(db *pg.DB) (refundingOrder Order, wasRefunded bool, err error) {
	err = db.Model(&refundingOrder).Where("refund_order_id = ?", o.ID).Select()
	if err == pg.ErrNoRows {
		err = nil
		wasRefunded = false
		return
	} else if err != nil {
		return
	}

	wasRefunded = true
	return
}
