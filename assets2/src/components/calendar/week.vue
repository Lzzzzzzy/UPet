<template>
  <div class="week">
    <nut-grid :column-num="7" square :border="false">
        <nut-grid-item v-for="(elem, i) in weekName" :key="i"> {{ elem }}</nut-grid-item>
    </nut-grid>
    <!-- 轮播组件 -->
    <swiper class="swiper" :current="current" :circular="true" @change="handleSlide">
      <swiper-item v-for="(week, index) in weeksT" :key="index">
        <nut-grid :column-num="7" square :border="false" clickable>
            <nut-grid-item v-for="(elem, i) in week" :key="i" class="date" :class="{ active: elem.isActive, today: elem.isToday }" @click="onHandleChangeDate(elem.date)">
                <div class="date-num">{{ elem.d }}</div>
                <div class="dot-container">
                    <div v-for="(color, i) in elem.dotColors" :key="i" :style="{ background: color }" class="dot"></div>
                </div>
            </nut-grid-item>
        </nut-grid>
      </swiper-item>
    </swiper>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, PropType } from "vue";
import dayjs from "dayjs";
import type { Dayjs } from 'dayjs';
import { getDays, getWeek } from '@/utils/common/datetime';
import { weekName } from '@/utils/common/const';

const props = defineProps({
  modelValue: {
    type: Object as PropType<Date | Dayjs>,
    required: true,
  },
  getDotInfoFunc: {
    type: Function,
    required: false,
  }
});

// 初始化引用
const selectedDate = computed(() => {
    return dayjs(props.modelValue);
});
const slideIndex = ref(1);
const previousIndex = ref(1);
const current = ref(1);
const weekIndexes = ref([-1, 0, 1]);

// 计算每周的日期
const weeks = computed(() => {
    return weekIndexes.value.map((item) => getDays(item, selectedDate.value, "week"));
});

const dotData = computed(() => {
    if (!props.getDotInfoFunc) return [];
    const dates = [];

    weeks.value.forEach((week) => {
        week.forEach((item) => {
            dates.push(item.date);
        });
    });
    return props.getDotInfoFunc(dates);
});

// 计算每周的日期，并设置是否为活动日期
const weeksT = computed(() => {
  return weeks.value.map(week => week.map((item, index) => ({
    ...item,
    isActive: selectedDate.value.format('YYYY-MM-DD') === item.date,
    isToday: dayjs().format('YYYY-MM-DD') === item.date,
    dotColors: dotData.value[item.date] || [],
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
    : [1, -1, 0];
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
  emit('update:modelValue', copiedDate.toDate());
};
</script>

<style lang="scss">
.week {

  .swiper {
    height: 55px;

    .date {
        .nut-grid-item__content {
            border-radius: 50%;
            justify-content: end;
        }

        &.active {
            .nut-grid-item__content {
                background-color: #fde98d;
            }
        }
    
        &.today {
            .nut-grid-item__content {
                border: 1px solid #fde98d;
            }
        }

        .date-num {
            display: flex;
            align-items: center;
            justify-content: center;
            height: 100%;
        }
    }

    .dot-container {
        display: flex;
        align-items: center;
        justify-content: center;
        height: 5px;

        .dot {
            width: 5px;
            height: 5px;
            border-radius: 50%;
            margin: 0 1px;
        }
    }
  }
}
</style>
