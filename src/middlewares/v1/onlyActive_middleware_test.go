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

func TestOnlyActiveMiddlewareTokenDoesNotExists(t *testing.T) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, "")
	}, OnlyActive)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnauthorized, res.Code)
}

func TestOnlyActiveMiddlewareTokenIsEmpty(t *testing.T) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, "")
	}, OnlyActive)

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

func TestOnlyActiveFailToGetUserByToken(t *testing.T) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, "")
	}, OnlyActive)

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

func TestOnlyActiveUserIsNotActive(t *testing.T) {
	getUserFunc = func(token string) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID:     uint(1),
			Active: false,
		}, nil
	}

	services.UserService = &UserServiceMock{}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, "")
	}, OnlyActive)

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

func TestOnlyActive(t *testing.T) {
	getUserFunc = func(token string) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID:     uint(1),
			Active: true,
		}, nil
	}

	services.UserService = &UserServiceMock{}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, "")
	}, OnlyActive)

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
