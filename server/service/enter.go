package service

import (
	"github.com/Lzzzzzzy/UPet/server/service/auth"
	"github.com/Lzzzzzzy/UPet/server/service/family"
	"github.com/Lzzzzzzy/UPet/server/service/pet"
	petTodo "github.com/Lzzzzzzy/UPet/server/service/pet_todo"
	"github.com/Lzzzzzzy/UPet/server/service/system"
	"github.com/Lzzzzzzy/UPet/server/service/user"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	// ExampleServiceGroup example.ServiceGroup
	PetServiceGroup     pet.ServiceGroup
	PetTodoServiceGroup petTodo.ServiceGroup
	AuthServiceGroup    auth.ServiceGroup
	FamilyServiceGroup  family.ServiceGroup
	UserServiceGroup    user.ServiceGroup
}
