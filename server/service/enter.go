package service

import (
	"github.com/Lzzzzzzy/UPet/server/service/pet"
	petTodo "github.com/Lzzzzzzy/UPet/server/service/pet_todo"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	// SystemServiceGroup  system.ServiceGroup
	// ExampleServiceGroup example.ServiceGroup
	PetServiceGroup     pet.ServiceGroup
	PetTodoServiceGroup petTodo.ServiceGroup
}
