<template>
<VideoPlayer v-if="selectedFileSuffix === 'mp4'"></VideoPlayer>
<!-- <DocumentReader v-else-if="selectedFileSuffix === 'pptx'"></DocumentReader> -->
<onlyoffice v-else-if="selectedFileSuffix === 'pptx'"></onlyoffice>
<defaulfSHOW v-else></defaulfSHOW>
</template>

<script setup>
import { ref ,reactive ,computed } from 'vue';
import { useStore } from 'vuex';
import defaulfSHOW from "../defaulfSHOW.vue"
import VideoPlayer from "../VideoPlayer/index.vue"
import DocumentReader from "../DocumentReader/index.vue"
import onlyoffice from "../DocumentReader/onlyoffice.vue"

const store = useStore()

const selectedFileSuffix = computed(() => {

    return store.state.fileModule.selectedFile.suffix
})

</script>