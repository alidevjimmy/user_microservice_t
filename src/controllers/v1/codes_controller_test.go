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
	sendCodeFunc func(body domains.SendCodeRequest) rest_errors.RestErr
)

type CodeServiceMock struct{}

func (*CodeServiceMock) Send(body domains.SendCodeRequest) rest_errors.RestErr {
	return sendCodeFunc(body)
}

func (*CodeServiceMock) Verify(phone string, code, reason int) (bool, rest_errors.RestErr) {
	return false, nil
}

func TestSendCodeServiceReturnedError(t *testing.T) {
	sendCodeFunc = func(body domains.SendCodeRequest) rest_errors.RestErr {
		return rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	services.CodeService = &CodeServiceMock{}

	body := domains.SendCodeRequest{
		Phone:  "09233",
		Reason: services.VERIFICATION,
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "sendCode"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = CodesController.SendCode(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, restErr.Status)
	assert.EqualValues(t, errors.InternalServerErrorMessage, restErr.Message)
}

func TestSendCodePhoneRequired(t *testing.T) {
	body := domains.SendCodeRequest{
		Reason: services.RESETPASSWORD,
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "sendCode"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = CodesController.SendCode(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, restErr.Status)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}

func TestSendCodeReasonRequired(t *testing.T) {
	body := domains.SendCodeRequest{
		Phone: "0901923",
	}
	j, err := json.Marshal(body)
	rb := bytes.NewReader(j)
	req := httptest.NewRequest(http.MethodPost, "/", rb)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c = echo.New().NewContext(req, rec)
	c.SetPath(fmt.Sprintf(v1prefix, "sendCode"))
	c.Echo().Validator = &Validator{validator: validator.New()}
	err = CodesController.SendCode(c)
	var restErr RestErrStruct
	assert.Nil(t, err)
	err = json.Unmarshal(rec.Body.Bytes(), &restErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, restErr.Status)
	assert.EqualValues(t, errors.InvalidInputErrorMessage, restErr.Message)
}
