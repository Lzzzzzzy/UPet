<script setup lang="ts">
import { ref } from 'vue';
import { useDidShow } from '@tarojs/taro';
import fullCalendar from '@/components/calendar/full-calendar/index.vue';
import petTodayTodos from '@/components/pet-today-todos/index.vue';

definePageConfig({
  navigationBarTitleText: '日历'
});

const pets = ref<Array<Pet.PetInfo>>([]);
const firstPet = ref<Pet.PetInfo>();

const getPetsInfo = () => {
  pets.value = [{
    id: 1,
    petName: '123',
    petAvatar: '',
    userId: 1,
  },
  {
    id: 2,
    petName: '1234',
    petAvatar: '',
    userId: 1,
  },{
    id: 3,
    petName: '1123',
    petAvatar: '',
    userId: 1,
  },{
    id: 4,
    petName: '1123',
    petAvatar: '',
    userId: 1,
  }
]}

useDidShow(()=>{
  getPetsInfo()
  if (pets.value.length) {
    firstPet.value = pets.value[0];
  }
})
</script>
<template>
  <basic-layout show-tab-bar>
    <custom-navbar title="日历" />
    <full-calendar></full-calendar>
    <nut-divider />
    <pet-today-todos :pets="pets"></pet-today-todos>
  </basic-layout>
</template>

<style lang="scss"></style>
