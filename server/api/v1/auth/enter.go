package auth

import "github.com/Lzzzzzzy/UPet/server/service"

var (
	jwtService      = service.ServiceGroupApp.AuthServiceGroup.JwtService
	authService     = service.ServiceGroupApp.AuthServiceGroup.AuthService
	registerService = service.ServiceGroupApp.UserServiceGroup.RegisterService
)

type ApiGroup struct {
	AuthApi
}
