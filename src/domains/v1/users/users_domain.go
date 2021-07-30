package users

import (
	"time"

	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/go-playground/validator/v10"
)

type (
	UserValidator struct {
		validator *validator.Validate
	}

	User struct {
		ID        uint      `json:"id" gorm:"primaryKey"`
		Phone     string    `json:"phone"`
		Username  string    `json:"username"`
		Name      string    `json:"name"`
		Family    string    `json:"family"`
		Age       uint      `json:"age"`
		Active    bool      `json:"active"`
		Blocked   bool      `json:"blocked"`
		Password  bool      `json:"password"`
		IsAdmin   bool      `json:"is_admin"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}

	RegisterRequest struct {
		Phone    string `json:"phone" validate:"required,max=12"`
		Username string `json:"username" validate:"required,max=250"`
		Name     string `json:"name" validate:"required,max=255"`
		Family   string `json:"family" validate:"required,max=255"`
		Age      uint   `json:"age" validate:"required,min=1"`
		Password string `json:"password" validate:"required,max=300"`
	}

	LoginRequest struct {
		PhoneOrUsername string `json:"phone" validate:"required"`
		Password        string `json:"password" validate:"required"`
	}

	RegisterResponse struct {
		Token string `json:"token"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}

	EditUserRequest struct {
		Username string `json:"username" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Family   string `json:"family" validate:"required"`
		Age      uint   `json:"age" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)

func (uv *UserValidator) Validate(i interface{}) error {
	if err := uv.validator.Struct(i); err != nil {
		return rest_errors.NewBadRequestError(err.Error())
	}

	return nil
}
