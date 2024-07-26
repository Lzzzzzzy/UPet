package response

type WechatAuthInfo struct {
	Openid     string `json:"openid" form:"openid"`           // 微信小程序用户open id
	SessionKey string `json:"session_key" form:"session_key"` // 微信小程序返回的session
	Unionid    string `json:"unionid" form:"unionid"`         // 微信小程序用户union id
	Errcode    int32  `json:"errcode" form:"errcode"`         // 微信小程序返回的错误码
	ErrMsg     string `json:"errmsg" form:"errmsg"`           // 微信小程序返回的错误信息
}
