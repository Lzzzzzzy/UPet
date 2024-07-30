<script lang="ts" setup>
import { navigateTo } from '@tarojs/taro';
const addPetText = "为宠物创建一个档案，立刻开启健康记录吧~"

const handleToAddPet = () => {
  navigateTo({
    url: '/package/package-pet/index'
  });
}
</script>

<template>
    <div class="w-full">
        <div class="flex flex-col items-center justify-center">
            <div class="text-160px i-local-cat"></div>
            <div class="text-sm my-4 text-coolgray">{{ addPetText }}</div>
            <nut-button color="#f7daa1" @click="handleToAddPet" class="!text-#000000" shape="round">
                <template #icon>
                    <div class="text-20px i-local-base-pet-icon"></div>
                </template>
                <span class="ml-10px">立即创建</span>
            </nut-button>
        </div>
    </div>
</template>

<style lang="scss">

</style>
