package v1

import "github.com/alidevjimmy/go-rest-utils/rest_errors"

var (
	CodeService codeServiceInterface = &codeService{}
)

type codeServiceInterface interface {
	SendCode(phone string) rest_errors.RestErr
	VerifyCode(phone string, code string) rest_errors.RestErr
}

type codeService struct{}

func (*codeService) SendCode(phone string) rest_errors.RestErr {
	return nil
}

func (*codeService) VerifyCode(phone string,code string) rest_errors.RestErr {
	return nil
}