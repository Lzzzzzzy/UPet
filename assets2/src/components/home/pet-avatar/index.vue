<script setup lang="ts">
import { ref } from 'vue';
const avatarImgUrl = ref('');
</script>

<template>
  <div class="pet-avatar">
    <div v-if="avatarImgUrl">
      <nut-avatar size="normal">
        <img :src="avatarImgUrl" />
      </nut-avatar>
    </div>
    <div v-else>
      <nut-avatar size="normal" class="items-center justify-center">
        <div class="text-20px i-local-base-pet-icon"></div>
      </nut-avatar>
    </div>
  </div>
</template>

<style lang="scss">
.pet-avatar {
  .nut-avatar {
    display: flex;
    background-color: rgba(237, 228, 149, 1) !important;
  }
}
</style>
