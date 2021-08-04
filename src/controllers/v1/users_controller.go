package controllers

import (
	"net/http"

	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1"
	"github.com/alidevjimmy/user_microservice_t/errors/v1"
	"github.com/alidevjimmy/user_microservice_t/services/v1"
	"github.com/labstack/echo/v4"
)

var UsersController usersControllerInterface = &usersController{}

type usersControllerInterface interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
	UpdateUserActiveState(c echo.Context) error
	UpdateUserBlockState(c echo.Context) error
	UpdateUser(c echo.Context) error
	ChangePassword(c echo.Context) error
	Verify(c echo.Context) error
}

type usersController struct{}

func (*usersController) Register(c echo.Context) error {
	rq := new(domains.RegisterRequest)
	if err := c.Bind(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	if err := c.Validate(rq); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	user, err := services.UserService.Register(*rq)
	if err != nil {
		c.JSON(err.Status(), err)
	}
	return c.JSON(200, user)
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

func (*usersController) UpdateUserActiveState(c echo.Context) error {
	return c.JSON(200, "you are not registered!")
}

func (*usersController) UpdateUserBlockState(c echo.Context) error {
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
