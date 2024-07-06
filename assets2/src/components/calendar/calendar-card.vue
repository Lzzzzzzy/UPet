<template>
  <div class="card">
    <nut-grid :column-num="7" :border="false" class="week-text-container">
        <nut-grid-item v-for="(elem, i) in weekName" :key="i" class="cell"> {{ elem }}</nut-grid-item>
    </nut-grid>
    <swiper
      class="swiper"
      :current="current"
      :circular="true"
      @change="handleSlide"
      :class="{ 'week-height': isWeek, 'month-height': !isWeek }"
    >
      <swiper-item v-for="(batch, index) in swipersDays" :key="index">
        <nut-grid :column-num="7" :border="false">
            <nut-grid-item v-for="index in ((batch[0].day.$W === 0 ? 7 : batch[0].day.$W)-1)" :key="index"></nut-grid-item>
            <nut-grid-item v-for="(elem, i) in batch" :key="i" @click="onHandleChangeDate(elem.date)" class="date">  
              <template #default>
                <div class="date-num" :class="{ active: elem.isActive, today: elem.isToday }">{{ elem.d }}</div>
                <div class="dot-container">
                    <div v-for="(color, i) in elem.dotColors" :key="i" :style="{ background: color }" class="dot"></div>
                </div>
              </template>
            </nut-grid-item>
        </nut-grid>
      </swiper-item>
    </swiper>
  </div>
</template>
<script setup lang="ts">
import { ref, computed, PropType } from "vue";
import type { Dayjs } from 'dayjs';
import dayjs from "dayjs";
import { getDays} from '@/utils/common/datetime';
import { weekName } from '@/utils/common/const';

const props = defineProps({
  modelValue: {
    type: Object as PropType<Date | Dayjs>,
    required: true,
  },
  getDotInfoFunc: {
    type: Function,
    required: false,
  },
  isWeek: {
    type: Boolean,
    default: false,
  }
});

const selectedDate = computed(() => {
  return dayjs(props.modelValue);
});

const calendarType = computed(()=>{
  return props.isWeek ? "week" : "month";
})
const slideIndex = ref(1); // 当前slide索引
const current = ref(1);
const previousIndex = ref(1);
const dateIndexes = ref([-1, 0, 1]);

const dotData = computed(() => {
  if (!props.getDotInfoFunc) return [];
  const dates: Array<any> = [];

  days.value.map((m) => {
      m.map((item, index) => {
          dates.push(item.date);
      });
  });
  return props.getDotInfoFunc(dates);
});

const days = computed(() => {
  return dateIndexes.value.map((item) => getDays(item, selectedDate.value, calendarType.value));
});

const swipersDays = computed(() => {
  const dates: Array<any> = [];
  days.value.map(m => {
    const swiperDays: Array<any> = [];
    m.map((item, index) => {
      const dateInfo = {
        ...item,
        isActive: selectedDate.value.format('YYYY-MM-DD') === item.date,
        isToday: dayjs().format('YYYY-MM-DD') === item.date,
        dotColors: dotData.value[item.date] || [],
      }
      swiperDays.push(dateInfo);
    })
    dates.push(swiperDays);
  });
  return dates;
});

// 处理轮播切换事件
const handleSlide = ({ detail: { current } }) => {
  slideIndex.value = current;
  const direct = getSlideDirect();
  previousIndex.value = current;

  let needSwitchedDate: Dayjs;
  if (direct === "left") {
    needSwitchedDate = selectedDate.value.subtract(1, calendarType.value);
  } else {
    needSwitchedDate = selectedDate.value.add(1, calendarType.value);
  }

  onHandleChangeDate(needSwitchedDate.format('YYYY-MM-DD'));
  dateIndexes.value = current === 0
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
.card {
  font-size: 14px;
  .nut-grid-item__content--center {
    justify-content: flex-start;
    padding-top: 0;
    padding-bottom: 0;
  }

  .swiper {

    &.week-height {
      height: 40px;
      transition: height 0.3s ease-in-out;  
    }

    &.month-height {
      height: 210px;
      transition: height 0.3s ease-in-out;
    }

    .date {

      .date-num {
          display: flex;
          align-items: center;
          justify-content: center;
          border-radius: 50%;
          width: calc(((100vw - 20px) / 7) * 0.5);
          height: calc(((100vw - 20px) / 7) * 0.5);

          &.active {
            background-color: #fde98d;
          }

          &.today {
            border: 1px solid #fde98d;
          }
      }
    }

    .dot-container {
        display: flex;
        align-items: center;
        justify-content: center;
        height: 5px;
        margin-top: 3px;

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