/** 用户相关模块 */
declare namespace Auth {
  /** 用户信息 */
  interface UserInfo {
    /** id */
    id: number;
    /** 用户头像 */
    avatar: string;
    /** 用户名 */
    nickname: string;
    /** 是否管理员 */
    isAdmin: boolean;
  }
}
