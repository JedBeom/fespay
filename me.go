package main

import (
	"net/http"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

func getSellerMe(c echo.Context) error {
	seller, ok := c.Get("seller").(models.Seller)
	if !ok {
		return echo.ErrInternalServerError
	}

	return c.JSONPretty(http.StatusOK, KeyValue{
		"login_id": seller.LoginID,
		"booth": KeyValue{
			"name": seller.Booth.Name,
			"coin": seller.Booth.Coin,
		},
	}, JSONIndent)
}

func getBoothMe(c echo.Context) error {
	s, ok := c.Get("seller").(models.Seller)
	if !ok {
		return echo.ErrInternalServerError
	}

	booth, err := models.BoothBySeller(db, &s)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSONPretty(http.StatusOK, booth, JSONIndent)
}

func getMine(c echo.Context) error {
	s, ok := c.Get("seller").(models.Seller)
	if !ok {
		return echo.ErrInternalServerError
	}

	if err := s.FillAll(db); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSONPretty(http.StatusOK, s, JSONIndent)

}
