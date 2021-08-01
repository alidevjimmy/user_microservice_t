package app

import (
	"fmt"
	"github.com/alidevjimmy/user_microservice_t/controllers/v1"
)

const (
	v1Prefix = "/v1/%s"
)

func urlMapper() {
	// v1
	e.POST(fmt.Sprintf(v1Prefix, "register"), v1.Register)
	e.POST(fmt.Sprintf(v1Prefix, "login"), v1.Login)
	e.POST(fmt.Sprintf(v1Prefix, "sendcode"), v1.SendCode)
	e.POST(fmt.Sprintf(v1Prefix, "verifycode"), v1.VerifyCode)
	e.GET(fmt.Sprintf(v1Prefix, "users"), v1.GetUsers) // TODO: implement OnlyAdmin middleware
	e.PATCH(fmt.Sprintf(v1Prefix, "active:user_id"), v1.ActiveUser)
	e.PATCH(fmt.Sprintf(v1Prefix, "blockuser:user_id"), v1.BlockUser)
	e.PATCH(fmt.Sprintf(v1Prefix, "edituser:user_id"), v1.EditUser) // TODO: implement OnlyActive middleware
}
