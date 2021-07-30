package app

import (
	"github.com/labstack/echo/v4"
)

var (
	e *echo.Echo
)

func StartApp(port string) {
	e = echo.New()
	urlMapper()
	e.Logger.Fatal(e.Start(port))
}
