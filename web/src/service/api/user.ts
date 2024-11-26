import { request } from '../request';
import { login, redirectTo } from '@tarojs/taro';
import { localStg } from '@/utils';

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

export function userLogin(redirectUrl?: string) {
  login({
    success: function (res: any) {
      if (res.code) {
        //发起网络请求
        userAuth(res.code).then((data: any) => {
          localStg.set("token", data.token, data.expiresAt);
          localStg.set("userInfo", data.user, null);
          if (!(data?.user.avatar && data?.user.nickname)) {
            // 跳转到用户注册页面
            redirectTo({
              url: `/package/package-register/index?redirectTo=${redirectUrl}`
            });
          }
        })
      } else {
        console.log('登录失败！' + res.errMsg)
      }
    }
  })
}

export async function getUserInfo(userId: number | string) {
  const resp = await request.get(`/api/user/${userId}`, {}, {
    useErrMsg: false
  });
  return resp.success;
}
