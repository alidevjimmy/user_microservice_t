package v1

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1/users"
	"gorm.io/gorm"
)

// const (
// 	selectUserByID = "SELECT * FROM users WHERE"
// 	selectUserByPhoneOrUsername = ""
// 	selectUserByPhoneAndUsername = ""
// )

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB : DB}
}

func (u *UserRepository) FindAll() ([]users.User, rest_errors.RestErr) {
	var users []users.User
	// err :=
	return users,nil
}

// func GetUser(token string) (*users.User, rest_errors.RestErr) {
// 	return nil, nil
// }
