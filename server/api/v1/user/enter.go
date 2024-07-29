package user

import "github.com/Lzzzzzzy/UPet/server/service"

type ApiGroup struct {
	UserApi
}

var (
	userService = service.ServiceGroupApp.UserServiceGroup.UserService
)
