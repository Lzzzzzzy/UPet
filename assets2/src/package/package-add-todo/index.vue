<script setup lang="ts">
import { ref, onBeforeMount, reactive } from "vue";
import { formatDatetime, getMinAndMaxDate } from "@/utils/common/datetime";
import { eventCenter } from "@tarojs/taro";
import RichText from "@/components/rich-text/index.vue";


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

const formData = reactive({
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
            <!-- <editor id="editor" class="editor" placeholder="备注，支持上传图片" @ready="editorReady"></editor> -->
             <rich-text :data="formData.remark" placeholder="备注，支持上传图片" @update-data="onUpdateRemark"></rich-text>
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

<style lang="scss"></style>
