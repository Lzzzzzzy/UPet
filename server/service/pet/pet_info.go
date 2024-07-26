package pet

import (
	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/common/request"
	"github.com/Lzzzzzzy/UPet/server/model/pet"
)

type PetInfoService struct{}

// @function: CreatePet
// @description: 创建宠物
// @param: e model.PetInfo
// @return: err error
func (petInfo *PetInfoService) CreatePet(e pet.PetInfo) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

// @function: DeletePet
// @description: 删除宠物
// @param: e model.PetInfo
// @return: err error
func (petInfo *PetInfoService) DeletePet(e pet.PetInfo) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

// @function: UpdatePet
// @description: 更新宠物
// @param: e model.PetInfo
// @return: err error
func (petInfo *PetInfoService) UpdatePet(e *pet.PetInfo) (err error) {
	err = global.GVA_DB.Save(&e).Error
	return err
}

// @function: GetPetInfo
// @description: 获取宠物信息
// @param: id uint
// @return: customer model.PetInfo, err error
func (petInfo *PetInfoService) GetPetInfo(id uint) (e pet.PetInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&e).Error
	return
}

// @function: GetPetInfoList
// @description: 分页获取宠物信息列表
// @param: familyID uint, info request.PageInfo
// @return: list interface{}, total int64, err error
func (petInfo *PetInfoService) GetPetInfoList(familyID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&pet.PetInfo{})

	var PetInfoList []pet.PetInfo
	err = db.Where("family_id = ?", familyID).Count(&total).Error
	if err != nil {
		return PetInfoList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Where("family_id = ?", familyID).Find(&PetInfoList).Error
	}
	return PetInfoList, total, err
}
