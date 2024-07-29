package router

import (
	"github.com/Lzzzzzzy/UPet/server/router/auth"
	"github.com/Lzzzzzzy/UPet/server/router/pet"
	petTodo "github.com/Lzzzzzzy/UPet/server/router/pet_todo"
	"github.com/Lzzzzzzy/UPet/server/router/system"
	"github.com/Lzzzzzzy/UPet/server/router/user"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Pet     pet.RouterGroup
	PetTodo petTodo.RouterGroup
	Auth    auth.RouterGroup
	User    user.RouterGroup
}
