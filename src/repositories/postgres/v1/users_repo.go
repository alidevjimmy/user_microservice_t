package repositories

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1"
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
	GetUserByID(id uint) (*domains.PublicUser, rest_errors.RestErr)
	GetUserByPhone(phone string) (*domains.PublicUser, rest_errors.RestErr)
	GetUserByUsername(username string) (*domains.PublicUser, rest_errors.RestErr)
	GetUserByPhoneOrUsernameAndPassword(pou, password string) (*domains.User, rest_errors.RestErr)
	UpdateUser(userId uint, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr)
	UpdatePasswordByPhone(newPass , phone string) (*domains.PublicUser, rest_errors.RestErr)
	UpdateActiveStateByPhone(phone string) (*domains.PublicUser, rest_errors.RestErr)
	UpdateActiveStateById(userId uint) (*domains.PublicUser, rest_errors.RestErr)
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{DB: db}
}


func (*userRepository) GetUserByID(id uint) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (*userRepository) GetUserByPhone(phone string) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (*userRepository) GetUserByUsername(username string) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (*userRepository) GetUserByPhoneOrUsernameAndPassword(pou, password string) (*domains.User, rest_errors.RestErr) {
	return nil, nil
}
func (u *userRepository) GetUsers(params map[string]interface{}) ([]domains.PublicUser, rest_errors.RestErr) {
	return []domains.PublicUser{}, nil
}

func (u *userRepository) UpdateActiveStateById(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}
func (u *userRepository) UpdateActiveStateByPhone(phone string) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (u *userRepository) UpdateBlockState(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (u *userRepository) UpdateUser(userId uint, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (u *userRepository) UpdatePasswordByPhone(newPass , phone string) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}


