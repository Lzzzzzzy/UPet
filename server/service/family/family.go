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
