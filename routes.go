package main

import (
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

func routes(e *echo.Echo) {
	e.Use(echoMw.Recover())
	e.Use(echoMw.RequestID())
	e.Use(MiddlewareLogger)

	e.GET("/api/v1/register/available", getAvailable)
	e.PATCH("/api/v1/register", patchRegister)
	e.POST("/api/v1/login", postLogin)

	api := e.Group("/api/v1", MiddlewareTokenCheck)
	{
		api.GET("/logout", getLogout)

		api.GET("/user", getMine)

		api.GET("/booths/:id", getBoothByID)
		api.GET("/booths/:id/records", getRecordsByBoothID)

		api.GET("/records/:id", getRecordByID)
		api.POST("/records", postRecord)
		api.PATCH("/records/:id", patchRecordByID)
		api.DELETE("/records/:id", deleteRecordByID)

		api.GET("/users/:id", getUserByID)
		api.GET("/users/card/:code", getUserByCardCode)
		api.GET("/users/:id/records", getRecordsByUserID)
	}

}
