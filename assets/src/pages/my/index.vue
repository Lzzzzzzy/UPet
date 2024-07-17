<script setup lang="ts">
import { navigateTo } from '@tarojs/taro';

definePageConfig({
  navigationBarTitleText: '个人中心'
});
const menus = [{"name": "爱宠档案", "url": "/package/package-pet-management/index"}, {"name": "爱宠共养", "url": "/package/package-pet-ownership/index"}, {"name": "联系我们", "url": "/package/package-contact/index"}]
const gotoMenu = (url: string) => {
  navigateTo({url})
}
</script>
<template>
  <basic-layout show-tab-bar>
    <custom-navbar title="个人中心" />
    <div v-for="(menu, index) in menus" :key="index" @click="gotoMenu(menu.url)">
      <nut-cell :title="menu.name">
         <template #desc>
           <span class="i-local-more text-25px"></span>
         </template>
       </nut-cell>
    </div>
  </basic-layout>
</template>
