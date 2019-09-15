package main

import (
	"net/http"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

func getStudentByBarcodeID(c echo.Context) error {
	barcodeID := c.Param("barcode")
	if barcodeID == "" {
		return echo.ErrBadRequest
	}

	student, err := models.StudentByBarcodeID(db, barcodeID)
	if err != nil {
		return ErrStudentNotFound
	}

	return c.JSONPretty(http.StatusOK, student, JSONIndent)
}
