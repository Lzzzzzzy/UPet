<script lang="ts" setup>
defineProps({
  todo: {
    type: Object,
    default: () => {},
  }
});
</script>

<template>
  <nut-row :gutter="10" class="mb-20px" type="flex" justify="space-between">
    <nut-col :span="5">
      <div class="bg-#ffffff py-10px border-rd-md text-#665D21 text-center">{{ todo.time }}</div>
    </nut-col>
    <nut-col :span="19">
      <nut-row :gutter="10" type="flex" justify="space-between" align="center">
        <nut-col :span="21">
          <div class="flex items-center justify-between p-10px border-rd-md bg-#ffffff">
          <div class="flex flex-col">
            <div class="font-size-35">{{ todo.title }}</div>
              <div class="font-size-23 text-#787263 mt-10px break-words" v-if="todo.remark">{{ todo.remark }}</div>
            </div>
          </div>
        </nut-col>
        <nut-col :span="3">
          <nut-checkbox v-model="todo.status" class="!mr-0"></nut-checkbox>
        </nut-col>
      </nut-row>
    </nut-col>
  </nut-row>
</template>

<style lang="scss">

</style>