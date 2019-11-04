package main

import (
	"net/http"
	"time"

	"github.com/JedBeom/fespay/models"
	"github.com/labstack/echo"
)

func getBooths(c echo.Context) error {
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

	var bs []models.Booth
	switch p.Column {
	case "name":
		if p.Like == "" {
			return ErrField.Send(c)
		}
		bs, err = models.BoothsSearchName(db, p.Like, p.Limit, p.Page)
	case "":
		return ErrField.Send(c)
	default:
		bs, err = models.Booths(db, p.Column, p.Limit, p.Page)
	}

	if err != nil {
		return err2ApiErr(err).Send(c)
	}

	return c.JSONPretty(http.StatusOK, bs, JSONIndent)
}

func getBoothByID(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	boothID := c.Param("id")

	if u.BoothID != AdminBoothID && boothID != u.BoothID {
		return echo.ErrForbidden
	}

	b, err := models.BoothByID(db, boothID, true)
	if err != nil {
		return err2ApiErr(err).Send(c)
	}

	return c.JSONPretty(http.StatusOK, b, JSONIndent)
}

func patchBoothByID(c echo.Context) error {
	u, ok := c.Get("user").(models.User)
	if !ok {
		return ErrInterface.Send(c)
	}

	if u.BoothID != AdminBoothID {
		return echo.ErrForbidden
	}

	tbID := c.Param("id")
	p := models.Booth{}
	if err := c.Bind(&p); err != nil {
		return echo.ErrBadRequest
	}

	if p.ID != "" && tbID != p.ID {
		return ErrBoothMismatch.Send(c)
	}

	tb, err := models.BoothByID(db, tbID, false)
	if err != nil {
		return echo.ErrNotFound
	}

	if p.Name != "" {
		tb.Name = p.Name
	}

	if p.Description != "" {
		tb.Description = p.Description
	}

	if p.Location != "" {
		tb.Location = p.Location
	}

	if p.Status != 0 {
		tb.Status = p.Status
	}

	now := time.Now()
	tb.UpdatedAt = &now
	if err := db.Update(&tb); err != nil {
		return err2ApiErr(err).Send(c)
	}

	return c.JSONPretty(http.StatusOK, tb, JSONIndent)
}

func postBooth(c echo.Context) error {
	return nil
}
