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

const showWeek = ref("month");
</script>
<template>
  <!-- <basic-layout show-tab-bar>
    <custom-navbar title="日历" />
     <vue-hash-calendar :show-week-view="showWeek" />
    <nut-button @click="showWeek = !showWeek">点击切换周、月显示</nut-button>
  </basic-layout> -->
  <mpCalendar
    backgroundText
    class-name="select-mode"
    :mode="showWeek"
  />
  <nut-button @click="showWeek =='month' ? 'week' : 'month'">点击切换周、月显示</nut-button>
</template>

<style lang="scss"></style>
