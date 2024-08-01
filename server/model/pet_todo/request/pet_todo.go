package request

import (
	"github.com/Lzzzzzzy/UPet/server/model/common"
)

type PetTodoCondition struct {
	Date  string `json:"date" form:"date"`   // 日期
	PetId uint   `json:"petId" form:"petId"` // 宠物id
}

type PetTodoInfo struct {
	Title      string             `json:"title" form:"title"`           // 标题
	Remark     string             `json:"remark" form:"remark"`         // 备注
	Remind     bool               `json:"remind" form:"remind"`         // 是否提醒， 0：否，1：是
	RemindDate []string           `json:"remindDate" form:"remindDate"` // 提醒日期
	RemindTime string             `json:"remindTime" form:"remindTime"` // 提醒时间
	Complete   bool               `json:"complete" form:"complete"`     // 是否完成，0：否，1：是
	Type       uint               `json:"type" form:"type"`             // 提醒类型， 0：日常记录， 1：待办事项
	Color      uint               `json:"color" form:"color" `          // 背景颜色
	TodoTime   *common.CustomTime `json:"todoTime" form:"todoTime" `    // 待办时间
	PetId      uint               `json:"petId" form:"petId" `          // 宠物ID
}
