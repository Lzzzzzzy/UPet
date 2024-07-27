package user

import (
	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/family"
	"github.com/Lzzzzzzy/UPet/server/model/user"
	"go.uber.org/zap"

	familyServ "github.com/Lzzzzzzy/UPet/server/service/family"
)

type RegisterService struct{}

func (registerService *RegisterService) RegisterUser(openid, unionid string) (*user.User, error) {
	var familyService = familyServ.FamilyService{}
	var userService = UserService{}
	var familyInfo family.Family
	familyService.CreateFamily(&familyInfo)

	var user user.User
	user.FamilyId = familyInfo.ID
	user.OpenId = openid
	user.UnionId = unionid
	user.NickName = "微信用户"
	user.IsAdmin = true
	err := userService.CreateUser(user)
	if err != nil {
		global.GVA_LOG.Error("创建用户失败!", zap.Error(err))
		return &user, err
	}
	return &user, nil
}