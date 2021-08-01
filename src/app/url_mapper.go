package app

import (
	"fmt"
	"github.com/alidevjimmy/user_microservice_t/controllers/v1/codes"
	"github.com/alidevjimmy/user_microservice_t/controllers/v1/users"
)

const (
	v1Prefix = "/v1/%s"
)

func urlMapper() {
	// v1
	e.POST(fmt.Sprintf(v1Prefix, "register"), users.Register)
	e.POST(fmt.Sprintf(v1Prefix, "login"), users.Login)
	e.POST(fmt.Sprintf(v1Prefix, "sendcode"), codes.SendCode)
	e.POST(fmt.Sprintf(v1Prefix, "verifycode"), codes.VerifyCode)
	e.GET(fmt.Sprintf(v1Prefix, "users"), users.GetUsers) // TODO: implement OnlyAdmin middleware
	e.PATCH(fmt.Sprintf(v1Prefix, "active:user_id"), users.ActiveUser)
	e.PATCH(fmt.Sprintf(v1Prefix, "blockuser:user_id"), users.BlockUser)
	e.PATCH(fmt.Sprintf(v1Prefix, "edituser:user_id"), users.EditUser) // TODO: implement OnlyActive middleware
}
