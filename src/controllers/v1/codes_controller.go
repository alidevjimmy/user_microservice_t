package controllers

import "github.com/labstack/echo/v4"

var CodesController codesControllerInterface = &codesController{}

type codesControllerInterface interface {
	SendCode(c echo.Context) error
}

type codesController struct {}

func (*codesController) SendCode(c echo.Context) error {
	return c.JSON(200, "you are not registered!")
}
