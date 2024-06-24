<template>
  <nut-tabbar v-model="activeName" bottom safe-area-inset-bottom placeholder @tab-switch="tabSwitch">
    <nut-tabbar-item v-for="(item, index) in List" :key="index" :tab-title="item.title" :icon="item.icon" :name="item.name">
    </nut-tabbar-item>
  </nut-tabbar>
</template>
<script setup>
import { h, ref } from 'vue'
import { Date, Category, My, Home } from '@nutui/icons-vue-taro'

const props = defineProps({
  active: {
    type: String,
    required: true,
  }
});

const List = [
  {
    title: '首页',
    icon: h(Home),
    name: "home"
  },
  {
    title: '日历',
    icon: h(Date),
    name: "calendar"
  },
  {
    title: '自定义',
    icon: h(Category),
    name: "self"
  },
  {
    title: '我的',
    icon: h(My),
    name: "my"
  }
]
console.log("props.active:", props.active);
const activeName = ref(props.active);

const emit = defineEmits(["changeTab"]);
const tabSwitch = (item, index)=>{
    emit("changeTab", index);
}

</script>

<style lang="scss">
.nut-tabbar {
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  background: rgba(240, 237, 228, 1);
  box-shadow: 0 0 10rpx rgba(102, 93, 33, 1);
}
</style>
