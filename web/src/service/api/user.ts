import { request } from '../request';
import { login, redirectTo } from '@tarojs/taro';
import { localStg } from '@/utils';

interface UserInfo {
  avatar: string,
  nickname: string,
  isAdmin: boolean,
  id: number,
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
  const resp = await request.put<UserInfo>('/api/user', data, {
    useErrMsg: false
  });
  return resp.success;
}

export async function userLogin(redirectUrl?: string) {
  const res = await login();
  if (res.code) {
    const authResp = await userAuth(res.code);
    localStg.set("token", authResp!.token, authResp!.expireAt);
    localStg.set("userInfo", authResp!.user, authResp!.expireAt);
  }
}

export async function getUserInfo(userId: number | string) {
  const resp = await request.get(`/api/user/${userId}`, {}, {
    useErrMsg: false
  });
  return resp.success;
}
