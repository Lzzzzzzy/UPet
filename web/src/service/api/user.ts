import { request } from '../request';

interface UserInfo {
  avatar: string,
  nickname: string,
}

interface RegisterResp {
  user: UserInfo;
  token: string;
  expireAt: number;
}

/** 用户注册 */
export async function userAuth(code: string) {
  const resp = await request.post<RegisterResp>('/api/auth', { code }, {
    useErrMsg: false
  });
  return resp.success;
}

/** 完善用户信息 */
export async function userInfoComplete(data: any) {
  const resp = await request.put('/api/user', data, {
    useErrMsg: false
  });
  return resp.success;
}
