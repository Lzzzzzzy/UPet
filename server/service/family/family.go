package family

import (
	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/family"
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
