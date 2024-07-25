package petTodo

import (
	api "github.com/Lzzzzzzy/UPet/server/api/v1"
)

type RouterGroup struct {
	PetTodoRouter
}

var (
	petTodoApi = api.ApiGroupApp.PetTodoApiGroup
)
