<script setup lang="ts">
import { ref, computed, watch, PropType } from "vue"
import calendarCard from './calendar-card.vue'
import type { Dayjs } from 'dayjs'

const props = defineProps({
  modelValue: {
    type: Object as PropType<Date>,
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
  }
});
const currentDate = ref(props.modelValue);

// 计算当前月份和年份
const formattedMonth = computed(() => {
  const month = currentDate.value.getMonth() + 1;
  const year = currentDate.value.getFullYear();
  return `${year}年${month < 10 ? `0${month}` : month}月`;
});

const changeDate = ref(false);
const minDate = new Date('1990-01-01')
const maxDate = new Date('2099-12-31')
const switchedDate = ref(currentDate.value);
const confirm = () => {
    changeDate.value = false;
    currentDate.value = switchedDate.value;
}

const emit = defineEmits(["update:modelValue"]);
watch(currentDate, (newVal) => {
  emit('update:modelValue', newVal);
})

watch(() => props.modelValue, (newValue) => {
  currentDate.value = newValue;
});

const swiperClass = computed(()=> props.swiperClass)
</script>
<template>
    <div>
        <div class="calendar-container">
            <div class="header-container">
                <div @click="changeDate = true">{{ formattedMonth }}</div>
                <slot name="header"></slot>
            </div>
            <calendar-card v-model="currentDate" :get-dot-info-func="getDotInfoFunc" :is-week="showWeek" :swiper-class="swiperClass"></calendar-card>
        </div>
        <nut-popup v-model:visible="changeDate" position="bottom" round safe-area-inset-bottom>
            <nut-date-picker
                v-model="switchedDate"
                :min-date="minDate"
                :max-date="maxDate"
                :three-dimensional="false"
                @confirm="confirm"
                @cancel="changeDate = false"
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