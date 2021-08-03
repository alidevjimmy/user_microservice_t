package services

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1"
)

var (
	UserService userServiceInterface = &userService{}
)

type userServiceInterface interface {
	Register(body domains.RegisterRequest) (*domains.RegisterResponse, rest_errors.RestErr)
	Login(body domains.LoginRequest) (*domains.LoginResponse, rest_errors.RestErr)
	GetUser(token string) (*domains.PublicUser, rest_errors.RestErr)
	GetUsers(params map[string]interface{}) ([]domains.PublicUser, rest_errors.RestErr)
	UpdateUserActiveState(userId uint) (*domains.PublicUser, rest_errors.RestErr)
	UpdateUserBlockState(userId uint) (*domains.PublicUser, rest_errors.RestErr)
	UpdateUser(token string, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr)
	ChangeForgotPassword(newPassword, phone string, code int) (*domains.PublicUser, rest_errors.RestErr)
	ActiveUser(phone string, code int) (*domains.PublicUser, rest_errors.RestErr)
}

type userService struct{}

func (*userService) Register(body domains.RegisterRequest) (*domains.RegisterResponse, rest_errors.RestErr) {
	return nil, nil
}

func (*userService) Login(body domains.LoginRequest) (*domains.LoginResponse, rest_errors.RestErr) {
	return nil, nil
}

// GetUser returns single user by its jwt token
func (*userService) GetUser(token string) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

// GetUsers returns all users by filter
func (*userService) GetUsers(params map[string]interface{}) ([]domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

// UpdateUserActiveState makes state of active field of user opposite
func (*userService) UpdateUserActiveState(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

// UpdateUserBlockState makes state of blocked field of user opposite
func (*userService) UpdateUserBlockState(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (*userService) UpdateUser(token string, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

// ChangeForgotPassword helps people who forgot their password using verification code
func (*userService) ChangeForgotPassword(newPassword, phone string, code int) (*domains.PublicUser, rest_errors.RestErr) {
	// return token
	return nil, nil
}

// ActiveUser Change user active state to true using verification code
func (*userService) ActiveUser(phone string, code int) (*domains.PublicUser, rest_errors.RestErr) {
	// return token
	return nil, nil
}
