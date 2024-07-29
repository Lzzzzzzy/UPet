package auth

import "github.com/Lzzzzzzy/UPet/server/service"

var (
	jwtService      = service.ServiceGroupApp.AuthServiceGroup.JwtService
	authService     = service.ServiceGroupApp.AuthServiceGroup.AuthService
	registerService = service.ServiceGroupApp.UserServiceGroup.RegisterService
	userService     = service.ServiceGroupApp.UserServiceGroup.UserService
)

type ApiGroup struct {
	AuthApi
}
