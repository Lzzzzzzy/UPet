<script setup lang="ts">
import { ref } from 'vue';
import calendar from '@/components/calendar/index.vue';

const currentDate = ref(new Date());
const dotInfos = (date: Array<string>) => {
    return {"2024-07-02": ["red", "black"], "2024-07-03": ["green"]}
}
const showWeek = ref(false);
</script>
<template>
    <!-- <basic-layout show-tab-bar>
        <custom-navbar title="首页" />
        <calendar></calendar>
  </basic-layout> -->
  <div class="mt-100px">
      <calendar v-model="currentDate" :get-dot-info-func="dotInfos" :show-week="showWeek">
        <template #header>
          <div @click="showWeek = !showWeek">切换日历</div>
        </template>
      </calendar>
  </div>
  <p>选中时间是: {{ currentDate }}</p>
</template>
