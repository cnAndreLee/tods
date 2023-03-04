<template>
<div class="container">

  <iframe class="iframe" :src="filkkeurl" scrolling="auto" allowfullscreen="true" ></iframe>

</div>
</template>

<script setup>
import config from "../../config.js"
import { computed } from 'vue';
import { useStore } from 'vuex'
const store = useStore()


const filkkeurl = computed(() => {
  const kkurl = config.kkfile + "/onlinePreview?url="
  const backend = config.backend
  const suffix = store.state.fileModule.selectedFile.suffix
  const id = store.state.fileModule.selectedFile.id
  const fileSource = backend + config.filespath + id + '.' + suffix
  return kkurl + encodeURIComponent(window.btoa(fileSource))
})

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