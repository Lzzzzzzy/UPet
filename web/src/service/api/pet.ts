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
export async function getPetsInfo() {
  const resp = await request.get<Pet.PetInfo>('/api/pet', {
    useErrMsg: false
  });
  return resp.success;
}
