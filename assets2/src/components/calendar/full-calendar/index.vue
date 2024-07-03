<template>
  <div>
    <div class="index-page">
      <nut-calendar-card v-model="currentDate" class="card-info" ref="calendarRef" firstDayOfWeek="1">
        <template #default="{ day }">
          {{ day.date <= 9 ? `0${day.date}` : day.date }}
        </template>
        <template #bottom="{ day }">
          <div :ref="(el)=>setDayRefs(el, day)">{{ isToday(day) ? "今天": "" }}</div>
        </template>
      </nut-calendar-card>
    </div>
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
  return day.date === today.getDate() && day.month === today.getMonth()+1 && day.year === today.getFullYear();
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
.card-info {
  background: inherit;
}


.step-container {
  padding: 3rpx 10rpx;
}

.nut-calendarcard-day-top {
  height: 0 !important;
  line-height: 0 !important;
}

.nut-calendarcard-day {
  &.next {
    display: none;
  }

  &.prev {
    visibility: hidden;
  }
  
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
</style>
