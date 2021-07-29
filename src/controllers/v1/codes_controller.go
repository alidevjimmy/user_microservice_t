package v1

import "github.com/labstack/echo/v4"

func SendCode(c echo.Context) error {
	return c.JSON(200,"you are not registered!")
}

func VerifyCode(c echo.Context) error {
	// TODO: check that code is for forget or verification and then do what should be done
	return c.JSON(200,"you are not registered!")
}