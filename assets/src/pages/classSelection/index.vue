<template>
  <view>
    <view class="self-title">
      <nut-navbar title="希望提醒哪些消息？" desc="完成" @click-right="onClick"></nut-navbar>
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
  </view>
</template>

<script setup>
import { ref, onBeforeMount } from 'vue';
import { IconFont } from '@nutui/icons-vue-taro'
import { Check } from '@nutui/icons-vue-taro'

const selectionList = [{
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
  icon: 'https://img11.360buyimg.com/imagetools/jfs/t1/137646/13/7132/1648/5f4c748bE43da8ddd/a3f06d51dcae7b60.png'}]


const addToRecommand = (item) => {
  console.log(item)
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
</style>
