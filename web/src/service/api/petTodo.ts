import { request } from '../request';
import { Pet } from "@/typings/pet";


interface petTodosReq {
  petId: number,
  date: string,
}

interface petTodo {
  id?: number,
  title: string,
  todoTime: string,
  remark: string,
  remind: boolean,
  remindDate: Array<string>,
  remindTime: string,
  complete: boolean,
  type: number,
  color: number,
}

/** 添加宠物待办 */
export async function addPetTodo(data: petTodo) {
  const resp = await request.post('/api/pet-todo', data, {
    useErrMsg: false
  });
  return resp.success;
}

/** 查询宠物待办 */
export async function getPetTodosOnPagenation(condition: petTodosReq) {
  const resp = await request.get<Array<Pet.PetTodo>>('/api/pet-todos', condition, {
    useErrMsg: false
  });
  return resp.success;
}

/** 查询宠物待办标记 */
export async function getPetTodosMark(dates: Pet.PetMarkParam) {
  const resp = await request.post<Array<Pet.PetTodo>>('/api/pet-todos/mark', dates, {
    useErrMsg: false
  });
  return resp.success;
}

/**
 * Updates the complete status of a todo item.
 *
 * @param {number} todoId - The ID of the todo item to update.
 * @param {boolean} complete - The new complete status of the todo item.
 * @return {Promise} A promise that resolves to the success status of the update operation.
 */
export async function updateTodoCompleteStatus(todoId: number, complete: boolean) {
  const resp = await request.put(`/api/pet-todo/${todoId}/complete`, { complete }, {
    useErrMsg: false
  });
  return resp.success;
}
