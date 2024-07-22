package service

import (
	"github.com/Lzzzzzzy/UPet/server/service/example"
	"github.com/Lzzzzzzy/UPet/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}
