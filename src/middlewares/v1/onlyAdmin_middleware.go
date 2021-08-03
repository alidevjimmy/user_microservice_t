package middlewares

import "github.com/labstack/echo/v4"

func OnlyAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error{
		return next(c)
	}
}
