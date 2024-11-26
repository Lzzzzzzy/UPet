<script setup lang="ts">
import { onBeforeMount, ref } from "vue";
import { switchTab, redirectTo } from '@tarojs/taro';
import { localStg } from '@/utils';
import { userInfoComplete } from '@/service/api/user';
import { uploadFileToSystem } from "@/service/api";

const baseAvatarUrl = "https://mmbiz.qpic.cn/mmbiz/icTdbqWNOwNRna42FI242Lcia07jQodd2FJGIYQfG0LAJGFxM4FbnQP6yfMxBgJ0F3YRqJCJ1aPAK2dQagdusBZg/0";
const userAvatarUrl = ref();
const onChooseAvatar = (e: any) => {
  const { avatarUrl } = e.detail
  userAvatarUrl.value = avatarUrl;
}
const userNickname = ref("");

const formData = ref();

const formRules = ref({
  avatar: [
    { required: true, message: '请完善头像' }
  ],
  nickname: [
    { required: true, message: '请填写昵称' }
  ],
});

const redirectUrl = ref("");

onBeforeMount(() => {
  const tokenUserInfo = localStg.get("userInfo")
  userAvatarUrl.value = tokenUserInfo?.avatar || baseAvatarUrl;
  userNickname.value = tokenUserInfo?.nickname || "";
  formData.value = tokenUserInfo;
  const instance = getCurrentInstance();
  const params = instance?.router?.params;
  redirectUrl.value = params?.redirectTo;
})

const formRef = ref();

const isLoading = ref(false);

const confirmRegister = async () => {
  formData.value.avatar = userAvatarUrl.value === baseAvatarUrl ? "" : userAvatarUrl.value;
  formData.value.nickname = userNickname.value;

  const { valid } = await formRef.value?.validate();
  if (valid) {
    isLoading.value = true;
    const avatarUrl = await uploadFileToSystem(userAvatarUrl.value);
    formData.value.avatar = avatarUrl;
    await userInfoComplete(formData.value);
    localStg.set("userInfo", formData.value);
    isLoading.value = false;
    if (redirectUrl.value) {
      redirectTo({
        url: redirectUrl.value
      });
    } else {
      switchTab({
        url: '/pages/index/index'
      })
    }
  }
};
</script>

<template>
  <basic-layout>
    <custom-navbar title="用户注册" />
    <nut-form :rules="formRules" :model-value="formData" ref="formRef">
      <nut-form-item label="" prop="avatar">
        <div class="flex-center">
          <button open-type="chooseAvatar" @chooseavatar="onChooseAvatar">
            <image class="w-60px h-60px" :src="userAvatarUrl"></image>
          </button>
        </div>
      </nut-form-item>
      <nut-form-item label="昵称" prop="nickname">
        <input type="nickname" class="weui-input" placeholder="请输入昵称" v-model="userNickname" />
      </nut-form-item>
      <nut-space class="m-10px flex justify-center w-full">
        <nut-button color="#f7daa1" @click="confirmRegister" class="!text-#000000" :loading="isLoading">确定</nut-button>
      </nut-space>
    </nut-form>
  </basic-layout>
</template>


<style lang="scss"></style>
