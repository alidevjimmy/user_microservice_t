package domains

import (
	"time"

	"gorm.io/gorm"
)

type (
	Code struct {
		gorm.Model
		Phone          string    `json:"phone"`
		Code           int       `json:"code"`
		CodePurpose    int       `json:"code_purpose"`
		CodeExpiration time.Time `json:"code_expiration"`
	}

	SendCodeRequest struct {
		Phone  string `json:"phone"`
		Reason int    `json:"reason"`
	}
)

func (c *Code) TableName() string {
	return "codes"
}
