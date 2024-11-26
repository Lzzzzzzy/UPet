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

	currentUserId := utils.GetUserID(c)

	err = userService.UpdateUserAdmin(currentUserId, uint(memberId))
	if err != nil {
		response.FailWithMessage("转让管理权给家庭成员ID错误", c)
		return
	}
	response.Ok(c)
}

// Family
// @Tags      Family
// @Summary   加入家庭
// @Success   200   {object}  response.Response{msg=string}  "加入结果"
// @Router    /api/family/join/:familyId [post]
func (e *FamilyApi) JoinFamily(c *gin.Context) {
	familyIdStr := c.Param("familyId")
	familyId, err := strconv.ParseUint(familyIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage("加入家庭错误", c)
		return
	}

	userId := utils.GetUserID(c)
	userInfo, err := userService.GetUserById(userId)
	if err != nil {
		global.GVA_LOG.Error("获取用户信息失败!", zap.Error(err))
		response.FailWithMessage("获取用户信息失败"+err.Error(), c)
		return
	}

	// 判断原家庭是否还有成员，如果有，且该成员是管理员时，将管理员权限转让给其他成员
	otherUsers, err := familyService.GetFamilyOtherMembers(userInfo.FamilyId, userId)
	if err != nil {
		global.GVA_LOG.Error("获取家庭其他成员失败!", zap.Error(err))
		response.FailWithMessage("获取家庭其他成员失败"+err.Error(), c)
		return
	}

	needRemoveFamily := false
	if len(*otherUsers) > 0 {
		// 转让管理员权限给其他成员
		users := *otherUsers
		err = userService.UpdateUserAdmin(userId, users[0].ID)
		if err != nil {
			global.GVA_LOG.Error("转让管理员权限失败!", zap.Error(err))
			response.FailWithMessage("转让管理员权限失败"+err.Error(), c)
		} else {
			needRemoveFamily = true
		}
	} 

	// 更新用户家庭
	userInfo.FamilyId = uint(familyId)
	userService.UpdateUser(userInfo)
	if needRemoveFamily {
		// 删除原家庭
		familyService.DeleteFamily(userInfo.FamilyId)
	}
	response.Ok(c)
}
