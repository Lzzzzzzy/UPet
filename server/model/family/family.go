package family

import (
	"github.com/Lzzzzzzy/UPet/server/global"
)

type Family struct {
	global.GVA_MODEL
	CreatedBy uint `json:"createdBy" form:"createdBy" gorm:"comment:创建人"` // 创建人
}
