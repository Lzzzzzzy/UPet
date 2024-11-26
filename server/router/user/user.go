package user

import (
	"github.com/Lzzzzzzy/UPet/server/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (e *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	authRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		authRouter.PUT("", userApi.UserInfoComplete) // 完善用户信息
		authRouter.GET("/:id", userApi.GetUser)      // 查询用户信息
	}
}
