package system

import api "github.com/Lzzzzzzy/UPet/server/api/v1"

type RouterGroup struct {
	InitRouter
	DictionaryDetailRouter
	FileRouter
	SystemRouter
}

var (
	dbApi               = api.ApiGroupApp.SystemApiGroup.DBApi
	dictionaryDetailApi = api.ApiGroupApp.SystemApiGroup.DictionaryDetailApi
	dictionaryApi       = api.ApiGroupApp.SystemApiGroup.DictionaryApi
	fileApi             = api.ApiGroupApp.SystemApiGroup.FileApi
	systemApi           = api.ApiGroupApp.SystemApiGroup.SystemApi
)
