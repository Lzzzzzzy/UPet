<script lang="ts" setup>
import { navigateTo } from "@tarojs/taro";

const props = defineProps({
  currentDate: {
    type: String,
    required: true,
  }
});

const handleToAddTodo = () => {
  navigateTo({
    url: `/package/package-add-todo/index?selectedDate=${props.currentDate}`
  });
}
</script>
<template>
    <div>
        <nut-avatar class="!flex justify-center items-center !bg-#98c7ce !position-fixed position-bottom-100px position-right-30px" size="large" @click="handleToAddTodo">
            <div class="text-30px i-local-add text-black"></div>
        </nut-avatar>
    </div>
</template>
<style lang="scss">
.form-item-border {
    border-bottom: 1px solid var(--nut-input-border-bottom, #eaf0fb)
  }
</style>