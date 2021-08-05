package repositories

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alidevjimmy/user_microservice_t/errors/v1"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"testing"
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
	user := struct {
		ID       uint
		Phone    string
		Username string
		Active   bool
		Blocked  bool
		Name     string
		Family   string
		Age      uint
		Password string
	}{
		ID:       uint(1),
		Phone:    "0923123",
		Active:   true,
		Blocked:  false,
		Name:     "name",
		Family:   "family",
		Username: "username",
		Age:      uint(20),
		Password: "password",
	}
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (id,phone,username,name,family,age,password,active,blocked) VALUES (?,?,?,?,?,?,?,?)").
		WithArgs(user.ID, user.Phone, user.Username, user.Name, user.Family, user.Age, user.Password, user.Active, user.Blocked).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db, false)
	u, err := up.GetUserByID(uint(1))
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
}

func TestUserRepository_GetUserByIdNotFoundShouldReturnNil(t *testing.T) {
	user := struct {
		ID       uint
		Phone    string
		Username string
		Active   bool
		Blocked  bool
		Name     string
		Family   string
		Age      uint
		Password string
	}{
		ID:       uint(1),
		Phone:    "0923123",
		Active:   true,
		Blocked:  false,
		Name:     "name",
		Family:   "family",
		Username: "username",
		Age:      uint(20),
		Password: "password",
	}
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (id,phone,username,name,family,age,password,active,blocked) VALUES (?,?,?,?,?,?,?,?)").
		WithArgs(user.ID, user.Phone, user.Username, user.Name, user.Family, user.Age, user.Password, user.Active, user.Blocked).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db, false)
	u, err := up.GetUserByID(uint(1))
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.NotFound, err.Status())
	assert.Equal(t, errors.UserNotFoundError, err.Message())
}

func TestUserRepository_GetUserByID(t *testing.T) {
	user := struct {
		ID       uint
		Phone    string
		Username string
		Active   bool
		Blocked  bool
		Name     string
		Family   string
		Age      uint
		Password string
	}{
		ID:       uint(1),
		Phone:    "0923123",
		Active:   true,
		Blocked:  false,
		Name:     "name",
		Family:   "family",
		Username: "username",
		Age:      uint(20),
		Password: "password",
	}
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (id,phone,username,name,family,age,password,active,blocked) VALUES (?,?,?,?,?,?,?,?)").
		WithArgs(user.ID, user.Phone, user.Username, user.Name, user.Family, user.Age, user.Password, user.Active, user.Blocked).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db, false)
	u, err := up.GetUserByID(uint(1))
	assert.NotNil(t, u)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, user.ID)
	assert.Equal(t, u.Phone, user.Phone)
	assert.Equal(t, u.Username, user.Username)
	assert.Equal(t, u.Name, user.Name)
	assert.Equal(t, u.Family, user.Family)
	assert.Equal(t, u.Age, user.Age)
}

func TestUserRepository_FailToGetUserByPhone(t *testing.T) {
	user := struct {
		ID       uint
		Phone    string
		Username string
		Active   bool
		Blocked  bool
		Name     string
		Family   string
		Age      uint
		Password string
	}{
		ID:       uint(1),
		Phone:    "0923123",
		Active:   true,
		Blocked:  false,
		Name:     "name",
		Family:   "family",
		Username: "username",
		Age:      uint(20),
		Password: "password",
	}
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (id,phone,username,name,family,age,password,active,blocked) VALUES (?,?,?,?,?,?,?,?)").
		WithArgs(user.ID, user.Phone, user.Username, user.Name, user.Family, user.Age, user.Password, user.Active, user.Blocked).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db, false)
	u, err := up.GetUserByPhone(user.Phone)
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
}

func TestUserRepository_GetUserByPhoneNotFoundShouldReturnNil(t *testing.T) {
	user := struct {
		ID       uint
		Phone    string
		Username string
		Active   bool
		Blocked  bool
		Name     string
		Family   string
		Age      uint
		Password string
	}{
		ID:       uint(1),
		Phone:    "0923123",
		Active:   true,
		Blocked:  false,
		Name:     "name",
		Family:   "family",
		Username: "username",
		Age:      uint(20),
		Password: "password",
	}
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (id,phone,username,name,family,age,password,active,blocked) VALUES (?,?,?,?,?,?,?,?)").
		WithArgs(user.ID, user.Phone, user.Username, user.Name, user.Family, user.Age, user.Password, user.Active, user.Blocked).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db, false)
	u, err := up.GetUserByPhone(user.Phone)
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.NotFound, err.Status())
	assert.Equal(t, errors.UserNotFoundError, err.Message())
}

func TestUserRepository_GetUserByPhone(t *testing.T) {
	user := struct {
		ID       uint
		Phone    string
		Username string
		Active   bool
		Blocked  bool
		Name     string
		Family   string
		Age      uint
		Password string
	}{
		ID:       uint(1),
		Phone:    "0923123",
		Active:   true,
		Blocked:  false,
		Name:     "name",
		Family:   "family",
		Username: "username",
		Age:      uint(20),
		Password: "password",
	}
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (id,phone,username,name,family,age,password,active,blocked) VALUES (?,?,?,?,?,?,?,?)").
		WithArgs(user.ID, user.Phone, user.Username, user.Name, user.Family, user.Age, user.Password, user.Active, user.Blocked).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db, false)
	u, err := up.GetUserByPhone(user.Phone)
	assert.NotNil(t, u)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, user.ID)
	assert.Equal(t, u.Phone, user.Phone)
	assert.Equal(t, u.Username, user.Username)
	assert.Equal(t, u.Name, user.Name)
	assert.Equal(t, u.Family, user.Family)
	assert.Equal(t, u.Age, user.Age)
}

func TestUserRepository_FailToGetUserByUsername(t *testing.T) {
	user := struct {
		ID       uint
		Phone    string
		Username string
		Active   bool
		Blocked  bool
		Name     string
		Family   string
		Age      uint
		Password string
	}{
		ID:       uint(1),
		Phone:    "0923123",
		Active:   true,
		Blocked:  false,
		Name:     "name",
		Family:   "family",
		Username: "username",
		Age:      uint(20),
		Password: "password",
	}
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (id,phone,username,name,family,age,password,active,blocked) VALUES (?,?,?,?,?,?,?,?)").
		WithArgs(user.ID, user.Phone, user.Username, user.Name, user.Family, user.Age, user.Password, user.Active, user.Blocked).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db, false)
	u, err := up.GetUserByUsername(user.Username)
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
}

func TestUserRepository_GetUserByUsernameNotFoundShouldReturnNil(t *testing.T) {
	user := struct {
		ID       uint
		Phone    string
		Username string
		Active   bool
		Blocked  bool
		Name     string
		Family   string
		Age      uint
		Password string
	}{
		ID:       uint(1),
		Phone:    "0923123",
		Active:   true,
		Blocked:  false,
		Name:     "name",
		Family:   "family",
		Username: "username",
		Age:      uint(20),
		Password: "password",
	}
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (id,phone,username,name,family,age,password,active,blocked) VALUES (?,?,?,?,?,?,?,?)").
		WithArgs(user.ID, user.Phone, user.Username, user.Name, user.Family, user.Age, user.Password, user.Active, user.Blocked).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db, false)
	u, err := up.GetUserByUsername(user.Username)
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.NotFound, err.Status())
	assert.Equal(t, errors.UserNotFoundError, err.Message())
}

func TestUserRepository_GetUserByUsername(t *testing.T) {
	user := struct {
		ID       uint
		Phone    string
		Username string
		Active   bool
		Blocked  bool
		Name     string
		Family   string
		Age      uint
		Password string
	}{
		ID:       uint(1),
		Phone:    "0923123",
		Active:   true,
		Blocked:  false,
		Name:     "name",
		Family:   "family",
		Username: "username",
		Age:      uint(20),
		Password: "password",
	}
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (id,phone,username,name,family,age,password,active,blocked) VALUES (?,?,?,?,?,?,?,?)").
		WithArgs(user.ID, user.Phone, user.Username, user.Name, user.Family, user.Age, user.Password, user.Active, user.Blocked).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db, false)
	u, err := up.GetUserByUsername(user.Username)
	assert.NotNil(t, u)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, user.ID)
	assert.Equal(t, u.Phone, user.Phone)
	assert.Equal(t, u.Username, user.Username)
	assert.Equal(t, u.Name, user.Name)
	assert.Equal(t, u.Family, user.Family)
	assert.Equal(t, u.Age, user.Age)
}

func TestUserRepository_FailToGetUserByPhoneOrUsernameAndPassword(t *testing.T) {
	user := struct {
		ID       uint
		Phone    string
		Username string
		Active   bool
		Blocked  bool
		Name     string
		Family   string
		Age      uint
		Password string
	}{
		ID:       uint(1),
		Phone:    "0923123",
		Active:   true,
		Blocked:  false,
		Name:     "name",
		Family:   "family",
		Username: "username",
		Age:      uint(20),
		Password: "password",
	}

	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (id,phone,username,name,family,age,password,active,blocked) VALUES (?,?,?,?,?,?,?,?)").
		WithArgs(user.ID, user.Phone, user.Username, user.Name, user.Family, user.Age, user.Password, user.Active, user.Blocked).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db, false)
	u, err := up.GetUserByPhoneOrUsernameAndPassword(user.Phone, "password")
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
	assert.Equal(t, errors.InternalServerErrorMessage, err.Message())
}

func TestUserRepository_GetUserByPhoneOrUsernameAndPasswordNotFoundShouldReturnNil(t *testing.T) {
	user := struct {
		ID       uint
		Phone    string
		Username string
		Active   bool
		Blocked  bool
		Name     string
		Family   string
		Age      uint
		Password string
	}{
		ID:       uint(1),
		Phone:    "0923123",
		Active:   true,
		Blocked:  false,
		Name:     "name",
		Family:   "family",
		Username: "username",
		Age:      uint(20),
		Password: "password",
	}
	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (id,phone,username,name,family,age,password,active,blocked) VALUES (?,?,?,?,?,?,?,?)").
		WithArgs(user.ID, user.Phone, user.Username, user.Name, user.Family, user.Age, user.Password, user.Active, user.Blocked).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db, false)
	u, err := up.GetUserByPhoneOrUsernameAndPassword(user.Phone, "password")
	assert.NotNil(t, err)
	assert.Nil(t, u)
	assert.Equal(t, http.NotFound, err.Status())
	assert.Equal(t, errors.UserNotFoundError, err.Message())
}

func TestUserRepository_GetUserByPhoneOrUsernameAndPassword(t *testing.T) {
	user := struct {
		ID       uint
		Phone    string
		Username string
		Active   bool
		Blocked  bool
		Name     string
		Family   string
		Age      uint
		Password string
	}{
		ID:       uint(1),
		Phone:    "0923123",
		Active:   true,
		Blocked:  false,
		Name:     "name",
		Family:   "family",
		Username: "username",
		Age:      uint(20),
		Password: "password",
	}

	s := MockDbConnection(t)
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO users (id,phone,username,name,family,age,password,active,blocked) VALUES (?,?,?,?,?,?,?,?)").
		WithArgs(user.ID, user.Phone, user.Username, user.Name, user.Family, user.Age, user.Password, user.Active, user.Blocked).
		WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
	s.mock.ExpectCommit()

	up := NewUserRepository(s.db, false)
	u, err := up.GetUserByPhoneOrUsernameAndPassword(user.Phone, "password")
	assert.NotNil(t, u)
	assert.Nil(t, err)
	assert.Equal(t, u.ID, user.ID)
	assert.Equal(t, u.Phone, user.Phone)
	assert.Equal(t, u.Username, user.Username)
	assert.Equal(t, u.Name, user.Name)
	assert.Equal(t, u.Family, user.Family)
	assert.Equal(t, u.Age, user.Age)
}

func TestUserRepository_GetUsers(t *testing.T) {
	s := MockDbConnection(t)
	users := []struct {
		ID       uint
		Phone    string
		Username string
		Active   bool
		Blocked  bool
		Name     string
		Family   string
		Age      uint
		Password string
	}{
		{
			ID:       uint(1),
			Phone:    "0923123",
			Active:   true,
			Blocked:  false,
			Name:     "name",
			Family:   "family",
			Username: "username",
			Age:      uint(20),
			Password: "password",
		},
		{
			ID:       uint(2),
			Phone:    "09123123",
			Active:   true,
			Blocked:  true,
			Name:     "name1",
			Family:   "family2",
			Username: "username2",
			Age:      uint(30),
			Password: "password",
		},
	}

	for _, user := range users {
		s.mock.ExpectBegin()
		s.mock.ExpectExec("INSERT INTO users (id,phone,username,name,family,age,password,active,blocked) VALUES (?,?,?,?,?,?,?,?)").
			WithArgs(user.ID, user.Phone, user.Username, user.Name, user.Family, user.Age, user.Password, user.Active, user.Blocked).
			WillReturnResult(sqlmock.NewResult(int64(user.ID), 1))
		s.mock.ExpectCommit()
	}

}
