package petTodo

import (
	"time"

	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/common/request"
	"github.com/Lzzzzzzy/UPet/server/model/pet"
	petTodo "github.com/Lzzzzzzy/UPet/server/model/pet_todo"
)

type PetTodoService struct{}

// @function: CreatePetTodo
// @description: 创建宠物待办
// @param: e model.PetTodoInfo
// @return: err error
func (petTodoInfo *PetTodoService) CreatePetTodo(e petTodo.PetTodoInfo) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

// @function: DeletePetTodo
// @description: 删除宠物待办
// @param: e model.PetTodoInfo
// @return: err error
func (petTodoInfo *PetTodoService) DeletePetTodo(e petTodo.PetTodoInfo) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

// @function: UpdatePetTodo
// @description: 更新宠物待办
// @param: e model.PetTodoInfo
// @return: err error
func (petTodoInfo *PetTodoService) UpdatePetTodo(e *petTodo.PetTodoInfo) (err error) {
	err = global.GVA_DB.Save(&e).Error
	return err
}

// @function: GetPetTodoInfo
// @description: 获取宠物待办信息
// @param: id uint
// @return: customer model.PetInfo, err error
func (petInfo *PetTodoService) GetPetTodoInfo(id uint) (e petTodo.PetTodoInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&e).Error
	return
}

// @function: GetPetTodoInfoList
// @description: 分页获取宠物待办信息列表
// @param: familyID uint, info request.PageInfo
// @return: list interface{}, total int64, err error
func (petTodoInfo *PetTodoService) GetPetInfoList(TodoTime time.Time, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&pet.PetInfo{})

	var PetTodoInfoList []petTodo.PetTodoInfo
	err = db.Limit(limit).Offset(offset).Where("todo_time = ?", TodoTime).Find(&PetTodoInfoList).Error
	return PetTodoInfoList, total, err
}
