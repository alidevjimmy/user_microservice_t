package app

import (
	"fmt"

	"github.com/alidevjimmy/user_microservice_t/controllers/v1"
	"github.com/alidevjimmy/user_microservice_t/middlewares/v1"
)

const (
	V1Prefix = "/v1/%s"
)

func urlMapper() {
	// v1
	e.POST(fmt.Sprintf(V1Prefix, "register"), controllers.UsersController.Register)
	e.POST(fmt.Sprintf(V1Prefix, "login"), controllers.UsersController.Login)
	e.POST(fmt.Sprintf(V1Prefix, "sendCode"), controllers.CodesController.SendCode)
	e.PUT(fmt.Sprintf(V1Prefix, "changePassword"), controllers.UsersController.ChangePassword)
	e.PUT(fmt.Sprintf(V1Prefix, "verifyUser"), controllers.UsersController.Verify)
	e.GET(fmt.Sprintf(V1Prefix, "getUser"), controllers.UsersController.GetUser)

	e.PATCH(fmt.Sprintf(V1Prefix, "updateUser:user_id"), controllers.UsersController.UpdateUser, middlewares.OnlyActive)
	// admin
	e.Group(fmt.Sprintf(V1Prefix, "admin"), middlewares.OnlyAdmin)
	e.GET(fmt.Sprintf(V1Prefix, "admin/users"), controllers.UsersController.GetUsers)
	e.PATCH(fmt.Sprintf(V1Prefix, "admin/toggleActive:user_id"), controllers.UsersController.ActiveUser)
	e.PATCH(fmt.Sprintf(V1Prefix, "admin/toggleBlock:user_id"), controllers.UsersController.BlockUser)
}
