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
	e.GET("/api/v1/register/available", getAvailable)
	e.PATCH("/api/v1/register", patchRegister)

	api := e.Group("/api/v1", MiddlewareTokenCheck, MiddlewareLogger)
	{
		// But /api/logout does need auth process
		api.GET("/logout", getLogout)

		api.GET("/me", getMine)

		api.GET("/orders/me", getMyOrders)
		api.GET("/orders/:id", getOrderByID)
		api.POST("/orders", postOrder)
		api.PATCH("/orders/:id", patchOrderByID)

		api.GET("/users/:id", getUserByID)
		api.GET("/users/card/:code", getUserByCardCode)
	}

}
