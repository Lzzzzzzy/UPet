<script setup lang="ts">
import todoCardHeader from "@/components/home/pet-today-todos/components/pet-todo-card-header/index.vue";
import todoCardContent from "@/components/home/pet-today-todos/components/pet-todo-card-content/index.vue";
import addTodoButton from "@/components/add-pet-todo-button/index.vue";
import { ref, onBeforeMount } from "vue";
import { eventCenter } from "@tarojs/taro";

defineProps({
  pets: {
    type: Array<Pet.PetInfo>,
    default: () => []
  }
});

const petTodos = ref()

const getPetTodos = () => {
    petTodos.value = [{"id": 1, "time": "08:00", "title": "换水", "remark": "全部两个碗里的倒矿泉水纯净水蒸馏水都行,就是不能用自来水", "status": false},{"id": 2, "time": "12:00", "title": "换粮", "remark": "全换", "status": false},{"id": 3, "time": "18:00", "title": "铲屎", "remark": "", "status": false}] 
}

getPetTodos();

const selectedPet = ref<Pet.PetInfo>();

onBeforeMount(() => {
    eventCenter.on("selectpet", (pet: Pet.PetInfo) => {
        selectedPet.value = pet
    });
});

</script>
<template>
    <div>
        <todo-card-header :pets="pets"></todo-card-header>
        <div>
            <div class="fw-600 text-18px text-#665d21 mb-25px" v-if="selectedPet">{{ selectedPet.petName }}的今日安排</div>
            <todo-card-content v-for="todo in petTodos" :key="todo.id" :todo="todo"></todo-card-content>
        </div>
        <add-todo-button></add-todo-button>
    </div>
</template>

<style lang="scss">
.step-container {
    .nut-step-main {
        width: 100%;
    }
}
</style>
