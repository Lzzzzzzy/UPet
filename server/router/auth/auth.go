package auth

import (
	"github.com/Lzzzzzzy/UPet/server/middleware"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (e *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	authRouter := Router.Group("auth").Use(middleware.OperationRecord())
	{
		authRouter.POST("", authApi.UserAuth) // 用户登录
	}
}
