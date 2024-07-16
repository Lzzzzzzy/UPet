<script setup lang="ts">
import { computed } from 'vue';
import { switchTab } from '@tarojs/taro';
import { useAppStore, useThemeStore } from '@/store';

const tabBar = {
  custom: true,
  color: '#000000',
  selectedColor: '#FF0000',
  list: [
    {
      pagePath: '/pages/index/index',
      text: '',
      icon: 'i-local-home',
    },
    {
      pagePath: '/pages/search/index',
      text: '',
      icon: 'i-local-search',
    },
    {
      pagePath: '/pages/my/index',
      text: '',
      icon: 'i-local-my',
    }
  ]
};

const themeStore = useThemeStore();
const theme = computed(() => themeStore.theme);
const themeVars = computed(() => themeStore.themeVars);

const appStore = useAppStore();
const activeTab = computed(() => appStore.getActiveTab);
function tabSwitch(item: any, url: string) {
  appStore.setActiveTab(url);
  switchTab({ url });
}
</script>
<script lang="ts">
export default {
  options: {
    addGlobalClass: true
  }
};
</script>
<template>
  <nut-config-provider :theme="theme" :theme-vars="themeVars">
    <nut-tabbar :model-value="activeTab" class="tabbar-bg" bottom safe-area-inset-bottom @tab-switch="tabSwitch" unactive-color="#222222" active-color="#222222">
      <nut-tabbar-item v-for="item in tabBar.list" :key="item.pagePath" :name="item.pagePath" :tab-title="item.text" :class="{'activate-tab': activeTab === item.pagePath}">
        <template #icon>
          <div class="text-20px" :class="item.icon"></div>
        </template>
      </nut-tabbar-item>
    </nut-tabbar>
  </nut-config-provider>
</template>

<style lang="scss">
.tabbar-bg {
  .nut-tabbar {
    // background: #E7E0CC;
    box-shadow: 0px -2px 6px rgba(0, 0, 0, 0.25);
    padding-top: 10px;
    height: auto;

    .nut-tabbar-item {
      .activate-tab {
        padding: 5px 20px;
        background: #f7daa1;
        border-radius: 20px;
      }

    }
  }
}
</style>
