package petTodo

import (
	"github.com/Lzzzzzzy/UPet/server/middleware"

	"github.com/gin-gonic/gin"
)

type PetTodoRouter struct{}

func (e *PetTodoRouter) InitPetTodoRouter(Router *gin.RouterGroup) {
	petRouter := Router.Group("petTodo").Use(middleware.OperationRecord())
	{
		petRouter.POST("pet-todo", petTodoApi.CreatePetInfo)              // 创建宠物
		petRouter.PUT("pet-todo/:petTodoID", petTodoApi.UpdatePetInfo)    // 更新宠物
		petRouter.DELETE("pet-todo/:petTodoID", petTodoApi.DeletePetInfo) // 删除宠物
		petRouter.GET("pet-todo/:petTodoID", petTodoApi.GetPetInfo)       // 获取单一宠物信息
		petRouter.GET("pet-todos", petTodoApi.GetPetInfoList)             // 获取宠物列表
	}
}
