<template>
    <div class="border-rd-50% b-1px b-solid b-color-#222222" :style="styleObj">
        <div class="i-local-checked" v-show="checked" :style="{'font-size': size}"></div>
    </div>
</template>

<script setup lang="ts">
import {computed} from "vue";

const props = defineProps({
  size: {
    type: String,
    default: "20px",
  },
  checked: {
    type: Boolean,
    default: false,
  },
  bgColor: {
    type: String,
    default: "#ffffff"
  },
  checkedBgColor: {
    type: String,
    default: "#ffffff"
  },
});
const styleObj = computed(() => {
    return {
        height: props.size, 
        width: props.size, 
        background: props.checked ? props.checkedBgColor : props.bgColor,
    }
})
</script>