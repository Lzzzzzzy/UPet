<script setup lang="ts">
import todoCardContent from "@/components/pet-todos/components/pet-todo-card-content/index.vue";
import todoCardHeader from "@/components/pet-todos/components/pet-todo-card-header/index.vue";
import addTodoButton from "@/components/add-pet-todo-button/index.vue";
import dayjs from "dayjs";
import { computed, PropType, ref } from "vue";
import { Pet } from "@/typings/pet";
import { deletePetTodo } from "@/service/api";
import { eventCenter } from "@tarojs/taro";

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

const needDeleteTodoId = ref<number | null>(null);

const loading = ref(false);

const handleConfirmDelete = async () => {
  loading.value = true;
  await deletePetTodo(needDeleteTodoId.value);
  eventCenter.trigger('refreshTodo');
  eventCenter.trigger('refreshDotData');
  loading.value = false;
  showDeleteConfirmPopup.value = false;
}

const showDeleteConfirmPopup = ref(false);

const handleDelete = () => {
  showDeleteConfirmPopup.value = true;
}

eventCenter.on('deleteTodo', (todoId: number) => {
  needDeleteTodoId.value = todoId;
  handleDelete()
})

</script>
<template>
  <div>
    <div class="pl-15px">
      <todo-card-header :pets="pets"></todo-card-header>
      <todo-card-content v-for="todo in todos" :key="todo.id" :todo="todo"></todo-card-content>
    </div>
    <add-todo-button :current-date="currentDateString" :pet-id="currentPetId"></add-todo-button>
    <nut-popup v-model:visible="showDeleteConfirmPopup" round class="popup-container">
      <div class="flex-center mb-20px ">
        确定要删除吗?
      </div>
      <div class="flex items-center justify-around">
        <nut-button @click="handleConfirmDelete" class="!text-#f56c6c" :loading="loading">
          删除
        </nut-button>
        <nut-button @click="showDeleteConfirmPopup = false">
          取消
        </nut-button>
      </div>
    </nut-popup>
  </div>
</template>

<style lang="scss">
.popup-container {
  .nut-popup {
    width: 60%;
    padding: 20px;
  }
}
</style>
