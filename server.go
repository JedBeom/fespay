package main

import (
	"os"

	"github.com/labstack/echo"
	"golang.org/x/crypto/acme/autocert"
)

func autoCertConfig(e *echo.Echo) {
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(Namespace)
	e.AutoTLSManager.Cache = autocert.DirCache("~/sites/fespay/.cache")
}

func run() {
	e := echo.New()
	routes(e)

	sslMode := os.Getenv("SSL_MODE")
	if sslMode == "AUTO" {
		autoCertConfig(e)
		e.Logger.Fatal(e.StartAutoTLS(":443"))
	} else if sslMode == "MANUAL" {
		e.Logger.Fatal(e.StartTLS(":443", os.Getenv("SSL_CRT"), os.Getenv("SSL_PRI")))
	} else {
		e.Logger.Fatal(e.Start(":80"))
	}
}
