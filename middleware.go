package main

import (
	"net/http"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

const KeyHeaderKey = "Rabbit-Fur"

func MiddlewareKeyCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.Request().Header.Get(KeyHeaderKey)
		if key == "" {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		sess, seller, err := models.SessionAndSellerByUUID(db, key)
		if err != nil || seller.ID == 0 {
			return echo.NewHTTPError(http.StatusUnauthorized, "bad key")
		}

		c.Set("sess", sess)
		c.Set("seller", seller)

		return next(c)
	}
}
