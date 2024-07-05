<script setup lang="ts">
import { ref } from "vue"
import month from './month.vue'
import week from './week.vue'
import dayjs from 'dayjs';
import type { Dayjs } from 'dayjs'

const currentDate = ref(dayjs())

</script>
<template>
    <div class="mt-50px">
        <div class="calendar">
            <week v-model="currentDate"></week>
            <month></month>
        </div>
        <p>选中时间是: {{ currentDate }}</p>
    </div>
</template>