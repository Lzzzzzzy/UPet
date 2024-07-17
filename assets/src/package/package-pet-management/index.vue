<script setup lang="ts">
import { onBeforeMount } from 'vue';

const getPetsInfo = () => {
  return [{
    id: 1,
    petName: '旺财',
    petAvatar: '',
    userId: 1,
  },
  {
    id: 2,
    petName: '帅帅',
    petAvatar: '',
    userId: 1,
  },{
    id: 3,
    petName: '1234',
    petAvatar: '',
    userId: 1,
  },{
    id: 4,
    petName: '1123',
    petAvatar: '',
    userId: 1,
  }
]}

onBeforeMount(() => {
  getPetsInfo();
})
</script>
<template>
  <basic-layout>
    <custom-navbar title="爱宠档案" left-show />
    <div class="w-375px text-30px text-primary">档案</div>
  </basic-layout>
</template>

<style lang="scss"></style>
