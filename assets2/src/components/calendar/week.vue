<template>
  <div class="week">
    <h3>{{ curMonth.$y }}年{{ curMonth.$M + 1 }}月</h3>
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
      <swiper-item v-for="(item, index) in weeksT" :key="index">
        <ul>
          <li
            v-for="(elem, i) in item"
            :key="i"
            class="cell date"
            :class="{
              active: elem.isActive,
            }"
          >
            <p>{{ elem.dateT }}</p>
          </li>
        </ul>
      </swiper-item>
    </swiper>
  </div>
</template>
<script setup>
import { ref, computed } from "vue";
import dayjs from 'dayjs';
import { getWeekDays, getWeek} from '@/utils/common/datetime';
import { weekName } from '@/utils/common/const';
// 周list
const date = dayjs().format('YYYY-MM-DD');
const selectedDate = ref(date);
const slideIndex = ref(1); // 当前slide索引
const current = ref(1)
const weekIndexs = ref([-1, 0, 1]);
const curMonth = computed(() => {
    return getWeek(weekIndexs.value[slideIndex.value], selectedDate.value).start;
});
const weeks = computed(() => {
    return weekIndexs.value.map((item) => getWeekDays(item, selectedDate.value));
});
let curDay = dayjs().format('d')
const weeksT = computed(() => {
    return weeks.value.map((week) => {
        return week.map((item, index) => {
            return {
                ...item,
                isActive: (curDay - index) === 1,
                dateT: item.d,
                // showMonth: item.d === 1,
            };
        });
    });
});

const handleSlide = ({ detail: { current } }) => {
    const curVal = weekIndexs.value[current];
    slideIndex.value = current;
    if (current === 0) {
        weekIndexs.value = [curVal, curVal + 1, curVal - 1];
    } else if (current === 1) {
        weekIndexs.value = [curVal - 1, curVal, curVal + 1];
    } else {
        weekIndexs.value = [curVal + 1, curVal - 1, curVal];
    }
};

</script>
<style lang="scss">
.week {
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
    li {
      padding: 10px 0;
      display: flex;
      justify-content: center;
      align-items: center;
      flex-direction: column;
      width: calc((100vw - 20px) / 7);
      margin: 0 4px;
    }
  }
  .swiper {
    width: calc(100vw - 20px);
    height: 55px;
    li {
      position: relative;
      padding-top: 7px;
      height: 55px;
      justify-content: center;
      background-color: rgba(255, 255, 255, 1);
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