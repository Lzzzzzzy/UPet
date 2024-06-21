<template>
  <view>
    <view class="self-title">
      <nut-navbar title="希望关注哪些消息？" desc="完成" @click-right="save"></nut-navbar>
    </view>
    <view>
      <nut-grid :gutter="10" :column-num="4">
        <nut-grid-item :text="item.name" v-for="(item, index) in selectionList" :key="index" @click="addToRecommand(item)">
          <nut-badge :key="index" :hidden="!item.selected">
            <template #icon>
              <Check size="8px"></Check>
            </template>
            <IconFont
              size="50"
              name="https://img11.360buyimg.com/imagetools/jfs/t1/137646/13/7132/1648/5f4c748bE43da8ddd/a3f06d51dcae7b60.png"
            ></IconFont>
          </nut-badge>
        </nut-grid-item>
      </nut-grid>
    </view>
    <nut-popup v-model:visible="showPopup" round>
      <view class="popup-content">
        <span class="popup-sub-title">您还可以随时进入</span>
        <span class="popup-main-title">我的 - 修改提醒事项</span>
        <span class="popup-sub-title">进行私人定制哦</span>
      </view>
    </nut-popup>
  </view>
</template>

<script setup>
import { ref, onBeforeMount } from 'vue';
import { IconFont } from '@nutui/icons-vue-taro'
import { Check } from '@nutui/icons-vue-taro'

const selectionList = ref([{
  name: 'text',
  icon: 'https://img11.360buyimg.com/imagetools/jfs/t1/137646/13/7132/1648/5f4c748bE43da8ddd/a3f06d51dcae7b60.png',
  selected: true},{
  name: 'text',
  icon: 'https://img11.360buyimg.com/imagetools/jfs/t1/137646/13/7132/1648/5f4c748bE43da8ddd/a3f06d51dcae7b60.png',selected: true},{
  name: 'text',
  icon: 'https://img11.360buyimg.com/imagetools/jfs/t1/137646/13/7132/1648/5f4c748bE43da8ddd/a3f06d51dcae7b60.png'},{
  name: 'text',
  icon: 'https://img11.360buyimg.com/imagetools/jfs/t1/137646/13/7132/1648/5f4c748bE43da8ddd/a3f06d51dcae7b60.png'},{
  name: 'text',
  icon: 'https://img11.360buyimg.com/imagetools/jfs/t1/137646/13/7132/1648/5f4c748bE43da8ddd/a3f06d51dcae7b60.png'},{
  name: 'text',
  icon: 'https://img11.360buyimg.com/imagetools/jfs/t1/137646/13/7132/1648/5f4c748bE43da8ddd/a3f06d51dcae7b60.png'},{
  name: 'text',
  icon: 'https://img11.360buyimg.com/imagetools/jfs/t1/137646/13/7132/1648/5f4c748bE43da8ddd/a3f06d51dcae7b60.png'}])


const showPopup = ref(false);


const addToRecommand = (item) => {
  console.log(item);
  if (!item.selected) {
    item.selected = false;
  }
  item.selected = !item.selected;
  console.log("after:", item);
}

const save = () => {
  console.log('save');
  showPopup.value = true;
  console.log(showPopup.value);
}

</script>

<style lang="scss">
.self-title{
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100rpx;
  font-size: 36rpx;
  font-weight: bold;
  color: #333333;
}

.nutui-iconfont.nut-icon.nut-icon-Check {
  color: #ffffff;
}

.popup-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #333333;
  padding: 30rpx;
}

.popup-main-title {
  color: var(--main-color);
}

.popup-sub-title {
  font-size: small;
}
.nut-popup--center.round {
  width: 80%;
}
</style>
