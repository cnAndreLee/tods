<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span>课程列表</span>
      </div>
    </template>
    <!-- <p v-if="files.length == 0">无</p> -->
    <div 
      v-for="file in files" 
      style="width: 100%;"
      @click="select(file)" 
      class="item"
      >
        <el-icon v-if="file.suffix == 'mp4'"><VideoCamera /></el-icon>
        <el-icon v-else><Document /></el-icon>
        {{file.title}}
  </div>
  </el-card>
</template>

<script setup>
import { useStore } from 'vuex'
import { ref ,reactive ,computed } from 'vue';
import { Document,VideoCamera } from '@element-plus/icons-vue'



const store = useStore()


const files = computed(()=>{
    // let localFiles = [];
    // for ( let v of store.state.fileModule.files ) {
    //     if ( v.filebelong == store.state.fileModule.selectdSecondaryCategory) {
    //         localFiles.push(v)
    //     }
    // }
    // return localFiles
    return store.state.fileModule.files
})

const select = (file) => {
  store.commit('fileModule/SET_SelectedFile',file)
}


</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.item {
  margin-bottom: 18px;
  
}

.item:hover {
  background-color: aqua;
}

.box-card {
  width: 100%;
  min-height: 100%;
}
</style>
