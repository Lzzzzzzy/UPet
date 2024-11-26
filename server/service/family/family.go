package family

import (
	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/family"
	"github.com/Lzzzzzzy/UPet/server/model/user"
)

type FamilyService struct{}

// @function: CreateFamily
// @description: 创建家庭
// @param: e model.Family
// @return: err error
func (familyService *FamilyService) CreateFamily(e *family.Family) (err error) {
	result := global.GVA_DB.Create(e)
	err = result.Error
	return err
}

// @function: GetFamilyMembers
// @description: 获取家庭成员列表
// @param: familyId 家庭id
// @return: err error
func (familyService *FamilyService) GetFamilyMembers(familyId uint) (UserInfoList *[]user.User, err error) {
	db := global.GVA_DB.Model(&user.User{})

	err = db.Where("family_id = ?", familyId).Find(&UserInfoList).Error
	return
}

// @function: GetFamilyOtherMembers
// @description: 获取除自己外的家庭成员列表
// @param: familyId 家庭id
// @return: err error
func (familyService *FamilyService) GetFamilyOtherMembers(familyId, userId uint) (UserInfoList *[]user.User, err error) {
	db := global.GVA_DB.Model(&user.User{})

	err = db.Where("family_id = ?", familyId).Where("id != ?", userId).Find(&UserInfoList).Error
	return
}

// @function: DeleteFamily
// @description: 删除家庭
// @param: familyId 家庭id
// @return: err error
func (familyService *FamilyService) DeleteFamily(familyId uint) (err error) {
	family := family.Family{}
	family.ID = familyId
	err = global.GVA_DB.Delete(&family).Error
	return
}
