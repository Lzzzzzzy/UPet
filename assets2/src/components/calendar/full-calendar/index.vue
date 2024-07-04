<template>
    <nut-calendar-card v-model="currentDate" class="!bg-inherit" ref="calendarRef" :firstDayOfWeek="1" @page-change="onHandleChangeMonth" @day-click="onHandleChangeDate">
        <template #default="{ day }">
            {{ day.date <= 9 ? `0${day.date}` : day.date }}
        </template>
        <template #bottom="{ day }">
            <div>{{ isToday(day) ? "今天": "" }}</div>
        </template>
    </nut-calendar-card>
</template>

<script setup lang="ts">
import { ref, nextTick } from 'vue';
import { formatDate } from '@/utils/common/datetime';
import type { CalendarCardDay } from "@nutui/nutui-taro";

interface ElementMap {
  [key: string]: HTMLElement;
}

const currentDate = ref();

const isToday = (day: CalendarCardDay) => {
  const today = new Date();
  return day.date === today.getDate() && day.month === today.getMonth()+1 && day.year === today.getFullYear();
}

const calendarRef = ref();
let dayRefs: ElementMap = {};

const getColorDate = (year: number, month: number) => {
    const data = [{date: formatDate(new Date(year, month-1, 10)), color: "green"},{date: formatDate(new Date(year, month-1, 2)), color: "red"}];
    return data;
}

const addBackgroundColor = (colorDate: Array<{date: string, color: string}>) => {
  colorDate.forEach(item => {
    const date = item.date;
    const color = item.color;
    const day = dayRefs[date];
    if (day) {
      day.classList.add(color);
    }
  })
}

const onHandleChangeMonth = async ({ year, month}: { year: number, month: number}) =>{
  if (!calendarRef.value){
    return
  }
  await nextTick();  // 需要等待DOM加载完成，否则calendarRef没加载完成，会获取到错误子节点
  dayRefs = {};
  const calendarChildNodes = calendarRef.value.$el.childNodes;
  const calendarContentNodes = calendarChildNodes[calendarChildNodes.length - 1];
  const daysContentNodes = calendarContentNodes.childNodes[calendarContentNodes.childNodes.length - 1].childNodes;

  daysContentNodes.forEach((day: HTMLElement) => {
    if (day.className && day.className.includes('nut-calendarcard-day') && day.className.includes('current')) {
      const date = day.childNodes[2].childNodes[1].data;
      const key = formatDate(new Date(year, month-1, date));
      dayRefs[key] = day;
    }
  }) 
  const colorDate = getColorDate(year, month)
  addBackgroundColor(colorDate);
  const today = new Date();
  if (month-1 === today.getMonth() && year === today.getFullYear()) {
    currentDate.value = today;
  } else {
    currentDate.value = new Date(year, month-1, 1);
  }
};

const onHandleChangeDate = (date: CalendarCardDay)=>{
    console.log("changeDate:", date);
};
</script>

<style lang="scss">
.nut-calendarcard-day {
  border-radius: 16px;

  &.next {
    display: none;
  }

  &.prev {
    visibility: hidden;
  }
  
  &.active{
    background-color: initial;
    border: 2px solid var(--nut-primary-color);
    box-sizing: border-box;
    border-radius: 16px;
    
    &:not(weekend) {
      color: inherit;
    }

    &.weekend {
      color: var(--nut-primary-color);
    }
  }
}

.green {
  background-color: green;
}

.red {
  background-color: red;
}
</style>
