package request

type AuthInfo struct {
	Code string `json:"code" form:"code"` // 微信签发的code
}
