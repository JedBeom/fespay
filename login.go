package main

import (
	"net/http"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

func postLogin(c echo.Context) error {
	p := struct {
		LoginID  string `json:"loginID"`
		Password string `json:"password"`
	}{}
	if err := c.Bind(&p); err != nil {
		return ErrLoginFailed.Send(c)
	}

	u, err := models.UserByLoginID(db, p.LoginID)
	if err != nil {
		return ErrLoginFailed.Send(c)
	}

	if u.Password == models.Encrypt(p.Password) {
		sess, err := u.NewSession(db, c.Request().UserAgent())
		if err == nil {
			return c.JSONPretty(200, Map{
				"token": sess.ID,
			}, JSONIndent)
		}
	}

	return ErrLoginFailed.Send(c)

}

func getLogout(c echo.Context) error {
	sessID, ok := c.Get("sess_id").(string)
	if !ok {
		return echo.ErrInternalServerError
	}

	sess := models.Session{
		ID: sessID,
	}

	if err := sess.Delete(db); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSONPretty(http.StatusOK, Map{
		"message": "log out success",
	}, JSONIndent)

}
