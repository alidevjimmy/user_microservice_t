package jwt

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	jwt2 "github.com/alidevjimmy/user_microservice_t/domains/v1/jwt"
	"github.com/golang-jwt/jwt"
)


type jwtService struct {}

type jwtInterface interface {
	Generate(data jwt.MapClaims) (string , rest_errors.RestErr)
	Verify(token string) (*jwt2.Jwt, bool , rest_errors.RestErr)
}

var (
	JwtService jwtInterface = &jwtService{}
)

func(*jwtService) Generate(data jwt.MapClaims) (string, rest_errors.RestErr) {
	return "" , nil
}

func(*jwtService) Verify(token string) (*jwt2.Jwt, bool, rest_errors.RestErr) {
	return nil,false, nil
}