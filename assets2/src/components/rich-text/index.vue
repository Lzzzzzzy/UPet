<script setup lang="ts">
import { ref } from 'vue';
import { createSelectorQuery } from '@tarojs/taro';

const props = defineProps({
  data: {
    type: String,
    required: true,
  },
  placeholder: {
    type: String,
  }
});

const editorCtx = ref();
const onEditorReady = () => {
  createSelectorQuery().select('#editor').context((res) => {
    console.log(res.context);
    editorCtx.value = res.context;
  }).exec()
  editorCtx.value.setContents({
    html: props.data,
  });
};

interface editorChangeDetail {
  html: string;
  delta: any;
  text: string;
}
const emit = defineEmits(["updateData"]);
const onEditorInput = (detail: editorChangeDetail) => {
  emit("updateData", detail.html);
};
</script>

<template>
  <div>
    <editor id="editor" class="editor" :placeholder="placeholder" @ready="onEditorReady" @input="onEditorInput"></editor>
    <div class="toolbar">
        <div class="i-local-image"></div>
    </div>
  </div>
</template>

<style scoped lang="less">
.rich-text {
  font-size: 14px;
  line-height: 1.5;
  color: #333;
  word-wrap: break-word;
  word-break: break-all;
  white-space: pre-wrap;
}
</style>
