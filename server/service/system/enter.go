package system

type ServiceGroup struct {
	InitDBService
	DictionaryService
	SystemConfigService
	OperationRecordService
	DictionaryDetailService
	FileUploadAndDownloadService
}
