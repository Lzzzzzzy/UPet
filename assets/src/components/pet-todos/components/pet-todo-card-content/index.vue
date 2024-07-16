<script lang="ts" setup>
import { PropType } from "vue";
import checkedRadio from "@/components/checked-radio/index.vue";

defineProps({
  todo: {
    type: Array as PropType<Array<Pet.PetTodo>>,
    default: () => {},
  }
});

</script>

<template>
  <nut-row :gutter="10" class="my-10px" type="flex" justify="space-between">
    <nut-col :span="4">
      <div class="bg-#ffffff py-10px border-rd-md text-center">{{ todo.time }}</div>
    </nut-col>
    <nut-col :span="20">
      <div class="flex items-center justify-start">
        <div class="flex flex-col p-10px border-rd-md bg-#ffffff">
          <div class="text-16px">{{ todo.title }}</div>
          <div class="text-12px mt-10px break-words" v-if="todo.remark">{{ todo.remark }}</div>
        </div>
        <div class="mx-10px">
          <checked-radio bg-color="unset" checked-bg-color="#f7daa1" :checked="todo.complete" size="20px" @click="todo.complete = !todo.complete"></checked-radio>
        </div>
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