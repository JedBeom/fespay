package main

import (
	"fmt"
	"net/http"

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

func (e *ApiError) Send(c echo.Context) error {
	return c.JSONPretty(e.StatusCode, e, JSONIndent)
}

func (e ApiError) Error() string {
	return fmt.Sprintf("Status: %d, Error: %d: %s", e.StatusCode, e.ErrorCode, e.Message)
}

var (
	ErrUnknown = NewApiError(502, -10, "unknown error was aborted")

	ErrLoginFailed = NewApiError(http.StatusUnauthorized, -100, "loginID or password is invalid")
	ErrInvalidKey  = NewApiError(http.StatusUnauthorized, -101, "bad uuid key")

	ErrUserNotInBooth = NewApiError(http.StatusForbidden, -200, "user isn't in a booth")

	ErrInvalidAmount        = NewApiError(http.StatusBadRequest, -301, "amount should be in the unit of 100")
	ErrRefundOrderIDInvalid = NewApiError(http.StatusNotFound, -302, "refundOrderID is invalid")
	ErrRefundRefunding      = NewApiError(http.StatusBadRequest, -303, "refunding order can't be refunded")
	ErrRefundedAgain        = NewApiError(http.StatusBadRequest, -304, "refunded order is being refunded again")
	ErrUserWalletIDInvalid  = NewApiError(http.StatusBadRequest, -305, "userWalletID is invalid")
	ErrBoothToBoothOrder    = NewApiError(http.StatusBadRequest, -306, "booth2booth order isn't supported")
)
