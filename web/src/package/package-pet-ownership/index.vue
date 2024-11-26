<script setup lang="ts">
import { onBeforeMount, ref, computed } from 'vue';
import type { UserInfo } from "@/typings/family";
import { getFamilyMemberList } from '@/service/api';
import { localStg } from '@/utils';

const getAllFamilyMember = async () => {
  const member = await getFamilyMemberList()
  return member
}

const familyMembers = ref<Array<UserInfo>>();

const mainSpanWidth = ref(12);

onBeforeMount(() => {
  familyMembers.value = getAllFamilyMember();
  if (!me.isAdmin) {
    mainSpanWidth.value = 20;
  }
})

const me: UserInfo = localStg.get("userInfo");

const showDeleteConfirmPopup = ref(false);
const needDeleteMember = ref<UserInfo>();
const deleteConfirmText = computed(() => {
  if (needDeleteMember.value.id === me.id) {
    return "确定要退出家庭吗？"
  }
  return `确定要移除家庭成员 ${needDeleteMember.value.name} 吗？`
});

const deleteMember = (member: UserInfo) => {
  showDeleteConfirmPopup.value = true;
  needDeleteMember.value = member;
}
const closeMemberPopup = () => {
  showDeleteConfirmPopup.value = false;
  needDeleteMember.value = null;
}
const confirmDelete = () => {
  closeMemberPopup();
}

const addMember = () => { };


const showGrantConfirmPopup = ref(false);
const needGrantMember = ref<UserInfo>();
const grantConfirmText = computed(() => {
  return `确定要转让管理权限给 ${needGrantMember.value.name} 吗？`
});

const grantMember = (member: UserInfo) => {
  showGrantConfirmPopup.value = true;
  needGrantMember.value = member;
}
const closeGrantPopup = () => {
  showGrantConfirmPopup.value = false;
  needGrantMember.value = null;
}
const confirmGrant = () => {
  closeGrantPopup();
}

onShareAppMessage(res) {
  if (res.from === 'button') {
    // 来自页面内转发按钮
    console.log(res.target)
  }
  return {
    title: '我们一起照顾毛孩子吧！',
    path: `/pages/index/index?inviteUserId=${me.id}`,
    imageUrl: '',
  }
}

</script>
<template>
  <basic-layout>
    <custom-navbar title="爱宠共养" left-show />
    <div class="bg-white py-30px svg-bg pr-10px">
      <nut-row :gutter="10" type="flex" justify="space-around" align="center">
        <nut-col :span="4">
          <div class="flex items-center justify-end">
            <pet-avatar :avatar-img-url="me.avatar" />
          </div>
        </nut-col>
        <nut-col :span="20">
          <div class="flex flex-col items-start justify-center">
            <div class="text-18px">
              {{ me.name }}
            </div>
            <div class="mt-10px">
              <nut-avatar size="small" class="!flex-center !bg-#98c7ce" @click="deleteMember(me.id)" v-if="!me.isAdmin">
                <div class="text-20px i-local-exit"></div>
              </nut-avatar>
              <div class="text-12px text-coolgray" v-if="me.isAdmin">管理员</div>
            </div>
          </div>
        </nut-col>
      </nut-row>
    </div>
    <div class="m-20px b-rd-16px py-30px bg-white" v-for="member in familyMembers" :key="member.id">
      <nut-row :gutter="10" type="flex" justify="space-around" align="center">
        <nut-col :span="4">
          <div class="flex items-center justify-end">
            <pet-avatar :avatar-img-url="member.avatar" />
          </div>
        </nut-col>
        <nut-col :span="mainSpanWidth">
          <div class="flex flex-col items-start justify-center">
            <div class="text-18px">
              {{ member.name }}
            </div>
            <div class="flex-center mt-10px" v-if="member.isAdmin">
              <div class="text-12px text-coolgray">管理员</div>
            </div>
          </div>
        </nut-col>
        <nut-col :span="4" v-if="me.isAdmin">
          <nut-avatar size="small" class="!flex-center !bg-#98c7ce" @click="grantMember(member)">
            <div class="text-20px i-local-grant"></div>
          </nut-avatar>
        </nut-col>
        <nut-col :span="4" v-if="me.isAdmin">
          <nut-avatar size="small" class="!flex-center !bg-#98c7ce" @click="deleteMember(member)">
            <div class="text-20px i-local-delete"></div>
          </nut-avatar>
        </nut-col>
      </nut-row>
    </div>

    <div class="flex-center mt-10px" v-if="me.isAdmin">
      <nut-button @click="addMember" color="#f7daa1" class="!text-black" openType="share">
        <div class="flex-center">
          <div class="text-15px i-local-add pr-5px"></div>
          <div>邀请家庭成员</div>
        </div>
      </nut-button>
    </div>

    <nut-popup v-model:visible="showDeleteConfirmPopup" round :style="{ 'width': '70%' }">
      <div class="px-20px py-40px">
        <div class="flex-col-center">
          <div>{{ deleteConfirmText }}</div>
        </div>
        <div class="mt-20px flex justify-evenly">
          <nut-button plain @click="confirmDelete">确定</nut-button>
          <nut-button @click="closeMemberPopup" color="#f7daa1" class="!text-black">取消</nut-button>
        </div>
      </div>
    </nut-popup>

    <nut-popup v-model:visible="showGrantConfirmPopup" round :style="{ 'width': '70%' }">
      <div class="px-20px py-40px">
        <div class="flex-col-center">
          <div>{{ grantConfirmText }}</div>
          <div class="mt-20px text-12px text-coolgray">转让后您将无法管理家庭和宠物档案</div>
        </div>
        <div class="mt-20px flex justify-evenly">
          <nut-button plain @click="confirmGrant">确定</nut-button>
          <nut-button @click="closeGrantPopup" color="#f7daa1" class="!text-black">取消</nut-button>
        </div>
      </div>
    </nut-popup>
  </basic-layout>
</template>


<style lang="scss">
.svg-bg {
  background-image: url('../../assets/svg/house.svg');
  background-size: contain;
  background-repeat: no-repeat;
  background-position: right;
}
</style>
