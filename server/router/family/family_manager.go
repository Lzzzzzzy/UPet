package family

import (
	"github.com/Lzzzzzzy/UPet/server/middleware"

	"github.com/gin-gonic/gin"
)

type FamilyRouter struct{}

func (e *FamilyRouter) InitFamilyRouter(Router *gin.RouterGroup) {
	familyRouter := Router.Group("family").Use(middleware.OperationRecord())
	{
		familyRouter.POST("/member/:memberId", familyApi.GrantFamilyManagePermission) // 转让家庭管理权限
		familyRouter.GET("/members", familyApi.GetFamilyMembers)                      // 获取家庭成员信息
		familyRouter.DELETE("/family/member/:memberId", familyApi.DeleteFamilyMember) // 删除家庭成员
	}
}
