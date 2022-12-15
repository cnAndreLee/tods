<template>
<div class="container">

  <iframe class="iframe" :src="fileurl" scrolling="auto"></iframe>
  <el-button @click="openNew">new</el-button>

</div>
</template>

<script setup>
import config from "../../config.js"
import { ref ,reactive ,computed } from 'vue';
import { useStore } from 'vuex'
const store = useStore()

const selectedFileUrl = computed(() => {
  return store.state.fileModule.selectedFile.url
})

const fileurl = computed(() => {
    let url = config.kkfile +btoa(selectedFileUrl.value)  
    return url
})

const openNew = () => {
  console.log("fileurl.value is : "+fileurl.value)
  window.open(fileurl.value)
}

</script>

<style scoped>
.container {
  width: 100%;
  height: 650px;
}

.iframe {
  width:100%;
  height: 100%;
}
</style>