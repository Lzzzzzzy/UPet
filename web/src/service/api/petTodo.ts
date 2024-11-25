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
 * 更新宠物待办完成情况
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

/**
 * 搜索宠物待办信息
 *
 * @param {string} content - Need search content.
 * @return {Promise} A promise that resolves to the success status of the update operation.
 */
export async function searchPetTodoInfo(content: string) {
  const resp = await request.get(`/api/pet-todos/infos`, { content }, {
    useErrMsg: false
  });
  return resp.success;
}

/** 编辑宠物待办 */
export async function editPetTodo(data: petTodo, id: number) {
  const resp = await request.put(`/api/pet-todo/${id}`, data, {
    useErrMsg: false
  });
  return resp.success;
}

/** 删除宠物待办 */
export async function deletePetTodo(id: number) {
  const resp = await request.delete(`/api/pet-todo/${id}`, {}, {
    useErrMsg: false
  });
  return resp.success;
}

/**
 * 根据id查询宠物待办信息
 *
 * @param {number} todoId - Need search todo id.
 * @return {Promise} A promise that resolves to the success status of the update operation.
 */
export async function getPetTodoInfoById(todoId: number) {
  const resp = await request.get(`/api/pet-todo/${todoId}`, null, {
    useErrMsg: false
  });
  return resp.success;
}


/**
 * 删除宠物待办信息
 *
 * @param {number} todoId - Need search todo id.
 * @return {Promise} A promise that resolves to the success status of the update operation.
 */
export async function deletePetTodoInfoById(todoId: number) {
  const resp = await request.delete(`/api/pet-todo/${todoId}`, null, {
    useErrMsg: false
  });
  return resp.success;
}
