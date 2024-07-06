<template>
  <div class="month">
    <nut-grid :column-num="7" square :border="false" :gutter="5">
        <nut-grid-item v-for="(elem, i) in weekName" :key="i" class="cell"> {{ elem }}</nut-grid-item>
    </nut-grid>
    <swiper
      class="swiper"
      :current="current"
      :circular="true"
      @change="handleSlide"
    >
      <swiper-item v-for="(month, index) in monthsT" :key="index">
        <nut-grid :column-num="7" square :border="false" :gutter="5">
            <nut-grid-item v-for="index in ((month[0].day.$W === 0 ? 7 : month[0].day.$W)-1)" :key="index"></nut-grid-item>
            <nut-grid-item v-for="(elem, i) in month" :key="i" class="date" :class="{ active: elem.isActive, today: elem.isToday }" @click="onHandleChangeDate(elem.date)">  
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
import { ref, computed } from "vue";
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
  }
});
// 周list
const selectedDate = computed(() => {
    return dayjs(props.modelValue);
});
const slideIndex = ref(1); // 当前slide索引
const current = ref(1);
const previousIndex = ref(1);
const monthIndexs = ref([-1, 0, 1]);

const dotData = computed(() => {
  if (!props.getDotInfoFunc) return [];
  const dates: Array<any> = [];

  months.value.map((m) => {
      m.map((item, index) => {
          dates.push(item.date);
      });
  });
  return props.getDotInfoFunc(dates);
});

const months = computed(() => {
    return monthIndexs.value.map((item) => getDays(item, selectedDate.value, "month"));
});

const monthsT = computed(() => {
  const dates: Array<any> = [];
  months.value.map(m => {
    const monthDates: Array<any> = [];
    m.map((item, index) => {
      const dateInfo = {
        ...item,
        isActive: selectedDate.value.format('YYYY-MM-DD') === item.date,
        isToday: dayjs().format('YYYY-MM-DD') === item.date,
        dotColors: dotData.value[item.date] || [],
      }
      monthDates.push(dateInfo);
    })
    dates.push(monthDates);
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
    needSwitchedDate = selectedDate.value.subtract(1, "month");
  } else {
    needSwitchedDate = selectedDate.value.add(1, "month");
  }

  onHandleChangeDate(needSwitchedDate.format('YYYY-MM-DD'));
  monthIndexs.value = current === 0
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
.month {

  .swiper {
    height: 400px;

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