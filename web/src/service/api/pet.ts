import { request } from '../request';
import { Pet } from "@/typings/pet";

/** 宠物注册 */
export async function addPet(data: Pet.PetInfo) {
  const resp = await request.post('/api/pet', data, {
    useErrMsg: false
  });
  return resp.success;
}

/** 宠物查询 */
export async function getAllPetsInfo() {
  const resp = await request.get<Array<Pet.PetInfo>>('/api/pets', {}, {
    useErrMsg: false
  });
  return resp.success;
}

/** 删除宠物 */
export async function deletePetInfo(id: number) {
  const resp = await request.delete(`/api/pet/${id}`, {}, {
    useErrMsg: false
  });
  return resp.success;
}
