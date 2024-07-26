package auth

import (
	api "github.com/Lzzzzzzy/UPet/server/api/v1"
)

type RouterGroup struct {
	AuthRouter
}

var (
	authApi = api.ApiGroupApp.AuthApiGroup.AuthApi
)
