package family

import "github.com/Lzzzzzzy/UPet/server/service"

var (
	familyService = service.ServiceGroupApp.FamilyServiceGroup.FamilyService
	userService   = service.ServiceGroupApp.UserServiceGroup.UserService
)

type ApiGroup struct {
	FamilyApi
}
