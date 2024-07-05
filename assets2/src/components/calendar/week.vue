<template>
  <div class="week">
    <!-- 显示当前月份和年份 -->
    <h3>{{ formattedMonth }}</h3>
    <ul>
      <!-- 显示星期名称 -->
      <li v-for="(elem, i) in weekName" :key="i" class="cell date">
        <p>{{ elem }}</p>
      </li>
    </ul>
    <!-- 轮播组件 -->
    <swiper class="swiper" :current="current" :circular="true" @change="handleSlide">
      <swiper-item v-for="(week, index) in weeksT" :key="index">
        <ul>
          <!-- 显示每周的日期 -->
          <li v-for="(elem, i) in week" :key="i" class="cell date" :class="{ active: elem.isActive }" @click="onHandleChangeDate(elem.date)">
            <p>{{ elem.d }}</p>
          </li>
        </ul>
      </swiper-item>
    </swiper>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, PropType, watch } from "vue";
import { useDidShow } from '@tarojs/taro';
import dayjs from "dayjs";
import type { Dayjs } from 'dayjs';
import { getWeekDays, getWeek } from '@/utils/common/datetime';
import { weekName } from '@/utils/common/const';

// 定义组件的props，currentDate为Dayjs对象类型并且是必须的
const props = defineProps({
  modelValue: {
    type: Object as PropType<Dayjs>,
    required: true,
  }
});

// 初始化引用
const selectedDate = ref(props.modelValue);
const slideIndex = ref(1);
const previousIndex = ref(1);
const current = ref(1);
const weekIndexes = ref([-1, 0, 1]);
const weeks = ref([]); 

// 计算当前月份和年份
const formattedMonth = computed(() => {
  const month = curMonth.value.$M + 1;
  return `${curMonth.value.$y}年${month < 10 ? `0${month}` : month}月`;
});

// 计算当前月
const curMonth = computed(() => {
  return getWeek(weekIndexes.value[slideIndex.value], selectedDate.value).start;
});

// 计算每周的日期
watch(weekIndexes, (newIndexes, oldIndexes) => {  
  // 当weekIndexes变化时，重新计算weeks  
  weeks.value = newIndexes.map(index => getWeekDays(index, selectedDate.value));
  // 注意：这里仍然使用了selectedDate.value，但仅在weekIndexes变化时  
}, {  
  deep: true, // 如果weekIndexes是对象或数组，需要深度监听  
}); 

weekIndexes.value = [-1, 0, 1]

// 计算每周的日期，并设置是否为活动日期
const weeksT = computed(() => {
  return weeks.value.map(week => week.map((item, index) => ({
    ...item,
    isActive: selectedDate.value.format('YYYY-MM-DD') === item.date,
  })));
});

// 处理轮播切换事件
const handleSlide = ({ detail: { current } }) => {
  slideIndex.value = current;
  const direct = getSlideDirect();
  previousIndex.value = current;

  let needSwitchedDate: Dayjs;
  if (direct === "left") {
    needSwitchedDate = selectedDate.value.subtract(1, "week");
  } else {
    needSwitchedDate = selectedDate.value.add(1, "week");
  }

  onHandleChangeDate(needSwitchedDate.format('YYYY-MM-DD'));
  weekIndexes.value = current === 0
    ? [0, 1, -1]
    : current === 1
    ? [-1, 0, 1]
    : [1, 1, 0];
};

const getSlideDirect = () => {
  if (previousIndex.value - 1 === slideIndex.value || slideIndex.value === previousIndex.value + 2) {
    return "left";
  } else {
    return "right";
  }
};

const emit = defineEmits(["update:modelValue"]);
const onHandleChangeDate = (date: string | Dayjs) => {
  const copiedDate = dayjs(date);
  selectedDate.value = copiedDate;
  emit('update:modelValue', copiedDate);
};
</script>

<style lang="scss">
.week {
  margin: 10px 10px 15px;

  & > h3 {
    padding: 10px;
    width: max-content;
    font-size: 24px;
    font-weight: 400;
    line-height: 20px;
  }

  ul {
    display: flex;

    li {
      padding: 10px 0;
      display: flex;
      justify-content: center;
      align-items: center;
      flex-direction: column;
      width: calc((100vw - 20px) / 7);
      margin: 0 4px;
    }
  }

  .swiper {
    width: calc(100vw - 20px);
    height: 55px;

    li {
      position: relative;
      padding-top: 7px;
      height: 55px;
      justify-content: center;
      background-color: rgba(255, 255, 255, 1);
      border-radius: 5px;

      &.active {
        background: #fde98d;
      }

      p {
        margin-bottom: 1.5px;
        height: 14px;
        font-size: 10px;
        font-weight: 400;
        line-height: 14px;
        color: rgba(85, 85, 85, 1);
        text-overflow: ellipsis;
        white-space: nowrap;
        overflow: hidden;
        text-align: center;
        width: 100%;

        &:last-child {
          height: 16.5px;
          font-size: 12px;
          line-height: 16.5px;
          color: rgba(85, 85, 85, 1);
        }
      }

      span {
        position: absolute;

        &.brand {
          top: 0;
          right: 0;
          padding: 0 2px;
          background: #a4c3ff;
          text-align: center;
          line-height: 15px;
          min-width: 15px;
          font-size: 7px;
          font-weight: 400;
          color: rgba(255, 255, 255, 1);
          border-radius: 0 2px 0 6px;
        }
      }
    }
  }
}
</style>
