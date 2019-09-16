package main

import (
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

func routes(e *echo.Echo) {
	e.Use(echoMw.Recover())
	e.Use(echoMw.RequestID())

	// /api/login does not need auth process
	e.POST("/api/v1/login", postLogin)

	api := e.Group("/api/v1", MiddlewareTokenCheck, MiddlewareLogger)
	{
		// But /api/logout does need auth process
		api.GET("/logout", getLogout)

		api.GET("/me", getMine)

		api.GET("/sellers/me", getSellerMe)
		api.GET("/booths/me", getBoothMe)

		api.POST("/products", postProducts)
		// api.GET("/product/me", getProductsMe)
		// api.GET("/product/:id", getProductByID)
		// 구현 계획 미정: api.PUT("/product/:id", putProduct)

		api.GET("/orders/me", getOrdersMe)
		api.GET("/orders/:id", getOrderByID)
		api.POST("/orders", postOrder)
		api.DELETE("/orders/:id", deleteOrderByID)

		api.GET("/students/:barcode", getStudentByBarcodeID)
	}

}
