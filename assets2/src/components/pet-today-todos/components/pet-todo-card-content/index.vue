<script lang="ts" setup>
defineProps({
  todo: {
    type: Object,
    default: () => {},
  }
});
</script>

<template>
  <nut-row :gutter="10" class="my-10px" type="flex" justify="space-between">
    <nut-col :span="5">
      <div class="bg-#ffffff py-10px border-rd-md text-center">{{ todo.time }}</div>
    </nut-col>
    <nut-col :span="19">
      <div class="flex items-center justify-start">
        <div class="flex flex-col p-10px border-rd-md bg-#ffffff">
          <div class="font-size-35">{{ todo.title }}</div>
          <div class="font-size-23 text-#787263 mt-10px break-words" v-if="todo.remark">{{ todo.remark }}</div>
        </div>
        <nut-checkbox v-model="todo.status" class="!mr-0 ml-10px">
          <template #icon> <div class="i-local-unchecked text-25px"></div> </template>
          <template #checkedIcon> <div class="i-local-checked text-25px"></div> </template>
        </nut-checkbox>
      </div>
    </nut-col>
  </nut-row>
</template>

<style lang="scss">
.todo-checkbox {
  .nut-checkbox__icon {
    // color: #665D21;
    color: #B59F72;
  }
}
</style>