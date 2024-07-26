package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Lzzzzzzy/UPet/server/global"
	authResp "github.com/Lzzzzzzy/UPet/server/model/auth/response"
	"go.uber.org/zap"
)

type AuthService struct{}

func (authService *AuthService) MiniprogramAuth(code string) (*authResp.WechatAuthInfo, error) {
	var resStruct authResp.WechatAuthInfo

	baseUrl := "https://api.weixin.qq.com/sns/jscode2session"
	wechatConfig := global.GVA_CONFIG.Wechat.Miniprogram
	appid := wechatConfig.AppId
	secret := wechatConfig.AppSecret
	url := baseUrl + "?appid=" + appid + "&secret=" + secret + "&js_code=" + code + "&grant_type=authorization_code"
	global.GVA_LOG.Info(fmt.Sprintf("微信code获取用户信息url: %s", url))
	resp, err := http.Get(url)
	if err != nil {
		global.GVA_LOG.Error("使用微信code获取用户信息失败!", zap.Error(err))
		return &resStruct, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("解析微信resp body失败!", zap.Error(err))
		return &resStruct, err
	}
	err = json.Unmarshal(body, &resStruct)
	if err != nil {
		global.GVA_LOG.Error("微信resp返回内容解析失败!", zap.Error(err))
		return &resStruct, err
	}
	if resStruct.Errcode != 0 {
		global.GVA_LOG.Error(fmt.Sprintf("微信小程序登录失败! 返回信息: %v\n", resStruct))
		return &resStruct, errors.New("微信登录失败")
	}
	return &resStruct, nil
}
