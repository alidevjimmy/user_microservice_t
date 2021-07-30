package v1

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1/users"
)

// const (
// 	selectUserByID = "SELECT * FROM users WHERE"
// 	selectUserByPhoneOrUsername = ""
// 	selectUserByPhoneAndUsername = ""
// )

func GetUser(token string) (*users.User, rest_errors.RestErr) {
	return nil, nil
}
