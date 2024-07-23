package petTodo

import "github.com/Lzzzzzzy/UPet/server/service"

type ApiGroup struct {
	PetTodoApi
}

var (
	petTodoService = service.ServiceGroupApp.PetTodoServiceGroup
)
