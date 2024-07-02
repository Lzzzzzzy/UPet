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
const confirmSelectTime = () => {}

const showDatePicker = ref(false);
</script>
<template>
    <div>
        <nut-avatar class="!flex justify-center items-center !bg-#F7DAA1 text-#665D21 !position-fixed position-bottom-100px position-right-30px" size="large" @click="showAddTodoPopup=true">
            <div class="text-30px i-local-add"></div>
        </nut-avatar>
        <nut-popup v-model:visible="showAddTodoPopup" position="bottom">
          <nut-form
            ref="formRef"
            :model-value="formData"
            :rules="formRules"
            star-position="right"
          >
            <div>
                <span>日期：</span>
                <span @click="showDatePicker=true">{{ formData.time }}</span>
                <nut-popup v-model:visible="showDatePicker" position="bottom">
                    <nut-date-picker
                        v-model="selectedTime"
                        type="time"
                        :three-dimensional="false"
                        :min-date="selectionMinDate"
                        :max-date="selectionMaxDate"
                        @confirm="confirmSelectTime"
                    ></nut-date-picker>
                </nut-popup>
            </div>
            <nut-form-item label="标题" prop="title" class="form-item-border">
                <nut-input v-model="formData.title" placeholder="Placeholder" />
            </nut-form-item>
            <nut-form-item label="备注" prop="remark" class="form-item-border">
                <nut-textarea v-model="formData.remark" />
            </nut-form-item>

            <nut-space class="m-10px flex justify-center w-full">
                <nut-button type="primary" @click="handleSubmit">提交</nut-button>
            </nut-space>
          </nut-form>
        </nut-popup>
    </div>
</template>