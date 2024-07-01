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
    eventCenter.trigger("selectpet", pet);
    showSelectPetPopup.value = false;
};

const getBatchPets = () => {  
    // 计算结束索引，如果数组长度小于3，则取数组长度作为结束索引  
    var endIndex = Math.min(props.pets.length, 1);  
    // 使用slice方法获取元素  
    return props.pets.slice(0, endIndex);  
}

const handleToAddPet = () => {
  navigateTo({
    url: '/package/package-add-pet/index'
  });
}

onMounted(() => {
    eventCenter.trigger("selectpet", props.pets[0]);
})
</script>

<template>
<div>
    <div class="flex justify-center my-20px">
        <nut-avatar-group max-count="2" span="-18" zIndex="right" @click="showSelectPetPopup = true" max-content="...">
            <pet-avatar :avatar-img-url="pet.petAvatar" v-for="pet in getBatchPets()" :key="pet.id" />
            <nut-avatar class="!flex justify-center items-center">
                <div class="text-20px i-local-add"></div>
            </nut-avatar>
        </nut-avatar-group>
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
</style>
