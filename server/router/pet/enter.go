package pet

import (
	api "github.com/Lzzzzzzy/UPet/server/api/v1"
)

type RouterGroup struct {
	PetRouter
}

var (
	petApi = api.ApiGroupApp.PetApiGroup
)
