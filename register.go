package main

import (
	"net/http"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

func getAvailable(c echo.Context) error {
	res := struct {
		CardCode    string `json:"cardCode"`
		IsAvailable bool   `json:"isAvailable"`
	}{}

	res.CardCode = c.QueryParam("code")

	var err error
	res.IsAvailable, err = models.CanCardRegistered(db, res.CardCode)
	if err != nil {
		return echo.ErrInternalServerError
	}

	status := http.StatusOK
	if !res.IsAvailable {
		status = http.StatusNotFound
	}

	return c.JSONPretty(status, res, JSONIndent)
}

func patchRegister(c echo.Context) error {
	p := struct {
		LoginID  string `json:"loginID"`
		Password string `json:"password"`
		CardCode string `json:"cardCode"`

		Number int    `json:"number"`
		Name   string `json:"name"`
	}{}

	if err := c.Bind(&p); err != nil ||
		p.LoginID == "" || p.Password == "" || p.CardCode == "" || p.Name == "" {
		return echo.ErrBadRequest
	}

	isAvailable, err := models.CanCardRegistered(db, p.CardCode)
	if err != nil || !isAvailable {
		return echo.ErrBadRequest
	}

	u, err := models.UserByCardCode(db, p.CardCode)
	if err != nil {
		return err2ApiErr(err).Send(c)
	}

	if u.Number != p.Number || u.Name != p.Name {
		return echo.ErrNotFound
	}

	u.LoginID = p.LoginID
	u.Password = p.Password
	if err := u.Register(db); err != nil {
		return err2ApiErr(err).Send(c)
	}

	return c.JSONPretty(http.StatusOK, Map{
		"message": "register successful",
	}, JSONIndent)
}
