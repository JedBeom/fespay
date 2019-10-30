package main

import (
	"net/http"

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
		return models.NewFieldError("No rows")
	} else if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSONPretty(http.StatusOK, rs, JSONIndent)
}

func getRecordsByBoothID(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	boothID := c.Param("id")
	if u.BoothID != AdminBoothID || u.BoothID != boothID {
		return echo.ErrForbidden
	}

	rs, err := u.Booth.Records(db, true)
	if err == pg.ErrNoRows {
		return models.NewFieldError("No rows")
	} else if err != nil {
		return echo.ErrInternalServerError
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

// 아직은 유저-유저 결제 미구현
func postRecord(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	if u.BoothID == "" {
		return ErrUserNotInBooth.Send(c)
	}

	p := struct {
		Type     int    `json:"type"`
		CardCode string `json:"cardCode"`
		Amount   int    `json:"amount"`
	}{}

	if err := c.Bind(&p); err != nil {
		return echo.ErrBadRequest
	}

	if models.RecordType(p.Type) == models.RecordCharge && u.BoothID != AdminBoothID {
		return echo.ErrForbidden
	}

	r := models.Record{
		Type:        models.RecordType(p.Type),
		StaffID:     u.ID,
		BoothID:     u.BoothID,
		Amount:      p.Amount,
		AccessLogID: c.Response().Header().Get(echo.HeaderXRequestID),
	}

	if !models.IsInvalidAmount(p.Amount) {
		return ErrInvalidAmount.Send(c)
	}

	if p.CardCode != "" {
		u, err := models.UserByCardCode(db, p.CardCode)
		if err != nil {
			return ErrInvalidCardCode.Send(c)
		}

		r.UserID = u.ID

		if err := r.PayAndCreate(db); err != nil {
			return err2ApiErr(err).Send(c)
		}
	}

	if err := r.Create(db); err != nil {
		return ErrUnknown.Send(c)
	}

	return c.JSONPretty(http.StatusCreated, r, JSONIndent)
}

func deleteRecordByID(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	u = u

	return nil
}

func patchRecordByID(c echo.Context) error {
	return nil
}
