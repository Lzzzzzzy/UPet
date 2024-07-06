<script setup lang="ts">
import { ref, computed, watch } from "vue"
import month from './month.vue'
import week from './week.vue'
import type { Dayjs } from 'dayjs'

const props = defineProps({
  modelValue: {
    type: Object as PropType<Date | Dayjs>,
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
</script>
<template>
    <div class="calendar-container">
        <div class="header-container">
            <h3 @click="changeDate = true">{{ formattedMonth }}</h3>
            <slot name="header"></slot>
        </div>
        <div class="calendar">
            <week v-model="currentDate" :get-dot-info-func="getDotInfoFunc" v-if="showWeek"></week>
            <month v-else v-model="currentDate" :get-dot-info-func="getDotInfoFunc"></month>
        </div>
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
</template>

<style lang="scss">
.calendar-container {
  .header-container {
    padding: 10px;
    display: flex;
    justify-content: space-between;
    align-items: center;

    h3 {
        width: max-content;
        font-size: 24px;
        font-weight: 400;
        line-height: 20px;
    }
  }
}
</style>