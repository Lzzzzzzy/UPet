package api

import "github.com/Lzzzzzzy/UPet/server/plugin/announcement/service"

var (
	Api         = new(api)
	serviceInfo = service.Service.Info
)

type api struct{ Info info }
