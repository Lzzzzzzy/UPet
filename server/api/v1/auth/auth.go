package auth

import (
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
// @Param     data  body      auth.AuthInfo            true  "微信签发的code"
// @Success   200   {object}  response.Response{msg=string}  "用户登录或注册"
// @Router    /api/auth [post]
func (e *AuthApi) UserAuth(c *gin.Context) {
	var authInfo authReq.AuthInfo
	err := c.ShouldBindJSON(&authInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(authInfo, utils.AuthVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	resp, err := authService.MiniprogramAuth(authInfo.Code)
	if err != nil {
		response.FailWithMessage("微信用户登录失败", c)
		return
	}
	user, err := registerService.RegisterUser(resp.Openid, resp.Unionid)
	if err != nil {
		response.FailWithMessage("微信用户注册失败", c)
		return
	}
	token, expire, err := jwtService.CreateToken(user)
	if err != nil {
		response.FailWithMessage("jwt创建失败", c)
		return
	}
	response.OkWithDetailed(authResp.LoginResponse{
		User:      authResp.UserInfo{NickName: user.NickName, Avatar: user.Avatar},
		Token:     token,
		ExpiresAt: expire,
	}, "登录成功", c)
}
