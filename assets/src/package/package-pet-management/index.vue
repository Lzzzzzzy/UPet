<script setup lang="ts">
import petAvatar from '@/components/home/pet-avatar/index.vue';
import { onBeforeMount, ref } from 'vue';
import { eventCenter, navigateTo } from "@tarojs/taro";
import dayjs from 'dayjs';
import { Pet } from "@/typings/pet";
import type { Dayjs } from 'dayjs';


const getPetsInfo = () => {
  return [{
    id: 1,
    name: '旺财',
    familyId: 1,
    avatar: '',
    birthday: "2022-10-02",
    gendar: 0,
    sterilizedState: 0,
    category: 0,
  },
  {
    id: 2,
    name: '帅帅',
    familyId: 1,
    avatar: '',
    birthday: "2021-03-02",
    gendar: 0,
    sterilizedState: 0,
    category: 1,
  },{
    id: 3,
    name: '旺财2',
    familyId: 1,
    avatar: '',
    birthday: "2012-10-02",
    gendar: 0,
    sterilizedState: 1,
    category: 1,
  },{
    id: 4,
    name: '1234',
    familyId: 1,
    avatar: '',
    birthday: "2002-10-02",
    gendar: 1,
    sterilizedState: 1,
    category: 1,
  }
]}

onBeforeMount(() => {
  pets.value = getPetsInfo();
})

const pets = ref<Array<Pet.PetInfo>>([]);

const editPet = (pet: Pet.PetInfo) => {
  navigateTo({url: `/package/package-add-pet/index`, success: () => {
    console.log("pet:", pet, typeof pet);
    eventCenter.trigger("selectEditPet", pet);
  }})
};

const calculateAge = (birthday: Date | string | Dayjs | null | undefined) => {
    if (!birthday) {
        return '未设置';
    }
    var today = dayjs();
    var birthDate = dayjs(birthday).toDate();
    
    const years = today.diff(birthDate, 'year');
    const months = today.diff(birthDate, 'month') % 12;
    
    if (!years) {
        return `${months}个月`;
    }
    return `${years}岁${months}个月`;
}

</script>
<template>
  <basic-layout>
    <custom-navbar title="爱宠档案" left-show />
    <div class="m-20px b-rd-16px py-30px bg-white" v-for="pet in pets" :key="pet.id">
        <nut-row :gutter="10" type="flex" justify="space-around" align="center">
            <nut-col :span="8">
                <div class="flex items-center justify-center">
                    <pet-avatar :avatar-img-url="pet.avatar" size="large" />
                </div>
            </nut-col>
            <nut-col :span="12">
                <div class="flex flex-col items-center justify-between">
                    <div class="text-16px">
                        {{ pet.name }}
                    </div>
                    <div class="mt-20px text-14px">
                        {{ calculateAge(pet.birthday) }}
                    </div>
                </div>
            </nut-col>
            <nut-col :span="4">
                <nut-avatar size="small" class="!flex justify-center items-center !bg-#98c7ce" @click="editPet(pet)">
                    <div class="text-20px i-local-edit"></div>
                </nut-avatar>
            </nut-col>
            <nut-col :span="4">
                <nut-avatar size="small" class="!flex justify-center items-center !bg-#98c7ce" @click="">
                    <div class="text-20px i-local-delete"></div>
                </nut-avatar>
            </nut-col>
        </nut-row>
    </div>
  </basic-layout>
</template>

<style lang="scss">
</style>
