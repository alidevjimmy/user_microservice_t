package app

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	e *echo.Echo
)

type Validator struct {
	validator *validator.Validate
}

func (uv *Validator) Validate(i interface{}) error {
	if err := uv.validator.Struct(i); err != nil {
		return rest_errors.NewBadRequestError(err.Error())
	}
	return nil
}

func StartApp(port string) {
	e = echo.New()
	e.Validator = &Validator{validator: validator.New()}
	urlMapper()
	e.Logger.Fatal(e.Start(port))
}
