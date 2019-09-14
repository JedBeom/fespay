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

	api := e.Group("/api/v1", MiddlewareKeyCheck)
	{
		// But /api/logout does need auth process
		api.GET("/logout", getLogout)

		api.GET("/me", getMine)

		api.GET("/seller/me", getSellerMe)
		api.GET("/booth/me", getBoothMe)

		api.POST("/product", postProducts)
		// api.GET("/product/me", getProductsMe)
		// api.GET("/product/:id", getProductByID)
		// 구현 계획 없음: api.PUT("/product", putProduct)

		// api.GET("/order/me", getOrdersMe)
		// api.GET("/order/:id", getOrderByID)
		// api.POST("/order", postOrder)
		// api.DELETE("/order/:id", getOrderByID)

		api.GET("/student/:barcode", getStudentByBarcodeID)
	}

}
