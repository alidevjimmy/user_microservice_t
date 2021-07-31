package users

import (
	"github.com/alidevjimmy/go-rest-utils/rest_errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type (
	UserValidator struct {
		validator *validator.Validate
	}

	User struct {
		gorm.Model
		Phone    string `json:"phone" gorm:"column:phone"`
		Username string `json:"username" gorm:"column:username"`
		Name     string `json:"name" gorm:"column:name"`
		Family   string `json:"family" gorm:"column:family"`
		Age      uint   `json:"age" gorm:"column:age"`
		Active   bool   `json:"active" gorm:"column:active"`
		Blocked  bool   `json:"blocked" gorm:"column:blocked"`
		Password bool   `gorm:"column:password"`
		IsAdmin  bool   `json:"is_admin" gorm:"column:is_admin"`
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
		PhoneOrUsername string `json:"phone" validate:"required,max=50"`
		Password        string `json:"password" validate:"required,max=300"`
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

func (u *User) TableName() string {
	return "users"
}
