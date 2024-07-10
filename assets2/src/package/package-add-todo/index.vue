<script setup lang="ts">
import { ref, onBeforeMount, reactive } from "vue";
import { formatDatetime, getMinAndMaxDate } from "@/utils/common/datetime";
import { eventCenter } from "@tarojs/taro";
import richTextContent from "@/components/rich-text/index.vue";


definePageConfig({
  navigationBarTitleText: '添加待办'
});

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

const tagList = [
  { text: "一般", value: 0 },
  { text: "需要注意", value: 1 },
  { text: "紧急情况", value: 2 }
]

const formData = reactive({
  title: '',
  time: new Date(),
  remark: '',
  remind: false,
  repeat: [repeatList.value[0].value],
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
    formData.time = selectedTime.value;
    showDatePicker.value = false;
}

// 重复周期选择
const showRepeatSelection = ref(false);
const getRepeatText = (value: number) => {
  const repeat = repeatList.value.find(item => item.value === value);
  return repeat?.text || '';
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
</script>
<template>
  <basic-layout>
    <custom-navbar title="添加待办" left-show />
    <div class="w-full text-20px">
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
          <rich-text-content :data="formData.remark" placeholder="备注，您也可以上传图片,例如确诊单或者希望给医生展示的症状" @update-data="onUpdateRemark"></rich-text-content>
        </nut-form-item>
        <nut-form-item label="标签" prop="tag" class="form-item-border">
          <div @click="showTagSelection=true" class="flex items-center">
            <div>{{ getTagText(formData.tag[0]) }}</div>
            <div class="w-15px h-15px border-rd-50% ml-10px" :class='`tag-color-${formData.tag[0]}`'></div>
          </div>
          <nut-popup v-model:visible="showTagSelection" position="bottom" round safe-area-inset-bottom>
            <nut-picker 
              v-model="formData.tag" 
              :columns="tagList" 
              title="选择标签" 
              @confirm="showTagSelection=false" 
              cancel-text=" "
            />
          </nut-popup>
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
    </div>
  </basic-layout>
</template>

<style lang="scss">
.tag-color-1 {
  background: #FFC300;
}

.tag-color-2 {
  background: #FF5733;
}
</style>
