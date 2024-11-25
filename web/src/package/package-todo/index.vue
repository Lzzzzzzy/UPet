<script setup lang="ts">
import { ref, onBeforeMount, reactive, computed } from "vue";
import { formatDatetime, formatTime, isSameDate } from "@/utils";
import { eventCenter, getCurrentInstance, switchTab } from "@tarojs/taro";
import richTextContent from "@/components/rich-text/index.vue";
import dayjs from "dayjs";
import checkedRadio from "@/components/checked-radio/index.vue";
import { addPetTodo, editPetTodo, deletePetTodo } from "@/service/api";

const selectedDate = ref(new Date());

const typeList = [
  { text: "日常记录", value: 0 },
  { text: "待办事项", value: 1 },
];

const showDeleteConfirmPopup = ref(false);

const colorList = ["#ffffff", "#E6A23C", "#F56C6C"]

const pageMode = ref("add");

const formData = reactive({
  title: '',
  todoTime: formatDatetime(new Date()),
  remark: '',
  remind: false,
  remindDate: [dayjs().format('YYYY-MM-DD')],
  remindTime: formatTime(new Date()),
  complete: false,
  type: typeList[0].value,
  color: 0,
  petId: 0,
  id: 0,
})

eventCenter.on('editTodoInfo', async (todoId: number) => {
  editMode.value = true;
  
  const data: any = await getPetTodoInfoById(todoId);
  formData.title = data.title;
  formData.todoTime = data.todoTime;
  formData.remark = data.remark;
  formData.remind = data.remind;
  formData.remindDate = data.remindDate;
  formData.remindTime = data.remindTime;
  formData.complete = data.complete;
  formData.type = data.type;
  formData.color = data.color;
  formData.petId = data.petId;
  formData.id = todoId;
})

const formRef = ref()

const formRules = ref({
  title: [
    { required: true, message: '请填写标题' }
  ]
})


const handleSubmit = async () => {
  const { valid } = await formRef.value?.validate();
  if (valid) {
    if (pageMode.value === 'add') {
      await addPetTodo(formData)
    } else {
      await editPetTodo(formData, formData.id)
    }
    eventCenter.trigger('refreshTodo');
    eventCenter.trigger('refreshDotData');
    switchTab({
      url: '/pages/index/index'
    })
  }
}

const handleDeleteTodo = async () => {
  await deletePetTodoInfoById(formData.id);
  switchTab({
    url: '/pages/index/index'
  });
}


onBeforeMount(() => {
  const instance = getCurrentInstance();
  const params = instance?.router?.params;
  selectedDate.value = params ? dayjs(params.selectedDate).toDate() : new Date();
  if (!isSameDate(selectedDate.value, new Date())) {
    formData.todoTime = formatDatetime(selectedDate.value);
  }
  formData.petId = parseInt(params?.petId || '0');

  eventCenter.on("editTodoData", (data) => {
    pageMode.value = 'edit';
    formData.title = data.title;
    formData.todoTime = formatDatetime(dayjs(data.todoTime).toDate());
    formData.remark = data.remark;
    formData.remind = data.remind;
    formData.remindDate = data.remindDate || [dayjs().format('YYYY-MM-DD')];
    formData.remindTime = formatTime(dayjs(data.remindTime).toDate());
    formData.complete = data.complete;
    formData.type = data.type;
    formData.color = data.color;
    formData.petId = data.petId;
    formData.id = data.id;
    console.log("formData:", formData);
  })
})

// 日期选择
const showDatePicker = ref(false);
const selectedTime = ref(new Date());
const confirmSelectTime = () => {
  formData.todoTime = formatDatetime(selectedTime.value);
  showDatePicker.value = false;
}

// 重复周期选择
const showRepeatDate = ref(false);

const currentDate = ref([new Date()])

const onHandleConfirmSelectDate = () => {
  const remindDate: Array<string> = [];
  currentDate.value.sort((a: Date, b: Date) => a.getTime() - b.getTime());
  currentDate.value.forEach((date: Date)=>{
    remindDate.push(dayjs(date).format('YYYY-MM-DD'))
  })
  formData.remindDate = remindDate;
  showRepeatDate.value = false;
}

const onUpdateRemark = (data: string) => {
  formData.remark = data;
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

const title = computed(() => {
  if (pageMode.value === 'add') {
    return '添加'
  }
  return '编辑'
})

const handleConfirmDelete = async () => {
  await deletePetTodo(formData.id);
  eventCenter.trigger('refreshTodo');
  eventCenter.trigger('refreshDotData');
  switchTab({
    url: '/pages/index/index'
  });
}

const showDeleteConfirmPopup = ref(false);

const handleDelete = () => {
  showDeleteConfirmPopup.value = true;
}
</script>
<template>
  <basic-layout>
    <custom-navbar :title=title left-show />
    <div class="w-full text-20px">
      <nut-form
        ref="formRef"
        :model-value="formData"
        :rules="formRules"
        star-position="right"
      >
        <nut-form-item label="时间" class="form-item-border">
          <div @click="showDatePicker=true" class="w-full">{{ formData.todoTime }}</div>
        </nut-form-item>
        <nut-form-item label="标题" prop="title" class="form-item-border">
          <nut-input v-model="formData.title" placeholder="请输入标题" />
        </nut-form-item>
        <nut-form-item label="备注" prop="remark" class="form-item-border">
          <rich-text-content :data="formData.remark" placeholder="备注，您也可以上传图片,例如确诊单或者希望给医生展示的症状" @update-data="onUpdateRemark" :showUploader="true"></rich-text-content>
        </nut-form-item>
        <nut-form-item label="类型" prop="type" class="form-item-border">
          <nut-radio-group v-model="formData.type">
            <div class="flex items-center">
              <div v-for="(type, index) in typeList" :key="index" class="mr-15px" @click="formData.type=type.value">
                <div :class="{'bg-#f7daa1': formData.type === type.value}" class="mr-5 border-1px b-solid border-color-#f7daa1 px-10px py-2px b-rd-12px">{{ type.text }}</div>
              </div>
            </div>
          </nut-radio-group>
        </nut-form-item>
        <nut-form-item label="颜色" prop="color" class="form-item-border">
          <nut-radio-group v-model="formData.color">
            <div class="flex items-center">
              <div v-for="(color, index) in colorList" :key="index">
                <nut-radio :label="index">
                  <template #icon> <checked-radio :bg-color="color" :checked="false" size="25px" :checked-bg-color="color"></checked-radio> </template>
                  <template #checkedIcon> <checked-radio :bg-color="color" :checked="true" size="25px" :checked-bg-color="color"></checked-radio> </template>
                </nut-radio>
              </div>
            </div>
          </nut-radio-group>
        </nut-form-item>
        <nut-form-item label="需要提醒" prop="remind" class="form-item-border" v-if="formData.type===1">
            <nut-switch v-model="formData.remind" active-color="#f7daa1" />
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
          <nut-button color="#f7daa1" @click="handleSubmit" class="!text-#000000">提交</nut-button>
          <nut-button color="#f56c6c" @click="handleDelete" class="!text-#ffffff" v-if="pageMode === 'edit'">
            删除
          </nut-button>
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
      <nut-popup v-model:visible="showDeleteConfirmPopup" position="bottom" round safe-area-inset-bottom>
        <div class="flex-center my-20px ">
          确定要删除吗?
        </div>
        <div class="flex items-center justify-around mx-20%">
          <nut-button @click="handleConfirmDelete" class="!text-#f56c6c">
            删除
          </nut-button>
          <nut-button @click="showDeleteConfirmPopup = false">
            取消
          </nut-button>
        </div>
      </nut-popup>
    </div>
  </basic-layout>
</template>

<style lang="scss">

.ellipsis-style {
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
}
</style>
