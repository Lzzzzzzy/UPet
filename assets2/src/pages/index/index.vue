<script setup lang="ts">
import { ref, onBeforeMount, watch } from 'vue';
import { eventCenter } from '@tarojs/taro';
import noPetRemind from '@/components/home/add-pet-remind/index.vue';
import petTodosPage from '@/components/pet-todos/index.vue';
import calendar from '@/components/calendar/index.vue';

/** 设置页面属性 */
definePageConfig({
  navigationBarTitleText: '首页'
});

/** 日历相关参数和方法 */
const currentDate = ref(new Date());
const getDotInfos = (date: Array<string>) => {
    return {"2024-07-02": ["red", "black"], "2024-07-03": ["green"], "2024-07-30": ["green"], "2024-07-31": ["orange"]}
}
const showWeek = ref(true);

const swiperDirection = ref('');
const changeCurrentDateAnimation = () => {
  const today = new Date().setHours(0, 0, 0, 0);
  if (currentDate.value.setHours(0, 0, 0, 0) ==  today) {
    return
  }
  if (currentDate.value.setHours(0, 0, 0, 0) > today) {
    swiperDirection.value = 'slide-left'
  } else {
    swiperDirection.value = 'slide-right'
  }
  setTimeout(() => swiperDirection.value = '', 300);
}

const goToday = () => {
  changeCurrentDateAnimation();
  currentDate.value = new Date();
}


/** 宠物相关参数和方法 */
const currentPet = ref<Pet.PetInfo>();

onBeforeMount(() => {
    eventCenter.on("selectpet", (pet: Pet.PetInfo) => {
        currentPet.value = pet
    });
});

const pets = ref<Array<Pet.PetInfo>>([]);
const getPetsInfo = () => {
  return [{
    id: 1,
    petName: '旺财',
    petAvatar: '',
    userId: 1,
  },
  {
    id: 2,
    petName: '帅帅',
    petAvatar: '',
    userId: 1,
  },{
    id: 3,
    petName: '1234',
    petAvatar: '',
    userId: 1,
  },{
    id: 4,
    petName: '1123',
    petAvatar: '',
    userId: 1,
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
  if (pet?.petName === "1234") {
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
</script>

<template>
  <basic-layout show-tab-bar>
    <custom-navbar title="首页" />
    <div v-if="pets.length">
      <div class="fixed-header" :class="{ expanded: !showWeek }">
        <calendar v-model="currentDate" :get-dot-info-func="getDotInfos" :show-week="showWeek" :swiper-class="swiperDirection">
          <template #header>
            <div>
              <div class="text-20px pr-20px" @click="showWeek = !showWeek" :class="{ 'i-local-calendar-month': showWeek, 'i-local-calendar-week':  !showWeek }"></div>
              <div class="text-20px i-local-goto-today pl-20px" @click="goToday"></div>
            </div>
          </template>
        </calendar>
      </div>
      <pet-todos-page :pets="pets" :todos="petTodos" class="todo-container"></pet-todos-page>
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

@keyframes slide-left {
  0% {
    transform: translateX(0);
  }
  50% {
    transform: translateX(50%);
  }
  100% {
    transform: translateX(0);
  }
}

@keyframes slide-right {
  0% {
    transform: translateX(0);
  }
  50% {
    transform: translateX(-50%);
  }
  100% {
    transform: translateX(0);
  }
}

.slide-left {
  animation: slide-left 0.3s forwards;
}

.slide-right {
  animation: slide-right 0.3s forwards;
}

.fixed-header {
  position: fixed;
  top: 90px;
  width: 100%;
  background-color: #ffffff;
  // transition: height 0.3s;
  overflow: hidden;
  height: 105px; /* 初始高度 */
  z-index: 99;
}

.fixed-header.expanded {
  height: 270px; /* 展开后的高度 */
}

.todo-container {
  padding: 105px 0; /* 初始高度，等于固定头部初始高度 */
  transition: padding-top 0.3s; /* 确保过渡效果 */
}

.fixed-header.expanded + .todo-container {
  padding-top: 270px; /* 展开后，内容容器下移 */
}
</style>
