package v1

import (
	"github.com/Lzzzzzzy/UPet/server/api/v1/example"
	"github.com/Lzzzzzzy/UPet/server/api/v1/pet"
	"github.com/Lzzzzzzy/UPet/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	PetApiGroup     pet.ApiGroup
}
