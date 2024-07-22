package request

import (
	"github.com/Lzzzzzzy/UPet/server/model/common/request"
	"github.com/Lzzzzzzy/UPet/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
