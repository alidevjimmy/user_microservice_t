package middlewares

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/errors/v1"
	"github.com/alidevjimmy/user_microservice_t/services/v1"
	"github.com/labstack/echo/v4"
	"net/http"
)

// OnlyActive Middleware make sure only who is active has access to continue
func OnlyActive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error{
		// get token from request header
		token := c.Request().Header.Get(echo.HeaderAuthorization)
		if token == "" {
			er := rest_errors.NewUnauthorizedError(errors.InvalidInputErrorMessage)
			return c.JSON(http.StatusUnauthorized, er)
		}
		// get user by token
		user , err := services.UserService.GetUser(token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		if !user.Active {
			er := rest_errors.NewUnauthorizedError(errors.UnAuthorizedAdminErrorMessage)
			return c.JSON(http.StatusUnauthorized, er)
		}
		return next(c)
	}
}
