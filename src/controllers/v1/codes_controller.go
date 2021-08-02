package controllers

import "github.com/labstack/echo/v4"

func SendCode(c echo.Context) error {
	return c.JSON(200, "you are not registered!")
}
