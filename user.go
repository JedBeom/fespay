package main

import (
	"net/http"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

func getUserByID(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	if u.BoothID != AdminBoothID {
		return echo.ErrForbidden
	}

	userID := c.Param("id")
	tu, err := models.UserByID(db, userID, true)
	if err != nil {
		return err2ApiErr(err).Send(c)
	}

	return c.JSONPretty(http.StatusOK, tu, JSONIndent)
}

func getUserByCardCode(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	if u.BoothID != AdminBoothID {
		return echo.ErrForbidden
	}

	code := c.Param("code")
	tu, err := models.UserByCardCode(db, code)
	if err != nil {
		return err2ApiErr(err).Send(c)
	}

	return c.JSONPretty(http.StatusOK, tu, JSONIndent)
}
