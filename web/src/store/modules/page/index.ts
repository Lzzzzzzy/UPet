import { defineStore } from 'pinia';

interface PageParams {
  /** todo搜索时点击的date */
  todoDate: string,
}

export const usePageStore = defineStore('page-store', {
  state: (): PageParams => ({
    // userInfo: getUserInfo(),
    // token: getToken()
    todoDate: ""
  }),
  getters: {
    /** 获取todo date */
    getTodoDate: state => state.todoDate,
    /**是否有todo date */
    hasTodoDate: state => !!state.todoDate,
  },
  actions: {
    /** 设置todo date */
    setTodoDate(date: string) {
      this.todoDate = date;
    },
    /**清空todo date */
    removeTodoDate() {
      this.todoDate = "";
    },
  },
});
