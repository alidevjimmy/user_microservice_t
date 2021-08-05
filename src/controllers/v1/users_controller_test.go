package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
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
	LoginRequest domains.LoginRequest = domains.LoginRequest{
		PhoneOrUsername: "09122334344",
		Password:        "password",
	}
	GetUserRequest domains.GetUserRequest = domains.GetUserRequest{
		Token: "thisis.some.token",
	}

	c echo.Context
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
	return getUsersFunc(params)
}

// UpdateUserActiveState makes state of active field of user opposite
func (*UserServiceMock) UpdateUserActiveState(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
	return updateUserActiveStateFunc(userId)
}

// UpdateUserBlockState makes state of blocked field of user opposite
func (*UserServiceMock) UpdateUserBlockState(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
	return updateUserBlockStateFunc(userId)
}

func (*UserServiceMock) UpdateUser(userId, token string, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr) {
	return updateUserFunc(userId,token,body)
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

type RestErrStruct struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  int    `json:"status"`
}
type Validator struct {
	validator *validator.Validate
}

func (uv *Validator) Validate(i interface{}) error {
	if err := uv.validator.Struct(i); err != nil {
		return rest_errors.NewBadRequestError(err.Error())
	}
	return nil
}

func TestRegisterPhoneRequired(t *testing.T) {
	body := domains.RegisterRequest{
		Phone:    "",
		Username: "test_user",
		Name:     "ali",
		Family:   "hamrani",
		Age:      uint(20),
		Password: "password",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.Register(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
}

func TestRegisterUsernameRequired(t *testing.T) {
	body := domains.RegisterRequest{
		Phone:    "09122334344",
		Username: "",
		Name:     "ali",
		Family:   "hamrani",
		Age:      uint(20),
		Password: "password",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.Register(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestRegisterNameRequired(t *testing.T) {
	body := domains.RegisterRequest{
		Phone:    "09122334344",
		Username: "test_user",
		Name:     "",
		Family:   "hamrani",
		Age:      uint(20),
		Password: "password",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.Register(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestRegisterFamilyRequired(t *testing.T) {
	body := domains.RegisterRequest{
		Phone:    "09122334344",
		Username: "test_user",
		Name:     "ali",
		Family:   "",
		Age:      uint(20),
		Password: "password",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.Register(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}
func TestRegisterAgeRequired(t *testing.T) {
	body := domains.RegisterRequest{
		Phone:    "09122334344",
		Username: "test_user",
		Name:     "ali",
		Family:   "hamrani",
		Age:      0,
		Password: "password",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.Register(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}
func TestRegisterPasswordRequired(t *testing.T) {
	body := domains.RegisterRequest{
		Phone:    "09122334344",
		Username: "test_user",
		Name:     "ali",
		Family:   "hamrani",
		Age:      uint(20),
		Password: "",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.Register(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestRegisterFailToBindReqBody(t *testing.T) {
	body := struct {
		Phon string
	}{
		Phon: "0091212",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.Register(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestRegisterServiceReturnedError(t *testing.T) {
	registerFunc = func(body domains.RegisterRequest) (*domains.RegisterResponse, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}
	services.UserService = &UserServiceMock{}

	body := domains.RegisterRequest{
		Phone:    "09122334344",
		Username: "test_user",
		Name:     "ali",
		Family:   "hamrani",
		Age:      uint(20),
		Password: "password",
	}
	j, err := json.Marshal(body)
	rd := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rd)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "register"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.Register(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, rec.Code)
	assert.EqualValues(t, errors.InternalServerErrorMessage, restErr.Message)
}

func TestLoginServiceReturnedError(t *testing.T) {
	loginFunc = func(body domains.LoginRequest) (*domains.LoginResponse, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	services.UserService = &UserServiceMock{}

	body := LoginRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "login"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.Login(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, rec.Code)
	assert.EqualValues(t, errors.InternalServerErrorMessage, restErr.Message)
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
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "login"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.Login(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestLoginPasswordRequired(t *testing.T) {
	LoginRequest.Password = ""
	body := LoginRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "login"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.Login(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestLoginPhoneOrUsernameRequired(t *testing.T) {
	LoginRequest.PhoneOrUsername = ""
	body := LoginRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "login"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.Login(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
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
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "getUser"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.GetUser(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
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
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "getUser"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.GetUser(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, rec.Code)
	assert.EqualValues(t, errors.InternalServerErrorMessage, restErr.Message)
}

func TestGetUserTokenRequired(t *testing.T) {
	GetUserRequest.Token = ""
	body := GetUserRequest
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "getUser"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.GetUser(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestGetUsersServiceReturnedError(t *testing.T) {
	getUsersFunc = func(param domains.GetUsersRequest) ([]domains.PublicUser, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	services.UserService = &UserServiceMock{}

	body := domains.GetUsersRequest{}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "admin/users"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.GetUsers(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, rec.Code)
	assert.EqualValues(t, errors.InternalServerErrorMessage, restErr.Message)
}
//func TestGetUsersFailToBindReqBody(t *testing.T) {
//	body := struct {
//		Phone string
//	}{
//		Phone: "0091212",
//	}
//	j, err := json.Marshal(body)
//	rb := bytes.NewReader(j)
//	req := httptest.NewRequest(http.MethodPost, "/", rb)
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
//	rec := httptest.NewRecorder()
//	c = echo.New().NewContext(req, rec)
//	c.SetPath(fmt.Sprintf(v1prefix, "admin/users"))
//	c.Echo().Validator = &Validator{validator: validator.New()}
//	err = UsersController.GetUsers(c)
//	var restErr RestErrStruct
//	assert.Nil(t, err)
//	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
//	assert.Nil(t, err)
//	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
//	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
//}

func TestUpdateUserActiveStateFailToBindReqBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "admin/toggleActive:user_id"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	c.SetParamNames("user_id")
	c.SetParamValues("1")
	err := UsersController.UpdateUserActiveState(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
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
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "admin/toggleActive:user_id"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.UpdateUserActiveState(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, rec.Code)
	assert.EqualValues(t, errors.InternalServerErrorMessage, restErr.Message)
}

func TestUpdateUserActiveStateUserIdRequired(t *testing.T) {
	body := domains.UpdateActiveUserStateRequest{}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "admin/toggleActive:user_id"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.UpdateUserActiveState(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestUpdateUserBlockStateFailToBindReqBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "admin/toggleBlock:user_id"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	c.SetParamNames("user_id")
	c.SetParamValues("1")
	err := UsersController.UpdateUserBlockState(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
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
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "admin/toggleBlock:user_id"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.UpdateUserBlockState(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, rec.Code)
	assert.EqualValues(t, errors.InternalServerErrorMessage, restErr.Message)
}

func TestUpdateUserBlockStateUserIdRequired(t *testing.T) {
	body := domains.UpdateBlockUserStateRequest{}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "admin/toggleBlock:user_id"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = UsersController.UpdateUserBlockState(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestChangePasswordFailToBindReqBody(t *testing.T) {
	body := domains.UpdateActiveUserStateRequest{
		UserID: uint(1),
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "changePassword"))
	err = UsersController.ChangePassword(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
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
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "changePassword"))
	err = UsersController.ChangePassword(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
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
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "changePassword"))
	err = UsersController.ChangePassword(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestChangePasswordCodeRequired(t *testing.T) {
	body := domains.ChangePasswordRequest{
		NewPassword: "1234556",
		Phone:       "09912323",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "changePassword"))
	err = UsersController.ChangePassword(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestChangePasswordNewPasswordRequired(t *testing.T) {
	body := domains.ChangePasswordRequest{
		Phone: "09912323",
		Code:  23123,
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "changePassword"))
	err = UsersController.ChangePassword(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

// verify user test

func TestVerifyUserFailToBindReqBody(t *testing.T) {
	body := domains.UpdateActiveUserStateRequest{
		UserID: uint(1),
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "verifyUser"))
	err = UsersController.Verify(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
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
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "verifyUser"))
	err = UsersController.Verify(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestVerifyUserPhoneRequired(t *testing.T) {
	body := domains.VerifyUserRequest{
		Code: 23123,
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "verifyUser"))
	err = UsersController.Verify(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestVerifyUserCodeRequired(t *testing.T) {
	body := domains.VerifyUserRequest{
		Phone: "09912323",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "verifyUser"))
	err = UsersController.Verify(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, rec.Code)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestUpdateUserServiceReturnedError(t *testing.T) {
	updateUserFunc = func(userId, token string, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	services.UserService = &UserServiceMock{}

	body := domains.UpdateUserRequest{
		Username: "user2",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPatch, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "updateUser:user_id"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	c.SetParamNames("user_id")
	c.SetParamValues("1")
	assert.EqualValues(t, http.MethodPatch, c.Request().Method)
	err = UsersController.UpdateUser(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, rec.Code)
	assert.EqualValues(t, errors.InternalServerErrorMessage, restErr.Message)
}
