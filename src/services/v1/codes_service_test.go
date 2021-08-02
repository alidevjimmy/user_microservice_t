package services

import (
	"fmt"
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1"
	"github.com/alidevjimmy/user_microservice_t/errors/v1"
	"github.com/alidevjimmy/user_microservice_t/repositories/postgres/v1"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"os"
	"testing"
	"time"
)


var (
	findCodeFunc func(phone string , code , reason int) (*domains.Code , rest_errors.RestErr)
)

type CodeRepoMock struct {
	DB *gorm.DB
}

func (c *CodeRepoMock) FindCode(phone string , code , reason int) (*domains.Code, rest_errors.RestErr) {
	return findCodeFunc(phone , code , reason)
}

func TestSendCodeFailToGetDataFromRepo(t *testing.T) {
	getUserFunc = func(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
		return nil , rest_errors.NewInternalServerError(errors.InternalServerErrorMessage,nil)
	}
	repositories.UserRepository = &UserRespositoryMock{}

	err := CodeService.Send(registerRequest.Phone , 0)
	assert.NotNil(t, err)
	assert.Equal(t, errors.InternalServerErrorMessage , err.Message())
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestSendCodeToUnExistsUser(t *testing.T) {
	getUserFunc = func(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
		return nil , nil
	}
	repositories.UserRepository = &UserRespositoryMock{}

	err := CodeService.Send(registerRequest.Phone , 0)
	assert.NotNil(t, err)
	assert.Equal(t, errors.UserNotFoundError , err.Message())
	assert.Equal(t, http.NotFound, err.Status())
}

func TestSendCodeToActiveUser(t *testing.T) {
	getUserFunc = func(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID: uint(1),
			Phone: registerRequest.Phone,
			Active: true,
		} , nil
	}
	repositories.UserRepository = &UserRespositoryMock{}

	err := CodeService.Send(registerRequest.Phone , 0)
	assert.NotNil(t, err)
	assert.Equal(t, errors.UserAlreadyActiveErrorMessage , err.Message())
	assert.Equal(t, http.StatusBadRequest, err.Status())
}

func TestSendCodeReasonNotZeroOrOne(t *testing.T) {
	reason := 2
	err := CodeService.Send(registerRequest.Phone , reason)
	assert.NotNil(t, err)
	assert.Equal(t, errors.InternalServerErrorMessage , err.Message())
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestFailToSendVerificationCode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodPost , fmt.Sprintf("https://api.kavenegar.com/v1/%s/sms/send.json" , os.Getenv("KAVENEGAR_API_CODE")),
		httpmock.NewStringResponder(http.StatusInternalServerError,"{}"))

	getUserFunc = func(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID: uint(1),
			Phone: registerRequest.Phone,
			Active: false,
		} , nil
	}


	err := CodeService.Send(registerRequest.Phone,0)

	assert.NotNil(t, err)
	assert.Equal(t, errors.InternalServerErrorMessage , err.Message())
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestSendVerificationCodeSuccessfully(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodPost , fmt.Sprintf("https://api.kavenegar.com/v1/%s/sms/send.json" , os.Getenv("KAVENEGAR_API_CODE")),
		httpmock.NewStringResponder(http.StatusOK,"{\"return\":{\"status\":200,\"message\":\"تایید شد\"},\"entries\":[{\"messageid\":1673299043,\"message\":\"salam this is test\",\"status\":5,\"statustext\":\"ارسال به مخابرات\",\"sender\":\"1000596446\",\"receptor\":\"09211231602\",\"date\":1627901748,\"cost\":570}]}"))

	getUserFunc = func(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID: uint(1),
			Phone: registerRequest.Phone,
			Active: false,
		} , nil
	}

	err := CodeService.Send(registerRequest.Phone,0)
	assert.Nil(t, err)
}

func TestFailToSendForgetPasswordCode(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodPost , fmt.Sprintf("https://api.kavenegar.com/v1/%s/sms/send.json" , os.Getenv("KAVENEGAR_API_CODE")),
		httpmock.NewStringResponder(http.StatusInternalServerError,"{}"))

	getUserFunc = func(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID: uint(1),
			Phone: registerRequest.Phone,
			Active: false,
		} , nil
	}


	err := CodeService.Send(registerRequest.Phone,1)

	assert.NotNil(t, err)
	assert.Equal(t, errors.InternalServerErrorMessage , err.Message())
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}


func TestSendForgetPasswordCodeSuccessfully(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodPost , fmt.Sprintf("https://api.kavenegar.com/v1/%s/sms/send.json" , os.Getenv("KAVENEGAR_API_CODE")),
		httpmock.NewStringResponder(http.StatusOK,"{\"return\":{\"status\":200,\"message\":\"تایید شد\"},\"entries\":[{\"messageid\":1673299043,\"message\":\"salam this is test\",\"status\":5,\"statustext\":\"ارسال به مخابرات\",\"sender\":\"1000596446\",\"receptor\":\"09211231602\",\"date\":1627901748,\"cost\":570}]}"))

	getUserFunc = func(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID: uint(1),
			Phone: registerRequest.Phone,
			Active: false,
		} , nil
	}

	err := CodeService.Send(registerRequest.Phone,1)
	assert.Nil(t, err)
}


func TestVerifyCodeFailToGetDataFromRepo(t *testing.T) {
	findCodeFunc = func(phone string , code , reason int) (*domains.Code, rest_errors.RestErr) {
		return nil , rest_errors.NewInternalServerError(errors.InternalServerErrorMessage,nil)
	}
	repositories.CodeRepository = &CodeRepoMock{}

	ok , err := CodeService.Verify(registerRequest.Phone ,23233, 0)
	assert.NotNil(t, err)
	assert.Equal(t, false , ok)
	assert.Equal(t, errors.InternalServerErrorMessage , err.Message())
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestVerifyCodeNotFound(t *testing.T) {
	findCodeFunc = func(phone string , code , reason int) (*domains.Code, rest_errors.RestErr) {
		return nil , nil
	}
	repositories.CodeRepository = &CodeRepoMock{}

	ok , err := CodeService.Verify(registerRequest.Phone ,23233, 0)
	assert.NotNil(t, err)
	assert.Equal(t, false , ok)
	assert.Equal(t, errors.UserNotFoundError , err.Message())
	assert.Equal(t, http.NotFound, err.Status())
}

func TestVerifyCodeSuccessfully(t *testing.T) {
	findCodeFunc = func(phone string , code , reason int) (*domains.Code, rest_errors.RestErr) {
		return &domains.Code{
			Code: 2313123,
			Phone: "09231212",
			CodeExpiration: time.Unix(0, time.Now().UnixNano() + 10000),
			CodePurpose: 0,
		} , nil
	}
	repositories.CodeRepository = &CodeRepoMock{}

	ok , err := CodeService.Verify(registerRequest.Phone ,23233, 0)
	assert.Nil(t, err)
	assert.Equal(t, true , ok)
}


func TestIsExpired(t *testing.T) {
	e := time.Unix(0, time.Now().UnixNano() - 1000)
	expired := IsExpired(e)
	assert.Equal(t, expired, true)
}

func TestIsNotExpired(t *testing.T) {
	e := time.Unix(0, time.Now().UnixNano() + 10000)
	expired := IsExpired(e)
	assert.Equal(t, expired, false)
}

func TestVerificationCodeIsExpired(t *testing.T) {
	findCodeFunc = func(phone string , code , reason int) (*domains.Code, rest_errors.RestErr) {
		return &domains.Code{
			Code: 2313123,
			Phone: "09231212",
			CodeExpiration: time.Unix(0, time.Now().UnixNano() - 10000),
			CodePurpose: 0,
		} , nil
	}
	repositories.CodeRepository = &CodeRepoMock{}

	ok , err := CodeService.Verify(registerRequest.Phone ,23233, 0)

	assert.Nil(t, err)
	assert.Equal(t, false , ok)
	assert.Equal(t, errors.CodeIsExpiredErrorMessage , err.Message())
	assert.Equal(t, http.StatusBadRequest, err.Status())
}