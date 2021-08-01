package services

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1"
)

var (
	UserService userServiceInterface = &userSerive{}
)

type userServiceInterface interface {
	Register(body domains.RegisterRequest) (*domains.RegisterResponse, rest_errors.RestErr)
	Login(body domains.LoginRequest) (*domains.LoginResponse, rest_errors.RestErr)
	GetUser(token string) (*domains.PublicUser, rest_errors.RestErr)
	GetUsers(params map[string]interface{}) ([]domains.PublicUser, rest_errors.RestErr)
	UpdateUserActiveState(userId uint) (*domains.PublicUser, rest_errors.RestErr)
	UpdateUserBlockState(userId uint) (*domains.PublicUser, rest_errors.RestErr)
	UpdateUser(userId uint, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr)
}

type userSerive struct{}

func (*userSerive) Register(body domains.RegisterRequest) (*domains.RegisterResponse, rest_errors.RestErr) {
	return nil, nil
}

func (*userSerive) Login(body domains.LoginRequest) (*domains.LoginResponse, rest_errors.RestErr) {
	return nil, nil
}

// GetUser returns single user by its jwt token
func (*userSerive) GetUser(token string) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

// GetUsers returns all users by filter
func (*userSerive) GetUsers(params map[string]interface{}) ([]domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

// UpdateUserActiveState makes state of active field of user opposite
func (*userSerive) UpdateUserActiveState(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

// UpdateUserBlockState makes state of blocked field of user opposite
func (*userSerive) UpdateUserBlockState(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (*userSerive) UpdateUser(userId uint, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}
