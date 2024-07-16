<script setup lang="ts">
import { ref } from "vue";
import { switchTab, eventCenter } from '@tarojs/taro';
import dayjs from "dayjs";
import { useAppStore } from '@/store';
import { localStg } from '@/utils';

definePageConfig({
  navigationBarTitleText: '搜索'
});

const searchText = ref("")

const searchedTodos = ref<Array<Pet.PetTodo>>([]);
const search = () => {
  if (!searchText.value) {
    return;
  }
  searchedTodos.value = [
      {"id": 1, "time": "2024-07-13 08:00", "title": "换水", "remark": "全部两个碗里的倒矿泉水纯净水蒸馏水都行,就是不能用自来水", "complete": false},
      {"id": 2, "time": "2024-07-13 12:00", "title": "换粮", "remark": "全换", "complete": false},
      {"id": 3, "time": "2024-07-13 18:00", "title": "铲屎", "remark": "", "complete": false},
      {"id": 4, "time": "2024-07-13 12:43", "title": "换水", "remark": "123123123123123123123123aasdasdasdsasadsadsdsdsadsdasqeqweqweqeqewqeqweqeqweq", "complete": false},
      {"id": 5, "time": "2024-07-13 18:00", "title": "铲屎", "remark": "", "complete": true},
      {"id": 6, "time": "2024-07-13 18:00", "title": "铲屎", "remark": "", "complete": false},
      {"id": 7, "time": "2024-07-13 18:00", "title": "铲屎", "remark": "", "complete": false},
      {"id": 8, "time": "2024-07-13 18:00", "title": "铲屎", "remark": "", "complete": false},
      {"id": 9, "time": "2024-07-13 18:00", "title": "铲屎", "remark": "", "complete": false},
  ] 
}

const skipToIndex = (date: Date) => {
  const todoDate = dayjs(date).format("YYYY-MM-DD")
  localStg.set("todoDate", todoDate)
  const appStore = useAppStore();
  const url = "/pages/index/index"
  appStore.setActiveTab(url);
  switchTab({ url });
}
</script>
<template>
  <basic-layout show-tab-bar>
    <custom-navbar title="搜索" />
    <div class="w-full text-30px fixed-header">
      <nut-searchbar v-model="searchText" @search="search" autofocus :clearable="false" placeholder="请输入检索内容"></nut-searchbar>
    </div>
    <div class="pb-45px pt-49px">
      <div v-for="(todo, index) in searchedTodos" :key="index">
        <div class="flex flex-col p-10px bg-#ffffff search-content-border" @click="skipToIndex(todo.time)">
          <div class="text-16px">{{ todo.title }}</div>
          <div class="text-12px mt-5px break-words" v-if="todo.remark">{{ todo.remark }}</div>
          <div class="mt-5px text-#717171 text-12px"> {{ todo.time }} </div>
        </div>
      </div>
    </div>
  </basic-layout>
</template>

<style lang="scss">
.search-content-border {
  border-top: 0.5px solid #ebebeb;
}

.fixed-header {
  position: fixed;
  top: 90px;
  width: 100%;
  background-color: #ffffff;
  overflow: hidden;
  z-index: 99;
}
</style>
