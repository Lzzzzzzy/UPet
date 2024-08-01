import type { Dayjs } from "dayjs";
/** 宠物相关模块 */
declare namespace Pet {
  /** 宠物信息 */
  interface PetInfo {
    /** 宠物id */
    id?: number;
    /** 宠物名 */
    name: string;
    /** 关联用户id */
    familyId?: number;
    /** 宠物头像 */
    avatar: string;
    /** 出生日期 */
    birthday?: Date | string | Dayjs;
    /** 性别 */
    gender: number;
    /** 绝育状态 */
    sterilizedState: number;
    /** 种类 */
    category: number;
  }

  /** 待办信息 */
  interface PetTodo {
    id?: number;
    title: string;
    todoTime: Date | string | Dayjs;
    remark: string;
    complete: boolean;
  }
}
