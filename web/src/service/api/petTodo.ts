import { request } from '../request';
import { Pet } from "@/typings/pet";


interface paginationPetTodosReq {
  page: number,
  pageSize: number,
  petId: number,
  date: string,
}

interface paginationPetTodosResp {
  page: number,
  pageSize: number,
  total: number,
  list: Array<Pet.PetTodo>,
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

/** 分页查询宠物待办 */
export async function getPetTodosOnPagenation(pagination: paginationPetTodosReq) {
  const resp = await request.get<paginationPetTodosResp>('/api/pet-todos', pagination, {
    useErrMsg: false
  });
  return resp.success;
}
