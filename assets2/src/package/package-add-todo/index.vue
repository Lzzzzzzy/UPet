<script setup lang="ts">
import { ref, onBeforeMount, reactive, computed } from "vue";
import { formatDatetime, formatTime, isSameDate, isBeforeDate } from "@/utils/common/datetime";
import { eventCenter, getCurrentInstance } from "@tarojs/taro";
import richTextContent from "@/components/rich-text/index.vue";
import dayjs from "dayjs";

const selectedDate = ref(new Date());

const isEarlierThanToday = computed(() => {
  return isBeforeDate(selectedDate.value, new Date())
})

const pageTitle = computed(()=>{
  return isEarlierThanToday.value ? "添加记录" : "添加待办"
})

const tagList = [
  { text: "日常", value: 0 },
  { text: "需要注意", value: 1 },
  { text: "紧急情况", value: 2 }
]

const formData = reactive({
  title: '',
  time: formatDatetime(new Date()),
  remark: '',
  remind: false,
  remindDate: [dayjs().format('YYYY-MM-DD')],
  remindTime: formatTime(new Date()),
  complete: false,
  tag: [tagList[0].value],
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
      eventCenter.trigger('refreshTodo', formData);
    } else {
      console.warn('error:', errors)
    }
  })
}


onBeforeMount(() => {
  const instance = getCurrentInstance();
  const params = instance?.router?.params;
  selectedDate.value = params ? dayjs(params.selectedDate).toDate() : new Date();
  if (!isSameDate(selectedDate.value, new Date())) {
    formData.time = formatDatetime(selectedDate.value);
  }
})

// 日期选择
const showDatePicker = ref(false);
const selectedTime = ref(new Date());
const confirmSelectTime = () => {
  formData.time = formatDatetime(selectedTime.value);
  showDatePicker.value = false;
}

// 重复周期选择
const showRepeatDate = ref(false);

const currentDate = ref([new Date()])

const onHandleConfirmSelectDate = () => {
  const remindDate: Array<string> = [];
  currentDate.value.sort((a, b) => a.getTime() - b.getTime());
  currentDate.value.forEach((date)=>{
    remindDate.push(dayjs(date).format('YYYY-MM-DD'))
  })
  formData.remindDate = remindDate;
  showRepeatDate.value = false;
}

const onUpdateRemark = (data: string) => {
  console.log("data:", data);
  formData.remark = data;
}

const showTagSelection = ref(false);
const getTagText = (value: number) => {
  const tag = tagList.find(item => item.value === value);
  return tag?.text || '';
}

const showRepeatTimePicker = ref(false);
const remindTime = ref(new Date());
const confirmRepeatTime = () => {
  formData.remindTime = formatTime(remindTime.value);
  showRepeatTimePicker.value = false;
}

const remindDatesString = computed(() => {
  if (formData.remindDate.length === 1) {
    if (isSameDate(formData.remindDate[0], dayjs())) {
      return `${formData.remindDate[0]} (今天)`
    }
    return formData.remindDate[0]
  }
  return `已选 ${formData.remindDate.length} 天`
});
</script>
<template>
  <basic-layout>
    <custom-navbar :title="pageTitle" left-show />
    <div class="w-full text-20px">
      <nut-form
        ref="formRef"
        :model-value="formData"
        :rules="formRules"
        star-position="right"
      >
        <nut-form-item label="时间" class="form-item-border">
          <div @click="showDatePicker=true" class="w-full">{{ formData.time }}</div>
        </nut-form-item>
        <nut-form-item label="标题" prop="title" class="form-item-border">
          <nut-input v-model="formData.title" placeholder="请输入标题" />
        </nut-form-item>
        <nut-form-item label="备注" prop="remark" class="form-item-border">
          <rich-text-content :data="formData.remark" placeholder="备注，您也可以上传图片,例如确诊单或者希望给医生展示的症状" @update-data="onUpdateRemark"></rich-text-content>
        </nut-form-item>
        <nut-form-item label="标签" prop="tag" class="form-item-border">
          <div @click="showTagSelection=true" class="flex items-center">
            <div>{{ getTagText(formData.tag[0]) }}</div>
            <div class="w-15px h-15px border-rd-50% ml-10px" :class='`tag-color-${formData.tag[0]}`'></div>
          </div>
        </nut-form-item>
        <nut-form-item label="需要提醒" prop="remind" class="form-item-border" v-if="!isEarlierThanToday">
            <nut-switch v-model="formData.remind" />
        </nut-form-item>
        <nut-form-item label="提醒日期" prop="remindDate" class="form-item-border" v-if="formData.remind">
            <div @click="showRepeatDate=true" class="text-#315efb">
              {{ remindDatesString }}
            </div>
        </nut-form-item>
        <nut-form-item label="提醒时间" prop="remindTime" class="form-item-border" v-if="formData.remind">
          <div @click="showRepeatTimePicker=true" class="text-#315efb">{{ formData.remindTime }}</div>
        </nut-form-item>

        <nut-space class="m-10px flex justify-center w-full">
          <nut-button type="primary" @click="handleSubmit">提交</nut-button>
        </nut-space>
      </nut-form>
      <nut-popup v-model:visible="showDatePicker" position="bottom" round safe-area-inset-bottom>
        <nut-date-picker
          v-model="selectedTime"
          type="hour-minute"
          :three-dimensional="false"
          @confirm="confirmSelectTime"
          cancel-text=" "
        ></nut-date-picker>
      </nut-popup>
      <nut-popup v-model:visible="showTagSelection" position="bottom" round safe-area-inset-bottom>
        <nut-picker 
          v-model="formData.tag" 
          :columns="tagList" 
          title="选择标签" 
          @confirm="showTagSelection=false" 
          cancel-text=" "
        />
      </nut-popup>
      <nut-popup v-model:visible="showRepeatDate" position="bottom" round safe-area-inset-bottom>
        <div class="flex items-center justify-between h-45px font-size-14px">
          <div class="date-picker__left"></div>
          <div class="date-picker__center">可以选择多个提醒日期</div>
          <div class="date-picker__right px-15px" @click="onHandleConfirmSelectDate">确认</div>
        </div>
        <calendar v-model="currentDate" :show-week="false" :show-change-mode-button="false" class="font-size-16px">
        </calendar>
      </nut-popup>
      <nut-popup v-model:visible="showRepeatTimePicker" position="bottom" round safe-area-inset-bottom>
        <nut-date-picker
          v-model="remindTime"
          type="hour-minute"
          :three-dimensional="false"
          @confirm="confirmRepeatTime"
          cancel-text=" "
        ></nut-date-picker>
      </nut-popup>
    </div>
  </basic-layout>
</template>

<style lang="scss">
.tag-color-0 {
  border: 1px solid #222222;
}

.tag-color-1 {
  background: #FFC300;
}

.tag-color-2 {
  background: #FF5733;
}

.ellipsis-style {
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
}
</style>
