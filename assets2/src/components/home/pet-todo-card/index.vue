<script lang="ts" setup>
import petAvatar from './pet-avatar/index.vue';
import { ref } from 'vue';

const props = defineProps({
  pet: {
    type: Object,
    default: () => {}
  }
});

const petName = ref('');
const todos = ref([]);

if (props.pet) {
  petName.value = props.pet.name;
//   const petId = props.pet.id;
//   todos.value = [
//     {
//       id: 1,
//       petId: petId,
//       content: '疫苗',
//       date: '2023-01-01',
//       remind: true,
//       remindTime: '09:00',
//       remindType: '1',
//       remindContent: '疫苗提醒'
//     }]
}
const noTodosText = "可以添加宠物的日程，例如疫苗接种等，您会在指定的时间收到提醒。也可以添加一条不需要提醒的记录。";
</script>
<template>
    <div class="w-full rd-8px default-border-style">
        <div class="flex justify-start items-center p-12px border-b-1 border-b-solid border-color-#665D21">
            <pet-avatar></pet-avatar>
            <span class="ml-15px fw-600 text-18px text-color-#665d21">{{ petName }}的今日安排</span>
        </div>
        <div class="p-3">
            <div v-if="todos.length"></div>
            <div v-else>{{ noTodosText }}</div>
        </div>
        <div class="flex justify-center items-center w-full my-10px">
            <nut-button shape="round" class="bg-transparent border-1px b-solid border-color-#665D21">
                <template #icon>
                    <div class="text-20px i-local-base-pet-icon"></div>
                </template>
                <span class="ml-10px">添加</span>
            </nut-button>
        </div>
    </div>
</template>
<style lang="scss">
.border-default-color {
    border-color: #665D21;
}
</style>
