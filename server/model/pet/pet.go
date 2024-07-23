package pet

import (
	"github.com/Lzzzzzzy/UPet/server/global"
)

type PetInfo struct {
	global.GVA_MODEL
	Avatar          string `json:"avatar" form:"avatar" gorm:"comment:宠物头像"`                     // 宠物头像
	Name            string `json:"name" form:"name" gorm:"comment:宠物名"`                          // 宠物名
	Birthday        string `json:"birthday" form:"birthday" gorm:"comment:宠物生日"`                 // 宠物生日
	Gender          string `json:"gender" form:"gender" gorm:"comment:宠物性别"`                     // 宠物性别
	Category        uint   `json:"category" form:"category" gorm:"comment:宠物类型"`                 // 宠物类型
	SterilizedState uint   `json:"sterilizedState" form:"sterilizedState" gorm:"comment:宠物绝育状态"` // 宠物绝育状态
	FamilyId        uint   `json:"familyId" form:"familyId" gorm:"comment:所属家庭id"`               // 所属家庭id
	CreatedBy       uint   `json:"createdBy" form:"createdBy" gorm:"comment:创建人"`                // 创建人
	UpdatedBy       uint   `json:"updatedBy" form:"updatedBy" gorm:"comment:更新人"`                // 更新人
}
