package user

import (
	"github.com/Lzzzzzzy/UPet/server/model/common/response"
	userReq "github.com/Lzzzzzzy/UPet/server/model/user/request"
	"github.com/Lzzzzzzy/UPet/server/utils"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// User
// @Tags      User
// @Summary   完善用户信息
// @accept    application/json
// @Produce   application/json
// @Param     data  body      user.UserInfo            true  "用户信息"
// @Success   200   {object}  response.Response{msg=string}  "用户信息保存结果"
// @Router    /api/user [put]
func (e *UserApi) UserInfoComplete(c *gin.Context) {
	var userInfo userReq.UserInfo
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(userInfo, utils.UserInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	user, err := userService.GetUserById(userId)
	if err != nil {
		response.FailWithMessage("用户信息完善失败, 查询用户数据错误", c)
		return
	}
	user.Avatar = userInfo.Avatar
	user.NickName = userInfo.NickName
	err = userService.UpdateUser(user)

	if err != nil {
		response.FailWithMessage("用户信息完善失败, 更新用户数据错误", c)
		return
	}
	response.Ok(c)
}
