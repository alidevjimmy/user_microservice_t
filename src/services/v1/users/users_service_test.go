package users

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alidevjimmy/go-rest-utils/crypto"
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	jwt3 "github.com/alidevjimmy/user_microservice_t/domains/v1/jwt"
	users2 "github.com/alidevjimmy/user_microservice_t/domains/v1/users"
	"github.com/alidevjimmy/user_microservice_t/errors"
	"github.com/alidevjimmy/user_microservice_t/repositories/postgres/v1/users"
	"github.com/alidevjimmy/user_microservice_t/services/v1/codes"
	jwt2 "github.com/alidevjimmy/user_microservice_t/services/v1/jwt"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
)

var (
	registerRequest users2.RegisterRequest = users2.RegisterRequest{
		Phone:    "09122334344",
		Username: "test_user",
		Name:     "ali",
		Family:   "hamrani",
		Age:      uint(20),
		Password: "password",
	}
	loginRequest users2.LoginRequest = users2.LoginRequest{
		PhoneOrUsername:    "09122334344",
		Password: "password",
	}
	// for mock different states of CodeService methods
	sendCodeFunc func(phone string) rest_errors.RestErr
	// for mock different states of JwtService methods
	generateJwtFunc func(data jwt.MapClaims) (string, rest_errors.RestErr)
	verifyJwtFunc func(token string) (*jwt3.Jwt, bool , rest_errors.RestErr)
	getUserFunc func(id uint) (*users2.PublicUser, rest_errors.RestErr)
)

type Suite struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

type JwtServiceMock struct{}

func(*JwtServiceMock) Generate(data jwt.MapClaims) (string, rest_errors.RestErr) {
	return generateJwtFunc(data)
}

func(*JwtServiceMock) Verify(token string) (*jwt3.Jwt, bool, rest_errors.RestErr) {
	return verifyJwtFunc(token)
}

type CodeServiceMock struct{}

func (*CodeServiceMock) Send(phone string) rest_errors.RestErr {
	return sendCodeFunc(phone)
}

func (*CodeServiceMock) Verify(phone string, code int) rest_errors.RestErr {
	return nil
}

type 	UserRespositoryMock struct {
	DB *gorm.DB
}

func (u *UserRespositoryMock) SetDB(db *gorm.DB) {}

func (*UserRespositoryMock) GetUserByID(id uint) (*users2.PublicUser, rest_errors.RestErr) {
	return getUserFunc(id)
}

func (*UserRespositoryMock) GetUserByPhone(phone string) (*users2.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (*UserRespositoryMock) GetUserByUsername(username string) (*users2.PublicUser, rest_errors.RestErr) {
	return nil, nil
}

func (*UserRespositoryMock) GetUserByPhoneOrUsernameAndPassword(pou , password string) (*users2.User, rest_errors.RestErr) {
	return nil, nil
}

func (u *UserRespositoryMock) New(db *gorm.DB) *UserRespositoryMock{
	u.DB = db
	return &UserRespositoryMock{DB: db}
}

func MockDbConnection(t *testing.T) (*Suite){
	s := &Suite{}
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	s.db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}
	if s.db == nil {
		t.Error("gorm db is null")
	}
	return s
}

func TestRegisterPessesValidated(t *testing.T) {
	rr, r := UserService.Register(registerRequest)
	assert.Nil(t, rr)
	assert.Nil(t, r)
}

func TestRegisterPhoneRequired(t *testing.T) {
	registerRequest.Phone = ""
	rr, err := UserService.Register(registerRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, err)
	assert.Equal(t,http.StatusBadRequest , err.Status())
	assert.Equal(t,errors.PhoneIsRequiredErrorMessage, err.Message())
}

func TestRegisterUsernameRequired(t *testing.T) {
	registerRequest.Username = ""
	rr, err := UserService.Register(registerRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, err)
	assert.Equal(t,http.StatusBadRequest , err.Status())
	assert.Equal(t,errors.UsernameIsRequiredErrorMessage, err.Message())
}

func TestRegisterNameRequired(t *testing.T) {
	registerRequest.Name = ""
	rr, err := UserService.Register(registerRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, err)
	assert.Equal(t,http.StatusBadRequest , err.Status())
	assert.Equal(t,errors.NameIsRequiredErrorMessage, err.Message())
}

func TestRegisterFamilyRequired(t *testing.T) {
	registerRequest.Family = ""
	rr, err := UserService.Register(registerRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, err)
	assert.Equal(t,http.StatusBadRequest , err.Status())
	assert.Equal(t,errors.FamilyIsRequiredErrorMessage, err.Message())
}
func TestRegisterAgeRequired(t *testing.T) {
	registerRequest.Age = 0
	rr, err := UserService.Register(registerRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, err)
	assert.Equal(t,http.StatusBadRequest , err.Status())
	assert.Equal(t,errors.AgeIsRequiredErrorMessage, err.Message())
}
func TestRegisterPasswordRequired(t *testing.T) {
	registerRequest.Password = ""
	rr, err := UserService.Register(registerRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, err)
	assert.Equal(t,http.StatusBadRequest , err.Status())
	assert.Equal(t,errors.PasswordIsRequiredErrorMessage, err.Message())
}

func TestRegisterCanInsertDuplicatedPhone(t *testing.T) {
	s := MockDbConnection(t)

	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (phone,username,name,family,age,password) VALUES (?,?,?,?,?,?)").
		WithArgs(registerRequest.Phone, registerRequest.Username, registerRequest.Name, registerRequest.Family, registerRequest.Age, registerRequest.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	//Respo

	rr, err := UserService.Register(registerRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest , err.Status())
	assert.Equal(t, errors.DuplicatePhoneErrorMessage , err.Message())
}

func TestRegisterCanInsertDuplicatedUsername(t *testing.T) {
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (phone,username,name,family,age,password) VALUES (?,?,?,?,?,?)").
		WithArgs(registerRequest.Phone, registerRequest.Username, registerRequest.Name, registerRequest.Family, registerRequest.Age, registerRequest.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	registerRequest.Phone = "092837232"

	users.UserRepository = &UserRespositoryMock{DB: s.db}

	rr, error := UserService.Register(registerRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, error)
	assert.Equal(t, http.StatusBadRequest , error.Status())
	assert.Equal(t, errors.DuplicateUsernameErrorMessage , error.Message())
}

func TestRegisterFailToGetVerificationCode(t *testing.T) {
	// mock Send service method
	sendCodeFunc = func(phone string) rest_errors.RestErr {
		return rest_errors.NewInternalServerError("cannot send verification code", nil)
	}

	codes.CodeService = &CodeServiceMock{}

	rr, err := UserService.Register(registerRequest)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError , err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage , err.Message())
	assert.Nil(t, rr)
}

func TestRegisterFailToGenerateJwtToken(t *testing.T) {
	// mock jwt service
	generateJwtFunc = func(data jwt.MapClaims) (string, rest_errors.RestErr) {
		return "" , rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	jwt2.JwtService = &JwtServiceMock{}

	rr, err := UserService.Register(registerRequest)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError , err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage , err.Message())
	assert.Nil(t, rr)
}

func TestRegisterUsernameOnlyCanContainUnderlineAndEnglishWordsAndNumbers(t *testing.T) {
	registerRequest.Username = "sdff-dfd"
	rr, err := UserService.Register(registerRequest)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest , err.Status())
	assert.Equal(t, errors.UsernameOnlyCanContainUnderlineAndEnglishWordsAndNumbersErrorMessage , err.Message())
	assert.Nil(t, rr)
}

// login tests

func TestLoginSeccessfully(t *testing.T) {
	s := MockDbConnection(t)
	registerRequest.Password = crypto.GenerateSha256(registerRequest.Password)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (phone,username,name,family,age,password) VALUES (?,?,?,?,?,?)").
		WithArgs(registerRequest.Phone, registerRequest.Username, registerRequest.Name, registerRequest.Family, registerRequest.Age, registerRequest.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	loginRequest.Password = registerRequest.Password
	_ = repositries.NewUserRepository(s.db)

	lr, err := UserService.Login(loginRequest)
	assert.NotNil(t, lr)
	assert.Nil(t, err)
}

func TestLoginPhoneOrUsernameRequired(t *testing.T) {
	loginRequest.PhoneOrUsername = ""
	rr, err := UserService.Login(loginRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, err)
	assert.Equal(t,http.StatusBadRequest , err.Status())
	assert.Equal(t,errors.PhoneOrUsernameIsRequiredErrorMessage, err.Message())
}

func TestLoginPasswordRequired(t *testing.T) {
	loginRequest.Password = ""
	rr, err := UserService.Login(loginRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, err)
	assert.Equal(t,http.StatusBadRequest , err.Status())
	assert.Equal(t,errors.PasswordIsRequiredErrorMessage, err.Message())
}

func TestLoginWithPhone(t *testing.T) {
	s := MockDbConnection(t)

	registerRequest.Password = crypto.GenerateSha256(registerRequest.Password)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (phone,username,name,family,age,password) VALUES (?,?,?,?,?,?)").
		WithArgs(registerRequest.Phone, registerRequest.Username, registerRequest.Name, registerRequest.Family, registerRequest.Age, registerRequest.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	loginRequest.Password = registerRequest.Password
	_ = repositries.NewUserRepository(s.db)

	lr, err := UserService.Login(loginRequest)
	assert.NotNil(t, lr)
	assert.Nil(t, err)
}

func TestLoginWithUsername(t *testing.T) {
	s := MockDbConnection(t)

	registerRequest.Password = crypto.GenerateSha256(registerRequest.Password)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (phone,username,name,family,age,password) VALUES (?,?,?,?,?,?)").
		WithArgs(registerRequest.Phone, registerRequest.Username, registerRequest.Name, registerRequest.Family, registerRequest.Age, registerRequest.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	loginRequest.Password = registerRequest.Password
	loginRequest.PhoneOrUsername = registerRequest.Username
	//repositries.UserRepository.

	lr, err := UserService.Login(loginRequest)
	assert.NotNil(t, lr)
	assert.Nil(t, err)
}

func TestLoginFailToGenerateJwtToken(t *testing.T) {
	// mock jwt service
	generateJwtFunc = func(data jwt.MapClaims) (string, rest_errors.RestErr) {
		return "" , rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil)
	}

	jwt2.JwtService = &JwtServiceMock{}

	rr, err := UserService.Register(registerRequest)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError , err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage , err.Message())
	assert.Nil(t, rr)
}

func TestGetUserFailToVerifyToken(t *testing.T) {
	token := "oiadshs23kj3h123j32.23kjhkjehdkjh.23kjhk2nbemnwbd"
	gur , err := UserService.GetUser(token)
	assert.NotNil(t, err)
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
	assert.Equal(t, http.StatusUnauthorized, err.Status())
	assert.Nil(t, gur)
}


func TestGetUserFailToGetDataFromRepository(t *testing.T) {
	verifyJwtFunc = func(token string) (*jwt3.Jwt, bool , rest_errors.RestErr) {
		return &jwt3.Jwt{
			Header:  jwt3.JwtHeader{Alg: "HS256",Typ: "JWT"},
			Payload: jwt3.JwtPayload{Sub: "1", Exp: time.Time{}},
		} , true , nil
	}
	jwt2.JwtService = &JwtServiceMock{}

	getUserFunc = func(id uint) (*users2.PublicUser, rest_errors.RestErr) {
		return nil , rest_errors.NewNotFoundError(errors.UserNotFoundError)
	}

	users.UserRepository = &UserRespositoryMock{}

	gu , err := UserService.GetUser("some token")
	assert.NotNil(t, err)
	assert.Nil(t, gu)
	assert.Equal(t, http.NotFound , err.Status())
	assert.Equal(t, errors.UserNotFoundError , err.Message())

}


func TestGetUserSuccessfully(t *testing.T) {
	verifyJwtFunc = func(token string) (*jwt3.Jwt, bool , rest_errors.RestErr) {
		return &jwt3.Jwt{
			Header:  jwt3.JwtHeader{Alg: "HS256",Typ: "JWT"},
			Payload: jwt3.JwtPayload{Sub: "1", Exp: time.Time{}},
		} , true , nil
	}

	jwt2.JwtService = &JwtServiceMock{}

	getUserFunc = func(id uint) (*users2.PublicUser, rest_errors.RestErr) {
		return &users2.PublicUser{
			ID:       uint(1),
			Phone:    registerRequest.Phone,
			Username: registerRequest.Username,
			Age:      registerRequest.Age,
		} , nil
	}

	users.UserRepository = &UserRespositoryMock{}

	pu , err := UserService.GetUser("some token")

	assert.NotNil(t, pu)
	assert.Nil(t, err)
}