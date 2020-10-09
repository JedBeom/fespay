package main

import (
	"fmt"
	"net/http"

	"github.com/JedBeom/fespay/models"

	"github.com/go-pg/pg"

	"github.com/labstack/echo"
)

type ApiError struct {
	StatusCode int    `json:"-"`
	ErrorCode  int    `json:"errorCode"`
	Message    string `json:"message"`
}

func NewApiError(status int, code int, message string) ApiError {
	return ApiError{
		StatusCode: status,
		ErrorCode:  code,
		Message:    message,
	}
}

func (e ApiError) Send(c echo.Context) error {
	return c.JSONPretty(e.StatusCode, e, JSONIndent)
}

func (e ApiError) Error() string {
	return fmt.Sprintf("Status: %d, Error: %d: %s", e.StatusCode, e.ErrorCode, e.Message)
}

var (
	ErrUnknown   = NewApiError(502, -10, "unknown error was aborted")
	ErrDBErr     = NewApiError(http.StatusInternalServerError, -11, "error occur")
	ErrInterface = NewApiError(http.StatusInternalServerError, -12, "err occur")

	ErrLoginFailed    = NewApiError(http.StatusUnauthorized, -100, "loginID or password is invalid")
	ErrInvalidKey     = NewApiError(http.StatusUnauthorized, -101, "bad uuid key")
	ErrEntityNotFound = NewApiError(http.StatusNotFound, -102, "entity not found")
	ErrField          = NewApiError(http.StatusUnprocessableEntity, -103, "field error")

	ErrUserNotInBooth = NewApiError(http.StatusForbidden, -200, "user isn't in a booth")
	ErrFrozenUser     = NewApiError(http.StatusBadRequest, -201, "user is frozen")
	ErrFrozenBooth    = NewApiError(http.StatusBadRequest, -202, "booth is frozen")
	ErrUserMismatch   = NewApiError(http.StatusUnprocessableEntity, -210, "user id is mismatch")
	ErrBoothMismatch  = NewApiError(http.StatusUnprocessableEntity, -211, "booth id is mismatch")

	ErrUnknownRecordType = NewApiError(http.StatusUnprocessableEntity, -300, "unknown record type")
	ErrInvalidAmount     = NewApiError(http.StatusBadRequest, -301, "amount should be in the unit of 100")
	ErrInvalidCardCode   = NewApiError(http.StatusNotFound, -305, "invalid cardCode")
	ErrCancelingAgain    = NewApiError(http.StatusUnprocessableEntity, -310, "record was canceled before")
	ErrExpiredRecord     = NewApiError(http.StatusUnprocessableEntity, -320, "record is expired")
)

func err2ApiErr(err error) ApiError {
	apiError, ok := err.(ApiError)
	if ok {
		return apiError
	}

	_, ok = err.(pg.Error)
	if ok {
		return ErrDBErr
	}

	fieldErr, ok := err.(models.FieldError)
	if ok {
		ef := ErrField
		ef.Message = fieldErr.Error()
		return ef
	}

	echoErr, ok := err.(*echo.HTTPError)
	if ok {
		return NewApiError(echoErr.Code, 0, http.StatusText(echoErr.Code))
	}

	if err == pg.ErrNoRows {
		return ErrEntityNotFound
	}

	return ErrUnknown
}
