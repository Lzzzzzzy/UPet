package family

import (
	"strconv"

	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/common/response"
	"github.com/Lzzzzzzy/UPet/server/model/family"
	"github.com/Lzzzzzzy/UPet/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FamilyApi struct{}

// Family
// @Tags      Family
// @Summary   获取家庭成员列表
// @Success   200   {object}  response.Response{msg=string}  "家庭成员列表"
// @Router    /api/family/members [get]
func (e *FamilyApi) GetFamilyMembers(c *gin.Context) {
	familyId := utils.GetUserFamilyID(c)
	members, err := familyService.GetFamilyMembers(familyId)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(members, "获取成功", c)
}

// Family
// @Tags      Family
// @Summary   删除家庭成员
// @Success   200   {object}  response.Response{msg=string}  "删除结果"
// @Router    /api/family/member/:memberId [delete]
func (e *FamilyApi) DeleteFamilyMember(c *gin.Context) {
	memberIdStr := c.Param("memberId")
	memberId, err := strconv.ParseUint(memberIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage("删除家庭成员ID错误", c)
		return
	}
	userInfo, err := userService.GetUserById(uint(memberId))
	if err != nil {
		global.GVA_LOG.Error("获取家庭成员信息失败!", zap.Error(err))
		response.FailWithMessage("获取家庭成员信息失败"+err.Error(), c)
		return
	}

	var familyInfo family.Family
	familyService.CreateFamily(&familyInfo)

	userInfo.FamilyId = familyInfo.ID
	userService.UpdateUser(userInfo)
}

// Family
// @Tags      Family
// @Summary   转让家庭管理权限
// @Success   200   {object}  response.Response{msg=string}  "转让结果"
// @Router    /api/family/member/:memberId [post]
func (e *FamilyApi) GrantFamilyManagePermission(c *gin.Context) {
	memberIdStr := c.Param("memberId")
	memberId, err := strconv.ParseUint(memberIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage("转让管理权给家庭成员ID错误", c)
		return
	}
	userInfo, err := userService.GetUserById(uint(memberId))
	if err != nil {
		global.GVA_LOG.Error("获取家庭成员信息失败!", zap.Error(err))
		response.FailWithMessage("获取家庭成员信息失败"+err.Error(), c)
		return
	}

	userInfo.IsAdmin = true
	userService.UpdateUser(userInfo)
}
