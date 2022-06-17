package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Create an instance
	e := echo.New()
	// Set middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Set rootings
	e.GET("/", hello)
	e.GET("/show", show)
	e.GET("/users/:name", getUerName)
	e.POST("/save", save)
	e.POST("/users", saveUser)
	// Start this server with port1232
	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Echo!")
}
func getUerName(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}
func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team: "+team+", member: "+member)
}
func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name: "+name+", email: "+email)
}
func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
