package user

import (
	api "github.com/Lzzzzzzy/UPet/server/api/v1"
)

type RouterGroup struct {
	UserRouter
}

var (
	userApi = api.ApiGroupApp.UserApiGroup.UserApi
)
