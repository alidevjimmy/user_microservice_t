package domains

import (
	"gorm.io/gorm"
)

type (
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

	PublicUser struct {
		ID       uint   `json:"id"`
		Phone    string `json:"phone"`
		Username string `json:"username"`
		Name     string `json:"name"`
		Family   string `json:"family"`
		Age      uint   `json:"age"`
		Active   bool   `json:"active"`
		Blocked  bool   `json:"blocked"`
		IsAdmin  bool   `json:"is_admin"`
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
		PhoneOrUsername string `json:"phoneOrUsername" validate:"required,max=50"`
		Password        string `json:"password" validate:"required,max=300"`
	}

	RegisterResponse struct {
		Token string `json:"token"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}

	UpdateUserRequest struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Family   string `json:"family"`
		Age      uint   `json:"age"`
		Password string `json:"password"`
	}

	GetUserRequest struct {
		Token string `json:"token" validate:"required"`
	}

	GetUsersRequest struct {
		Active  bool `json:"active"`
		Blocked bool `json:"blocked"`
	}

	UpdateActiveUserStateRequest struct {
		UserID uint `json:"user_id" validate:"required"`
	}

	UpdateBlockUserStateRequest struct {
		UserID uint `json:"user_id" validate:"required"`
	}

	ChangePasswordRequest struct {
		Phone       string `json:"phone" validate:"required,max=12"`
		Code        int    `json:"code"  validate:"required"`
		NewPassword string `json:"new_password"  validate:"required, max=12"`
	}

	VerifyUserRequest struct {
		Phone string `json:"phone" validate:"required,max=12"`
		Code  int    `json:"code"  validate:"required"`
	}
)

func (u *User) TableName() string {
	return "users"
}
