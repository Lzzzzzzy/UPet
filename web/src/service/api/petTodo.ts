import { request } from '../request';
import { Pet } from "@/typings/pet";


interface paginationPetTodo {
  page: number,
  pageSize: number,
  petId: number,
  date: string,
}

interface petTodo {
  id?: number,
  title: string,
  time: string,
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
export async function getPetTodosOnPagenation(pagination: paginationPetTodo) {
  const resp = await request.get<Array<Pet.PetTodo>>('/api/pet-todos', pagination, {
    useErrMsg: false
  });
  return resp.success;
}
