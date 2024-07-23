package petTodo

import (
	"time"

	"github.com/Lzzzzzzy/UPet/server/global"
)

type PetTodoInfo struct {
	global.GVA_MODEL
	Title      string    `json:"title" form:"title" gorm:"comment:标题"`             // 标题
	Remark     string    `json:"remark" form:"remark" gorm:"comment:备注"`           // 备注
	Remind     uint      `json:"remind" form:"remind" gorm:"comment:是否提醒"`         // 是否提醒， 0：否，1：是
	RemindTime time.Time `json:"remindTime" form:"remindTime" gorm:"comment:提醒时间"` // 提醒时间
	Complete   uint      `json:"complete" form:"complete" gorm:"comment:是否完成"`     // 是否完成，0：否，1：是
	Type       uint      `json:"type" form:"type" gorm:"comment:提醒类型"`             // 提醒类型， 0：日常记录， 1：待办事项
	Color      uint      `json:"color" form:"color" gorm:"comment:背景颜色"`           // 背景颜色
	TodoTime   time.Time `json:"todoTime" form:"todoTime" gorm:"comment:待办时间"`     // 待办时间
	PetId      uint      `json:"petId" form:"petId" gorm:"comment:宠物ID"`           // 宠物ID
	CreatedBy  uint      `json:"createdBy" form:"createdBy" gorm:"comment:创建人"`    // 创建人
	UpdatedBy  uint      `json:"updatedBy" form:"updatedBy" gorm:"comment:更新人"`    // 更新人
}
