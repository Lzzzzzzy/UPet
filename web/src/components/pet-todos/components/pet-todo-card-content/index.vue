<script lang="ts" setup>
import { PropType } from "vue";
import checkedRadio from "@/components/checked-radio/index.vue";
import { Pet } from "@/typings/pet";
import { formatTime } from "@/utils";
import dayjs from "dayjs";
import type { Dayjs } from "dayjs";
import { updateTodoCompleteStatus } from "@/service/api";
import richTextContent from "@/components/rich-text/index.vue";
import { navigateTo, eventCenter } from "@tarojs/taro";
import { RichText } from "@tarojs/components";

const props = defineProps({
  todo: {
    type: Object as PropType<Pet.PetTodo>,
    default: () => {},
  }
});

const getTimeStr = (timeStr: string | Date | Dayjs) => {
  return formatTime(dayjs(timeStr).toDate())
}

const colorList = ["#ffffff", "#E6A23C", "#F56C6C"]

const getBorderColor = (idx: number) => {
  return colorList[idx];
}

const changeComplete = () => {
  props.todo.complete = !props.todo.complete;
  updateTodoCompleteStatus(props.todo.id!, props.todo.complete);
}

const updateTodo = () => {
  navigateTo({
    url: `/package/package-todo/index`,
    success() {
      eventCenter.trigger('editTodoData', props.todo);
    },
  });
}
</script>

<template>
  <nut-row :gutter="10" class="my-10px" type="flex" justify="space-between">
    <nut-col :span="4">
      <div class="bg-#ffffff py-10px border-rd-md text-center">{{ getTimeStr(todo.todoTime) }}</div>
    </nut-col>
    <nut-col :span="20">
      <div class="flex items-center justify-start">
        <div class="flex flex-col border-rd-md bg-#ffffff pb-10px" @click="updateTodo">
          <div class="dot-container my-4px mr-2px" v-if="todo.color">
            <div :style="{ background: getBorderColor(todo.color) }" class="dot"></div>
          </div>
          <div class="text-16px px-10px fw-600" :class="{'mt-10px': !todo.color}">{{ todo.title }}</div>
          <rich-text :nodes="todo.remark" v-if="todo.remark" class="text-12px break-words pt-10px px-10px"></rich-text>
        </div>
        <div class="mx-10px">
          <checked-radio bg-color="unset" checked-bg-color="#f7daa1" :checked="todo.complete" size="20px" @click="changeComplete"></checked-radio>
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

.dot-container {
  display: flex;
  align-items: center;
  justify-content: flex-end;

  .dot {
      width: 7px;
      height: 7px;
      border-radius: 50%;
      margin: 0 1px;
  }
}
</style>