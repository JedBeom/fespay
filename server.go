package main

import (
	"context"
	"os"
	"os/signal"
	"time"

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

	go func() {
		if sslMode == "AUTO" {
			autoCertConfig(e)
			e.Logger.Fatal(e.StartAutoTLS(":443"))
		} else if sslMode == "MANUAL" {
			e.Logger.Fatal(e.StartTLS(":443", os.Getenv("SSL_CRT"), os.Getenv("SSL_PRI")))
		} else {
			e.Logger.Fatal(e.Start(":80"))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
