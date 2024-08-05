<template>
  <div class="card">
    <!-- 星期名称 -->
    <nut-grid :column-num="7" :border="false" class="week-text-container">
      <nut-grid-item v-for="(elem, i) in weekName" :key="i" class="cell"> {{ elem }}</nut-grid-item>
    </nut-grid>

    <!-- 轮播组件 -->
    <swiper
      class="swiper mode-change-transition"
      :current="current"
      :circular="true"
      @change="handleSlide"
      :style="{'height': calcCalendarHeight}"
    >
      <swiper-item v-for="(batch, index) in swipersDays" :key="index">
        <nut-grid :column-num="7" :border="false" :class="swiperClass">
          <!-- 从1号开始，空出1号之前的日期 -->
          <nut-grid-item v-for="index in ((batch[0].day.$W === 0 ? 7 : batch[0].day.$W)-1)" :key="index"></nut-grid-item>

          <!-- 当月日期 -->
          <nut-grid-item v-for="(elem, i) in batch" :key="i" @click="onHandleChangeSelectedDate(elem.date)" class="date">
            <template #default>
              <div class="date-num" :class="{ active: isActiveDay(elem.date), today: elem.isToday }">{{ elem.d }}</div>
              <div class="dot-container">
                <!-- 日期下面的标记小圆点 -->
                <div :style="{ background: elem.dotColor }" class="dot" v-if="elem.dotColor"></div>
              </div>
            </template>
          </nut-grid-item>
        </nut-grid>
      </swiper-item>
    </swiper>
  </div>
</template>
<script setup lang="ts">
import { ref, computed, watch, onBeforeMount } from "vue";
import type { Dayjs } from 'dayjs';
import dayjs from "dayjs";
import { getDays, isDateInArray, isSameDate, weekName, isBeforeDate } from '@/utils';
import { eventCenter } from "@tarojs/taro";

interface datInfos {
  d: Number,
  m: Number,
  date: String,
  day: Dayjs,
};
const props = defineProps({
  modelValue: {  // 当前指向的日期
    type: Date,
    required: true,
  },
  getDotInfoFunc: {
    type: Function || null || undefined,
    required: false,
  },
  isWeek: {
    type: Boolean,
    default: false,
  },
  swiperClass: {
    type: String,
    default: "",
  },
  selectedDates: {  // 当前选中的日期
    type: Array<Date>,
    default: false,
  }
});

onBeforeMount(() => {
  eventCenter.on("refreshDotData", () => {
    getDotData(true);
  })
})

const emit = defineEmits(["update:modelValue", "changeSelectedDates", "updateSwiperHeight"]);

const currentDate = computed(() => {
  return dayjs(props.modelValue);
});

const selectedDates = ref<Array<Date>>([]);

watch(() => props.selectedDates, (newValue) => {
  selectedDates.value = newValue;
}, { immediate: true });

const calendarType = computed(()=>{
  return props.isWeek ? "week" : "month";
})
const slideIndex = ref(1); // 当前slide索引
const current = ref(1);
const previousIndex = ref(1);
const dateIndexes = ref([-1, 0, 1]);

const minDate = ref();
const maxDate = ref();
const dotDataCache = ref([]); 

const getDotData = (refresh: Boolean = false) => {
  if (!props.getDotInfoFunc) return [];
  const dates = {"minDate": "", "maxDate":""};
  let min: Date;
  let max: Date;
  days.value.map((m: Array<datInfos>) => {
      m.map((item: datInfos) => {
        if (!min && !max) {
          min = item.day;
          max = item.day;
        } else {
          if (isBeforeDate(item.day, min)) {
            min = item.day;
          }
          if (isBeforeDate(max, item.day)) {
            max = item.day;
          }
        }
      });
  });

  if ((!minDate.value || !maxDate.value) || (!isSameDate(min, minDate.value) || !isSameDate(max, maxDate.value)) || refresh){
    minDate.value = min;
    maxDate.value = max;
    dates.minDate = dayjs(minDate.value).format("YYYY-MM-DD");
    dates.maxDate = dayjs(maxDate.value).format("YYYY-MM-DD");
    props.getDotInfoFunc(dates).then((result: any) => {
      dotDataCache.value = result;
      return dotDataCache.value;
    });
  }
  return dotDataCache.value;
}

const days = computed(() => {
  return dateIndexes.value.map((item) => getDays(item, currentDate.value, calendarType.value));
});

watch(days.value, (newVal) => {
  getDotData();
}, {
  immediate: true,
})

const isActiveDay = (date: string) => {
  let res = false;
  selectedDates.value.forEach((d) => {
    if (isSameDate(d, date)) {
      res = true;
    }
  });
  return res;
};

const getColorByIndex = (idx: number) => {
  const colorList = ["#94938D", "#E6CF00", "#E25342"];
  return colorList[idx];
}

const swipersDays = computed(() => {
  const dates: Array<any> = [];
  const dotDataValue = dotDataCache.value;
  days.value.map(m => {
    const swiperDays: Array<any> = [];
    m.map((item, index) => {
      const dateInfo = {
        ...item,
        isToday: isSameDate(dayjs(), item.date),
      }
      if (dotDataValue && (dotDataValue[item.date] || dotDataValue[item.date] === 0)) {
        dateInfo.dotColor = getColorByIndex(dotDataValue[item.date]);
      }
      swiperDays.push(dateInfo);
    })
    dates.push(swiperDays);
  });
  return dates;
});

interface SlideDetail {
  current: number;
}

// 处理轮播切换事件
const handleSlide = ({ detail: { current } }: {detail: SlideDetail}) => {
  slideIndex.value = current;
  const direct = getSlideDirect();
  previousIndex.value = current;

  let needSwitchedDate: Dayjs;
  if (direct === "left") {
    needSwitchedDate = currentDate.value.subtract(1, calendarType.value);
  } else {
    needSwitchedDate = currentDate.value.add(1, calendarType.value);
  }

  dateIndexes.value = current === 0
    ? [0, 1, -1]
    : current === 1
    ? [-1, 0, 1]
    : [1, -1, 0];
  
  emit('update:modelValue', needSwitchedDate.toDate())
};

const getSlideDirect = () => {
  if (previousIndex.value - 1 === slideIndex.value || slideIndex.value === previousIndex.value + 2) {
    return "left";
  } else {
    return "right";
  }
};

const onHandleChangeSelectedDate = (date: string | Dayjs) => {
  const day = dayjs(date);
  const isDaySelected = isDateInArray(day.toDate(), selectedDates.value);
  if (selectedDates.value.length === 1 && isDaySelected) {  // 只有一天选中时不能再取消
    return;
  }

  if (isDateInArray(day.toDate(), selectedDates.value)) {
    // 移除选中日期
    selectedDates.value = selectedDates.value.filter((d) => dayjs(d).format('YYYY-MM-DD') !== day.format('YYYY-MM-DD'));
  } else {
    // 添加选中日期
    selectedDates.value.push(day.toDate());
    selectedDates.value = selectedDates.value;
  }
  emit('changeSelectedDates', selectedDates.value);
};

const swiperClass = computed(()=> props.swiperClass);

const calcCalendarHeight = computed(() => {
  if (props.isWeek) {
    emit("updateSwiperHeight", 40)
    return "40px";
  }
  const dates = swipersDays.value[slideIndex.value];
  const firstLineCount = dates[0].day.$W === 0 ? 1 : (7-dates[0].day.$W+1);
  const otherDatesCount = dates.length - firstLineCount;
  let lines = Math.ceil(otherDatesCount / 7);
  lines += 1;
  const height = 43 * lines;
  emit("updateSwiperHeight", height)
  return `${height}px`;
})
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

    &.mode-change-transition {
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
            background-color: #F7DAA1;
          }

          &.today {
            border: 1px solid #F7DAA1;
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