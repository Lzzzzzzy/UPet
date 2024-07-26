package user

import (
	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/user"
)

type UserService struct{}

// @function: CreateUser
// @description: 创建用户
// @param: e model.User
// @return: err error
func (userService *UserService) CreateUser(e user.User) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}
