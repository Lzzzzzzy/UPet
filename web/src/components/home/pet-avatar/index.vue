<script setup lang="ts">
defineProps({
  avatarImgUrl: {
    type: String,
    default: ""
  },
  size : {
    type: String,
    default: "normal"
  },
});
</script>

<template>
  <div class="pet-avatar">
    <div v-if="avatarImgUrl">
      <nut-avatar :size="size">
        <img :src="avatarImgUrl" />
      </nut-avatar>
    </div>
    <div v-else>
      <nut-avatar :size="size" class="items-center justify-center">
        <div class="i-local-base-pet-icon" :class="{'text-20px': size === 'normal', 'text-30px': size === 'large', 'text-15px': size === 'small'}"></div>
      </nut-avatar>
    </div>
  </div>
</template>

<style lang="scss">
.pet-avatar {
  .nut-avatar {
    display: flex;
    background-color: #F7DAA1 !important;
  }
}
</style>
