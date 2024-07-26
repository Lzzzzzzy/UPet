import { request } from '../request';

interface RegisterResp {
  user: string;
  token: string;
  expireAt: number;
}

/** 示例 */
export function userAuth(code: string) {
  return request.post<RegisterResp>('/api/auth', { code }, {
    useErrMsg: true
  });
}
