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
	res.IsAvailable, err = models.CheckCardAvailable(db, res.CardCode)
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

		Grade int    `json:"grade"`
		Name  string `json:"name"`
	}{}

	if err := c.Bind(&p); err != nil ||
		p.LoginID == "" || p.Password == "" || p.CardCode == "" || p.Name == "" {
		return echo.ErrBadRequest
	}

	isAvailable, err := models.CheckCardAvailable(db, p.CardCode)
	if err != nil || !isAvailable {
		return echo.ErrBadRequest
	}

}
