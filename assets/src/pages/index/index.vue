<template>
  <view>
    <calendar v-if="currentTab === 'calendar'"></calendar>
    <classSelection v-if="currentTab === 'self'"></classSelection>
    <tabbar :active="currentTab" @changeTab="handleChangeTab"></tabbar>
  </view>
</template>

<script setup>
import { ref } from 'vue';
import tabbar from '../../components/tabbar.vue';
import calendar from '../../pages/calendar/index.vue';
import classSelection from '../../pages/classSelection/index.vue';

const currentTab = ref("calendar");
const handleChangeTab = (tabName) => {
  console.log(tabName)
  currentTab.value = tabName;
}
</script>

<style lang="scss">
:root,
page {
  --main-color: #1E90FF;
  --nut-steps-base-title-font-size: 35rpx;
  --nut-steps-base-title-color: #000000;
  --nut-primary-color: var(--main-color);
  --nut-steps-wait-icon-bg-color: var(--main-color);
  --nut-tabs-titles-background-color: #ffffff;
  --nut-tabs-titles-item-font-size: 30rpx;
  --nut-grid-item-text-font-size: 30rpx;
  --nut-grid-item-content-padding: 16rpx;
  --nut-grid-item-text-margin: 0;
  --nut-grid-item-content-bg-color: #F5F5F5;
  --nut-badge-background-color: var(--main-color);
}
</style>
