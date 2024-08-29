import { request } from '../request';


/** 获取用户家庭成员 */
export async function getFamilyMemberList() {
  const resp = await request.put('/api/family/member', {
    useErrMsg: false
  });
  return resp.success;
}
