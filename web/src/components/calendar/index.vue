<script setup lang="ts">
import { ref, computed, watch, PropType } from "vue"
import calendarCard from './calendar-card.vue'
import { isSameDate, isBeforeDate } from "@/utils";

const props = defineProps({
  modelValue: {
    type: Object as PropType<Date> | Object as PropType<Array<Date>>,
    required: true,
  },
  getDotInfoFunc: {
    type: Function,
    required: false,
  },
  showWeek: {
    type: Boolean,
    default: true,
  },
  showChangeModeButton: {
    type: Boolean,
    default: true,
  }
});

// 对于父组件来说，只关心选中的日期，不关心当前指向的日期，所以当前指向的日期只是组件内部的记录，向父组件只暴露选中的日期

const currentDate = ref(new Date());  // 当前指向的日期
const selectedDates = ref();  // 当前选中的日期
const multiple = ref(false);  // 是否多选
const switchedDate = ref(new Date());  // 切换日期时指向的初始日期

const showChangeDatePopup = ref(false);
const minDate = new Date('1990-01-01')
const maxDate = new Date('2099-12-31')

const confirm = () => {
  currentDate.value = switchedDate.value;
  showChangeDatePopup.value = false;
}

// 向父组件返回选中的日期
const emit = defineEmits(["update:modelValue", "changeMode", "updateSwiperHeight"]);
watch(selectedDates, (newVal) => {
  if (multiple.value) {
    emit('update:modelValue', selectedDates.value);
  } else {
    emit('update:modelValue', selectedDates.value[selectedDates.value.length - 1]);
  }
}, { deep: true })

watch(currentDate, (newVal) => {
  switchedDate.value = newVal;
  if (!multiple.value) {
    selectedDates.value = [newVal]
  }
})

watch(() => props.modelValue, (newValue) => {
  if (Array.isArray(newValue)) {
    selectedDates.value = newValue;
    multiple.value = true;
  } else {
    selectedDates.value = [newValue];
    currentDate.value = newValue;
  }
}, { immediate: true });

const formattedMonth = computed(() => {
  const month = currentDate.value.getMonth() + 1;
  const year = currentDate.value.getFullYear();
  return `${year}年${month < 10 ? `0${month}` : month}月`;
});

const onHandleSelectDate = (date: Array<Date>) => {
  selectedDates.value = date
}

const showWeekView = ref(props.showWeek);
const swiperClass = ref('');
const changeCurrentDateAnimation = () => {
  const today = new Date();
  if (isSameDate(currentDate.value, today)) {
    return
  }
  if (isBeforeDate(today, currentDate.value)) {
    swiperClass.value = 'slide-left'
  } else {
    swiperClass.value = 'slide-right'
  }
  setTimeout(() => swiperClass.value = '', 300);
}

const goToday = () => {
  changeCurrentDateAnimation();
  currentDate.value = new Date();
}

const changeCalendarMode = () => {
  showWeekView.value = !showWeekView.value
  const mode = showWeekView.value ? "week" : "month";
  emit("changeMode", mode);
}

const updateSwiperHeight = (height: number) => {
  emit("updateSwiperHeight", height)
}
</script>
<template>
    <div>
        <div class="calendar-container">
            <div class="header-container">
                <div @click="showChangeDatePopup = true">{{ formattedMonth }}</div>
                <div>
                  <div class="text-20px pr-20px" @click="changeCalendarMode" :class="{ 'i-local-calendar-month': showWeekView, 'i-local-calendar-week':  !showWeekView }" v-if="showChangeModeButton"></div>
                  <div class="text-20px i-local-goto-today pl-20px" @click="goToday"></div>
                  <slot name="header"></slot>
                </div>
            </div>
            <calendar-card v-model="currentDate" :get-dot-info-func="getDotInfoFunc" :is-week="showWeekView" :swiper-class="swiperClass" @change-selected-dates="onHandleSelectDate" :selected-dates="selectedDates" @update-swiper-height="updateSwiperHeight"></calendar-card>
        </div>
        <nut-popup v-model:visible="showChangeDatePopup" position="bottom" round safe-area-inset-bottom>
            <nut-date-picker
                v-model="switchedDate"
                :min-date="minDate"
                :max-date="maxDate"
                :three-dimensional="false"
                @confirm="confirm"
                @cancel="showChangeDatePopup = false"
            ></nut-date-picker>
        </nut-popup>
    </div>
</template>

<style lang="scss">
.calendar-container {
  .header-container {
    padding: 8px 15px;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
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
</style>