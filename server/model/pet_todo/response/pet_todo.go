package response

import (
	petTodo "github.com/Lzzzzzzy/UPet/server/model/pet_todo"
)

type PetTodoInfoResponse struct {
	PetTodo petTodo.PetTodoInfo `json:"pet"`
}
