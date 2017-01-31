package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func initServer() {
	e := echo.New()
	e.GET("/data", GetEmployees())
	e.Logger.Fatal(e.Start(":1323"))
}

func GetEmployees() echo.HandlerFunc {
	return func(c echo.Context) error {
		rnd := random(10000, 500000)
		emp := loadEmployees(rnd)
		return c.JSON(http.StatusCreated, emp)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
