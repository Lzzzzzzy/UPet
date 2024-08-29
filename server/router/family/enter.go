package family

import (
	api "github.com/Lzzzzzzy/UPet/server/api/v1"
)

type RouterGroup struct {
	FamilyRouter
}

var (
	familyApi = api.ApiGroupApp.FamilyApiGroup
)
