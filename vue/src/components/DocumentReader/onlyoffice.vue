<template>
    <DocumentEditor 
      id="docEditor" 
      documentServerUrl="http://10.0.0.10:8081/"
      :config="config"
      :events_onDocumentReady="onDocumentReady"
      /> 
</template>
  
<script lang="ts" setup>
import { ref,reactive,computed } from 'vue';
import { DocumentEditor } from "@onlyoffice/document-editor-vue";
import { useStore } from 'vuex'
const store = useStore()
const selectedFileUrl = computed(() => {
    return store.state.fileModule.selectedFile.url
})
const selectedFileSuffix = computed(() => {
    return store.state.fileModule.selectedFile.suffix
})
const selectedFileId = computed(() => {
    return store.state.fileModule.selectedFile.id
})
const selectedFileTitle = computed(() => {
    return store.state.fileModule.selectedFile.title
})

var config= reactive({
    document: {
        fileType: selectedFileSuffix.value,
        // key: selectedFileId.value,
        title: selectedFileTitle.value,
        url: selectedFileUrl.value
    },
    documentType: "slide",
    type:"desktop",
    token:"Bearer " + store.state.userModule.token,
    editorConfig: {
        mode:"view",
        lang:"zh-CN"
    }
})

const onDocumentReady = () => {
    console.log("Document is loaded");
}

console.log("2222222"+selectedFileUrl.value)
</script>