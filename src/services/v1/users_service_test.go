package v1

import (
	"github.com/alidevjimmy/user_microservice_t/domains/v1/users"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	registerRequest users.RegisterRequest = users.RegisterRequest{
		Phone:    "09122334344",
		Username: "testuser",
		Name:     "ali",
		Family:   "hamrani",
		Age:      uint(20),
		Password: "password",
	}
)

func TestRegisterPessesValidated(t *testing.T) {
	_, r := UserService.Register(registerRequest)
	assert.Nil(t, r)
}

func TestRegisterUsernameRequired(t *testing.T) {
	registerRequest.Username = ""
	_, r := UserService.Register(registerRequest)
	assert.NotNil(t, r)
	assert.Equal(t, "username is required", r.Error())
}

func TestRegisterNameRequired(t *testing.T) {
	registerRequest.Name = ""
	_, r := UserService.Register(registerRequest)
	assert.NotNil(t, r)
	assert.Equal(t, "name is required", r.Error())
}

func TestRegisterFamilyRequired(t *testing.T) {
	registerRequest.Family = ""
	_, r := UserService.Register(registerRequest)
	assert.NotNil(t, r)
	assert.Equal(t, "family is required", r.Error())
}
func TestRegisterAgeRequired(t *testing.T) {
	registerRequest.Age = 0
	_, r := UserService.Register(registerRequest)
	assert.NotNil(t, r)
	assert.Equal(t, "age is required", r.Error())
}
func TestRegisterPasswordRequired(t *testing.T) {
	registerRequest.Password = ""
	_, r := UserService.Register(registerRequest)
	assert.NotNil(t, r)
	assert.Equal(t, "password is required", r.Error())
}

func TestRegisterCanInsertDuplicatedPhone(t *testing.T) {
	// TODO: database mock
}

func TestRegisterCanInsertDuplicatedUsername(t *testing.T) {
	// TODO: database mock
}
