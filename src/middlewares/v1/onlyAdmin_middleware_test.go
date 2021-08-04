package middlewares

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1"
	"github.com/alidevjimmy/user_microservice_t/services/v1"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	getUserFunc func(token string) (*domains.PublicUser, rest_errors.RestErr)
)

type UserServiceMock struct{}

func (*UserServiceMock) Register(body domains.RegisterRequest) (*domains.RegisterResponse, rest_errors.RestErr) {
	return nil, nil
}

func (*UserServiceMock) Login(body domains.LoginRequest) (*domains.LoginResponse, rest_errors.RestErr) {
	return nil, nil
}

// GetUser returns single user by its jwt token
func (*UserServiceMock) GetUser(token string) (*domains.PublicUser, rest_errors.RestErr) {
	return getUserFunc(token)
}

// GetUsers returns all users by filter
func (*UserServiceMock) GetUsers(params domains.GetUsersRequest) ([]domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

// UpdateUserActiveState makes state of active field of user opposite
func (*UserServiceMock) UpdateUserActiveState(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

// UpdateUserBlockState makes state of blocked field of user opposite
func (*UserServiceMock) UpdateUserBlockState(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (*UserServiceMock) UpdateUser(userId, token string, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

// ChangeForgotPassword helps people who forgot their password using verification code
func (*UserServiceMock) ChangeForgotPassword(body domains.ChangePasswordRequest) (*domains.PublicUser, rest_errors.RestErr) {
	// return token
	return nil, nil
}

// ActiveUser Change user active state to true using verification code
func (*UserServiceMock) VerifyUser(body domains.VerifyUserRequest) (*domains.PublicUser, rest_errors.RestErr) {
	// return token
	return nil, nil
}

func TestOnlyAdminMiddlewareTokenDoesNotExists(t *testing.T) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, "")
	}, OnlyAdmin)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnauthorized, res.Code)
}

func TestOnlyAdminMiddlewareTokenIsEmpty(t *testing.T) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, "")
	}, OnlyAdmin)

	body := struct {
		Token string `json:"token"`
	}{
		Token: "",
	}
	j, _ := json.Marshal(body)
	nr := bytes.NewReader(j)

	req := httptest.NewRequest(http.MethodGet, "/", nr)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnauthorized, res.Code)
}

func TestOnlyAdminFailToGetUserByToken(t *testing.T) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, "")
	}, OnlyAdmin)

	body := struct {
		Token string `json:"token"`
	}{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ",
	}
	j, _ := json.Marshal(body)
	nr := bytes.NewReader(j)

	req := httptest.NewRequest(http.MethodGet, "/", nr)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnauthorized, res.Code)
}

func TestOnlyAdminUserIsNotAdmin(t *testing.T) {
	getUserFunc = func(token string) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID:      uint(1),
			IsAdmin: false,
		}, nil
	}

	services.UserService = &UserServiceMock{}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, "")
	}, OnlyAdmin)

	body := struct {
		Token string `json:"token"`
	}{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ",
	}
	j, _ := json.Marshal(body)
	nr := bytes.NewReader(j)

	req := httptest.NewRequest(http.MethodGet, "/", nr)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnauthorized, res.Code)
}

func TestOnlyAdmin(t *testing.T) {
	getUserFunc = func(token string) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID:      uint(1),
			IsAdmin: true,
		}, nil
	}

	services.UserService = &UserServiceMock{}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, "")
	}, OnlyAdmin)

	body := struct {
		Token string `json:"token"`
	}{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ",
	}
	j, _ := json.Marshal(body)
	nr := bytes.NewReader(j)

	req := httptest.NewRequest(http.MethodGet, "/", nr)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
