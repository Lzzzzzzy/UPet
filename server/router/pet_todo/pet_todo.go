package petTodo

import (
	"github.com/Lzzzzzzy/UPet/server/middleware"

	"github.com/gin-gonic/gin"
)

type PetTodoRouter struct{}

func (e *PetTodoRouter) InitPetTodoRouter(Router *gin.RouterGroup) {
	petTodoRouter := Router.Group("pet-todo").Use(middleware.OperationRecord())
	petTodoRouterWithoutRecord := Router.Group("pet-todo")
	petsTodoRouterWithoutRecord := Router.Group("pet-todos")
	{
		petTodoRouter.POST("", petTodoApi.CreatePetTodo)              // 创建宠物待办
		petTodoRouter.PUT("/:petTodoID", petTodoApi.UpdatePetTodo)    // 更新宠物待办
		petTodoRouter.DELETE("/:petTodoID", petTodoApi.DeletePetTodo) // 删除宠物待办
	}
	{
		petTodoRouterWithoutRecord.GET("/:petTodoID", petTodoApi.GetPetTodo) // 获取单一宠物待办信息
	}
	{
		petsTodoRouterWithoutRecord.GET("", petTodoApi.GetPetTodoList) // 获取宠物待办列表
	}
}
