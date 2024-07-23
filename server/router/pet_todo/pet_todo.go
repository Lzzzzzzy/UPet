package petTodo

import (
	"github.com/Lzzzzzzy/UPet/server/middleware"

	"github.com/gin-gonic/gin"
)

type PetTodoRouter struct{}

func (e *PetTodoRouter) InitPetTodoRouter(Router *gin.RouterGroup) {
	petRouter := Router.Group("petTodo").Use(middleware.OperationRecord())
	{
		petRouter.POST("pet-todo", petTodoApi.CreatePetInfo)              // 创建宠物待办
		petRouter.PUT("pet-todo/:petTodoID", petTodoApi.UpdatePetInfo)    // 更新宠物待办
		petRouter.DELETE("pet-todo/:petTodoID", petTodoApi.DeletePetInfo) // 删除宠物待办
		petRouter.GET("pet-todo/:petTodoID", petTodoApi.GetPetInfo)       // 获取单一宠物待办信息
		petRouter.GET("pet-todos", petTodoApi.GetPetInfoList)             // 获取宠物待办列表
	}
}
