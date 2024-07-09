<script setup lang="ts">
import todoCardContent from "@/components/pet-todos/components/pet-todo-card-content/index.vue";
import todoCardHeader from "@/components/pet-todos/components/pet-todo-card-header/index.vue";
import addTodoButton from "@/components/add-pet-todo-button/index.vue";

defineProps({
  todos: {
    type: Array<Pet.PetTodo>,
    default: () => []
  },
  pets: {
    type: Array<Pet.PetInfo>,
    default: () => []
  }
});
</script>
<template>
    <div>
        <div class="pl-15px">
            <todo-card-header :pets="pets"></todo-card-header>
            <todo-card-content v-for="todo in todos" :key="todo.id" :todo="todo"></todo-card-content>
        </div>
        <add-todo-button></add-todo-button>
    </div>
</template>

<style lang="scss">
</style>
