import { defineStore } from 'pinia';
import { useRoutePath } from '@/composables';

interface AppState {
  /** 用户信息 */
  activeTab: string;
}

export const useAppStore = defineStore('app-store', {
  state: (): AppState => ({
    activeTab: useRoutePath()
  }),
  getters: {
    /** 是否登录 */
    getActiveTab: state => state.activeTab
  },
  actions: {
    setActiveTab(tab: string) {
      this.activeTab = tab;
    },
  }
});
