package domains

import (
	"gorm.io/gorm"
	"time"
)

type (
	Code struct {
		gorm.Model
		Phone          string `json:"phone"`
		Code           int    `json:"code"`
		CodePurpose     int `json:"code_purpose"`
		CodeExpiration time.Time `json:"code_expiration"`
	}
)

func (c *Code) TableName() string {
	return "codes"
}
