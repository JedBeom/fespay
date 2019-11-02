package main

import (
	"net/http"
	"time"

	"github.com/go-pg/pg"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

func getRecordsByUserID(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	userID := c.Param("id")
	if u.BoothID != AdminBoothID || userID != u.ID {
		return echo.ErrForbidden
	}

	rs, err := models.RecordsByUserID(db, userID)
	if err == pg.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound, "no records")
	} else if err != nil {
		return err2ApiErr(err).Send(c)
	}

	return c.JSONPretty(http.StatusOK, rs, JSONIndent)
}

func getRecordsByBoothID(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	boothID := c.Param("id")
	if u.BoothID != AdminBoothID && u.BoothID != boothID {
		return echo.ErrForbidden
	}

	rs, err := models.RecordsByBoothID(db, boothID)
	if err == pg.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound, "no records")
	} else if err != nil {
		return err2ApiErr(err).Send(c)
	}

	return c.JSONPretty(http.StatusOK, rs, JSONIndent)
}

func getRecordByID(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	id := c.Param("id")
	r, err := models.RecordByID(db, id)
	if err == pg.ErrNoRows {
		return echo.ErrNotFound
	} else if err != nil {
		return echo.ErrInternalServerError
	}

	// 어드민이거나 결제가 유저나 부스에서 이루어졌는지 확인
	if u.BoothID == AdminBoothID || r.UserID == u.ID || u.BoothID == r.BoothID {
		return c.JSONPretty(http.StatusOK, r, JSONIndent)
	}

	return echo.ErrForbidden
}

func postRecord(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	// user can not create a record himself
	if u.BoothID == "" {
		return ErrUserNotInBooth.Send(c)
	}

	// frozen booth
	if u.Booth.Status == models.StatusFrozen {
		return ErrFrozenBooth.Send(c)
	}

	p := struct {
		Type     models.RecordType `json:"type"`
		CardCode string            `json:"cardCode"`
		Amount   int               `json:"amount"`
	}{}

	// parse params
	if err := c.Bind(&p); err != nil {
		return ErrField.Send(c)
	}

	// Check amount is valid
	if !models.IsInvalidAmount(p.Amount) {
		return ErrInvalidAmount.Send(c)
	}

	r := models.Record{
		StaffID:     u.ID,
		BoothID:     u.BoothID,
		Amount:      p.Amount,
		AccessLogID: c.Response().Header().Get(echo.HeaderXRequestID),
	}

	switch p.Type {
	case models.RecordCharge: // charging
		r.Type = models.RecordCharge
		if u.BoothID != AdminBoothID { // only admin can do it
			return echo.ErrForbidden
		}

		if p.CardCode == "" { // cardCode should be told
			return ErrField.Send(c)
		}

		_, err := models.UserByCardCode(db, p.CardCode) // get targetUser
		if err != nil {
			return ErrInvalidCardCode.Send(c)
		}

		if err := r.ChargeAndCreate(db); err != nil {
			return err2ApiErr(err).Send(c)
		}
	case models.RecordOrder: // order
		r.Type = models.RecordOrder
		if p.CardCode == "" { // 선결제 생성인 경우
			tu, err := models.UserByCardCode(db, p.CardCode) // get targetUser
			if err != nil || tu.Status == models.StatusSuspended {
				return ErrInvalidCardCode.Send(c)
			}

			if tu.Status == models.StatusFrozen { // 결제 불능 상태인 경우
				return ErrFrozenUser.Send(c)
			}

			r.UserID = tu.ID
			r.PaidAt = time.Now()
			if err := r.PayAndCreate(db); err != nil {
				return err2ApiErr(err).Send(c)
			}
		} else { // 선 생성 후 결제인 경우
			if err := r.Create(db); err != nil {
				return err2ApiErr(err).Send(c)
			}
		}
	default: // 뭐야 그거...
		return ErrUnknownRecordType.Send(c)
	}

	return c.JSONPretty(http.StatusCreated, Map{
		"id": r.ID,
	}, JSONIndent)
}

func deleteRecordByID(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	if u.BoothID == "" {
		return ErrUserNotInBooth.Send(c)
	}

	if u.Booth.Status == models.StatusFrozen {
		return ErrFrozenUser.Send(c)
	}

	recordID := c.Param("id")

	r, err := models.RecordByID(db, recordID)
	if err != nil {
		return err2ApiErr(err).Send(c)
	}

	if r.BoothID != u.BoothID && u.BoothID != AdminBoothID { // 내 부스가 아니면서 어드민도 아니면
		return echo.ErrForbidden
	}

	if !r.CanceledAt.IsZero() { // 이미 취소된 것이다
		return ErrCancelingAgain.Send(c)
	}

	return c.JSONPretty(http.StatusOK, Map{
		"message": "canceled successfully",
	}, JSONIndent)
}

func patchRecordByID(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	if u.Status == models.StatusFrozen {
		return ErrFrozenUser.Send(c)
	}

	recordID := c.Param("id")
	r, err := models.RecordByID(db, recordID)
	if err != nil {
		return err2ApiErr(err).Send(c)
	}

	if !r.PaidAt.IsZero() {
		return ErrExpiredRecord.Send(c)
	}

	r.UserID = u.ID
	err = db.RunInTransaction(func(tx *pg.Tx) error {
		if err := r.Pay(tx); err != nil {
			return err
		}

		return tx.Update(&r)
	})
	if err != nil {
		return err2ApiErr(err).Send(c)
	}

	return c.JSONPretty(http.StatusOK, Map{
		"message": "paid successfully",
	}, JSONIndent)
}
