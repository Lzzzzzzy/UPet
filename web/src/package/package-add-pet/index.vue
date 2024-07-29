<script setup lang="ts">
import { switchTab, eventCenter } from '@tarojs/taro';
import { computed, onBeforeMount, ref } from 'vue';
import { Pet } from "@/typings/pet";
import { uploadFileToSystem } from "@/service/api/file";

const title = ref("添加档案");

definePageConfig({
  navigationBarTitleText: title.value
});

onBeforeMount(()=>{
  eventCenter.on("selectEditPet", (pet: Pet.PetInfo) => {
    console.log("on selectEditPet:", typeof pet);
    formData.value.name = pet.name;
    formData.value.avatar = pet.avatar;
    formData.value.gendar = pet.gendar;
    formData.value.sterilizedState = pet.sterilizedState;
    formData.value.category = pet.category;
    formData.value.birthday = pet.birthday;
    formData.value.id = pet.id;
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

const formData = ref<Pet.PetInfo>({
  name: '',
  avatar: '',
  gendar: 0,  // 0 公 1 母
  sterilizedState: 0,  // 0 未绝育 1 已绝育
  category: 0, // petTypes
  birthday: ""
})

const formRef = ref();

const formRules = ref({
  name: [
    { required: true, message: '请填写名字' }
  ]
})

const gendarList = [{"text": "男孩", "value": 0}, {"text": "女孩", "value": 1}]
const sterilizedStateList = [{"text": "未绝育", "value": 0}, {"text": "已绝育", "value": 1}]

const imageUrl = ref("");
const onConfirmAvatar = (url: string) => {
  imageUrl.value = url
};

const editText = computed(() => {
  if (!imageUrl.value) {
    return "选择照片";
  }
  return "";
})

const handleSubmit = async () => {
  const { valid } = await formRef.value?.validate();
  if (valid) {
    const avatarUrl = await uploadFileToSystem(imageUrl.value);
    formData.value.avatar = avatarUrl;
    formData.value.category = selectedPetType.value[0];
    switchTab({
      url: '/pages/index/index'
    })
  }
}
</script>

<template>
  <basic-layout>
    <custom-navbar :title="title" left-show />
    <div class="bg-#ffffff">
      <div class="w-full flex justify-center items-center pt-80px pb-30px flex-col" :class="{'pet-avatar-container': imageUrl}">
        <nut-avatar-cropper shape="round" @confirm="onConfirmAvatar" :edit-text="editText">
          <nut-avatar size="100" shape="round">
            <img v-if="imageUrl" :src="imageUrl" />
          </nut-avatar>
        </nut-avatar-cropper>
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
    </div>
  </basic-layout>
</template>

<style lang="scss">
.form-item-border {
  border-bottom: 1px solid var(--nut-input-border-bottom, #eaf0fb)
}

.pet-avatar-container {
  .nut-avatar-cropper__edit-text {
    background: initial;
  }
}
</style>
