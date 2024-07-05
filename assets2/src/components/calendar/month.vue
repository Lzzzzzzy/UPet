<template>
  <div class="month">
    <h3>{{ curMonth.text}}</h3>
    <ul>
          <li
            v-for="(elem, i) in weekName"
            :key="i"
            class="cell date"
          >
            <p>{{ elem }}</p>
          </li>
        </ul>
    <swiper
      class="swiper"
      :current="current"
      :circular="true"
      @change="handleSlide"
    >
      <swiper-item v-for="(item, index) in monthsT" :key="index">
        <ul>
        <template v-if="item.start!==0">
          <li
            v-for="(i) in item.start-1"
            :key="i"
          >
          </li>
        </template>
          
          <li
            v-for="(elem, i) in item.days"
            :key="i"
          >
            <p>{{ i+1}}</p>
          </li>
        </ul>
      </swiper-item>
    </swiper>
  </div>
</template>
<script setup>
import { ref, computed } from "vue";
import dayjs from 'dayjs';
import { getMonth} from '@/utils/common/datetime';
import { weekName } from '@/utils/common/const';
// 周list
const date = dayjs().format('YYYY-MM-DD');
const selectedDate = ref(date);
const slideIndex = ref(1); // 当前slide索引
const current = ref(1)
const monthIndexs = ref([-1, 0, 1]);
const curMonth = computed(() => {
    return monthsT.value[slideIndex.value];
});
const months = computed(() => {
    return monthIndexs.value.map((item) => getMonth(item, selectedDate.value));
});

const monthsT = computed(() => {
    return months.value
});

const handleSlide = ({ detail: { current } }) => {
    const curVal = monthIndexs.value[current];
    slideIndex.value = current;
    if (current === 0) {
        monthIndexs.value = [curVal, curVal + 1, curVal - 1];
    } else if (current === 1) {
        monthIndexs.value = [curVal - 1, curVal, curVal + 1];
    } else {
        monthIndexs.value = [curVal + 1, curVal - 1, curVal];
    }
};
</script>
<style lang="scss">
.month {
  margin: 10px 10px 15px;
  &>h3{
    padding: 10px;
    width: max-content;
    font-size: 24px;
    font-weight: 400;
    line-height: 20px;
  }
  ul {
    display: flex;
    margin-bottom: 10px;
    li {
      padding: 10px 0;
      display: flex;
      justify-content: center;
      align-items: center;
      flex-direction: column;
      width: calc((100vw - 20px) / 7);
      padding: 0 4px;
    }
  }
  .swiper {
    width: calc(100vw - 20px);
    height: 400px;
    background-color: rgba(255, 255, 255, 1);
    ul{
      flex-wrap: wrap;
    }
    li {
      position: relative;
      padding-top: 7px;
      height: 55px;
      justify-content: center;
      border-radius: 5px;
      &.active {
        background: #fde98d;
      }
      p {
        margin-bottom: 1.5px;
        height: 14px;
        font-size: 10px;
        font-weight: 400;
        line-height: 14px;
        color: rgba(85, 85, 85, 1);
        text-overflow: ellipsis;
        white-space: nowrap;
        overflow: hidden;
        text-align: center;
        width: 100%;
        &:last-child {
          height: 16.5px;
          font-size: 12px;
          line-height: 16.5px;
          color: rgba(85, 85, 85, 1);
        }
      }
      span {
        position: absolute;
        &.brand {
          top: 0;
          right: 0;
          padding: 0 2px;
          background: #a4c3ff;
          text-align: center;
          line-height: 15px;
          min-width: 15px;
          font-size: 7px;
          font-weight: 400;
          color: rgba(255, 255, 255, 1);
          border-radius: 0 2px 0 6px;
        }
      }
    }
  }
}
</style>