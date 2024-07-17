<script setup lang="ts">
import { onBeforeMount, ref } from 'vue';
import { eventCenter, navigateTo } from "@tarojs/taro";

const getPetsInfo = () => {
  return [{
    id: 1,
    petName: '旺财',
    petAvatar: '',
    userId: 1,
  },
  {
    id: 2,
    petName: '帅帅',
    petAvatar: '',
    userId: 1,
  },{
    id: 3,
    petName: '1234',
    petAvatar: '',
    userId: 1,
  },{
    id: 4,
    petName: '1123',
    petAvatar: '',
    userId: 1,
  }
]}

onBeforeMount(() => {
  pets.value = getPetsInfo();
})

const pets = ref<Array<Pet.PetInfo>>([]);

const editPet = (pet: Pet.PetInfo) => {
  navigateTo({url: `/package/package-add-pet/index`, success: () => {
    eventCenter.trigger("selectPet", pet)
  }})
};
</script>
<template>
  <basic-layout>
    <custom-navbar title="爱宠档案" left-show />
    <div class="w-full" v-for="pet in pets" :key="pet.id">
      <div @click="editPet(pet)">
        <pet-avatar :avatar-img-url="pet.petAvatar" />
        <div>{{ pet.petName }}</div>
      </div>
    </div>
  </basic-layout>
</template>

<style lang="scss"></style>
