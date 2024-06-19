<template>
    <view>
        <span v-if="iconType === 'feed'" class="feed-icon">
            <My />
        </span>
        <span v-if="iconType === 'purchase'" class="purchase-icon">
            <Cart />
        </span>
        <span v-if="iconType === 'clean'" class="purchase-icon">
            <Del2 />
        </span>
    </view>
</template>

<script setup>
import { My, Cart, Del2 } from '@nutui/icons-vue-taro';

const props = defineProps({
  iconType: {
    type: String,
    required: true,
  }
});
</script>

<style lang="scss">
</style>
