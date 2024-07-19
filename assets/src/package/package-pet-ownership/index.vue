<script setup lang="ts">
import { onBeforeMount, ref } from 'vue';
import type { UserInfo } from "@/typings/family";

const getAllFamilyMember = () => {
  return [
    {
      "id": 1,
      "name": "爸爸",
      "isAdmin": false,
      "avatar": ""
    },
    {
      "id": 2,
      "name": "妈妈",
      "isAdmin": false,
      "avatar": ""
    },
    {
      "id": 3,
      "name": "对象",
      "isAdmin": true,
      "avatar": ""
    },
]
}

const familyMembers = ref<Array<UserInfo>>();

const mainSpanWidth = ref(8);

onBeforeMount(() => {
  familyMembers.value = getAllFamilyMember();
  if (!me.isAdmin) {
    mainSpanWidth.value = 16;
  }
})

const me = {
  "id": 4,
  "name": "我自己",
  "isAdmin": true,
  "avatar": ""
}

const showDeleteConfirmPopup = ref(false);
const needDeleteMemberId = ref();
const deleteMember = (memberId: number) => {
  showDeleteConfirmPopup.value = true;
  needDeleteMemberId.value = memberId;
}
const closeMemberPopup = () => {
  showDeleteConfirmPopup.value = false;
  needDeleteMemberId.value = null;
}
const confirmDelete = () => {
  closeMemberPopup();
}

const addMember = () => {};


const showGrantConfirmPopup = ref(false);
const needGrantMemberId = ref();
const grantMember = (memberId: number) => {
  showGrantConfirmPopup.value = true;
  needGrantMemberId.value = memberId;
}
const closeGrantPopup = () => {
  showGrantConfirmPopup.value = false;
  needGrantMemberId.value = null;
}
const confirmGrant = () => {
  closeGrantPopup();
}
</script>
<template>
<basic-layout>
  <custom-navbar title="爱宠共养" left-show />
  <div class="m-20px b-rd-16px py-30px bg-white" v-for="member in familyMembers" :key="member.id">
      <nut-row :gutter="10" type="flex" justify="space-around" align="center">
          <nut-col :span="8">
              <div class="flex items-center justify-center">
                  <pet-avatar :avatar-img-url="member.avatar" size="large" />
              </div>
          </nut-col>
          <nut-col :span="mainSpanWidth">
              <div class="flex flex-col items-start justify-center">
                  <div class="text-20px">
                      {{ member.name }}
                  </div>
                  <div class="flex-center mt-10px">
                    <div class="i-local-key text-15px ml-5px" v-if="member.isAdmin"></div>
                  </div>
              </div>
          </nut-col>
          <nut-col :span="4" v-if="me.isAdmin">
              <nut-avatar size="small" class="!flex-center !bg-#98c7ce" @click="grantMember(member.id)">
                  <div class="text-20px i-local-grant"></div>
              </nut-avatar>
          </nut-col>
          <nut-col :span="4" v-if="me.isAdmin">
            <nut-avatar size="small" class="!flex-center !bg-#98c7ce" @click="deleteMember(member.id)">
              <div class="text-20px i-local-exit"></div>
            </nut-avatar>
          </nut-col>
      </nut-row>
  </div>

  <div class="flex-center">
      <nut-button @click="addMember" color="#f7daa1" class="!text-black">
          <div class="flex-center">
              <div class="text-15px i-local-add pr-5px"></div>
              <div>邀请家庭成员</div>
          </div>
      </nut-button>
  </div>

  <nut-popup v-model:visible="showDeleteConfirmPopup" round :style="{'width': '70%'}">
    <div class="px-20px py-40px">
      <div class="flex-col-center">
        <div>确定要移除家庭成员吗?</div>
      </div>
      <div class="mt-20px flex justify-evenly">
        <nut-button plain @click="confirmDelete">确定</nut-button>
        <nut-button @click="closeMemberPopup" color="#f7daa1" class="!text-black">取消</nut-button>
      </div>
    </div>
  </nut-popup>

  <nut-popup v-model:visible="showGrantConfirmPopup" round :style="{'width': '70%'}">
    <div class="px-20px py-40px">
      <div class="flex-col-center">
        <div>确定要转让管理权限吗？</div>
        <div class="mt-20px text-12px text-coolgray">您将无法管理共养家庭</div>
      </div>
      <div class="mt-20px flex justify-evenly">
        <nut-button plain @click="confirmGrant">确定</nut-button>
        <nut-button @click="closeGrantPopup" color="#f7daa1" class="!text-black">取消</nut-button>
      </div>
    </div>
  </nut-popup>
</basic-layout>
</template>


<style lang="scss"></style>
