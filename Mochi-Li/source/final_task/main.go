package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const ()



func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	e.Logger.Fatal(e.Start(":2334"))
}
