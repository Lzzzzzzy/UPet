<script lang="ts" setup>
import { ref, onBeforeMount } from "vue";
import { formatDate, getMinAndMaxDate } from "@/utils/common/datetime";

const showAddTodoPopup = ref(false);
const formData = ref({
  title: '',
  time: formatDate(new Date()),
  remark: ''
})

const formRef = ref()

const formRules = ref({
  title: [
    { required: true, message: '请填写标题' }
  ]
})

const handleSubmit = () => {
  formRef.value?.validate().then(({ valid, errors }) => {
    if (valid) {
      
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

const selectedTime = ref('');
const confirmSelectTime = () => {
    formData.value.time = formatDate(selectedTime.value);
    showDatePicker.value = false;
}

const showDatePicker = ref(false);
</script>
<template>
    <div>
        <nut-avatar class="!flex justify-center items-center !bg-#F7DAA1 text-#665D21 !position-fixed position-bottom-100px position-right-30px" size="large" @click="showAddTodoPopup=true">
            <div class="text-30px i-local-add"></div>
        </nut-avatar>
        <nut-popup v-model:visible="showAddTodoPopup" position="bottom" round safe-area-inset-bottom pop-class="mb-50px">
          <nut-form
            ref="formRef"
            :model-value="formData"
            :rules="formRules"
            star-position="right"
          >
            <nut-form-item label="日期" class="form-item-border">
                <div @click="showDatePicker=true" class="w-full">{{ formData.time }}</div>
                <nut-popup v-model:visible="showDatePicker" position="bottom" round safe-area-inset-bottom>
                    <nut-date-picker
                        v-model="selectedTime"
                        type="hour-minute"
                        :three-dimensional="false"
                        :min-date="selectionMinDate"
                        :max-date="selectionMaxDate"
                        @confirm="confirmSelectTime"
                        @cancel="showDatePicker=false"
                    ></nut-date-picker>
                </nut-popup>
            </nut-form-item>
            <nut-form-item label="标题" prop="title" class="form-item-border">
                <nut-input v-model="formData.title" placeholder="请输入标题" />
            </nut-form-item>
            <nut-form-item label="备注" prop="remark" class="form-item-border">
                <nut-textarea v-model="formData.remark" placeholder="请输入备注" />
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