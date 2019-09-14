package main

import (
	"github.com/labstack/echo"
	"golang.org/x/crypto/acme/autocert"
)

func autoCertConfig(e *echo.Echo) {
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(Namespace)
	e.AutoTLSManager.Cache = autocert.DirCache("~/server/fespay/.cache")
}

func run() {
	e := echo.New()
	routes(e)
	if err := e.Start(":8080"); err != nil {
		panic(err)
	}
}
