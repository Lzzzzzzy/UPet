<script setup lang="ts">
import { ref } from 'vue';
import { useDidShow } from '@tarojs/taro';
import noPetRemind from '@/components/home/add-pet-remind/index.vue';
import petTodayTodos from '@/components/pet-today-todos/index.vue';
import calendar from '@/components/calendar/index.vue';

/** 设置页面属性 */
definePageConfig({
  navigationBarTitleText: '首页'
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

const currentDate = ref(new Date());
const dotInfos = (date: Array<string>) => {
    return {"2024-07-02": ["red", "black"], "2024-07-03": ["green"], "2024-07-30": ["green"], "2024-07-31": ["orange"]}
}
const showWeek = ref(true);
</script>

<script lang="ts">
export default {
  name: 'Index',
}
</script>

<template>
  <basic-layout show-tab-bar>
    <custom-navbar title="首页" />
    <div v-if="pets.length">
      <div class="bg-#fff">
        <calendar v-model="currentDate" :get-dot-info-func="dotInfos" :show-week="showWeek">
          <template #header>
            <div @click="showWeek = !showWeek">切换日历</div>
          </template>
        </calendar>
      </div>
      <pet-today-todos :pets="pets"></pet-today-todos>
    </div>
    <div v-else class="position-absolute pos-top-50% translate-middle full-width">
      <no-pet-remind></no-pet-remind>
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
