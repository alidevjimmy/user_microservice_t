package repositories

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/alidevjimmy/user_microservice_t/domains/v1"
	"github.com/alidevjimmy/user_microservice_t/errors/v1"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

var (
	FakeUser domains.PublicUser = domains.PublicUser{
		ID: uint(1),
		Phone:    "09122334344",
		Username: "test_user",
		Name:     "ali",
		Family:   "hamrani",
		Age:      uint(20),
	}
)

type Suite struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

func MockDbConnection(t *testing.T) *Suite {
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

func TestUserRepository_FailToGetUserById(t *testing.T) {
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("SELECT * FROM users WHERE id=1").
		WillReturnError(rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db)
	u, err := up.GetUserByID(uint(1))
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
}

func TestUserRepository_GetUserByIdNotFoundShouldReturnNil(t *testing.T) {
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT * FROM users WHERE id=1").
		WillReturnRows(sqlmock.NewRows([]string{"id","phone","username","name","family","age"}).AddRow())
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db)
	u, err := up.GetUserByID(uint(1))
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.NotFound, err.Status())
	assert.Equal(t, errors.UserNotFoundError, err.Message())
}

func TestUserRepository_GetUserByID(t *testing.T) {
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT * FROM users WHERE id=1").
		WillReturnRows(sqlmock.NewRows([]string{"id","phone","username","name","family","age"}).
			AddRow(FakeUser.ID,FakeUser.Phone,FakeUser.Username,FakeUser.Name,FakeUser.Age,FakeUser.Age))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db)
	u, err := up.GetUserByID(uint(1))
	assert.NotNil(t, u)
	assert.Nil(t, err)
	assert.Equal(t, u.ID,FakeUser.ID)
	assert.Equal(t, u.Phone,FakeUser.Phone)
	assert.Equal(t, u.Username,FakeUser.Username)
	assert.Equal(t, u.Name,FakeUser.Name)
	assert.Equal(t, u.Family,FakeUser.Family)
	assert.Equal(t, u.Age,FakeUser.Age)
}


func TestUserRepository_FailToGetUserByPhone(t *testing.T) {
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT * FROM users WHERE phone=$1").WithArgs(FakeUser.Phone).
		WillReturnError(rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db)
	u, err := up.GetUserByPhone(FakeUser.Phone)
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
}

func TestUserRepository_GetUserByPhoneNotFoundShouldReturnNil(t *testing.T) {
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT * FROM users WHERE phone=$1").WithArgs(FakeUser.Phone).
		WillReturnRows(sqlmock.NewRows([]string{"id","phone","username","name","family","age"}).AddRow())
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db)
	u, err := up.GetUserByPhone(FakeUser.Phone)
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.NotFound, err.Status())
	assert.Equal(t, errors.UserNotFoundError, err.Message())
}

func TestUserRepository_GetUserByPhone(t *testing.T) {
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT * FROM users WHERE phone=$1").WithArgs(FakeUser.Phone).
		WillReturnRows(sqlmock.NewRows([]string{"id","phone","username","name","family","age"}).
			AddRow(FakeUser.ID,FakeUser.Phone,FakeUser.Username,FakeUser.Name,FakeUser.Age,FakeUser.Age))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db)
	u, err := up.GetUserByPhone(FakeUser.Phone)
	assert.NotNil(t, u)
	assert.Nil(t, err)
	assert.Equal(t, u.ID,FakeUser.ID)
	assert.Equal(t, u.Phone,FakeUser.Phone)
	assert.Equal(t, u.Username,FakeUser.Username)
	assert.Equal(t, u.Name,FakeUser.Name)
	assert.Equal(t, u.Family,FakeUser.Family)
	assert.Equal(t, u.Age,FakeUser.Age)
}

func TestUserRepository_FailToGetUserByUsername(t *testing.T) {
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT * FROM users WHERE username=$1").WithArgs(FakeUser.Username).
		WillReturnError(rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db)
	u, err := up.GetUserByUsername(FakeUser.Username)
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
}

func TestUserRepository_GetUserByUsernameNotFoundShouldReturnNil(t *testing.T) {
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT * FROM users WHERE username=$1").WithArgs(FakeUser.Username).
		WillReturnRows(sqlmock.NewRows([]string{"id","phone","username","name","family","age"}).AddRow())
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db)
	u, err := up.GetUserByUsername(FakeUser.Username)
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.NotFound, err.Status())
	assert.Equal(t, errors.UserNotFoundError, err.Message())
}

func TestUserRepository_GetUserByUsername(t *testing.T) {
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT * FROM users WHERE username=$1").WithArgs(FakeUser.Username).
		WillReturnRows(sqlmock.NewRows([]string{"id","phone","username","name","family","age"}).
			AddRow(FakeUser.ID,FakeUser.Phone,FakeUser.Username,FakeUser.Name,FakeUser.Age,FakeUser.Age))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db)
	u, err := up.GetUserByUsername(FakeUser.Username)
	assert.NotNil(t, u)
	assert.Nil(t, err)
	assert.Equal(t, u.ID,FakeUser.ID)
	assert.Equal(t, u.Phone,FakeUser.Phone)
	assert.Equal(t, u.Username,FakeUser.Username)
	assert.Equal(t, u.Name,FakeUser.Name)
	assert.Equal(t, u.Family,FakeUser.Family)
	assert.Equal(t, u.Age,FakeUser.Age)
}

func TestUserRepository_FailToGetUserByPhoneOrUsernameAndPassword(t *testing.T) {
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT * FROM users WHERE username=$1 OR phone=$2 AND password=$3").WithArgs(FakeUser.Username, FakeUser.Phone,"password").
		WillReturnError(rest_errors.NewInternalServerError(errors.InternalServerErrorMessage, nil))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db)
	u, err := up.GetUserByPhoneOrUsernameAndPassword(FakeUser.Phone,"password")
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
}

func TestUserRepository_GetUserByPhoneOrUsernameAndPasswordNotFoundShouldReturnNil(t *testing.T) {
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT * FROM users WHERE username=$1 OR phone=$2 AND password=$3").WithArgs(FakeUser.Username, FakeUser.Phone,"password").
		WillReturnRows(sqlmock.NewRows([]string{"id","phone","username","name","family","age"}).AddRow())
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db)
	u, err := up.GetUserByPhoneOrUsernameAndPassword(FakeUser.Phone,"password")
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.NotFound, err.Status())
	assert.Equal(t, errors.UserNotFoundError, err.Message())
}

func TestUserRepository_GetUserByPhoneOrUsernameAndPassword(t *testing.T) {
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery("SELECT * FROM users WHERE username=$1 OR phone=$2 AND password=$3").WithArgs(FakeUser.Username, FakeUser.Phone,"password").
		WillReturnRows(sqlmock.NewRows([]string{"id","phone","username","name","family","age"}).
			AddRow(FakeUser.ID,FakeUser.Phone,FakeUser.Username,FakeUser.Name,FakeUser.Age,FakeUser.Age))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db)
	u, err := up.GetUserByPhoneOrUsernameAndPassword(FakeUser.Phone,"password")
	assert.NotNil(t, u)
	assert.Nil(t, err)
	assert.Equal(t, u.ID,FakeUser.ID)
	assert.Equal(t, u.Phone,FakeUser.Phone)
	assert.Equal(t, u.Username,FakeUser.Username)
	assert.Equal(t, u.Name,FakeUser.Name)
	assert.Equal(t, u.Family,FakeUser.Family)
	assert.Equal(t, u.Age,FakeUser.Age)
}
