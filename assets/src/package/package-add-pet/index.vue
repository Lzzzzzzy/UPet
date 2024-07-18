<script setup lang="ts">
import { switchTab, eventCenter } from '@tarojs/taro';
import { onBeforeMount, ref, reactive } from 'vue';
import { Pet } from "@/typings/pet";

const title = ref("添加档案");

definePageConfig({
  navigationBarTitleText: title.value
});

onBeforeMount(()=>{
  eventCenter.on("selectEditPet", (pet: Pet.PetInfo) => {
    console.log("on selectEditPet:", typeof pet);
    formData.name = pet.name;
    formData.avatar = pet.avatar;
    formData.gendar = pet.gendar;
    formData.sterilizedState = pet.sterilizedState;
    formData.category = pet.category;
    formData.birthday = pet.birthday;
    formData.id = pet.id;
    title.value = "修改档案"
    console.log("formData:", formData);
  });
})

const showTypePicker = ref(false);

const petTypes = ref([
  { text: '猫猫', value: 0 },
  { text: '狗狗', value: 1 },
])
const selectedPetType = ref([0])

const formData = reactive<Pet.PetInfo>({
  name: '',
  avatar: '',
  gendar: 0,  // 0 公 1 母
  sterilizedState: 0,  // 0 未绝育 1 已绝育
  category: 0, // petTypes
  birthday: ""
})

const formRef = ref()

const formRules = ref({
  name: [
    { required: true, message: '请填写名字' }
  ]
})

const handleSubmit = () => {
  formRef.value?.validate().then(({ valid, errors }: { valid: boolean, errors: any}) => {
    if (valid) {
      formData.category = selectedPetType.value[0];
      switchTab({
        url: '/pages/index/index'
      })
    } else {
      console.warn('error:', errors)
    }
  })
}

const uploadUrl = ref("");

const gendarList = [{"text": "男孩", "value": 0}, {"text": "女孩", "value": 1}]
const sterilizedStateList = [{"text": "未绝育", "value": 0}, {"text": "已绝育", "value": 1}]
</script>

<template>
  <basic-layout>
    <custom-navbar :title="title" left-show />
    <div class="w-full flex justify-center items-center mt-80px mb-30px flex-col">
      <nut-uploader
        :url="uploadUrl"
        is-preview
      >
      <div class="w-80px h-80px border-rd-50% flex justify-center items-center bg-#ffffff border-1px b-solid border-color-#f7daa1">上传照片</div>
      </nut-uploader>
      <div class="text-sm my-4 text-coolgray">上传一张孩子最可爱的照片作为头像</div>
    </div>
    <nut-form
      ref="formRef"
      :model-value="formData"
      :rules="formRules"
      star-position="right"
    >
      <nut-form-item label="名字" prop="name" class="form-item-border">
        <nut-input
          v-model="formData.name"
          placeholder="怎么称呼孩子?"
          type="text"
        />
      </nut-form-item>
      <nut-form-item label="种类" prop="type" class="form-item-border">
        <nut-cell :title="petTypes[selectedPetType[0]].text" @click="showTypePicker = true" class="!p-0">
          <template #desc>
            <div class="i-ph-caret-right-bold text-20px text-#000"></div>
          </template>
        </nut-cell>
        <nut-popup v-model:visible="showTypePicker" position="bottom">
          <nut-picker v-model="selectedPetType" :columns="petTypes" @confirm="showTypePicker = false" @cancel="showTypePicker = false" />
        </nut-popup>
      </nut-form-item>
      <nut-form-item label="性别" prop="type" class="form-item-border">
        <div class="flex">
          <div v-for="(gendar, index) in gendarList" :key="index" class="mr-15px" @click="formData.gendar=gendar.value">
            <div :class="{'bg-#f7daa1': formData.gendar === gendar.value}" class="mr-5 border-1px b-solid border-color-#f7daa1 px-10px py-2px b-rd-12px">{{ gendar.text }}</div>
          </div>
        </div>
      </nut-form-item>
      <nut-form-item label="绝育" prop="type" class="form-item-border">
        <div class="flex">
          <div v-for="(sterilizedState, index) in sterilizedStateList" :key="index" class="mr-15px" @click="formData.sterilizedState=sterilizedState.value">
            <div :class="{'bg-#f7daa1': formData.sterilizedState === sterilizedState.value}" class="mr-5 border-1px b-solid border-color-#f7daa1 px-10px py-2px b-rd-12px">{{ sterilizedState.text }}</div>
          </div>
        </div>
      </nut-form-item>
      <nut-form-item label="年龄" prop="type" class="form-item-border">
        <div class="flex">
          <div v-for="(sterilizedState, index) in sterilizedStateList" :key="index" class="mr-15px" @click="formData.sterilizedState=sterilizedState.value">
            <div :class="{'bg-#f7daa1': formData.sterilizedState === sterilizedState.value}" class="mr-5 border-1px b-solid border-color-#f7daa1 px-10px py-2px b-rd-12px">{{ sterilizedState.text }}</div>
          </div>
        </div>
      </nut-form-item>

      <nut-space class="m-10px flex justify-center w-full">
        <nut-button type="primary" @click="handleSubmit" class="!text-black" color="#f7daa1">提交</nut-button>
      </nut-space>
    </nut-form>
  </basic-layout>
</template>

<style lang="scss">
.form-item-border {
  border-bottom: 1px solid var(--nut-input-border-bottom, #eaf0fb)
}
</style>
