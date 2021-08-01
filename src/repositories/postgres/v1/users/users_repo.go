package users

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1/users"
	"gorm.io/gorm"
)

const (
	selectUserByID                         = "SELECT %s FROM %s WHERE id=%s"
	selectUserByPhone                      = "SELECT %s from %s WHERE phone=%s"
	selectUserByUsername                   = "SELECT %s from %s WHERE username=%s"
	selectUserByPhoneOrUsernameAndPassword = "SELECT %s from %s WHERE %s=%s AND password=%s"
)

var (
	UserRepository userRepositoryInterface = &userRepository{}
)

type userRepository struct {
	DB *gorm.DB
}

type userRepositoryInterface interface {
	GetUserByID(id uint) (*users.PublicUser, rest_errors.RestErr)
	GetUserByPhone(phone string) (*users.PublicUser, rest_errors.RestErr)
	GetUserByUsername(username string) (*users.PublicUser, rest_errors.RestErr)
	GetUserByPhoneOrUsernameAndPassword(pou , password string) (*users.User, rest_errors.RestErr)
	New(db *gorm.DB) userRepositoryInterface
}

func (u *userRepository) New(db *gorm.DB) userRepositoryInterface{
	u.DB = db
	var rp userRepositoryInterface = &userRepository{DB: db}
	return rp
}

// func GetUsers(// map for filter res) ([]users.User, rest_errors.RestErr) {
// 	var users []users.User
// 	// err :=
// 	return users,nil
// }

func (*userRepository) GetUserByID(id uint) (*users.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (*userRepository) GetUserByPhone(phone string) (*users.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (*userRepository) GetUserByUsername(username string) (*users.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (*userRepository) GetUserByPhoneOrUsernameAndPassword(pou , password string) (*users.User, rest_errors.RestErr) {
	return nil, nil
}