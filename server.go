package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Create an instance
	e := echo.New()
	// Set middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Set a root
	e.GET("/", hello)
	// Start this server with port1232
	e.Logger.Fatal(e.Start(":1232"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Echo!")
}
