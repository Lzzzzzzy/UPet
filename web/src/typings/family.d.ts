/** 家庭成员相关模块 */
declare namespace Familymember {
    /** 用户信息 */
    interface UserInfo {
      /** 用户id */
      id: number;
      /** 用户名 */
      name: string;
      /** 用户手机号 */
      isAdmin: boolean;
      /** 头像 */
      avatar: string;
    }
}
