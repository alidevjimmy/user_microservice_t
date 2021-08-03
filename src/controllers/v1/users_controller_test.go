package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/app"
	"github.com/alidevjimmy/user_microservice_t/domains/v1"
	"github.com/alidevjimmy/user_microservice_t/errors/v1"
	"github.com/alidevjimmy/user_microservice_t/services/v1"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	RegisterRequest domains.RegisterRequest = domains.RegisterRequest{
		Phone:    "09122334344",
		Username: "test_user",
		Name:     "ali",
		Family:   "hamrani",
		Age:      uint(20),
		Password: "password",
	}
	LoginRequest domains.LoginRequest = domains.LoginRequest{
		PhoneOrUsername: "09122334344",
		Password:        "password",
	}
	GetUserRequest domains.GetUserRequest = domains.GetUserRequest{
		Token: "thisis.some.token",
	}

	getUserFunc  func(token string) (*domains.PublicUser, rest_errors.RestErr)
	registerFunc func(body domains.RegisterRequest) (*domains.RegisterResponse, rest_errors.RestErr)
	loginFunc    func(body domains.LoginRequest) (*domains.LoginResponse, rest_errors.RestErr)
)

type UserServiceMock struct{}

func (*UserServiceMock) Register(body domains.RegisterRequest) (*domains.RegisterResponse, rest_errors.RestErr) {
	return registerFunc(body)
}

func (*UserServiceMock) Login(body domains.LoginRequest) (*domains.LoginResponse, rest_errors.RestErr) {
	return loginFunc(body)
}

// GetUser returns single user by its jwt token
func (*UserServiceMock) GetUser(token string) (*domains.PublicUser, rest_errors.RestErr) {
	return getUserFunc(token)
}

// GetUsers returns all users by filter
func (*UserServiceMock) GetUsers(params map[string]interface{}) ([]domains.PublicUser, rest_errors.RestErr) {
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

func (*UserServiceMock) UpdateUser(token string, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

// ChangeForgotPassword helps people who forgot their password using verification code
func (*UserServiceMock) ChangeForgotPassword(newPassword, phone string, code int) (*domains.PublicUser, rest_errors.RestErr) {
	// return token
	return nil, nil
}

// ActiveUser Change user active state to true using verification code
func (*UserServiceMock) ActiveUser(phone string, code int) (*domains.PublicUser, rest_errors.RestErr) {
	// return token
	return nil, nil
}

func TestRegisterPhoneRequired(t *testing.T) {
	RegisterRequest.Phone = ""
	body := RegisterRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(app.V1Prefix, "register"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestRegisterUsernameRequired(t *testing.T) {
	RegisterRequest.Username = ""
	body := RegisterRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(app.V1Prefix, "register"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestRegisterNameRequired(t *testing.T) {
	RegisterRequest.Name = ""
	body := RegisterRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(app.V1Prefix, "register"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestRegisterFamilyRequired(t *testing.T) {
	RegisterRequest.Family = ""
	body := RegisterRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(app.V1Prefix, "register"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}
func TestRegisterAgeRequired(t *testing.T) {
	RegisterRequest.Age = 0
	body := RegisterRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(app.V1Prefix, "register"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}
func TestRegisterPasswordRequired(t *testing.T) {
	RegisterRequest.Password = ""
	body := RegisterRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(app.V1Prefix, "register"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestRegisterFailToBindReqBody(t *testing.T) {
	body := struct {
		Phone string
	}{
		Phone: "0091212",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(app.V1Prefix, "register"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestRegisterServiceReturnedError(t *testing.T) {
	registerFunc = func(body domains.RegisterRequest) (*domains.RegisterResponse, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	services.UserService = &UserServiceMock{}

	body := RegisterRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(app.V1Prefix, "register"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestLoginServiceReturnedError(t *testing.T) {
	loginFunc = func(body domains.LoginRequest) (*domains.LoginResponse, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	services.UserService = &UserServiceMock{}

	body := RegisterRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(app.V1Prefix, "register"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestLoginFailToBindReqBody(t *testing.T) {
	body := struct {
		Phone string
	}{
		Phone: "0091212",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(app.V1Prefix, "login"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestLoginPasswordRequired(t *testing.T) {
	LoginRequest.Password = ""
	body := LoginRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(app.V1Prefix, "login"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestLoginPhoneOrUsernameRequired(t *testing.T) {
	LoginRequest.PhoneOrUsername = ""
	body := LoginRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(app.V1Prefix, "login"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}
