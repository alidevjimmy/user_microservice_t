package services

import (
	"math/rand"
	"time"

	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1"
)

var (
	CodeService codeServiceInterface = &codeService{}
)

const (
	VERIFICATION = iota
	RESETPASSWORD
)

type codeServiceInterface interface {
	Send(body domains.SendCodeRequest) rest_errors.RestErr
	Verify(phone string, code, reason int) (bool, rest_errors.RestErr)
}
type codeService struct{}

func (*codeService) Send(body domains.SendCodeRequest) rest_errors.RestErr {
	return nil
}

func (*codeService) Verify(phone string, code, reason int) (bool, rest_errors.RestErr) {
	return false, nil
}

func RandomCodeGenerator() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(99999-10000+1) + 10000
}

func IsExpired(exp time.Time) bool {
	return time.Now().UnixNano() > exp.UnixNano()
}
