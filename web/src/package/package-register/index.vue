<script setup lang="ts">
import {onBeforeMount, ref} from "vue";
import { switchTab } from '@tarojs/taro';
import { localStg } from '@/utils';
import { userInfoComplete } from '@/service/api/user';

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

onBeforeMount(()=>{
    const tokenUserInfo = localStg.get("userInfo")
    userAvatarUrl.value = tokenUserInfo?.avatar || baseAvatarUrl;
    userNickname.value = tokenUserInfo?.nickname || "";
    formData.value = tokenUserInfo;
})

const formRef = ref();

const confirmRegister = () => {
    formData.value.avatar = userAvatarUrl.value === baseAvatarUrl ? "" : userAvatarUrl.value;
    formData.value.nickname = userNickname.value;
    console.log("formData:", formData.value);
    formRef.value?.validate().then(({ valid, errors }: { valid: boolean, errors: any}) => {
    if (valid) {
        userInfoComplete(formData.value).then(() => {
            localStg.set("userInfo", formData.value);
            switchTab({
                url: '/pages/index/index'
            })
        })
    } else {
      console.warn('error:', errors)
    }
  })
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
        <nut-button color="#f7daa1" @click="confirmRegister" class="!text-#000000">确定</nut-button>
    </nut-space>
  </nut-form>
</basic-layout>
</template>


<style lang="scss">

</style>
