package services

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1"
	"github.com/golang-jwt/jwt"
)

type jwtService struct{}

type jwtInterface interface {
	GenerateJwtToken(data jwt.MapClaims) (string, rest_errors.RestErr)
	VerifyJwtToken(token string) (*domains.Jwt, rest_errors.RestErr)
}

var (
	JwtService jwtInterface = &jwtService{}
)

func (*jwtService) GenerateJwtToken(data jwt.MapClaims) (string, rest_errors.RestErr) {

	return "", nil
}

func (*jwtService) VerifyJwtToken(token string) (*domains.Jwt, rest_errors.RestErr) {
	return nil, nil
}
