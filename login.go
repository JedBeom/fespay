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
		return echo.ErrBadRequest
	}

	seller, err := models.UserByLoginIDWithBooth(db, payload.ID)
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

	return ErrLoginFailed

}

func getLogout(c echo.Context) error {
	sessID, ok := c.Get("sess_id").(string)
	if !ok {
		return echo.ErrInternalServerError
	}

	sess := models.Session{
		ID: sessID,
	}

	if err := sess.Delete(db); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSONPretty(http.StatusOK, KeyValue{
		"message": "logout success",
	}, JSONIndent)

}
