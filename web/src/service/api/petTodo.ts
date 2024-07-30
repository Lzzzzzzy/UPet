import { request } from '../request';
import { Pet } from "@/typings/pet";


interface paginationPetTodo {
  page: number,
  pageSize: number,
  petId: number,
  date: string,
}

/** 添加宠物待办 */
export async function addPetTodo(data: Pet.PetInfo) {
  const resp = await request.post('/api/pet', data, {
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
