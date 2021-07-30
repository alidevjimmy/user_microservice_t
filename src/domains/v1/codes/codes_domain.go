package codes

import "time"

type (
	Code struct {
		Id             int       `json:"id"`
		UserID         int       `json:"user_id"`
		Code           string    `json:"code"`
		CodeExpiration string    `json:"code_expiration"`
		CreatedAt      string    `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		DeletedAt      time.Time `json:"deleted_at"`
	}
)
