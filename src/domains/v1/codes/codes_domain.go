package codes
type Code struct {
	Id int `json:"id"`
	UserID int `json:"user_id"`
	Code string `json:"code"`
	CodeExpiration string `json:"code_expiration"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}