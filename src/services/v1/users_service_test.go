package services

import (
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1"
	"github.com/alidevjimmy/user_microservice_t/errors/v1"
	repositories "github.com/alidevjimmy/user_microservice_t/repositories/postgres/v1"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)

var (
	RegisterRequest domains.RegisterRequest = domains.RegisterRequest{
		Phone:    "09122334344",
		Username: "test_user",
		Name:     "ali",
		Family:   "hamrani",
		Age:      uint(20),
		Password: "password",
	}
	loginRequest domains.LoginRequest = domains.LoginRequest{
		PhoneOrUsername: "09122334344",
		Password:        "password",
	}
	sendCodeFunc                            func(phone string) rest_errors.RestErr
	generateJwtFunc                         func(data jwt.MapClaims) (string, rest_errors.RestErr)
	verifyJwtFunc                           func(token string) (*domains.Jwt, bool, rest_errors.RestErr)
	getUserFunc                             func(id uint) (*domains.PublicUser, rest_errors.RestErr)
	getUsersFunc                            func(params map[string]interface{}) ([]domains.PublicUser, rest_errors.RestErr)
	updateUserActiveStateByIdFunc           func(userId uint) (*domains.PublicUser, rest_errors.RestErr)
	updateUserActiveStateByPhoneFunc        func(phone string) (*domains.PublicUser, rest_errors.RestErr)
	updateUserBlockStateFunc                func(userId uint) (*domains.PublicUser, rest_errors.RestErr)
	updateUserFunc                          func(userId uint, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr)
	verifyCodeFunc                          func(phone string, code, reason int) (bool, rest_errors.RestErr)
	updatePasswordByPhoneFunc               func(newPass, phone string) (*domains.PublicUser, rest_errors.RestErr)
	getUserByPhoneFunc                      func(phone string) (*domains.PublicUser, rest_errors.RestErr)
	getUserByUsernameFunc                   func(username string) (*domains.PublicUser, rest_errors.RestErr)
	getUserByPhoneOrUsernameAndPasswordFunc func(pou, password string) (*domains.PublicUser, rest_errors.RestErr)
)

type Suite struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

type JwtServiceMock struct{}

func (*JwtServiceMock) GenerateJwtToken(data jwt.MapClaims) (string, rest_errors.RestErr) {
	return generateJwtFunc(data)
}

func (*JwtServiceMock) VerifyJwtToken(token string) (*domains.Jwt, bool, rest_errors.RestErr) {
	return verifyJwtFunc(token)
}

type CodeServiceMock struct{}

func (*CodeServiceMock) Send(phone string, reason int) rest_errors.RestErr {
	return sendCodeFunc(phone)
}

func (*CodeServiceMock) Verify(phone string, code, reason int) (bool, rest_errors.RestErr) {
	return verifyCodeFunc(phone, code, reason)
}

type UserRespositoryMock struct {
	DB *gorm.DB
}

func (*UserRespositoryMock) GetUserByID(id uint) (*domains.PublicUser, rest_errors.RestErr) {
	return getUserFunc(id)
}

func (*UserRespositoryMock) GetUserByPhone(phone string) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (*UserRespositoryMock) GetUserByUsername(username string) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (*UserRespositoryMock) GetUserByPhoneOrUsernameAndPassword(pou, password string) (*domains.User, rest_errors.RestErr) {
	return nil, nil
}

func (u *UserRespositoryMock) GetUsers(params map[string]interface{}) ([]domains.PublicUser, rest_errors.RestErr) {
	return getUsersFunc(params)
}

func (u *UserRespositoryMock) UpdateBlockState(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
	return updateUserBlockStateFunc(userId)
}

func (u *UserRespositoryMock) UpdateUser(userId uint, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr) {
	return updateUserFunc(userId, body)
}

func (u *UserRespositoryMock) ChangeForgotPassword(newPassword, phone string, code int) (*domains.PublicUser, rest_errors.RestErr) {
	return nil, nil
}
func (u *UserRespositoryMock) UpdatePasswordByPhone(newPass, phone string) (*domains.PublicUser, rest_errors.RestErr) {
	return updatePasswordByPhoneFunc(newPass, phone)
}

func (u *UserRespositoryMock) UpdateActiveStateById(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
	return updateUserActiveStateByIdFunc(userId)
}
func (u *UserRespositoryMock) UpdateActiveStateByPhone(phone string) (*domains.PublicUser, rest_errors.RestErr) {
	return updateUserActiveStateByPhoneFunc(phone)
}

func TestRegisterCanInsertDuplicatedPhone(t *testing.T) {
	getUserByPhoneFunc = func(phone string) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID:    uint(1),
			Phone: phone,
		}, nil
	}

	repositories.UserRepository = &UserRespositoryMock{}

	rr, err := UserService.Register(RegisterRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status())
	assert.Equal(t, errors.DuplicatePhoneErrorMessage, err.Message())
}

func TestRegisterCanInsertDuplicatedUsername(t *testing.T) {
	getUserByUsernameFunc = func(username string) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID:       uint(1),
			Username: username,
		}, nil
	}

	repositories.UserRepository = &UserRespositoryMock{}

	rr, error := UserService.Register(RegisterRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, error)
	assert.Equal(t, http.StatusBadRequest, error.Status())
	assert.Equal(t, errors.DuplicateUsernameErrorMessage, error.Message())
}

func TestRegisterFailToSendVerificationCode(t *testing.T) {
	sendCodeFunc = func(phone string) rest_errors.RestErr {
		return rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	CodeService = &CodeServiceMock{}

	rr, err := UserService.Register(RegisterRequest)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
	assert.Nil(t, rr)
}

func TestRegisterFailToGenerateJwtToken(t *testing.T) {
	// mock jwt service
	generateJwtFunc = func(data jwt.MapClaims) (string, rest_errors.RestErr) {
		return "", rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	JwtService = &JwtServiceMock{}

	rr, err := UserService.Register(RegisterRequest)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
	assert.Nil(t, rr)
}

func TestRegisterUsernameOnlyCanContainUnderlineAndEnglishWordsAndNumbers(t *testing.T) {
	RegisterRequest.Username = "sdff-dfd"
	rr, err := UserService.Register(RegisterRequest)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status())
	assert.Equal(t, errors.UsernameOnlyCanContainUnderlineAndEnglishWordsAndNumbersErrorMessage, err.Message())
	assert.Nil(t, rr)
}

// login tests

func TestLoginSeccessfully(t *testing.T) {
	getUserByPhoneOrUsernameAndPasswordFunc = func(pou, password string) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID:    uint(1),
			Phone: pou,
		}, nil
	}

	repositories.UserRepository = &UserRespositoryMock{}

	lr, err := UserService.Login(loginRequest)
	assert.NotNil(t, lr)
	assert.Nil(t, err)
}

func TestLoginPhoneOrUsernameRequired(t *testing.T) {
	loginRequest.PhoneOrUsername = ""
	rr, err := UserService.Login(loginRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status())
	assert.Equal(t, errors.PhoneOrUsernameIsRequiredErrorMessage, err.Message())
}

func TestLoginPasswordRequired(t *testing.T) {
	loginRequest.Password = ""
	rr, err := UserService.Login(loginRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status())
	assert.Equal(t, errors.PasswordIsRequiredErrorMessage, err.Message())
}

func TestLoginWithPhone(t *testing.T) {
	getUserByPhoneOrUsernameAndPasswordFunc = func(pou, password string) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID:    uint(1),
			Phone: pou,
		}, nil
	}

	repositories.UserRepository = &UserRespositoryMock{}
	lr, err := UserService.Login(loginRequest)
	assert.NotNil(t, lr)
	assert.Nil(t, err)
}

func TestLoginWithUsername(t *testing.T) {
	getUserByPhoneOrUsernameAndPasswordFunc = func(pou, password string) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID:    uint(1),
			Phone: pou,
		}, nil
	}

	loginRequest.Password = RegisterRequest.Password
	loginRequest.PhoneOrUsername = RegisterRequest.Username
	lr, err := UserService.Login(loginRequest)
	assert.NotNil(t, lr)
	assert.Nil(t, err)
}

func TestLoginFailToGenerateJwtToken(t *testing.T) {
	// mock jwt service
	generateJwtFunc = func(data jwt.MapClaims) (string, rest_errors.RestErr) {
		return "", rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	JwtService = &JwtServiceMock{}

	rr, err := UserService.Register(RegisterRequest)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
	assert.Nil(t, rr)
}

func TestGetUserFailToVerifyToken(t *testing.T) {
	token := "oiadshs23kj3h123j32.23kjhkjehdkjh.23kjhk2nbemnwbd"
	gur, err := UserService.GetUser(token)
	assert.NotNil(t, err)
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
	assert.Equal(t, http.StatusUnauthorized, err.Status())
	assert.Nil(t, gur)
}

func TestGetUserFailToGetDataFromRepository(t *testing.T) {
	verifyJwtFunc = func(token string) (*domains.Jwt, bool, rest_errors.RestErr) {
		return &domains.Jwt{
			Sub: "1",
			Exp: time.Now().Add(time.Hour),
		}, true, nil
	}
	JwtService = &JwtServiceMock{}

	getUserFunc = func(id uint) (*domains.PublicUser, rest_errors.RestErr) {
		return nil, rest_errors.NewNotFoundError(errors.InternalServerErrorMessage)
	}

	repositories.UserRepository = &UserRespositoryMock{}

	gu, err := UserService.GetUser("some token")
	assert.NotNil(t, err)
	assert.Nil(t, gu)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
}

func TestGetUserNotFound(t *testing.T) {
	verifyJwtFunc = func(token string) (*domains.Jwt, bool, rest_errors.RestErr) {
		return &domains.Jwt{
			Sub: "1",
			Exp: time.Now().Add(time.Hour),
		}, true, nil
	}
	JwtService = &JwtServiceMock{}

	getUserFunc = func(id uint) (*domains.PublicUser, rest_errors.RestErr) {
		return nil, nil
	}

	repositories.UserRepository = &UserRespositoryMock{}

	gu, err := UserService.GetUser("some token")
	assert.NotNil(t, err)
	assert.Nil(t, gu)
	assert.Equal(t, http.NotFound, err.Status())
	assert.Equal(t, errors.UserNotFoundError, err.Message())
}

func TestGetUserSuccessfully(t *testing.T) {
	verifyJwtFunc = func(token string) (*domains.Jwt, bool, rest_errors.RestErr) {
		return &domains.Jwt{
			Sub: "1",
			Exp: time.Now().Add(time.Hour),
		}, true, nil
	}

	JwtService = &JwtServiceMock{}

	getUserFunc = func(id uint) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID:       uint(1),
			Phone:    RegisterRequest.Phone,
			Username: RegisterRequest.Username,
			Age:      RegisterRequest.Age,
		}, nil
	}

	repositories.UserRepository = &UserRespositoryMock{}

	pu, err := UserService.GetUser("some token")
	assert.NotNil(t, pu)
	assert.Nil(t, err)
}

func TestFailToGetUsersFromRepository(t *testing.T) {
	getUsersFunc = func(params map[string]interface{}) ([]domains.PublicUser, rest_errors.RestErr) {
		return []domains.PublicUser{}, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	repositories.UserRepository = &UserRespositoryMock{}

	gu, err := UserService.GetUsers(map[string]interface{}{"test": "pass"})
	assert.NotNil(t, err)
	assert.Nil(t, gu)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
}

func TestGetUsersSuccessfully(t *testing.T) {
	getUsersFunc = func(params map[string]interface{}) ([]domains.PublicUser, rest_errors.RestErr) {
		return []domains.PublicUser{}, nil
	}

	repositories.UserRepository = &UserRespositoryMock{}

	gu, err := UserService.GetUsers(map[string]interface{}{"test": "pass"})
	assert.NotNil(t, gu)
	assert.Nil(t, err)
}

func TestFailToUpdateUserActiveState(t *testing.T) {
	updateUserActiveStateByIdFunc = func(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	repositories.UserRepository = &UserRespositoryMock{}

	gu, err := UserService.UpdateUserActiveState(uint(1))

	assert.NotNil(t, err)
	assert.Nil(t, gu)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
}

func TestSuccessfullyUpdateUserActiveState(t *testing.T) {
	updateUserActiveStateByIdFunc = func(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID: 1,
		}, nil
	}

	repositories.UserRepository = &UserRespositoryMock{}

	u, err := UserService.UpdateUserActiveState(uint(1))

	assert.NotNil(t, u)
	assert.Nil(t, err)
}

func TestFailToUpdateUserBlockState(t *testing.T) {
	updateUserBlockStateFunc = func(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	repositories.UserRepository = &UserRespositoryMock{}

	gu, err := UserService.UpdateUserBlockState(uint(1))

	assert.NotNil(t, err)
	assert.Nil(t, gu)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
}

func TestSuccessfullyUpdateUserBlockState(t *testing.T) {
	updateUserBlockStateFunc = func(userId uint) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID: 1,
		}, nil
	}

	repositories.UserRepository = &UserRespositoryMock{}

	u, err := UserService.UpdateUserBlockState(uint(1))

	assert.NotNil(t, u)
	assert.Nil(t, err)
}

func TestFailToUpdateUser(t *testing.T) {
	updateUserFunc = func(userId uint, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	repositories.UserRepository = &UserRespositoryMock{}
	body := domains.UpdateUserRequest{
		Username: RegisterRequest.Username,
	}
	gu, err := UserService.UpdateUser("", body)

	assert.NotNil(t, err)
	assert.Nil(t, gu)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
}

func TestSuccessfullyUpdateUser(t *testing.T) {
	updateUserFunc = func(userId uint, body domains.UpdateUserRequest) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{
			ID: 1,
		}, nil
	}

	repositories.UserRepository = &UserRespositoryMock{}

	body := domains.UpdateUserRequest{
		Username: RegisterRequest.Username,
	}

	u, err := UserService.UpdateUser("", body)

	assert.NotNil(t, u)
	assert.Nil(t, err)
}

func TestChangePasswordFailToVerifyCode(t *testing.T) {
	verifyCodeFunc = func(phone string, code, reason int) (bool, rest_errors.RestErr) {
		return false, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	CodeService = &CodeServiceMock{}

	u, err := UserService.ChangeForgotPassword("new pass", RegisterRequest.Phone, 123123)
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
}

func TestChangePasswordInvalidCode(t *testing.T) {
	verifyCodeFunc = func(phone string, code, reason int) (bool, rest_errors.RestErr) {
		return false, nil
	}

	CodeService = &CodeServiceMock{}

	u, err := UserService.ChangeForgotPassword("new pass", RegisterRequest.Phone, 123123)
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.StatusNotFound, err.Status())
	assert.Equal(t, errors.CodeOrPhoneDoesNotExistsErrorMessage, err.Message())
}

func TestChangePasswordFailToUpdatePassword(t *testing.T) {
	verifyCodeFunc = func(phone string, code, reason int) (bool, rest_errors.RestErr) {
		return true, nil
	}

	CodeService = &CodeServiceMock{}

	updatePasswordByPhoneFunc = func(newPass, phone string) (*domains.PublicUser, rest_errors.RestErr) {
		return nil, rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}
	repositories.UserRepository = &UserRespositoryMock{}
	u, err := UserService.ChangeForgotPassword("new pass", RegisterRequest.Phone, 123123)
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.StatusNotFound, err.Status())
	assert.Equal(t, errors.CodeOrPhoneDoesNotExistsErrorMessage, err.Message())
}

func TestChangePasswordSuccessfully(t *testing.T) {
	verifyCodeFunc = func(phone string, code, reason int) (bool, rest_errors.RestErr) {
		return true, nil
	}

	CodeService = &CodeServiceMock{}

	updatePasswordByPhoneFunc = func(newPass, phone string) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{}, nil
	}

	repositories.UserRepository = &UserRespositoryMock{}

	u, err := UserService.ChangeForgotPassword("new pass", RegisterRequest.Phone, 123123)
	assert.Nil(t, err)
	assert.NotNil(t, u)
}

func TestActiveUserInvalidCode(t *testing.T) {
	verifyCodeFunc = func(phone string, code, reason int) (bool, rest_errors.RestErr) {
		return false, nil
	}

	CodeService = &CodeServiceMock{}

	u, err := UserService.ActiveUser(RegisterRequest.Phone, 123123)
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.StatusNotFound, err.Status())
	assert.Equal(t, errors.CodeOrPhoneDoesNotExistsErrorMessage, err.Message())
}

func TestActiveUserFailToUpdate(t *testing.T) {
	verifyCodeFunc = func(phone string, code, reason int) (bool, rest_errors.RestErr) {
		return true, nil
	}

	CodeService = &CodeServiceMock{}

	updateUserActiveStateByPhoneFunc = func(phone string) (*domains.PublicUser, rest_errors.RestErr) {
		return nil, nil
	}
	repositories.UserRepository = &UserRespositoryMock{}
	u, err := UserService.ActiveUser(RegisterRequest.Phone, 123123)
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.StatusNotFound, err.Status())
	assert.Equal(t, errors.CodeOrPhoneDoesNotExistsErrorMessage, err.Message())
}

func TestActiveUserSuccessfully(t *testing.T) {
	verifyCodeFunc = func(phone string, code, reason int) (bool, rest_errors.RestErr) {
		return true, nil
	}

	CodeService = &CodeServiceMock{}

	updateUserActiveStateByPhoneFunc = func(phone string) (*domains.PublicUser, rest_errors.RestErr) {
		return &domains.PublicUser{}, nil
	}

	repositories.UserRepository = &UserRespositoryMock{}

	u, err := UserService.ActiveUser(RegisterRequest.Phone, 123123)
	assert.Nil(t, err)
	assert.NotNil(t, u)
}
