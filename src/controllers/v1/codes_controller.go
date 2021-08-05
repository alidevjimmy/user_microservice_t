package controllers

import (
	"net/http"

	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1"
	"github.com/alidevjimmy/user_microservice_t/errors/v1"
	"github.com/alidevjimmy/user_microservice_t/services/v1"
	"github.com/labstack/echo/v4"
)

var CodesController codesControllerInterface = &codesController{}

type codesControllerInterface interface {
	SendCode(c echo.Context) error
}

type codesController struct{}

func (*codesController) SendCode(c echo.Context) error {
	rq := new(domains.SendCodeRequest)
	if err := c.Bind(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	if err := c.Validate(rq); err != nil {
		er := rest_errors.NewBadRequestError(errors.InvalidInputErrorMessage)
		return c.JSON(http.StatusBadRequest, er)
	}
	err := services.CodeService.Send(*rq)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.NoContent(http.StatusOK)
}
