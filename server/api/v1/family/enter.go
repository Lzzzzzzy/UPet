package family

import "github.com/Lzzzzzzy/UPet/server/service"

var (
	familyService = service.ServiceGroupApp.FamilyServiceGroup.FamilyService
	userService   = service.ServiceGroupApp.UserServiceGroup.UserService
	petService = service.ServiceGroupApp.PetServiceGroup
)

type ApiGroup struct {
	FamilyApi
}
