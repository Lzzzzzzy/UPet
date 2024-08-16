<script setup lang="ts">
import petAvatar from '@/components/home/pet-avatar/index.vue';
import { onBeforeMount, ref } from 'vue';
import { eventCenter, navigateTo } from "@tarojs/taro";
import dayjs from 'dayjs';
import type { Pet } from "@/typings/pet";
import type { Dayjs } from 'dayjs';
import { getAllPetsInfo } from '@/service/api';


onBeforeMount(() => {
  getPetsInfo();
})

const pets = ref<Array<Pet.PetInfo>>([]);

const getPetsInfo = async () => {
  const petInfos = await getAllPetsInfo()
  if (petInfos) {
    pets.value = petInfos;
  }
}

const editPet = (pet: Pet.PetInfo) => {
  navigateTo({url: `/package/package-pet/index`, success: () => {
    eventCenter.trigger("selectEditPet", pet);
  }})
};

const addPet = () => {
    navigateTo({url: `/package/package-pet/index`})
}

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

const showDeleteConfirmPopup = ref(false);
const needDeletePetId = ref();
const deletePet = (petId: number | undefined) => {
  needDeletePetId.value = petId;
  showDeleteConfirmPopup.value = true;
}
const closePopup = () => {
  showDeleteConfirmPopup.value = false;
  needDeletePetId.value = null;
}
const confirmDelete = () => {
  closePopup();
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
            <nut-col :span="8">
                <div class="flex flex-col items-start justify-start">
                    <div class="text-20px">
                        {{ pet.name }}
                    </div>
                    <div class="mt-20px text-12px text-coolgray">
                        {{ calculateAge(pet.birthday) }}
                    </div>
                </div>
            </nut-col>
            <nut-col :span="4">
                <nut-avatar size="small" class="!flex-center !bg-#98c7ce" @click="editPet(pet)">
                    <div class="text-20px i-local-edit"></div>
                </nut-avatar>
            </nut-col>
            <nut-col :span="4">
                <nut-avatar size="small" class="!flex-center !bg-#98c7ce" @click="deletePet(pet.id)">
                    <div class="text-20px i-local-delete"></div>
                </nut-avatar>
            </nut-col>
        </nut-row>
    </div>

    <div class="flex-center">
        <nut-button @click="addPet" color="#f7daa1" class="!text-black">
            <div class="flex-center">
                <div class="text-15px i-local-add pr-5px"></div>
                <div>添加</div>
            </div>
        </nut-button>
    </div>

    <nut-popup v-model:visible="showDeleteConfirmPopup" round :style="{'width': '70%'}">
      <div class="px-20px py-40px">
        <div class="flex-col-center">
          <div>删除档案后无法恢复</div>
          <div>确定要删除吗?</div>
        </div>
        <div class="mt-20px flex justify-evenly">
          <nut-button plain @click="confirmDelete">确定</nut-button>
          <nut-button @click="closePopup" color="#f7daa1" class="!text-black">取消</nut-button>
        </div>
      </div>
  </nut-popup>
  </basic-layout>
</template>

<style lang="scss">
</style>
