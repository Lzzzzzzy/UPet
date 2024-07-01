<script lang="ts" setup>
const props = defineProps({
  todo: {
    type: Object<Pet.PetTodo>,
    default: () => {},
  }
});
</script>

<template>
    <div class="flex items-start mb-20px">
        <div class="bg-#ffffff p-10px border-rd-md mr-15px text-#665D21">{{ todo.time }}</div>
        <div class="flex items-center">
          <div class="flex items-center justify-between p-10px border-rd-md bg-#ffffff mr-15px">
              <div class="flex flex-col">
                  <div class="font-size-35">{{ todo.title }}</div>
                  <div class="font-size-23 text-#787263 mt-10px" v-if="todo.remark">{{ todo.remark }}</div>
              </div>
          </div>
          <nut-checkbox v-model="todo.status" class="!mr-0"></nut-checkbox>
        </div>
    </div>
</template>

<style lang="scss">

</style>