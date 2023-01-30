<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span>课程列表</span>
      </div>
    </template>
    <p v-for="file in localfiles" @click="select(file)" class="item text">
      <el-icon v-if="file.suffix == 'mp4'"><VideoCamera /></el-icon>
      <el-icon v-else><Document /></el-icon>
      {{file.title}}
    </p>
  </el-card>
</template>

<script setup>
import { useStore } from 'vuex'
import { ref ,reactive ,computed } from 'vue';
import { Document,VideoCamera } from '@element-plus/icons-vue'



const store = useStore()


const localfiles = computed(()=>{
    let localFiles = [];
    for ( let v of store.state.fileModule.files ) {
        if ( v.filebelong == store.state.fileModule.selectdSecondaryCategory) {
            localFiles.push(v)
        }
    }
    return localFiles
})

const select = (file) => {
  store.commit('fileModule/SET_SelectedFile',file)
  console.log(store.state.fileModule.selectedFile)
}


</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
  
}

.item:hover {
  background-color: aqua;
}

.box-card {
  width: 100%;
}
</style>
