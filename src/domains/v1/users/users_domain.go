package users

type User struct {
	Id int `json:"id"`
	Phone string `json:"phone"`
	Username string `json:"username"`
	Name string `json:"name"`
	Family string `json:"family"`
	Age uint `json:"age"`
	Active bool `json:"active"`
	Blocked bool `json:"blocked"`
	Password bool `json:"password"`
	IsAdmin bool `json:"is_admin"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type RegisterRequest struct {
	Phone string `json:"phone"`
	Username string `json:"username"`
	Name string `json:"name"`
	Family string `json:"family"`
	Age uint `json:"age"`
	Password bool `json:"password"`
}

type LoginRequest struct {
	Phone string `json:"phone"`
	Username string `json:"username"`
	Password bool `json:"password"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Token string `json:"token"`
}


type EditUserRequest struct{
	Username string `json:"username"`
	Name string `json:"name"`
	Family string `json:"family"`
	Age uint `json:"age"`
	Password string `json:"password"`
}