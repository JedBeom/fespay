package main

import (
	"net/http"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

func postProducts(c echo.Context) error {
	seller, ok := c.Get("seller").(models.Seller)
	if !ok {
		return echo.ErrInternalServerError
	}

	var bindProducts struct {
		Products []models.Product `json:"products"`
	}

	if err := c.Bind(&bindProducts); err != nil {
		return echo.ErrInternalServerError
	}

	if err := models.InsertProducts(db, seller.BoothID, &bindProducts.Products); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSONPretty(http.StatusOK, bindProducts.Products, JSONIndent)
}
