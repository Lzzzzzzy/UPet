<script setup lang="ts">
import fullCalendar from '@/components/calendar/full-calendar/index.vue';

definePageConfig({
  navigationBarTitleText: '日历'
});
</script>
<template>
  <basic-layout show-tab-bar>
    <custom-navbar title="日历" />
    <full-calendar></full-calendar>
  </basic-layout>
</template>

<style lang="scss"></style>
