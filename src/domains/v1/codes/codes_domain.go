package codes

import (
	"gorm.io/gorm"
)

type (
	Code struct {
		gorm.Model
		UserID         int    `json:"user_id"`
		Code           string `json:"code"`
		CodeExpiration string `json:"code_expiration"`
		CreatedAt      string `json:"created_at"`
	}
)

func (c *Code) TableName() string {
	return "codes"
}
