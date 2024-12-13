package auth

import (
	"encoding/json"
	"fmt"

	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/user"
	"gorm.io/gorm"

	authReq "github.com/Lzzzzzzy/UPet/server/model/auth/request"
	authResp "github.com/Lzzzzzzy/UPet/server/model/auth/response"
	"github.com/Lzzzzzzy/UPet/server/model/common/response"
	"github.com/Lzzzzzzy/UPet/server/utils"
	"github.com/gin-gonic/gin"
)

type AuthApi struct{}

// Auth
// @Tags      Auth
// @Summary   微信第三方登录
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authReq.AuthInfo            true  "微信签发的code"
// @Success   200   {object}  response.Response{msg=string}  "用户登录或注册"
// @Router    /auth [post]
func (e *AuthApi) UserAuth(c *gin.Context) {
	var authInfo authReq.AuthInfo
	err := c.ShouldBindJSON(&authInfo)
	global.GVA_LOG.Info(fmt.Sprintf("微信用户登录code: %s", authInfo.Code))
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("微信用户登录获取参数失败: %e", err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(authInfo, utils.AuthVerify)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("微信用户登录参数验证失败: %e", err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	resp, err := authService.MiniprogramAuth(authInfo.Code)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("微信用户登录失败: %e", err))
		response.FailWithMessage("微信用户登录失败", c)
		return
	}
	jsonStr, err := json.Marshal(resp)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("marshal error: %e", err))
		return
	}
	global.GVA_LOG.Info(fmt.Sprintf("微信用户登录返回struct: %s", string(jsonStr)))
	var user *user.User
	user, err = userService.GetUserByOpenId(resp.Openid)
	if err != nil {
		if err == gorm.ErrRecordNotFound { // 数据库没查询到数据时注册
			global.GVA_LOG.Info("新用户，开始注册")
			user, err = registerService.RegisterUser(resp.Openid, resp.Unionid)
			if err != nil {
				global.GVA_LOG.Error(fmt.Sprintf("微信用户注册失败: %e", err))
				response.FailWithMessage("微信用户注册失败", c)
				return
			}
		} else {
			global.GVA_LOG.Error(fmt.Sprintf("微信用户注册查询是否存在时失败, %e", err))
			response.FailWithMessage("微信用户注册查询是否存在时失败", c)
			return
		}
	}

	token, expire, err := jwtService.CreateToken(user)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("jwt创建失败, %e", err))
		response.FailWithMessage("jwt创建失败", c)
		return
	}
	global.GVA_LOG.Info(fmt.Sprintf("用户登录成功, token: %s", token))
	response.OkWithDetailed(authResp.LoginResponse{
		User:      authResp.UserInfo{NickName: user.NickName, Avatar: user.Avatar, ID: user.ID, IsAdmin: user.IsAdmin},
		Token:     token,
		ExpiresAt: expire,
	}, "登录成功", c)
}
