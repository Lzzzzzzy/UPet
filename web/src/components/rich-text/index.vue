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
  },
  showUploader: {
    type: Boolean,
    default: false,
  },
  readOnly: {
    type: Boolean,
    default: false,
  }
});


const editorCtx = ref();
const onEditorReady = () => {
  createSelectorQuery().select('#editor').context((res) => {
    console.log(res.context);
    editorCtx.value = res.context;
    setContent();
  }).exec()
};

const setContent = () => {
  editorCtx.value.setContents({
    html: props.data,
  });
}

interface editorChangeDetail {
  html: string;
  delta: any;
  text: string;
}
const emit = defineEmits(["updateData"]);
const onEditorInput = ({ detail: { html, delta, text } }: {detail: editorChangeDetail}) => {
  emit("updateData", html);
};

const uploadUrl = ref("");
</script>

<template>
  <div>
    <editor id="editor" class="editor break-words h-50px min-h-0" :placeholder="placeholder" @ready="onEditorReady" @input="onEditorInput" :readOnly="readOnly"></editor>
    <div class="flex items-center justify-end mt-5px" v-if="showUploader">
      <nut-uploader :url="uploadUrl">
        <div class="i-local-image text-25px text-#333333"></div>
      </nut-uploader>
    </div>
  </div>
</template>

<style lang="scss">

</style>
