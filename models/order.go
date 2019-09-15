package models

import (
	"net/http"
	"time"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/lithammer/shortuuid"
)

func OrdersByBoothID(db *pg.DB, boothID int, page int, limit int) (orders []Order, err error) {
	if limit <= 0 || limit > 100 {
		limit = 100
	}

	if page <= 0 {
		page = 1
	}
	page -= 1

	err = db.Model(&orders).
		Where("booth_id = ?", boothID).Order("date DESC").
		Offset(page * limit).Limit(limit).Select()
	return
}

func OrderByID(db *pg.DB, id string) (order Order, err error) {
	order.ID = id
	err = db.Model(&order).WherePK().Select()
	return
}

func (order *Order) Create(db *pg.DB) error {
	return db.RunInTransaction(func(tx *pg.Tx) error {
		order.ID = shortuuid.New()
		if err := challengePay(tx, order.BoothID, order.StudentID, order.GrandTotal, false); err != nil {
			return err
		}

		if err := tx.Insert(order); err != nil {
			return echo.ErrInternalServerError
		}

		if err := order.createOToPs(tx); err != nil {
			return echo.ErrInternalServerError
		}

		return nil
	})
}

func (order *Order) Cancel(db *pg.DB) error {
	return db.RunInTransaction(func(tx *pg.Tx) error {
		order.IsCanceled = true
		if err := challengePay(tx, order.BoothID, order.StudentID, order.GrandTotal, true); err != nil {
			return err
		}
		if err := tx.Update(order); err != nil {
			return echo.ErrInternalServerError
		}

		return nil
	})
}

func challengePay(tx *pg.Tx, boothID, studentID, price int, isRefund bool) error {
	booth, student, err := getBoothAndStudentByIDForUpdate(tx, boothID, studentID)
	if err != nil {
		return echo.ErrInternalServerError
	}

	if isRefund {
		if booth.Coin -= price; booth.Coin <= 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Not Enough Coin")
		}

		student.Coin += price
	} else {
		if student.Coin -= price; student.Coin <= 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Not Enough Coin")
		}

		booth.Coin += price
	}

	t := time.Now()
	student.UpdatedAt = t
	booth.UpdatedAt = t

	if err := tx.Update(&booth); err != nil {
		return echo.ErrInternalServerError
	}

	if err := tx.Update(&student); err != nil {
		return echo.ErrInternalServerError
	}

	return nil
}

func getBoothAndStudentByIDForUpdate(tx *pg.Tx, boothID, studentID int) (booth Booth, student Student, err error) {

	booth.ID = boothID
	student.ID = studentID

	if err = tx.Model(&booth).WherePK().For("UPDATE").Select(); err != nil {
		err = echo.ErrInternalServerError
		return
	}

	if err = tx.Model(&student).WherePK().For("UPDATE").Select(); err != nil {
		err = echo.ErrInternalServerError
		return
	}

	return
}
