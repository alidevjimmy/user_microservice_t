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
