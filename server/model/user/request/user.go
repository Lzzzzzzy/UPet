package request

type UserInfo struct {
	Id       uint   `json:"id" form:"id"`                                 // 用户ID
	Avatar   string `json:"avatar" form:"avatar"`                         // 用户头像url
	NickName string `json:"nickname" form:"nickname"`                     // 昵称
	IsAdmin  bool   `json:"isAdmin" form:"isAdmin" gorm:"comment:是否为管理员"` // 是否为管理员
}
