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
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	user, err := services.UserService.Register(*rq)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(http.StatusCreated, user)
}

func (*usersController) Login(c echo.Context) error {
	rq := new(domains.LoginRequest)
	if err := c.Bind(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	if err := c.Validate(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	user, err := services.UserService.Login(*rq)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(http.StatusOK, user)
}

func (*usersController) GetUser(c echo.Context) error {
	rq := new(domains.GetUserRequest)
	if err := c.Bind(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	if err := c.Validate(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	user, err := services.UserService.GetUser(rq.Token)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(http.StatusOK, user)
}

func (*usersController) GetUsers(c echo.Context) error {
	rq := new(domains.GetUsersRequest)
	if err := c.Bind(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	if err := c.Validate(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	user, err := services.UserService.GetUsers(*rq)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(http.StatusOK, user)
}

func (*usersController) UpdateUserActiveState(c echo.Context) error {
	rq := new(domains.UpdateActiveUserStateRequest)
	if err := c.Bind(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	if err := c.Validate(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	user, err := services.UserService.UpdateUserActiveState(rq.UserID)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(http.StatusOK, user)
}

func (*usersController) UpdateUserBlockState(c echo.Context) error {
	rq := new(domains.UpdateBlockUserStateRequest)
	if err := c.Bind(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	if err := c.Validate(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	user, err := services.UserService.UpdateUserBlockState(rq.UserID)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(http.StatusOK, user)
}

func (*usersController) UpdateUser(c echo.Context) error {
	rq := new(domains.UpdateUserRequest)
	if err := c.Bind(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	if err := c.Validate(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	token := c.QueryParam("token")
	userID := c.Param("user_id")
	user, err := services.UserService.UpdateUser(userID, token, *rq)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(http.StatusOK, user)
}

func (*usersController) ChangePassword(c echo.Context) error {
	rq := new(domains.ChangePasswordRequest)
	if err := c.Bind(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	if err := c.Validate(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	user, err := services.UserService.ChangeForgotPassword(*rq)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(http.StatusOK, user)
}

func (*usersController) Verify(c echo.Context) error {
	rq := new(domains.VerifyUserRequest)
	if err := c.Bind(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	if err := c.Validate(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	user, err := services.UserService.VerifyUser(*rq)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(http.StatusOK, user)
}
