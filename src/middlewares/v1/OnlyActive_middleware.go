package middlewares

import "github.com/labstack/echo/v4"

func OnlyActive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error{
		return next(c)
	}
}
