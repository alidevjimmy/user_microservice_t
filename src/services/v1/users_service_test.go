package services

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alidevjimmy/go-rest-utils/crypto"
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/errors"
	"github.com/alidevjimmy/user_microservice_t/repositories/postgres/v1"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"net/http"
	"testing"

	"github.com/alidevjimmy/user_microservice_t/domains/v1/users"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
)

var (
	registerRequest users.RegisterRequest = users.RegisterRequest{
		Phone:    "09122334344",
		Username: "test_user",
		Name:     "ali",
		Family:   "hamrani",
		Age:      uint(20),
		Password: "password",
	}
	loginRequest users.LoginRequest = users.LoginRequest{
		PhoneOrUsername:    "09122334344",
		Password: "password",
	}
	// for mock different states of CodeService methods
	sendCodeFunc func(phone string) rest_errors.RestErr
	// for mock different states of JwtService methods
	generateJwtFunc func(data jwt.MapClaims) (string, rest_errors.RestErr)
)

type Suite struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

type JwtServiceMock struct{}

func(*JwtServiceMock) Generate(data jwt.MapClaims) (string, rest_errors.RestErr) {
	return generateJwtFunc(data)
}

func(*JwtServiceMock) Verify(token string) (*jwt.MapClaims, bool, rest_errors.RestErr) {
	return nil,false, nil
}

type CodeServiceMock struct{}

func (*CodeServiceMock) Send(phone string) rest_errors.RestErr {
	return sendCodeFunc(phone)
}

func (*CodeServiceMock) Verify(phone string, code int) rest_errors.RestErr {
	return nil
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

	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (phone,username,name,family,age,password) VALUES (?,?,?,?,?,?)").
		WithArgs(registerRequest.Phone, registerRequest.Username, registerRequest.Name, registerRequest.Family, registerRequest.Age, registerRequest.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	repo := repositries.NewUserRepository(s.db)
	user, err := repo.GetUserByPhone(registerRequest.Phone)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	rr, error := UserService.Register(registerRequest)
	assert.Nil(t, rr)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest , error.Status())
	assert.Equal(t, errors.DuplicatePhoneErrorMessage , error.Message())
}

func TestRegisterCanInsertDuplicatedUsername(t *testing.T) {
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

	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (phone,username,name,family,age,password) VALUES (?,?,?,?,?,?)").
		WithArgs(registerRequest.Phone, registerRequest.Username, registerRequest.Name, registerRequest.Family, registerRequest.Age, registerRequest.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	registerRequest.Phone = "092837232"
	repo := repositries.NewUserRepository(s.db)
	user, err := repo.GetUserByUsername(registerRequest.Username)
	assert.Nil(t, err)
	assert.NotNil(t, user)

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

	CodeService = &CodeServiceMock{}

	rr, err := UserService.Register(registerRequest)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError , err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage , err.Message())
	assert.Nil(t, rr)
}

func TestRegisterFailToGenerateJwtToken(t *testing.T) {
	// mock jwt service
	generateJwtFunc = func(data jwt.MapClaims) (string, rest_errors.RestErr) {
		return "" , rest_errors.NewInternalServerError("cannot generate jwt token", nil)
	}

	JwtService = &JwtServiceMock{}

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
	s := &Suite{}
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()

	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}
	defer db.Close()
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
	registerRequest.Password = crypto.GenerateSha256(registerRequest.Password)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (phone,username,name,family,age,password) VALUES (?,?,?,?,?,?)").
		WithArgs(registerRequest.Phone, registerRequest.Username, registerRequest.Name, registerRequest.Family, registerRequest.Age, registerRequest.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	loginRequest.Password = registerRequest.Password
	repo := repositries.NewUserRepository(s.db)
	user, err := repo.GetUserByPhoneOrUsernameAndPassword(loginRequest.PhoneOrUsername , loginRequest.Password)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	lr, error := UserService.Login(loginRequest)
	assert.NotNil(t, lr)
	assert.Nil(t, error)
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

// test login with phone
// test login with username
// test login success

