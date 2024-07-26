package router

import (
	"github.com/Lzzzzzzy/UPet/server/router/auth"
	"github.com/Lzzzzzzy/UPet/server/router/pet"
	petTodo "github.com/Lzzzzzzy/UPet/server/router/pet_todo"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	// System  system.RouterGroup
	// Example example.RouterGroup
	Pet     pet.RouterGroup
	PetTodo petTodo.RouterGroup
	Auth    auth.RouterGroup
}
