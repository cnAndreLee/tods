<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span>课程列表</span>
      </div>
    </template>
    <!-- <p v-if="files.length == 0">无</p> -->
    <el-row>
      <el-col :span="20">
        <div 
          v-for="file in files" 
          @click="select(file)" 
          class="item"
          >
            <el-icon v-if="file.suffix == 'mp4'"><VideoCamera /></el-icon>
            <el-icon v-else><Document /></el-icon>
            {{file.title}}
        </div>
      </el-col>
      <el-col :span="4" >
        <!-- <a v-for="file in files" :href="'http://andrelee.f3322.net:28000/api/files/' + file.id + '.' + file.suffix" download><el-icon class="item" size="middle" ><Download /></el-icon></a> -->
        
        <el-icon v-for="file in files" class="item" size="middle" @click="downloadfile(file)"><Download /></el-icon>
      </el-col>
  </el-row>
  </el-card>
</template>

<script setup>
import { useStore } from 'vuex'
import { ref ,reactive ,computed } from 'vue';
import { Document,VideoCamera } from '@element-plus/icons-vue'
import config from '../../config.js'

const store = useStore()

const files = computed(()=>{
    return store.state.fileModule.files
})

const select = (file) => {
  store.commit('fileModule/SET_SelectedFile',file)
}


const downloadfile = (file) => {
  const url = config.backend + "/api/files/" + file.id + '.' + file.suffix
  const filename = file.title

  const x = new XMLHttpRequest()
  x.open('GET', url, true)
  x.responseType = 'blob'
  x.onload = e => {
    // 会创建一个 DOMString，其中包含一个表示参数中给出的对象的URL。这个 URL 的生命周期和创建它的窗口中的 document 绑定。这个新的URL 对象表示指定的 File 对象或 Blob 对象。
    const url = window.URL.createObjectURL(x.response)
    const a = document.createElement('a')
    a.href = url
    a.download = filename
    a.click()
  }
  x.send()
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
  width: 100%;
  height: 25px;
}

.item:hover {
  background-color: aqua;
}

.box-card {
  width: 100%;
  min-height: 100%;
}
</style>
