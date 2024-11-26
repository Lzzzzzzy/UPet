package user

import (
	"strconv"

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

// User
// @Tags      User
// @Summary   查询用户信息
// @Param     id  path      int            true  "用户id"
// @Success   200   {object}  response.Response{msg=string}  "用户信息"
// @Router    /api/user/:id [get]
func (e *UserApi) GetUser(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage("用户ID错误", c)
		return
	}
	user, err := userService.GetUserById(uint(userId))
	if err != nil {
		response.FailWithMessage("查询用户数据错误", c)
		return
	}
	response.OkWithDetailed(userReq.UserInfo{NickName: user.NickName, Avatar: user.Avatar, Id: user.ID, IsAdmin: user.IsAdmin}, "查询成功", c)
}
