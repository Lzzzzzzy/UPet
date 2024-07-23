package user

import "github.com/Lzzzzzzy/UPet/server/global"

type User struct {
	global.GVA_MODEL
	NickName string `json:"nickname" form:"nickname" gorm:"comment:昵称"`          // 昵称
	Phone    string `json:"phone" form:"phone" gorm:"comment:手机号"`               // 手机号
	Avatar   string `json:"avatar" form:"avatar" gorm:"comment:头像"`              // 头像
	OpenId   string `json:"openId" form:"openId" gorm:"comment:微信小程序openId"`     //微信小程序openId
	UnionId  string `json:"unionId" form:"unionId" gorm:"comment:微信公众平台UnionId"` //微信公众平台UnionId
	FamilyId uint   `json:"familyId" form:"familyId" gorm:"comment:familyId"`    //familyId
	IsAdmin  uint   `json:"isAdmin" form:"isAdmin" gorm:"comment:是否为管理员"`        // 是否为管理员
}
