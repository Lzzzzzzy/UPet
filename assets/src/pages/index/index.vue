<script setup lang="ts">
import { ref, onBeforeMount, watch, computed } from 'vue';
import { eventCenter } from '@tarojs/taro';
import noPetRemind from '@/components/home/add-pet-remind/index.vue';
import petTodosPage from '@/components/pet-todos/index.vue';
import calendar from '@/components/calendar/index.vue';
import { Pet } from "@/typings/pet";

/** 设置页面属性 */
definePageConfig({
  navigationBarTitleText: '首页'
});

/** 日历相关参数和方法 */
const currentDate = ref(new Date());
const calendarMode = ref("week");
const getDotInfos = (date: Array<string>) => {
    return {"2024-07-02": ["red", "black"], "2024-07-03": ["green"], "2024-07-30": ["green"], "2024-07-31": ["orange"]}
}

/** 宠物相关参数和方法 */
const currentPet = ref<Pet.PetInfo>();

onBeforeMount(() => {
  eventCenter.on("selectpet", (pet: Pet.PetInfo) => {
    currentPet.value = pet;
  });
  eventCenter.on("jumpToDate", (date: Date) => {
    console.log("jumpToDate");
    currentDate.value = date;
  })
});

// useDidShow(() => {
//   const pageStore = usePageStore();
//   if (pageStore.hasTodoDate) {
//     const todoDate = pageStore.getTodoDate;
//     currentDate.value = dayjs(todoDate).toDate();
//     pageStore.removeTodoDate();
//   }
// })

const pets = ref<Array<Pet.PetInfo>>([]);
const getPetsInfo = () => {
  return [{
    id: 1,
    name: '旺财',
    familyId: 1,
    avatar: '',
    birthday: "2022-10-02",
    gendar: 0,
    sterilizedState: 0,
    category: 0,
  },
  {
    id: 2,
    name: '帅帅',
    familyId: 1,
    avatar: '',
    birthday: "2021-03-02",
    gendar: 0,
    sterilizedState: 0,
    category: 1,
  },{
    id: 3,
    name: '旺财2',
    familyId: 1,
    avatar: '',
    birthday: "2012-10-02",
    gendar: 0,
    sterilizedState: 1,
    category: 1,
  },{
    id: 4,
    name: '1234',
    familyId: 1,
    avatar: '',
    birthday: "2002-10-02",
    gendar: 1,
    sterilizedState: 1,
    category: 1,
  }
]}
pets.value = getPetsInfo();


/** 待办事项相关参数和方法 */
const petTodos = ref();
watch([currentDate, currentPet], ([newVal1, newVal2], [oldVal1, oldVal2]) => {
  const todos = getPetTodos(currentDate.value, currentPet.value!);
  petTodos.value = todos;
});

const getPetTodos = (date: Date, pet: Pet.PetInfo) => {
  if (pet?.name === "1234") {
    return [
      {"id": 1, "time": "08:00", "title": "换水", "remark": "全部两个碗里的倒矿泉水纯净水蒸馏水都行,就是不能用自来水", "complete": false},
      {"id": 2, "time": "12:00", "title": "换粮", "remark": "全换", "complete": false},
      {"id": 3, "time": "18:00", "title": "铲屎", "remark": "", "complete": false},
    ]
  }
  if (date.getDate() === 2) {
    return [
      {"id": 1, "time": "08:00", "title": "换水", "remark": "全部两个碗里的倒矿泉水纯净水蒸馏水都行,就是不能用自来水", "complete": false},
      {"id": 2, "time": "12:00", "title": "换粮", "remark": "全换", "complete": false},
    ]
  }
  return [
      {"id": 1, "time": "08:00", "title": "换水", "remark": "全部两个碗里的倒矿泉水纯净水蒸馏水都行,就是不能用自来水", "complete": false},
      {"id": 2, "time": "12:00", "title": "换粮", "remark": "全换", "complete": false},
      {"id": 3, "time": "18:00", "title": "铲屎", "remark": "", "complete": false},
      {"id": 4, "time": "12:43", "title": "换水", "remark": "123123123123123123123123aasdasdasdsasadsadsdsdsadsdasqeqweqweqeqewqeqweqeqweq", "complete": false},
      {"id": 5, "time": "18:00", "title": "铲屎", "remark": "", "complete": true},
      {"id": 6, "time": "18:00", "title": "铲屎", "remark": "", "complete": false},
      {"id": 7, "time": "18:00", "title": "铲屎", "remark": "", "complete": false},
      {"id": 8, "time": "18:00", "title": "铲屎", "remark": "", "complete": false},
      {"id": 9, "time": "18:00", "title": "铲屎", "remark": "", "complete": false},
  ] 
}

const changeCalendarMode = (value: string) => {
  calendarMode.value = value
}

const swiperHeight = ref(40);
const updateSwiperHeight = (height: number) => {
  swiperHeight.value = height;
}
const calendarHeight = computed(()=>{
  const height = swiperHeight.value + 65;
  return `${height}px`;
})
</script>

<template>
  <basic-layout show-tab-bar>
    <custom-navbar title="首页" />
    <div v-if="pets.length">
      <div class="fixed-header" :style="{height: calendarHeight}">
        <calendar v-model="currentDate" :get-dot-info-func="getDotInfos" :show-week="true" @change-mode="changeCalendarMode" @update-swiper-height="updateSwiperHeight">
        </calendar>
      </div>
      <pet-todos-page :pets="pets" :todos="petTodos" class="todo-container" :style="{'padding-top': calendarHeight}" :current-date="currentDate"></pet-todos-page>
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

.fixed-header {
  position: fixed;
  top: 90px;
  width: 100%;
  background-color: #ffffff;
  overflow: hidden;
  z-index: 99;
}

.todo-container {
  transition: padding-top 0.3s; /* 确保过渡效果 */
  padding-bottom: 100px;
}
</style>
