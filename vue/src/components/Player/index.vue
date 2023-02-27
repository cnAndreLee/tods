<template>
<VideoPlayer v-if="selectedFileSuffix === 'mp4'"></VideoPlayer>
<kkfileReader v-else-if="officeFileType.indexOf(selectedFileSuffix) > -1 "></kkfileReader>
<!-- <onlyoffice v-else-if="false"></onlyoffice> -->
<defaulfSHOW v-else></defaulfSHOW>
</template>

<script setup>
import { ref ,reactive ,computed } from 'vue';
import { useStore } from 'vuex';
import config from '../../config.js'
import defaulfSHOW from "./defaulfSHOW.vue"
import VideoPlayer from "./videoPlayer.vue"
import kkfileReader from "./kkfile.vue"
import onlyoffice from "./onlyoffice.vue"

const store = useStore()

// office的后缀名列表
const officeFileType = config.officeFileType

const selectedFileSuffix = computed(() => {
    return store.state.fileModule.selectedFile.suffix
})

</script>