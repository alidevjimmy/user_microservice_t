package controllers

import "github.com/labstack/echo/v4"

var UsersController usersControllerInterface = &usersController{}

type usersControllerInterface interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
	ActiveUser(c echo.Context) error
	BlockUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	ChangePassword(c echo.Context) error
	Verify(c echo.Context) error
}

type usersController struct{}

func (*usersController) Register(c echo.Context) error {

	return c.JSON(200, "you are not registered!")
}

func (*usersController) Login(c echo.Context) error {
	return c.JSON(200, "you are not registered!")
}

func (*usersController) GetUser(c echo.Context) error {
	return c.JSON(200, "you are not registered!")
}

func (*usersController) GetUsers(c echo.Context) error {
	return c.JSON(200, "you are not registered!")
}

func (*usersController) ActiveUser(c echo.Context) error {
	return c.JSON(200, "you are not registered!")
}

func (*usersController) BlockUser(c echo.Context) error {
	return c.JSON(200, "you are not registered!")
}

func (*usersController) UpdateUser(c echo.Context) error {
	return c.JSON(200, "you are not registered!")
}

func (*usersController) ChangePassword(c echo.Context) error {
	return c.JSON(200, "you are not registered!")
}

func (*usersController) Verify(c echo.Context) error {
	return c.JSON(200, "you are not registered!")
}
