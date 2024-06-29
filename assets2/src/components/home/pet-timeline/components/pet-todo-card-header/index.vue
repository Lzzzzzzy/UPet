<script lang="ts" setup>
import petAvatar from '@/components/home/pet-avatar/index.vue';
import { Events } from "@tarojs/taro";
import { ref } from 'vue';

const props = defineProps({
  pets: {
    type: Array<Pet.PetInfo>,
    default: () => []
  }
});

const events = new Events();
const showSelectPetPopup = ref(false);
const selectedPet = ref<Pet.PetInfo>(props.pets[0]);

const selectPet = (pet: Pet.PetInfo) => {
    selectedPet.value = pet;
    events.trigger("selectPet", pet);
    showSelectPetPopup.value = false;
};
</script>

<template>
    <div class="flex justify-center">
        <nut-avatar-group max-count="3" z-index="right" max-content="..." span="-16" @click="showSelectPetPopup = true">
            <pet-avatar :avatar-img-url="pet.petAvatar" v-for="pet in pets" :key="pet.id" />
            <nut-avatar class="!flex justify-center items-center">
                <div class="text-20px i-local-add"></div>
            </nut-avatar>
        </nut-avatar-group>
    </div>
    <div class="fw-600 text-18px text-color-#665d21">{{ selectedPet.petName }}的今日安排</div>
    <nut-popup v-model:visible="showSelectPetPopup">
        <nut-grid :column-num="3">
            <nut-grid-item :text="pet.petName" v-for="pet in pets" :key="pet.id">
                <pet-avatar :avatar-img-url="pet.petAvatar" @click="selectPet(pet)" />
            </nut-grid-item>
            <nut-grid-item text="添加">
                <nut-avatar class="!flex justify-center items-center">
                    <div class="text-20px i-local-add"></div>
                </nut-avatar>
            </nut-grid-item>
        </nut-grid>
    </nut-popup>
</template>

<style lang="scss">
</style>
