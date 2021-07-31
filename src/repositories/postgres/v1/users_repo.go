package repositries

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1/users"
	"gorm.io/gorm"
)

const (
	selectUserByID                         = "SELECT %s FROM %s WHERE id=%s"
	selectUserByPhone                      = "SELEC %s from %s WHERE phone=%s"
	selectUserByUsername                   = "SELEC %s from %s WHERE username=%s"
	selectUserByPhoneOrUsernameAndPassword = "SELECT %s from %s WHERE %s=%s AND password=%s"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

// func GetUsers(// map for filter res) ([]users.User, rest_errors.RestErr) {
// 	var users []users.User
// 	// err :=
// 	return users,nil
// }

func (u *UserRepository) GetUserByID(id uint) (*users.User, rest_errors.RestErr) {
	return nil, nil
}

func (u *UserRepository) GetUserByPhone(phone string) (*users.User, rest_errors.RestErr) {
	return nil, nil
}

func (u *UserRepository) GetUserByUsername(username string) (*users.User, rest_errors.RestErr) {
	return nil, nil
}

func (u *UserRepository) GetUserByPhoneOrUsernameAndPassword(pou , password string) (*users.User, rest_errors.RestErr) {
	return nil, nil
}

