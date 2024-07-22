package router

import "github.com/Lzzzzzzy/UPet/server/plugin/announcement/api"

var (
	Router  = new(router)
	apiInfo = api.Api.Info
)

type router struct{ Info info }
