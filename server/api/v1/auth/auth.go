package auth

import (
	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/common/response"
	"github.com/Lzzzzzzy/UPet/server/model/pet"
	"github.com/Lzzzzzzy/UPet/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthApi struct{}

// Auth
// @Tags      Auth
// @Summary   微信第三方登录
// @accept    application/json
// @Produce   application/json
// @Param     data  body      pet.PetInfo            true  "宠物信息"
// @Success   200   {object}  response.Response{msg=string}  "创建宠物"
// @Router    /pet [post]
func (e *AuthApi) CreatePetInfo(c *gin.Context) {
	var petInfo pet.PetInfo
	err := c.ShouldBindJSON(&petInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(petInfo, utils.PetInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	petInfo.FamilyId = utils.GetUserFamilyID(c)
	petInfo.CreatedBy = utils.GetUserID(c)
	petInfo.UpdatedBy = utils.GetUserID(c)
	err = petService.CreatePet(petInfo)

	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
