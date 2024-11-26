<template>
    <basic-layout>
        <custom-navbar title="加入家庭" left-show />
        <div>
            确认加入{{ userInfo.nickname }}的家庭吗？
        </div>
        <div>
            <nut-button type="primary" @click="confirmJoinFamily">是</nut-button>
            <nut-button type="default" @click="gotoIndex">否</nut-button>
        </div>
    </basic-layout>
</template>

<script lang="ts" setup>
import { onBeforeMount, ref } from "vue";
import { getCurrentInstance, switchTab } from "@tarojs/taro";
import { userLogin, getUserInfo } from '@/service/api';

const userInfo = ref<any>();
onBeforeMount(async () => {
    const instance = getCurrentInstance();
    const params = instance?.router?.params;
    const inviteUserId = params?.inviteUserId;
    const url = `/package/package-pet-ownership/confirmJoinFamily?inviteUserId=${inviteUserId}`;
    userLogin(url);
    userInfo.value = await getUserInfo(inviteUserId);
})

const confirmJoinFamily = () => {

}

const gotoIndex = () => {
    switchTab({
        url: '/pages/index/index'
    })
}
</script>
