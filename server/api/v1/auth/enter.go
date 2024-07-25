package auth

import "github.com/Lzzzzzzy/UPet/server/service"

var (
	jwtService = service.ServiceGroupApp.AuthServiceGroup.JwtService
)

type ApiGroup struct {
	AuthApi
}
