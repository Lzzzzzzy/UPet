<script lang="ts" setup>
import { ref } from "vue";

const showAddTodoPopup = ref(false);

</script>
<template>
    <div>
        <nut-avatar class="!flex justify-center items-center !bg-#F7DAA1 text-#665D21 !position-fixed position-bottom-100px position-right-30px" size="large" @click="showAddTodoPopup=true">
            <div class="text-30px i-local-add"></div>
        </nut-avatar>
        <nut-popup v-model:visible="showAddTodoPopup">
            添加
        </nut-popup>
    </div>
</template>