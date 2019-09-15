package main

import (
	"net/http"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

func postOrder(c echo.Context) error {
	seller, ok := c.Get("seller").(models.Seller)
	if !ok {
		return echo.ErrInternalServerError
	}

	var order models.Order
	order.SellerID = seller.ID
	order.BoothID = seller.BoothID

	payload := struct {
		StudentBarcodeID string `json:"student_barcode_id"`

		ProductIDs []string `json:"product_ids"`

		SubTotal   int `json:"sub_total"`
		Discount   int `json:"discount"`
		GrandTotal int `json:"grand_total"`
	}{}

	if err := c.Bind(&payload); err != nil {
		return echo.ErrBadRequest
	}

	student, err := models.StudentByBarcodeID(db, payload.StudentBarcodeID)
	if err != nil {
		return ErrStudentNotFound
	}
	order.StudentID = student.ID

	subTotal, products, err := models.SumAndValidateProducts(db, seller.BoothID, payload.ProductIDs)
	if err != nil || products == nil {
		return err
	}
	order.Products = products

	if subTotal != payload.SubTotal {
		return ErrWrongCalculate
	}
	order.SubTotal = subTotal

	order.GrandTotal = subTotal - payload.Discount
	order.Discount = payload.Discount
	if payload.GrandTotal != order.GrandTotal {
		return ErrWrongCalculate
	}

	order.AccessLogID = c.Response().Header().Get(echo.HeaderXRequestID)
	if err := order.Create(db); err != nil {
		return err
	}

	return c.JSONPretty(http.StatusCreated, order, JSONIndent)
}

func getOrdersMe(c echo.Context) error {
	seller, ok := c.Get("seller").(models.Seller)
	if !ok {
		return echo.ErrInternalServerError
	}

	payload := struct {
		Page  int `query:"page"`
		Limit int `query:"limit"`
	}{}

	if err := c.Bind(&payload); err != nil {
		return echo.ErrBadRequest
	}

	orders, err := models.OrdersByBoothID(db, seller.BoothID, payload.Page, payload.Limit)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSONPretty(http.StatusOK, orders, JSONIndent)

}

func getOrderByID(c echo.Context) error {
	seller, ok := c.Get("seller").(models.Seller)
	if !ok {
		return echo.ErrInternalServerError
	}

	orderID := c.Param("id")
	order, err := models.OrderByID(db, orderID)
	if err != nil {
		return echo.ErrNotFound
	}

	// If no permission
	if order.BoothID != seller.BoothID {
		return echo.ErrNotFound
	}

	return c.JSONPretty(http.StatusOK, order, JSONIndent)
}

func deleteOrderByID(c echo.Context) error {
	seller, ok := c.Get("seller").(models.Seller)
	if !ok {
		return echo.ErrInternalServerError
	}

	orderID := c.Param("id")
	order, err := models.OrderByID(db, orderID)
	if err != nil {
		return echo.ErrNotFound
	}

	if order.BoothID != seller.BoothID {
		return echo.ErrNotFound
	}

	if err := order.Cancel(db); err != nil {
		return err
	}

	return c.JSONPretty(http.StatusAccepted, order, JSONIndent)
}
