package main

import (
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

func routes(e *echo.Echo) {
	e.Pre(echoMw.RemoveTrailingSlash())
	e.Use(echoMw.CORS())
	e.Use(echoMw.Recover())
	e.Use(echoMw.RequestID())
	e.Use(MiddlewareLogger)

	frontPrefix := "front/dist"
	e.File("/*", frontPrefix+"/index.html")
	e.File("/favicon.ico", frontPrefix+"/favicon.ico")
	e.Static("/css", frontPrefix+"/css")
	e.Static("/img", frontPrefix+"/img")
	e.Static("/js", frontPrefix+"/js")

	e.GET("/api/v1/register/available", getAvailable)
	e.PATCH("/api/v1/register", patchRegister)
	e.POST("/api/v1/login", postLogin)

	api := e.Group("/api/v1", MiddlewareTokenCheck)
	{
		api.GET("/logout", getLogout)

		api.GET("/user", getMine)

		api.GET("/booths", getBooths) // admin
		api.GET("/booths/:id", getBoothByID)
		api.GET("/booths/:id/records", getRecordsByBoothID)
		api.POST("/booths", postBooth)           // admin
		api.PATCH("/booths/:id", patchBoothByID) // admin

		api.GET("/records", getRecords) // admin
		api.GET("/records/:id", getRecordByID)
		api.POST("/records", postRecord)
		api.PATCH("/records/:id", patchRecordByID)
		api.DELETE("/records/:id", deleteRecordByID)

		api.GET("/users", getUsers) // admin
		api.GET("/users/:id", getUserByID)
		api.GET("/users/card/:code", getUserByCardCode) // admin
		api.GET("/users/:id/records", getRecordsByUserID)
		api.POST("/users", postUser)           // admin
		api.PATCH("/users/:id", patchUserByID) // admin
	}
}

type GetParam struct {
	Limit  int    `query:"limit"`
	Page   int    `query:"page"`
	Column string `query:"column"`
	Like   string `query:"like"`
}

func parseGetParam(c echo.Context) (GetParam, error) {
	p := GetParam{}
	if err := c.Bind(&p); err != nil {
		return p, echo.ErrBadRequest
	}

	if p.Limit == 0 || p.Limit > 50 {
		p.Limit = 20
	}

	if p.Page <= 0 {
		p.Page = 1
	}

	return p, nil
}
