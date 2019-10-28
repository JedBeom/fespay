package main

import (
	"net/http"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

func getMine(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return echo.ErrInternalServerError
	}

	return c.JSONPretty(http.StatusOK, u, JSONIndent)

}
