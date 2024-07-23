package pet

import "github.com/Lzzzzzzy/UPet/server/service"

type ApiGroup struct {
	PetApi
}

var (
	petService = service.ServiceGroupApp.PetServiceGroup
)
