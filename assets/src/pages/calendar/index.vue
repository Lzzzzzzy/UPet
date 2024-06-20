<template>
  <view>
    <view class="index-page">
      <nut-calendar-card v-model="currentDate" :current="data.length-1">
        <template #bottom="{ day }">
          {{ day.date ===  todayDate ? "今天": "" }}
        </template>
      </nut-calendar-card>
    </view>
    <view class="subtitle">
      <span>毛孩子的今天</span>
      <nut-tabs v-model="currentTab" align="left" :animated-time="0">
        <nut-tab-pane v-for="tab in tabList" :key="tab.paneKey" :title="tab.title" :pane-key="tab.paneKey"></nut-tab-pane>
      </nut-tabs>
    </view>
    <view class="step-container">
      <nut-steps direction="vertical">
        <nut-step v-for="info in data" :key="info.id" :title="info.title" :content="info.content">
          <template #icon>
            <actionIcons :icon-type="info.type"></actionIcons>
          </template>
        </nut-step>
      </nut-steps>
    </view>
    <view class="record-button-container">
      <nut-button plain type="info">记一笔</nut-button>
    </view>
    <!-- <tabbar activate="calendar"></tabbar> -->
  </view>
</template>

<script setup>
import { ref, onBeforeMount } from 'vue';
import actionIcons from '../../components/action-icons.vue';
// import tabbar from '../../components/tabbar.vue';

const currentDate = ref();
const todayDate = ref();
const currentTab = ref(1);

onBeforeMount(()=>{
  currentDate.value = new Date();
  todayDate.value = new Date().getDate();
})

const tabList = [{title: "日常", paneKey: 1}, {title: "花销", paneKey: 2}, {title: "清理", paneKey: 3}]
const data = ref([{"title": "喂食", "content": "喂了一半", "id": 1, "type": "feed"}, {"title": "铲屎", "content": "没铲干净没铲干净没铲干净没铲干净没铲干净没铲干净没铲干净没铲干净没铲干净没铲干净没铲干净没铲干净", "id": 2, "type": "clean"}, {"title": "买猫粮", "content": "买了300g分装猫粮", "id": 3, "type": "purchase"}, {"title": "喂食", "content": "喂了一半", "id": 1, "type": "feed"}, {"title": "铲屎", "content": "", "id": 2, "type": "clean"}, {"title": "买猫粮", "content": "买了300g分装猫粮", "id": 3, "type": "purchase"}])
</script>

<style lang="scss">
.index-page {
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.step-container {
  padding: 3rpx 10rpx;
}

.nut-calendarcard-day-top {
  height: 0 !important;
  line-height: 0 !important;
}

.nut-calendarcard-day {
  height: 80rpx !important;
}

.nut-step-main {
  margin-bottom: 40rpx;
}

.subtitle {
  padding: 0 10rpx;
  margin: 10rpx 0 20rpx;
  font-size: 38rpx;
  font-weight: 600;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.custom-title {
  color: black;
  cursor: pointer;
}

.custom-title.active {
  color: red;
}

.nut-tab-pane {
  padding: 0 !important;
}

.record-button-container {
  position: fixed;
  right: 30rpx;
  bottom: 60rpx;
}
</style>
