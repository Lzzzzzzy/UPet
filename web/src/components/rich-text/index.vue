<script setup lang="ts">
import { ref, watch } from 'vue';
import { createSelectorQuery, eventCenter, chooseImage } from '@tarojs/taro';
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
});


const editorCtx = ref();
const onEditorReady = () => {
  createSelectorQuery().select('#editor').context((res) => {
    editorCtx.value = res.context;
  }).exec()
};

watch(() => props.data, () => {
  setContent();
})

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

const htmlText = ref("");

const updateText = () => {
  emit("updateData", htmlText.value);
}

const onEditorInput = ({ detail: { html, delta, text } }: { detail: editorChangeDetail }) => {
  htmlText.value = html;
  updateText();
};

const imgsList = ref([]);
const imgMap = ref({});

const uploadImgsAndFormatHtmlText = async () => {
  const promises = imgsList.value.map(async (fileUrl: string) => {
    const pictureUrl = await uploadFileToSystem(fileUrl);
    imgMap.value[fileUrl] = pictureUrl;
  })
  await Promise.all(promises);
  imgsList.value.forEach((localUrl: string) => {
    const remoteUrl = imgMap.value[localUrl];
    htmlText.value = htmlText.value.replace(localUrl, remoteUrl);
  })
  updateText();
  eventCenter.trigger("imgsUploaded");
}

eventCenter.on("uploadPicture", async (res: any) => {
  await uploadImgsAndFormatHtmlText();
})

const handleUploadImage = async () => {
  try {
    const res = await chooseImage({
      count: 1,
      sizeType: ['original', 'compressed'],
      sourceType: ['album', 'camera']
    });

    const tempFilePath = res.tempFilePaths[0];
    imgsList.value.push(tempFilePath);

    // // 插入图片到编辑器
    editorCtx.value.insertImage({
      src: tempFilePath,
      alt: 'image'
    });

  } catch (error) {
    console.error('上传失败', error);
  }
};
</script>

<template>
  <div>
    <div v-if="imgsList" class="flex">
      <div v-for="item in imgsList" :key="item.uid" class="w-fit">
        <img :src="item.url" class="w-50px !h-50px" />
      </div>
    </div>
    <editor id="editor" class="editor break-words max-h-150px min-h-0" :placeholder="placeholder" :showImgToolbar="true"
      @ready="onEditorReady" @input="onEditorInput" />
    <div class="i-local-image text-25px text-#333333" @click="handleUploadImage" v-if="showUploader"></div>
  </div>
</template>

<style lang="scss">
.nut-uploader {
  .nut-uploader__preview {
    display: none;
  }
}
</style>
