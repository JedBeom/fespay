package main

import (
	"net/http"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

const KeyHeaderKey = "Rabbit-Fur"

func MiddlewareTokenCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.Request().Header.Get(KeyHeaderKey)
		if key == "" {
			return echo.ErrUnauthorized
		}
		s, u, err := models.SessionAndUserByID(db, key)
		if err != nil || u.ID == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "bad key")
		}

		c.Set("sess_id", s.ID)
		c.Set("user", u)

		return next(c)
	}
}

func MiddlewareLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, ok := c.Get("sess_id").(string)
		if !ok {
			return echo.ErrInternalServerError
		}

		access := models.AccessLog{
			ID:        c.Response().Header().Get(echo.HeaderXRequestID),
			Path:      c.Path(),
			SessionID: sess,
			IP:        c.RealIP(),
		}

		if err := access.Create(db); err != nil {
			return echo.ErrInternalServerError
		}

		return next(c)
	}
}
