package v1

import (
	"github.com/Lzzzzzzy/UPet/server/api/v1/auth"
	"github.com/Lzzzzzzy/UPet/server/api/v1/family"
	"github.com/Lzzzzzzy/UPet/server/api/v1/pet"
	petTodo "github.com/Lzzzzzzy/UPet/server/api/v1/pet_todo"
	"github.com/Lzzzzzzy/UPet/server/api/v1/system"
	"github.com/Lzzzzzzy/UPet/server/api/v1/user"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	PetApiGroup     pet.ApiGroup
	PetTodoApiGroup petTodo.ApiGroup
	AuthApiGroup    auth.ApiGroup
	UserApiGroup    user.ApiGroup
	SystemApiGroup  system.ApiGroup
	FamilyApiGroup  family.ApiGroup
}
