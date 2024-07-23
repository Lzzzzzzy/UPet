package request

import "github.com/Lzzzzzzy/UPet/server/model/common/request"

type PetTodoPageInfo struct {
	request.PageInfo
	Date  string `json:"date" form:"date"`   // 日期
	PetId uint   `json:"petId" form:"petId"` // 宠物id
}
