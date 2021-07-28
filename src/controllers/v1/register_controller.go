package v1

import "github.com/labstack/echo/v4"

// Register is controller for adding new users
func Register(c echo.Context) error{
	return c.JSON(200,"you are not registered!")
}