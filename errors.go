package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

var (
	ErrLoginFailed     = errors.New("ID or PIN does not match")
	ErrStudentNotFound = echo.NewHTTPError(http.StatusNotFound, "student not found")
	ErrWrongCalculate  = echo.NewHTTPError(http.StatusBadRequest, "total was wrong")
)
