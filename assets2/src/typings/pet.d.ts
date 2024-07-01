/** 宠物相关模块 */
declare namespace Pet {
  /** 宠物信息 */
  interface PetInfo {
    /** 宠物id */
    id: number;
    /** 宠物名 */
    petName: string;
    /** 关联用户id */
    userId: number;
    /** 宠物头像 */
    petAvatar: string;
  }

  /** 待办信息 */
  interface PetTodo {
    id: number;
    title: string;
    time: Date;
    remark: string;
    status: boolean;
  }
}
