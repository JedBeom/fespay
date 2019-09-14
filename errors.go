package main

import (
	"errors"
)

var (
	ErrorLoginFailed = errors.New("ID or PIN does not match")
)
