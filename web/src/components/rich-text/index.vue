<script setup lang="ts">
import { ref } from 'vue';
import { createSelectorQuery } from '@tarojs/taro';
import { uploadFileToSystem } from "@/service/api";

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
const uploadRef = ref();
const defaultFileList = ref([]);

const submitUpload = () => {
  console.log(defaultFileList.value);
}
</script>

<template>
  <div>
    <div v-if="defaultFileList" class="flex">
      <div v-for="item in defaultFileList" :key="item.uid" class="w-fit">
        <img :src="item.url" class="w-50px !h-50px" />
      </div>
    </div>
    <editor id="editor" class="editor break-words h-50px min-h-0" :placeholder="placeholder" @ready="onEditorReady" @input="onEditorInput" :readOnly="readOnly"></editor>
    <div class="flex items-center justify-end mt-5px" v-if="showUploader">

      <nut-uploader maximum="5" :auto-upload="false" ref="uploadRef" v-model:file-list="defaultFileList">
        <div class="i-local-image text-25px text-#333333"></div>
      </nut-uploader>
    </div>
  </div>
</template>

<style lang="scss">
.nut-uploader {
  .nut-uploader__preview {
    display: none;
  }
}
</style>
