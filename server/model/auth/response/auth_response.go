package response

type UserInfo struct {
	NickName string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像
}

type LoginResponse struct {
	User      UserInfo `json:"user"`
	Token     string   `json:"token"`
	ExpiresAt int64    `json:"expiresAt"`
}
