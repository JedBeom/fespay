package main

import (
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
		if err != nil || u.ID == "" || u.Status == models.StatusSuspended {
			return ErrInvalidKey.Send(c)
		}

		// 아예 부스 없는 척 해버리기~
		if u.Booth.Status == models.StatusSuspended {
			u.BoothID = ""
			u.Booth = nil
		}

		c.Set("sess_id", s.ID)
		c.Set("user", u)

		return next(c)
	}
}

func MiddlewareLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		access := models.AccessLog{
			ID:     c.Response().Header().Get(echo.HeaderXRequestID),
			Method: c.Scheme(),
			Path:   c.Path(),
			IP:     c.RealIP(),
		}
		err := next(c)
		sID, ok := c.Get("sess_id").(string)
		if ok {
			access.SessionID = sID
		}

		if err := access.Create(db); err != nil {
			echo.Logger.Error("Access Logging:", err)
		}
		return err

	}
}
