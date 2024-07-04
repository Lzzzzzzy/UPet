<script lang="ts" setup>
import { ref, onBeforeMount } from "vue";
import { formatDatetime, getMinAndMaxDate } from "@/utils/common/datetime";
import { eventCenter } from "@tarojs/taro";

const showAddTodoPopup = ref(false);

const repeatList = ref([
  { text: '不重复', value: 0 },
  { text: '每天', value: 1 },
  { text: '每周', value: 2 },
  { text: '每两周', value: 3 },
  { text: '每月', value: 4 },
  { text: '每三个月', value: 5 },
  { text: '每半年', value: 6 },
  { text: '每年', value: 7 },
])

const formData = ref({
  title: '',
  time: new Date(),
  remark: '',
  remind: false,
  repeat: [repeatList.value[0].value],
  complete: false,
})

const formRef = ref()

const formRules = ref({
  title: [
    { required: true, message: '请填写标题' }
  ]
})

const handleSubmit = () => {
  formRef.value?.validate().then(({ valid, errors }: { valid: boolean, errors: any}) => {
    if (valid) {
      eventCenter.trigger('refreshTodo', formData.value);
      showAddTodoPopup.value = false;
    } else {
      console.warn('error:', errors)
    }
  })
}

const selectionMinDate = ref();
const selectionMaxDate = ref();

onBeforeMount(() => {
  const { minDate, maxDate } = getMinAndMaxDate(new Date())
  selectionMinDate.value = minDate;
  selectionMaxDate.value = maxDate;
})

// 日期选择
const showDatePicker = ref(false);
const selectedTime = ref();
const confirmSelectTime = () => {
    formData.value.time = selectedTime.value;
    showDatePicker.value = false;
}

// 重复周期选择
const showRepeatSelection = ref(false);
const getRepeatText = (value: number) => {
  const repeat = repeatList.value.find(item => item.value === value);
  return repeat?.text || '';
}
</script>
<template>
    <div>
        <nut-avatar class="!flex justify-center items-center !bg-#F7DAA1 text-#665D21 !position-fixed position-bottom-100px position-right-30px" size="large" @click="showAddTodoPopup=true">
            <div class="text-30px i-local-add"></div>
        </nut-avatar>
        <nut-popup v-model:visible="showAddTodoPopup" position="bottom" round safe-area-inset-bottom pop-class="mb-50px">
          <div class="w-full flex justify-center p-20px">新增待办</div>
          <nut-form
            ref="formRef"
            :model-value="formData"
            :rules="formRules"
            star-position="right"
          >
            <nut-form-item label="时间" class="form-item-border">
                <div @click="showDatePicker=true" class="w-full">{{ formatDatetime(formData.time) }}</div>
                <nut-popup v-model:visible="showDatePicker" position="bottom" round safe-area-inset-bottom>
                    <nut-date-picker
                        v-model="selectedTime"
                        type="hour-minute"
                        :three-dimensional="false"
                        @confirm="confirmSelectTime"
                        cancel-text=" "
                        :min-date="selectionMinDate"
                        :max-date="selectionMaxDate"
                    ></nut-date-picker>
                </nut-popup>
            </nut-form-item>
            <nut-form-item label="标题" prop="title" class="form-item-border">
                <nut-input v-model="formData.title" placeholder="请输入标题" />
            </nut-form-item>
            <nut-form-item label="备注" prop="remark" class="form-item-border">
                <nut-textarea v-model="formData.remark" class="h-50px" placeholder="请输入备注"/>
            </nut-form-item>
            <nut-form-item label="需要提醒" prop="remind" class="form-item-border">
                <nut-switch v-model="formData.remind" />
            </nut-form-item>
            <nut-form-item label="重复" prop="repeat" class="form-item-border" v-if="formData.remind">
                <div @click="showRepeatSelection=true">{{ getRepeatText(formData.repeat[0]) }}</div>
                <nut-popup v-model:visible="showRepeatSelection" position="bottom" round safe-area-inset-bottom pop-class="mb-50px">
                    <nut-picker v-model="formData.repeat" :columns="repeatList" title="选择重复提醒周期" @confirm="showRepeatSelection=false" cancel-text=" " />
                </nut-popup>
            </nut-form-item>

            <nut-space class="m-10px flex justify-center w-full">
                <nut-button type="primary" @click="handleSubmit">提交</nut-button>
            </nut-space>
          </nut-form>
        </nut-popup>
    </div>
</template>
<style lang="scss">
.form-item-border {
    border-bottom: 1px solid var(--nut-input-border-bottom, #eaf0fb)
  }
</style>