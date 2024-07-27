package pet

import (
	"github.com/Lzzzzzzy/UPet/server/middleware"

	"github.com/gin-gonic/gin"
)

type PetRouter struct{}

func (e *PetRouter) InitPetRouter(Router *gin.RouterGroup) {
	petRouter := Router.Group("pet").Use(middleware.OperationRecord())
	petRouterWithoutRecord := Router.Group("pet")
	petsRouterWithoutRecord := Router.Group("pets")
	{
		petRouter.POST("", petApi.CreatePetInfo)          // 创建宠物
		petRouter.PUT("/:petID", petApi.UpdatePetInfo)    // 更新宠物
		petRouter.DELETE("/:petID", petApi.DeletePetInfo) // 删除宠物
	}
	{
		petRouterWithoutRecord.GET("/:petID", petApi.GetPetInfo) // 获取单一宠物信息
	}
	{
		petsRouterWithoutRecord.GET("", petApi.GetPetInfoList) // 获取宠物列表
	}
}
