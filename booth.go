package main

import (
	"net/http"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

func getBoothByID(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface
	}

	boothID := c.Param("id")

	if u.BoothID != AdminBoothID && boothID != u.BoothID {
		return echo.ErrForbidden
	}

	b, err := models.BoothByID(db, boothID)
	if err != nil {
		return err2ApiErr(err).Send(c)
	}

	return c.JSONPretty(http.StatusOK, b, JSONIndent)
}
