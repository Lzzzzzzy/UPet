<script setup lang="ts">
definePageConfig({
  navigationBarTitleText: '添加档案'
});

import { ref } from 'vue'

const showTypePicker = ref(false);

const petTypes = ref([
  { text: '猫猫', value: 0 },
  { text: '狗狗', value: 1 },
])

const selectedPetType = ref([0])
const selectedPetTypeText = ref(petTypes.value[selectedPetType.value[0]].text)

const formData = ref({
  name: '',
  file: '',
})

const formRef = ref()

const reset = () => {
  formRef.value?.reset()
}

const submit = () => {
  formRef.value?.validate().then(({ valid, errors }) => {
    if (valid) {
      console.log('success:', formData.value)
    } else {
      console.warn('error:', errors)
    }
  })
}

const formRules = ref({
  name: [
    { required: true, message: '请填写名字' }
  ]
})

const confirmUpdatePetType = ({ selectedValue, selectedOptions }) => {
  console.log(selectedValue[0], selectedOptions[0])
  showTypePicker.value = false
}

</script>
<template>
  <basic-layout>
    <custom-navbar title="添加档案" left-show />
    <nut-uploader
      url="http://服务地址"
      is-preview
    >
    </nut-uploader>
    <nut-form
      ref="formRef"
      :model-value="formData"
      :rules="formRules"
    >
      <nut-form-item label="名字" prop="name">
        <nut-input
          v-model="formData.name"
          placeholder="请输入宠物名字"
          type="text"
        />
      </nut-form-item>
      <nut-form-item label="宠物种类" prop="type">
        <nut-cell :desc="petTypes[selectedPetType[0]].text" @click="showTypePicker = true" class="!contents"></nut-cell>
        <nut-popup v-model:visible="showTypePicker" position="bottom">
          <nut-picker v-model="selectedPetType" :columns="petTypes" @confirm="confirmUpdatePetType" @cancel="showTypePicker = false" />
        </nut-popup>
      </nut-form-item>

      <nut-space style="margin: 10px">
        <nut-button type="primary" size="small" @click="submit">提交</nut-button>
      </nut-space>
    </nut-form>
  </basic-layout>
</template>

<style lang="scss"></style>
