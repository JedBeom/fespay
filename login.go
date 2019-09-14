package main

import (
	"net/http"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

func postLogin(c echo.Context) error {
	payload := struct {
		ID  string `json:"id" query:"id"`
		Pin string `json:"pin" query:"pin"`
	}{}
	if err := c.Bind(&payload); err != nil {
		return echo.ErrInternalServerError
	}

	seller, err := models.SellerByLoginID(db, payload.ID)
	if err != nil {
		return echo.ErrInternalServerError
	}

	if seller.Pin == payload.Pin {
		sess, err := seller.NewSession(db)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSONPretty(200, KeyValue{
			"uuid": sess.ID,
		}, JSONIndent)
	}

	return ErrorLoginFailed

}

func getLogout(c echo.Context) error {
	sess, ok := c.Get("sess").(models.Session)
	if !ok {
		return echo.ErrInternalServerError
	}

	if err := sess.Delete(db); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSONPretty(http.StatusOK, KeyValue{
		"message": "logout success",
	}, JSONIndent)

}
