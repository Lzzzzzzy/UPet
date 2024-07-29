package system

import "github.com/Lzzzzzzy/UPet/server/service"

type ApiGroup struct {
	DBApi
	DictionaryDetailApi
	FileApi
	DictionaryApi
}

var (
	initDBService           = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	dictionaryService       = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	systemConfigService     = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	operationRecordService  = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
	dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
	fileService             = service.ServiceGroupApp.SystemServiceGroup.FileUploadAndDownloadService
)
