package v1

import (
	"github.com/Lzzzzzzy/UPet/server/api/v1/pet"
	petTodo "github.com/Lzzzzzzy/UPet/server/api/v1/pet_todo"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	PetApiGroup     pet.ApiGroup
	petTodoApiGroup petTodo.ApiGroup
}
