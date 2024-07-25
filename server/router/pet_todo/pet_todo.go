package petTodo

import (
	"github.com/Lzzzzzzy/UPet/server/middleware"

	"github.com/gin-gonic/gin"
)

type PetTodoRouter struct{}

func (e *PetTodoRouter) InitPetTodoRouter(Router *gin.RouterGroup) {
	petRouter := Router.Group("api").Use(middleware.OperationRecord())
	{
		petRouter.POST("pet-todo", petTodoApi.CreatePetTodo)              // 创建宠物待办
		petRouter.PUT("pet-todo/:petTodoID", petTodoApi.UpdatePetTodo)    // 更新宠物待办
		petRouter.DELETE("pet-todo/:petTodoID", petTodoApi.DeletePetTodo) // 删除宠物待办
		petRouter.GET("pet-todo/:petTodoID", petTodoApi.GetPetTodo)       // 获取单一宠物待办信息
		petRouter.GET("pet-todos", petTodoApi.GetPetTodoList)             // 获取宠物待办列表
	}
}
