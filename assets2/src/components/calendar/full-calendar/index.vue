<template>
  <div>
    <div class="index-page">
      <nut-calendar-card v-model="currentDate" class="card-info" ref="calendarRef">
        <template #bottom="{ day }">
          <div :ref="(el)=>setDayRefs(el, day)">{{ isToday(day) ? "今天": "" }}</div>
        </template>
      </nut-calendar-card>
    </div>
    <div class="subtitle">
      <span>毛孩子的今天</span>
      <nut-tabs v-model="currentTab" align="left" :animated-time="0">
        <nut-tab-pane v-for="tab in tabList" :key="tab.paneKey" :title="tab.title" :pane-key="tab.paneKey"></nut-tab-pane>
      </nut-tabs>
    </div>
    <div class="step-container">
      <nut-steps direction="vertical">
        <nut-step v-for="info in data" :key="info.id" :title="info.title" :content="info.content">
          <template #icon>
            123
          </template>
        </nut-step>
      </nut-steps>
    </div>
    <div class="record-button-container">
      <nut-button plain type="info">记一笔</nut-button>
    </div>
    <!-- <tabbar activate="calendar"></tabbar> -->
  </div>
</template>

<script setup>
import { ref, onBeforeMount, onMounted } from 'vue';

const currentDate = ref();
const today = ref();
const currentTab = ref(1);

onBeforeMount(()=>{
  currentDate.value = new Date();
})

const isToday = (day) => {
  const today = new Date();
  return day.date === today.getDate() && day.month === today.getMonth()+1;
}

const calendarRef = ref();
const dayRefs = ref({});
onMounted(()=>{
  const calendarChildNodes = calendarRef.value.$el.childNodes;
  const calendarContentNodes = calendarChildNodes[calendarChildNodes.length - 1];
  const daysContentNodes = calendarContentNodes.childNodes[calendarContentNodes.childNodes.length - 1].childNodes;

  const today = new Date();
  const currentMonth = today.getMonth()+1;
  const currentYear = today.getFullYear();
  console.log(currentYear);
  daysContentNodes.forEach(day => {
    let month = currentMonth;
    let year = currentYear;
    if (day.classList && day.classList.contains('nut-calendarcard-day')) {
      if (day.classList.contains("prev")) {  // 本月如果是1月，那上个月应该是去年的12月
        if (currentMonth === 1) {
          month = 12;
          year = currentYear - 1;
        } else {
          month = currentMonth - 1;
        }
      } else if (day.classList.contains("next")) {  // 本月如果是12月，那下个月应该是明年的1月
        if (currentMonth === 12) {
          month = 1;
          year = currentYear + 1;
        } else {
          month = currentMonth + 1;
        }
      }

      const date = day.childNodes[2].childNodes[1].data;
      const key = `${year}-${month}-${date}`
      dayRefs.value[key] = day;
    }
  })

  addBackgroundColor();
  console.log(dayRefs)
})

const addBackgroundColor = () => {
  const data = [{"2024-7-10": "green", "2024-7-11": "red"}]
  data.forEach(item => {
    const date = Object.keys(item)[0];
    const color = Object.values(item)[0];
    const day = dayRefs.value[date];
    console.log("day:", day);
    if (day) {
      day.style.backgroundColor = color;
      day.classList.add('highlight');
    }
  })
}

const setDayRefs = (el, day) => {
  // console.log("el:", el.parentNode);
  return day.date + day.month;
}


// onMounted(()=>{
//   console.log("mounted");
//   const dayElements = document.querySelectorAll('.nut-calendarcard-day');
//   dayElements.forEach(function(element) {
//     console.log("element:", element);
//     console.log("day-info:", element.querySelector('.day-info'));
//       const number = parseInt(element.querySelector('.day-info').textContent);
      
//       if (number === 1) {
//           element.classList.add('highlight');
//       } else if (number === 2) {
//           element.classList.add('special');
//       }
//   });
// })

const tabList = [{title: "日常", paneKey: 1}, {title: "花销", paneKey: 2}, {title: "清理", paneKey: 3}]
const data = ref([{"title": "喂食", "content": "喂了一半", "id": 1, "type": "feed"}, {"title": "铲屎", "content": "没铲干净没铲干净没铲干净没铲干净没铲干净没铲干净没铲干净没铲干净没铲干净没铲干净没铲干净没铲干净", "id": 2, "type": "clean"}, {"title": "买猫粮", "content": "买了300g分装猫粮", "id": 3, "type": "purchase"}, {"title": "喂食", "content": "喂了一半", "id": 1, "type": "feed"}, {"title": "铲屎", "content": "", "id": 2, "type": "clean"}, {"title": "买猫粮", "content": "买了300g分装猫粮", "id": 3, "type": "purchase"}])
</script>

<style lang="scss">
.index-page {
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  .card-info {
    background: inherit;
  }
}

.step-container {
  padding: 3rpx 10rpx;
}

.nut-calendarcard-day-top {
  height: 0 !important;
  line-height: 0 !important;
}

.nut-calendarcard-day {
  
  &.active{
    background-color: initial;
    border: 1rpx solid #000000;
    box-sizing: border-box;
    
    &:not(weekend) {
      color: inherit;
    }

    &.weekend {
      color: var(--nut-primary-color);
    }
  }
}

.nut-step-main {
  margin-bottom: 40rpx;
}

.subtitle {
  padding: 0 10rpx;
  margin: 10rpx 0 20rpx;
  font-size: 38rpx;
  font-weight: 600;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.custom-title {
  color: black;
  cursor: pointer;
}

.custom-title.active {
  color: red;
}

.nut-tab-pane {
  padding: 0 !important;
}

.record-button-container {
  position: fixed;
  right: 30rpx;
  bottom: 60rpx;
}

.highlight {
  background-color: yellow;
}

.special {
  background-color: lightblue;
}
</style>
