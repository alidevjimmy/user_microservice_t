package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alidevjimmy/go-rest-utils/rest_errors"
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

	getUserFunc               func(token string) (*domains.PublicUser, rest_errors.RestErr)
	getUsersFunc              func(params domains.GetUsersRequest) ([]domains.PublicUser, rest_errors.RestErr)
	registerFunc              func(body domains.RegisterRequest) (*domains.RegisterResponse, rest_errors.RestErr)
	loginFunc                 func(body domains.LoginRequest) (*domains.LoginResponse, rest_errors.RestErr)
	updateUserActiveStateFunc func(userId uint) (*domains.PublicUser, rest_errors.RestErr)
	updateUserBlockStateFunc  func(userId uint) (*domains.PublicUser, rest_errors.RestErr)
	changePasswordFunc        func(body domains.ChangePasswordRequest) (*domains.PublicUser, rest_errors.RestErr)
	verifyUserFunc            func(body domains.VerifyUserRequest) (*domains.PublicUser, rest_errors.RestErr)
	updateUserFunc            func(userId, token string, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr)
)

const (
	v1prefix = "/v1/%s"
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
	return changePasswordFunc(body)
}

// ActiveUser Change user active state to true using verification code
func (*UserServiceMock) VerifyUser(body domains.VerifyUserRequest) (*domains.PublicUser, rest_errors.RestErr) {
	// return token
	return verifyUserFunc(body)
}

func TestRegisterPhoneRequired(t *testing.T) {
	RegisterRequest.Phone = ""
	body := RegisterRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
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
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
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
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
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
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
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
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
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
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
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
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
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
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
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
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
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
	c.SetPath(fmt.Sprintf(v1prefix, "login"))
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
	c.SetPath(fmt.Sprintf(v1prefix, "login"))
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
	c.SetPath(fmt.Sprintf(v1prefix, "login"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestGetUserFailToBindReqBody(t *testing.T) {
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
	c.SetPath(fmt.Sprintf(v1prefix, "getUser"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestGetUserServiceReturnedError(t *testing.T) {
	getUserFunc = func(token string) (*domains.PublicUser, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	services.UserService = &UserServiceMock{}

	body := GetUserRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "getUser"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestGetUserTokenRequired(t *testing.T) {
	GetUserRequest.Token = ""
	body := GetUserRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "getUser"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestGetUsersServiceReturnedError(t *testing.T) {
	getUsersFunc = func(param domains.GetUsersRequest) ([]domains.PublicUser, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	services.UserService = &UserServiceMock{}

	body := GetUserRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "admin/users"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}
func TestGetUsersFailToBindReqBody(t *testing.T) {
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
	c.SetPath(fmt.Sprintf(v1prefix, "admin/users"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestUpdateUserActiveStateFailToBindReqBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "admin/toggleActive:user_id"))
	c.SetParamNames("user_id")
	c.SetParamValues("1")
	err := UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestUpdateUserActiveStateServiceReturnedError(t *testing.T) {
	updateUserActiveStateFunc = func(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	services.UserService = &UserServiceMock{}

	body := domains.UpdateActiveUserStateRequest{
		UserID: uint(1),
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "admin/toggleActive:user_id"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestUpdateUserActiveStateUserIdRequired(t *testing.T) {
	body := domains.UpdateActiveUserStateRequest{}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "admin/toggleActive:user_id"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestUpdateUserBlockStateFailToBindReqBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "admin/toggleBlock:user_id"))
	c.SetParamNames("user_id")
	c.SetParamValues("1")
	err := UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestUpdateUserBlockStateServiceReturnedError(t *testing.T) {
	updateUserBlockStateFunc = func(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	services.UserService = &UserServiceMock{}

	body := domains.UpdateBlockUserStateRequest{
		UserID: uint(1),
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "admin/toggleBlock:user_id"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestUpdateUserBlockStateUserIdRequired(t *testing.T) {
	body := domains.UpdateBlockUserStateRequest{}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "admin/toggleBlock:user_id"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestChangePasswordFailToBindReqBody(t *testing.T) {
	body := domains.UpdateActiveUserStateRequest{
		UserID: uint(1),
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "changePassword"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestChangePasswordServiceReturnedError(t *testing.T) {
	changePasswordFunc = func(body domains.ChangePasswordRequest) (*domains.PublicUser, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	services.UserService = &UserServiceMock{}

	body := domains.ChangePasswordRequest{
		NewPassword: "1234556",
		Phone:       "09912323",
		Code:        23123,
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "changePassword"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestChangePasswordPhoneRequired(t *testing.T) {
	body := domains.ChangePasswordRequest{
		NewPassword: "1234556",
		Phone:       "",
		Code:        23123,
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "changePassword"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestChangePasswordCodeRequired(t *testing.T) {
	body := domains.ChangePasswordRequest{
		NewPassword: "1234556",
		Phone:       "09912323",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "changePassword"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestChangePasswordNewPasswordRequired(t *testing.T) {
	body := domains.ChangePasswordRequest{
		Phone: "09912323",
		Code:  23123,
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "changePassword"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

// verify user test

func TestVerifyUserFailToBindReqBody(t *testing.T) {
	body := domains.UpdateActiveUserStateRequest{
		UserID: uint(1),
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "verifyUser"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestVerifyUserServiceReturnedError(t *testing.T) {
	verifyUserFunc = func(body domains.VerifyUserRequest) (*domains.PublicUser, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	services.UserService = &UserServiceMock{}

	body := domains.VerifyUserRequest{
		Phone: "09912323",
		Code:  23123,
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "verifyUser"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestVerifyUserPhoneRequired(t *testing.T) {
	body := domains.VerifyUserRequest{
		Code: 23123,
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "verifyUser"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestVerifyUserCodeRequired(t *testing.T) {
	body := domains.VerifyUserRequest{
		Phone: "09912323",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "verifyUser"))
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

// test update user

func TestUpdateUserFailToBindReqBody(t *testing.T) {
	body := domains.UpdateActiveUserStateRequest{
		UserID: uint(1),
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "updateUser:user_id"))
	c.SetParamNames("user_id")
	c.SetParamValues("1")
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}

func TestUpdateUsereReturnedError(t *testing.T) {
	updateUserFunc = func(userId, token string, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	services.UserService = &UserServiceMock{}

	body := domains.UpdateUserRequest{
		Username: "user2",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "updateUser:user_id"))
	c.SetParamNames("user_id")
	c.SetParamValues("1")
	err = UsersController.Register(c)
	assert.NotNil(t, err)
}
