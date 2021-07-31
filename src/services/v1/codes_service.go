package services

import "github.com/alidevjimmy/go-rest-utils/rest_errors"

var (
	CodeService codeServiceInterface = &codeService{}
)


type codeServiceInterface interface {
	Send(phone string) rest_errors.RestErr
	Verify(phone string, code int) rest_errors.RestErr
}

type codeService struct{}

func (*codeService) Send(phone string) rest_errors.RestErr {
	return nil
}

func (*codeService) Verify(phone string, code int) rest_errors.RestErr {
	return nil
}
