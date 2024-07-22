package request

import (
	"github.com/Lzzzzzzy/UPet/server/model/common/request"
	"github.com/Lzzzzzzy/UPet/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
