<script setup lang="ts">
import petTodoCard from '@/components/home/pet-todo-card/index.vue';
import petAddCard from '@/components/home/pet-add-card/index.vue';
import { ref } from 'vue';
import { useDidShow } from '@tarojs/taro';

/** 设置页面属性 */
definePageConfig({
  navigationBarTitleText: '首页'
});

const pets = ref<Array<Pet.PetInfo>>([]);

const getPetsInfo = () => {
  pets.value = [{
    id: 1,
    name: '123',
    file: '',
    gendar: 0,  // 0 公 1 母
    sterilizedState: 1,  // 0 未绝育 1 已绝育
    petType: 0, // petTypes
  }]
}

useDidShow(()=>{
  console.log("show index")
  getPetsInfo()
})
</script>

<template>
  <basic-layout show-tab-bar>
    <custom-navbar title="首页" />
    <div v-if="pets.length" class="">
      <pet-todo-card v-for="pet in pets" :key="pet.id" :pet="pet"></pet-todo-card>
    </div>
    <div v-else class="position-absolute pos-top-50% translate-middle full-width">
      <pet-add-card></pet-add-card>
    </div>
  </basic-layout>
</template>

<style lang="scss">
.translate-middle {
  transform: translateY(-50%);
}

.full-width {
  width: calc(100% - 20px);
}
</style>
