<script lang="ts" setup>
import { navigateTo } from "@tarojs/taro";

const handleToAddTodo = () => {
  navigateTo({
    url: '/package/package-add-todo/index'
  });
}
</script>
<template>
    <div>
        <nut-avatar class="!flex justify-center items-center !bg-#F7DAA1 !position-fixed position-bottom-100px position-right-30px" size="large" @click="handleToAddTodo">
            <div class="text-30px i-local-add"></div>
        </nut-avatar>
    </div>
</template>
<style lang="scss">
.form-item-border {
    border-bottom: 1px solid var(--nut-input-border-bottom, #eaf0fb)
  }
</style>