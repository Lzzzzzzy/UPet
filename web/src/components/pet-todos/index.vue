<script setup lang="ts">
import todoCardContent from "@/components/pet-todos/components/pet-todo-card-content/index.vue";
import todoCardHeader from "@/components/pet-todos/components/pet-todo-card-header/index.vue";
import addTodoButton from "@/components/add-pet-todo-button/index.vue";
import dayjs from "dayjs";
import { computed, PropType } from "vue";
import { Pet } from "@/typings/pet";

const props = defineProps({
  todos: {
    type: Array as PropType<Array<Pet.PetTodo>>,
    default: () => []
  },
  pets: {
    type: Array as PropType<Array<Pet.PetInfo>>,
    default: () => []
  },
  currentDate: {
    type: Date,
    default: () => new Date()
  },
  currentPetId: {
    type: Number,
    required: true,
  }
});

const currentDateString = computed(() => dayjs(props.currentDate).format('YYYY-MM-DD'));
</script>
<template>
    <div>
        <div class="pl-15px">
            <todo-card-header :pets="pets"></todo-card-header>
            <todo-card-content v-for="todo in todos" :key="todo.id" :todo="todo"></todo-card-content>
        </div>
        <add-todo-button :current-date="currentDateString" :pet-id="currentPetId"></add-todo-button>
    </div>
</template>

<style lang="scss">
</style>
