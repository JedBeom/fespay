package main

import (
	"net/http"
	"time"

	"github.com/go-pg/pg"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

func getMyOrders(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return echo.ErrUnauthorized
	}

	o, err := models.OrdersRelatedWallet(db, u.WalletID)
	if err == pg.ErrNoRows {
		return models.NewFieldError("user doesn't have a wallet")
	} else if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSONPretty(http.StatusOK, o, JSONIndent)
}

func getMyBoothOrders(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return echo.ErrUnauthorized
	}

	if u.BoothID == "" {
		return echo.ErrBadRequest
	}

	o, err := models.OrdersRelatedWallet(db, u.Booth.WalletID)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSONPretty(http.StatusOK, o, JSONIndent)
}

func getOrderByID(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrUnknown.Send(c)
	}

	id := c.Param("id")
	o, err := models.OrderByID(db, id)
	if err == pg.ErrNoRows {
		return echo.ErrNotFound
	} else if err != nil {
		return echo.ErrInternalServerError
	}

	if u.BoothID == AdminBoothID || u.WalletID == o.FromID || u.WalletID == o.ToID ||
		u.Booth.WalletID == o.From.ID || u.Booth.WalletID == o.To.ID {
		return c.JSONPretty(http.StatusOK, o, JSONIndent)
	}

	return echo.ErrForbidden
}

// 아직은 유저-유저 결제 미구현
func postOrder(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrUnknown.Send(c)
	}

	if u.BoothID == "" {
		return ErrUserNotInBooth.Send(c)
	}

	p := struct {
		UserWalletID string `json:"userWalletID"`
		Amount       int    `json:"amount"`
	}{}

	if err := c.Bind(&p); err != nil {
		return echo.ErrBadRequest
	}

	o := models.Order{
		StaffID:     u.ID,
		Amount:      p.Amount,
		AccessLogID: c.Response().Header().Get(echo.HeaderXRequestID),
	}

	if !models.IsInvalidAmount(p.Amount) {
		return ErrInvalidAmount.Send(c)
	}

	if p.UserWalletID != "" {
		w, err := models.WalletByID(db, p.UserWalletID)
		if err != nil {
			return ErrUserWalletIDInvalid.Send(c)
		}

		_, err = w.User(db)
		if err != nil {
			return ErrBoothToBoothOrder
		}
		o.ClosedAt = time.Now()
	}

	if err := o.Create(db); err != nil {
		return ErrUnknown.Send(c) // TODO: 오더 프로세스에 따른 에러 처리!!!!!
	}

}

func deleteOrder(c echo.Context) error {
	if p.RefundOrderID != "" {
		ro, err := models.OrderByID(db, p.RefundOrderID)
		if err != nil {
			return ErrRefundOrderIDInvalid.Send(c)
		}

		if ro.RefundOrderID != "" {
			return ErrRefundRefunding.Send(c)
		}

		_, wasRefunded, err := ro.WasRefunded(db)
		if wasRefunded || err != nil {
			return ErrRefundedAgain.Send(c)
		}
	} else {

	}
}
