package services

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/golang-jwt/jwt"
)


type jwtService struct {}

type jwtInterface interface {
	Generate(data jwt.MapClaims) (string , rest_errors.RestErr)
	Verify(token string) (*jwt.MapClaims, bool , rest_errors.RestErr)
}

var (
	JwtService jwtInterface = &jwtService{}
)

func(*jwtService) Generate(data jwt.MapClaims) (string, rest_errors.RestErr) {
	return "" , nil
}

func(*jwtService) Verify(token string) (*jwt.MapClaims, bool, rest_errors.RestErr) {
	return nil,false, nil
}