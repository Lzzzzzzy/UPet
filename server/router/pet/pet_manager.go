package pet

import (
	"github.com/Lzzzzzzy/UPet/server/middleware"

	"github.com/gin-gonic/gin"
)

type PetRouter struct{}

func (e *PetRouter) InitPetRouter(Router *gin.RouterGroup) {
	petRouter := Router.Group("pet").Use(middleware.OperationRecord())
	petRouterWithoutRecord := Router.Group("pet")
	{
		petRouter.POST("pet", petApi.CreatePetInfo)          // 创建宠物
		petRouter.PUT("pet/:petID", petApi.UpdatePetInfo)    // 更新宠物
		petRouter.DELETE("pet/:petID", petApi.DeletePetInfo) // 删除宠物
	}
	{
		petRouterWithoutRecord.GET("pet/:petID", petApi.GetPetInfo) // 获取单一宠物信息
		petRouterWithoutRecord.GET("pets", petApi.GetPetInfoList)   // 获取宠物列表
	}
}
