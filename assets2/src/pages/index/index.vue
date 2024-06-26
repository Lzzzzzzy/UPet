<script setup lang="ts">
import petTodoCard from '@/components/home/pet-todo-card/index.vue';
import { ref } from 'vue';

/** 设置页面属性 */
definePageConfig({
  navigationBarTitleText: '首页'
});

const pets = ref<Array<Pet.PetInfo>>([]);
</script>

<template>
  <basic-layout show-tab-bar>
    <custom-navbar title="首页" />
    <div v-if="pets.length">
      <pet-todo-card v-for="pet in pets" :key="pet.id" :pet="pet"></pet-todo-card>
    </div>
    <div v-else></div>
  </basic-layout>
</template>
