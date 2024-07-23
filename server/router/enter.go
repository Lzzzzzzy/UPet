package router

import (
	"github.com/Lzzzzzzy/UPet/server/router/example"
	"github.com/Lzzzzzzy/UPet/server/router/pet"
	"github.com/Lzzzzzzy/UPet/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Pet     pet.RouterGroup
}
