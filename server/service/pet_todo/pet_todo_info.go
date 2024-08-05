package petTodo

import (
	"fmt"
	"time"

	"github.com/Lzzzzzzy/UPet/server/global"
	petTodo "github.com/Lzzzzzzy/UPet/server/model/pet_todo"
	petTodoReq "github.com/Lzzzzzzy/UPet/server/model/pet_todo/request"
	"github.com/Lzzzzzzy/UPet/server/utils"
)

type PetTodoService struct{}

// @function: CreatePetTodo
// @description: 创建宠物待办
// @param: e model.PetTodoInfo
// @return: err error
func (petTodoInfo *PetTodoService) CreatePetTodos(e []*petTodo.PetTodoInfo) (err error) {
	err = global.GVA_DB.CreateInBatches(&e, 10).Error
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
func (petTodoInfo *PetTodoService) GetPetTodoInfoList(minTime, maxTime time.Time, petId uint) (PetTodoInfoList []petTodo.SimplePetTodoInfo, err error) {
	db := global.GVA_DB.Model(&petTodo.PetTodoInfo{})

	err = db.Where("remind_time >= ?", minTime).Where("remind_time <= ?", maxTime).Where("pet_id = ?", petId).Order("remind_time").Find(&PetTodoInfoList).Error
	return
}

func (petTodoInfo *PetTodoService) FormatRemindTime(e *petTodoReq.PetTodoInfo) []time.Time {
	remindTime := make([]time.Time, len(e.RemindDate))
	for i, v := range e.RemindDate {
		timeStr := fmt.Sprintf("%s %s:00", v, e.RemindTime)
		t, _ := utils.ParseStringToTime(timeStr)
		remindTime[i] = t
	}
	return remindTime
}

// @function: GetPetTodoInfoList
// @description: 分页获取宠物待办信息列表
// @param: familyID uint, info request.PageInfo
// @return: list interface{}, total int64, err error
func (petTodoInfo *PetTodoService) GetPetTodoInfoMarkList(minTime, maxTime time.Time, FamilyId uint) (PetTodoColorList []petTodo.SimplePetTodoColor, err error) {
	db := global.GVA_DB.Model(&petTodo.PetTodoInfo{})

	err = db.Where("remind_time >= ?", minTime).Where("remind_time <= ?", maxTime).Where("family_id = ?", FamilyId).Order("remind_time").Find(&PetTodoColorList).Error
	return
}

// @function: UpdatePetTodoComplete
// @description: 更新宠物待办完成情况
// @param: todoId uint, complete bool
// @return: err error
func (petTodoInfo *PetTodoService) UpdatePetTodoComplete(todoId uint, complete bool) (err error) {
	db := global.GVA_DB.Model(&petTodo.PetTodoInfo{})
	err = db.Where("id = ?", todoId).Select("Complete").Updates(petTodo.PetTodoInfo{Complete: complete}).Error
	return err
}
