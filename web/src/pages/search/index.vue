<script setup lang="ts">
import { ref } from "vue";
import { switchTab, eventCenter } from '@tarojs/taro';
import dayjs from "dayjs";
import { useAppStore } from '@/store';
import { Pet } from "@/typings/pet";
import { searchPetTodoInfo } from "@/service/api";
import { RichText } from "@tarojs/components";


definePageConfig({
  navigationBarTitleText: '搜索'
});

const searchText = ref("")

const searchedTodos = ref<Array<Pet.PetTodo>>([]);
const search = async () => {
  if (!searchText.value) {
    return;
  }
  searchedTodos.value = await searchPetTodoInfo(searchText.value);
}

const skipToIndex = (petTodo: Pet.PetTodo) => {
  const appStore = useAppStore();
  const url = "/pages/index/index"
  appStore.setActiveTab(url);
  switchTab({ url, success: () => {
    eventCenter.trigger("selectPet", petTodo.petId)
    eventCenter.trigger("jumpToDate", dayjs(petTodo.remindTime).toDate())
  } });
}
</script>

<template>
  <basic-layout show-tab-bar>
    <custom-navbar title="搜索" />
    <div class="w-full text-30px fixed-header">
      <nut-searchbar v-model="searchText" @search="search" autofocus :clearable="false" placeholder="请输入检索内容">
        <template #rightout> 
          <div class="i-local-search text-black" @click="search"></div>
         </template>
      </nut-searchbar>
    </div>
    <div class="pb-45px pt-50px">
      <div v-for="(todo, index) in searchedTodos" :key="index">
        <div class="flex flex-col p-10px bg-#ffffff search-content-border" @click="skipToIndex(todo)">
          <div class="text-16px">{{ todo.title }}</div>
          <!-- <div class="text-12px mt-5px break-words" v-if="todo.remark">{{ todo.remark }}</div> -->
          <rich-text :nodes="todo.remark" v-if="todo.remark" class="text-12px break-words pt-10px"></rich-text>
          <div class="mt-5px text-#717171 text-12px"> {{ todo.todoTime }} </div>
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
