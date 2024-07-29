package request

type UserInfo struct {
	Avatar   string `json:"avatar" form:"avatar"`     // 用户头像url
	NickName string `json:"nickname" form:"nickname"` // 昵称
}
