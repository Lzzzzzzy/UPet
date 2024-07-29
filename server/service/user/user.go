package user

import (
	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/user"
)

type UserService struct{}

// @function: CreateUser
// @description: 创建用户
// @param: e *model.User
// @return: err error
func (userService *UserService) CreateUser(e *user.User) (err error) {
	err = global.GVA_DB.Create(e).Error
	return
}

// @function: GetUserByOpenId
// @description: 通过openId查询用户
// @param: openId string
// @return: e *model.User err error
func (userService *UserService) GetUserByOpenId(openId string) (e *user.User, err error) {
	err = global.GVA_DB.Where("open_id = ?", openId).First(&e).Error
	return
}

// @function: GetUserById
// @description: 通过Id查询用户
// @param: id uint
// @return: e *model.User err error
func (userService *UserService) GetUserById(id uint) (e *user.User, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&e).Error
	return
}

// @function: UpdateUser
// @description: 更新用户信息
// @param: e model.User
// @return: err error
func (userService *UserService) UpdateUser(e *user.User) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}
