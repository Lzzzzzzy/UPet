<template>
  <div>
    <van-calendar
      title="日历"
      :poppable="false"
      :show-confirm="false"
      class="calendar"
    />

  </div>
</template>

<script>

export default {

  data () {
    return {
    }
  },

  created () {
  }
}
</script>

<style lang="scss" scoped>
.calendar {
  --calendar-height: 618px;
}
</style>
