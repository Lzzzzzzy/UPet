<script setup lang="ts">
import { ref, computed, watch, PropType } from "vue"
import calendarCard from './calendar-card.vue'

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
  swiperClass: {
    type: String,
    default: "",
  },
  // multiple: {
  //   type: Boolean,
  //   default: false,
  // }
});
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
const emit = defineEmits(["update:modelValue"]);
watch(selectedDates, (newVal) => {
  if (multiple.value) {
    emit('update:modelValue', selectedDates.value);
  } else {
    emit('update:modelValue', selectedDates.value[selectedDates.value.length - 1]);
  }
})

watch(currentDate, (newVal) => {
  switchedDate.value = newVal;
  if (!multiple.value) {
    selectedDates.value = [newVal]
  }
})

watch(() => props.modelValue, (newValue) => {
  if (Array.isArray(props.modelValue)) {
    selectedDates.value = props.modelValue;
    multiple.value = true;
  } else {
    selectedDates.value = [props.modelValue];
  }
}, { immediate: true });

const swiperClass = computed(()=> props.swiperClass)

const formattedMonth = computed(() => {
  const month = currentDate.value.getMonth() + 1;
  const year = currentDate.value.getFullYear();
  return `${year}年${month < 10 ? `0${month}` : month}月`;
});

const onHandleSelectDate = (date: Array<Date>) => {
  selectedDates.value = date
}
</script>
<template>
    <div>
        <div class="calendar-container">
            <div class="header-container">
                <div @click="showChangeDatePopup = true">{{ formattedMonth }}</div>
                <slot name="header"></slot>
            </div>
            <calendar-card v-model="currentDate" :get-dot-info-func="getDotInfoFunc" :is-week="showWeek" :swiper-class="swiperClass" @change-selected-dates="onHandleSelectDate" :selected-dates="selectedDates"></calendar-card>
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

</style>