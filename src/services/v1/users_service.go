package services

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1/users"
)

var (
	UserService userServiceInterface = &userSerive{}
)


type userServiceInterface interface {
	Register(body users.RegisterRequest) (*users.RegisterResponse, rest_errors.RestErr)
	Login(body users.LoginRequest) (*users.LoginResponse, rest_errors.RestErr)
	GetUser(token string) (*users.User, rest_errors.RestErr)
	GetUsers(params interface{}) ([]users.User, rest_errors.RestErr)
	ToggleActiveUser(userId int) rest_errors.RestErr
	ToggleBlockUser(userId int) rest_errors.RestErr
	EditUser(body users.EditUserRequest) rest_errors.RestErr
}

type userSerive struct{}

func (*userSerive) Register(body users.RegisterRequest) (*users.RegisterResponse, rest_errors.RestErr) {
	// check username only can contain _ and english words
	return nil, nil
}

func (*userSerive) Login(body users.LoginRequest) (*users.LoginResponse, rest_errors.RestErr) {
	return nil, nil
}

// GetUser returns single user by its jwt token
func (*userSerive) GetUser(token string) (*users.User, rest_errors.RestErr) {
	return nil, nil
}

// GetUsers returns all users by filter
func (*userSerive) GetUsers(params interface{}) ([]users.User, rest_errors.RestErr) {
	return nil, nil
}

// ToggleActiveUser makes state of active field of user opposite
func (*userSerive) ToggleActiveUser(userId int) rest_errors.RestErr {
	return nil
}

// ToggleBlockUser makes state of blocked field of user opposite
func (*userSerive) ToggleBlockUser(userId int) rest_errors.RestErr {
	return nil
}

func (*userSerive) EditUser(body users.EditUserRequest) rest_errors.RestErr {
	return nil
}
