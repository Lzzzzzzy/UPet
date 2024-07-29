package response

type UserInfo struct {
	NickName string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像
	ID       uint   `json:"id"`
	IsAdmin  bool   `json:"isAdmin"` //  是否为管理员
}

type LoginResponse struct {
	User      UserInfo `json:"user"`
	Token     string   `json:"token"`
	ExpiresAt int64    `json:"expiresAt"`
}
