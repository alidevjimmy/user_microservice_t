package services

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"math/rand"
	"time"
)

var (
	CodeService codeServiceInterface = &codeService{}
)

const (
	VERIFICATION = iota
	RESETPASSWORD
)

type codeServiceInterface interface {
	Send(phone string, reason int) rest_errors.RestErr
		Verify(phone string, code, reason int) (bool, rest_errors.RestErr)
}
type codeService struct{}

func (*codeService) Send(phone string, reason int) rest_errors.RestErr {
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
