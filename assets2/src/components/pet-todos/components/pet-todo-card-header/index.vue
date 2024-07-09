<script lang="ts" setup>
import petAvatar from '@/components/home/pet-avatar/index.vue';
import { eventCenter, navigateTo } from "@tarojs/taro";
import { ref, onMounted } from 'vue';

const props = defineProps({
  pets: {
    type: Array<Pet.PetInfo>,
    default: () => []
  }
});

const showSelectPetPopup = ref(false);

const selectPet = (pet: Pet.PetInfo) => {
    currentPet.value = pet;
    eventCenter.trigger("selectpet", pet);
    showSelectPetPopup.value = false;
};


const handleToAddPet = () => {
  navigateTo({
    url: '/package/package-add-pet/index'
  });
}

onMounted(() => {
    eventCenter.trigger("selectpet", currentPet.value);
})

const currentPet = ref<Pet.PetInfo>(props.pets[0]);
</script>

<template>
<div>
    <div class="flex items-center my-20px todo-header">
        <nut-button @click="showSelectPetPopup=true">
            <template #default>
                <div class="flex items-center text-16px">
                    <div>{{ currentPet.petName }}</div>
                    <div class="i-local-more-options"></div>
                </div>
            </template>
        </nut-button>
        <!-- <div @click="showSelectPetPopup=true" class="flex items-center bg-#DEDEDE p-7px rounded-50%">
            <div>{{ currentPet.petName }}</div>
            <div class="text-20px i-local-more-options"></div>
        </div> -->
        <div class="text-16px ml-5px">的今日安排</div>
    </div>
    <nut-popup v-model:visible="showSelectPetPopup" position="top">
        <nut-grid :column-num="4"  class="mt-25%">
            <nut-grid-item :text="pet.petName" v-for="pet in pets" :key="pet.id" @click="selectPet(pet)">
                <pet-avatar :avatar-img-url="pet.petAvatar" />
            </nut-grid-item>
            <nut-grid-item text="添加" @click="handleToAddPet">
                <nut-avatar class="!flex justify-center items-center">
                    <div class="text-20px i-local-add"></div>
                </nut-avatar>
            </nut-grid-item>
        </nut-grid>
    </nut-popup>
</div>
</template>

<style lang="scss">
.todo-header {
    .nut-button--normal {
        background-color: #DEDEDE;
        padding: 0 10px;
    }
}
</style>
