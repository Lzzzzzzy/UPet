package pet

import (
	"strconv"

	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/common/response"
	"github.com/Lzzzzzzy/UPet/server/model/pet"
	"github.com/Lzzzzzzy/UPet/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PetApi struct{}

// CreatePetInfo
// @Tags      PetInfo
// @Summary   创建宠物
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      pet.PetInfo            true  "宠物信息"
// @Success   200   {object}  response.Response{msg=string}  "创建宠物"
// @Router    /api/pet [post]
func (e *PetApi) CreatePetInfo(c *gin.Context) {
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

// DeletePetInfo
// @Tags      PetInfo
// @Summary   删除宠物
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     petID  path      int             true  "宠物ID"
// @Success   200   {object}  response.Response{msg=string}  "删除宠物"
// @Router    /api/pet/:petID [delete]
func (e *PetApi) DeletePetInfo(c *gin.Context) {
	var petInfo pet.PetInfo
	petIdStr := c.Param("petID")
	petID, err := strconv.ParseUint(petIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage("宠物ID错误", c)
		return
	}
	petInfo.ID = uint(petID)
	err = utils.Verify(petInfo.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = petService.DeletePet(petInfo)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdatePetInfo
// @Tags      PetInfo
// @Summary   更新宠物信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     petID  path      int             true  "宠物ID"
// @Param     data  body      pet.PetInfo            true  "宠物信息"
// @Success   200   {object}  response.Response{msg=string}  "更新宠物信息"
// @Router    /api/pet/:petID [put]
func (e *PetApi) UpdatePetInfo(c *gin.Context) {
	var petInfo pet.PetInfo
	petIdStr := c.Param("petID")
	petID, err := strconv.ParseUint(petIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage("宠物ID错误", c)
		return
	}

	err = c.ShouldBindJSON(&petInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	petInfo.ID = uint(petID)
	err = utils.Verify(petInfo.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(petInfo, utils.PetInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pet, err := petService.GetPetInfo(uint(petID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	petInfo.CreatedAt = pet.CreatedAt
	petInfo.UpdatedBy = utils.GetUserID(c)
	petInfo.DeletedAt = pet.DeletedAt
	petInfo.FamilyId = pet.FamilyId

	err = petService.UpdatePet(&petInfo)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetPetInfo
// @Tags      PetInfo
// @Summary   获取单一宠物信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     pet.PetInfo    true  "宠物ID"
// @Success   200   {object}  response.Response{data=pet.PetInfo,msg=string}  "获取单一宠物信息"
// @Router    /api/pet/:petID [get]
func (e *PetApi) GetPetInfo(c *gin.Context) {
	var petInfo pet.PetInfo
	petIdStr := c.Param("petID")
	petID, err := strconv.ParseUint(petIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage("宠物ID错误", c)
		return
	}
	petInfo.ID = uint(petID)
	err = utils.Verify(petInfo.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := petService.GetPetInfo(petInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

// GetPetInfoList
// @Tags      PetInfo
// @Summary   获取所有宠物列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=[]petResp.PetInfoResponse,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /api/pets [get]
func (e *PetApi) GetPetInfoList(c *gin.Context) {
	petList, err := petService.GetPetInfoList(utils.GetUserFamilyID(c))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(petList, "获取成功", c)
}
