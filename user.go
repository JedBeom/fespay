package main

import (
	"net/http"
	"time"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

func getUsers(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	if u.BoothID != AdminBoothID {
		return echo.ErrForbidden
	}

	p, err := parseGetParam(c)
	if err != nil {
		return err
	}

	var us []models.User
	switch p.Column {
	case "name":
		if p.Like == "" {
			return ErrField.Send(c)
		}
		us, err = models.UsersContainName(db, p.Like, p.Limit, p.Page)
	case "":
		return ErrField.Send(c)
	default:
		us, err = models.Users(db, p.Column, p.Like, p.Limit, p.Page)
	}

	if err != nil {
		return err2ApiErr(err).Send(c)
	}

	return c.JSONPretty(http.StatusOK, us, JSONIndent)
}

func getUserByID(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	if u.BoothID != AdminBoothID {
		return echo.ErrForbidden
	}

	userID := c.Param("id")
	tu, err := models.UserByID(db, userID, true)
	if err != nil {
		return err2ApiErr(err).Send(c)
	}

	return c.JSONPretty(http.StatusOK, tu, JSONIndent)
}

func getUserByCardCode(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	if u.BoothID != AdminBoothID {
		return echo.ErrForbidden
	}

	code := c.Param("code")
	tu, err := models.UserByCardCode(db, code)
	if err != nil {
		return err2ApiErr(err).Send(c)
	}

	return c.JSONPretty(http.StatusOK, tu, JSONIndent)
}

func patchUserByID(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	if u.BoothID != AdminBoothID {
		return echo.ErrForbidden
	}

	tuID := c.Param("id")
	p := models.User{}

	if err := c.Bind(&p); err != nil {
		return echo.ErrBadRequest
	}

	if p.ID != "" || p.ID != tuID {
		return ErrUserMismatch.Send(c)
	}

	tu, err := models.UserByID(db, tuID, false)
	if err != nil {
		return err2ApiErr(err).Send(c)
	}

	if p.LoginID != "" {
		tu.LoginID = p.LoginID
	}

	if p.Password != "" {
		tu.Password, err = models.Encrypt(p.LoginID)
		if err != nil {
			return err2ApiErr(err).Send(c)
		}
	}

	if p.Grade != 0 {
		tu.Grade = p.Grade
	}

	if p.Class != 0 {
		tu.Class = p.Class
	}

	if p.Number != 0 {
		tu.Number = p.Number
	}

	if p.Type != 0 {
		switch p.Type {
		case models.UserStudent:
			if tu.Grade == 0 || tu.Class == 0 || tu.Number == 0 {
				return ErrField.Send(c)
			}

			tu.Type = p.Type
		default:
			tu.Grade = 0
			tu.Class = 0
			tu.Number = 0
			tu.Type = p.Type
		}
	}

	if p.Name != "" {
		tu.Name = p.Name
	}

	if p.CardCode != "" {
		tu.CardCode = p.CardCode
	}

	if p.Status != 0 {
		tu.Status = p.Status
	}

	now := time.Now()
	tu.UpdatedAt = &now

	if err := db.Update(&tu); err != nil {
		return err2ApiErr(err).Send(c)
	}

	return c.JSONPretty(http.StatusOK, tu, JSONIndent)
}

func postUser(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	if u.BoothID != AdminBoothID {
		return echo.ErrForbidden
	}

	tu := models.User{}
	if err := c.Bind(&tu); err != nil {
		return echo.ErrBadRequest
	}

	tu.Coin = 0
	tu.Status = models.StatusWorking
	if err := tu.Create(db); err != nil {
		return err2ApiErr(err).Send(c)
	}
	return c.JSONPretty(http.StatusCreated, tu, JSONIndent)
}
