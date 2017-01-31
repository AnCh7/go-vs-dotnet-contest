package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func initServer() {
	e := echo.New()
	e.GET("/users", CreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}

func CreateUser(c echo.Context) error {
	emp := loadEmployees()
	return c.JSON(http.StatusCreated, emp)
}
