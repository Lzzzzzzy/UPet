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

const formData = ref({
  name: '',
  file: '',
  gendar: 0,
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
          placeholder="怎么称呼孩子?"
          type="text"
        />
      </nut-form-item>
      <nut-form-item label="宠物种类" prop="type">
        <nut-cell :title="petTypes[selectedPetType[0]].text" desc=">" @click="showTypePicker = true" class="!p-0"></nut-cell>
        <nut-popup v-model:visible="showTypePicker" position="bottom">
          <nut-picker v-model="selectedPetType" :columns="petTypes" @confirm="showTypePicker = false" @cancel="showTypePicker = false" />
        </nut-popup>
      </nut-form-item>
      <nut-form-item label="性别" prop="type">
        <!-- <nut-grid :gutter="10" clickable>
          <nut-grid-item text="男孩" :class="{'bg-#665D21': formData.gendar === 0,}"></nut-grid-item>
          <nut-grid-item text="女孩" :class="{'bg-#665D21': formData.gendar === 1}"></nut-grid-item>
        </nut-grid> -->
        <div :class="{'bg-#665D21': formData.gendar === 0, 'text-white': formData.gendar === 0}" class="w-20">男孩</div>
        <div :class="{'bg-#665D21': formData.gendar === 1, 'text-white': formData.gendar === 1}" class="w-20">女孩</div>
      </nut-form-item>
      <nut-form-item label="绝育状态" prop="type">
        <nut-cell :title="petGendar[selectedPetGendar[0]].text" desc=">" @click="showGendarPicker = true" class="!p-0"></nut-cell>
        <nut-popup v-model:visible="showGendarPicker" position="bottom">
          <nut-picker v-model="selectedPetGendar" :columns="petGendar" @confirm="showGendarPicker = false" @cancel="showGendarPicker = false" />
        </nut-popup>
      </nut-form-item>

      <nut-space style="margin: 10px">
        <nut-button type="primary" size="small" @click="submit">提交</nut-button>
      </nut-space>
    </nut-form>
  </basic-layout>
</template>

<style lang="scss"></style>
