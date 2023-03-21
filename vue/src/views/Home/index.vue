<template>
  <el-row>
    <el-col :span="24"><Header></Header></el-col>
  </el-row>
  <el-row>
    <el-col style="height:30px"></el-col>
  </el-row>
  <el-row justify="center">
    <el-col :span="4" >
      <el-card style="min-height: 600px">
        <template #header>
          <span>课程分类</span>
        </template>
        <el-tree 
          :data="TreeData" 
          :props="defaultProps" 
          @node-click="handleNodeClick" 
          highlight-current 
          icon="null" 
          class="aside-tree" 
          />
      </el-card>
    </el-col>
    <el-col :span="1"></el-col>
    <el-col :span="14"><Player /></el-col>
    <el-col :span="1"></el-col>
    <el-col :span="4"><FileList /></el-col>
  </el-row>
</template>

<script lang="ts" setup>
  import {useStore} from 'vuex'
  import Header from "../../components/Header/index.vue"
  import Player from "../../components/Player/index.vue"
  import FileList from "../../components/FileList/index.vue"
  import { ref,computed,onMounted,onBeforeMount } from 'vue'
  import storageService from '../../service/storageService';
  import { useRouter } from 'vue-router'
  import { ElMessage } from 'element-plus'

  const store = useStore()
  const router = useRouter()

  interface Tree {
    title: string
    id:string
    level:Number
    children?: Tree[]
  }

  const defaultProps = {
    children: 'children',
    label: 'title',
  }
  const TreeData: Tree[] = ref([])


  const handleNodeClick = (data: Tree) => {

    let userInfo = store.state.userModule.userInfo
    if ( data.level == 1 ) {
      
      if (userInfo.is_admin == false) {
        if (data.permissions.indexOf(userInfo.school) == -1) {
          ElMessage.error("无权限")
          router.go(0)
        }
      }

      store.commit('fileModule/SET_FILES', [])
    }
    
    if (data.level == 2) {
      store.dispatch('fileModule/getFiles',{"id":data.id}).then((res)=>{
      }).catch((res)=>{
        console.log(res.data.msg)
      })
    }
  }

  const GetCategories = () => {
    store.dispatch('fileModule/getCategories').then((res) => {
      const categories = res.data.data.categories
      TreeData.value = categories
    }).catch((err)=>{
      console.log("--"+err)
      if (err.response.data.status >= 4000) {
        ElMessage.error("错误："+err.data.msg)
      }
    })
  }

  const DataInit = () => {

    GetCategories()
    
  };
  DataInit()
</script>

<style>

  .aside-tree {
    margin-right: 6px;
  }


  .el-tree {
    background-color: #fff;
  }

  .el-tree-node__content {
    height: 40px;
    border-bottom: 1px solid #ebebeb;
  }
  

  .el-tree--highlight-current .el-tree-node.is-current>.el-tree-node__content {
    background-color: #b3cae4;
    color: #fff;
  }

</style>