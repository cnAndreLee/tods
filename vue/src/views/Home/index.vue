<template>
<el-row>
  <el-col :span="24"><Header></Header></el-col>
</el-row>
<el-row>
  <el-col :span="4" >
    <el-tree 
      :data="treedata" 
      :props="defaultProps" 
      @node-click="handleNodeClick" 
      highlight-current 
      icon="null" 
      class="aside-tree" 
      />
  </el-col>
  
  <el-col :span="16"><Player /></el-col>
  <el-col :span="4"><FileList /></el-col>
</el-row>

<!-- <div>
  <el-container>
    <el-header><Header></Header></el-header>
    <el-container>
      <el-aside width="200px">
        <el-tree 
          :data="treedata" 
          :props="defaultProps" 
          @node-click="handleNodeClick" 
          :highlight-current="true" 
          icon="null" 
        />
      </el-aside>
      <el-container>
        <el-aside><Player /></el-aside>
        <el-main><FileList /></el-main>
      </el-container>
    </el-container>
  </el-container>
</div> -->


</template>

<script lang="ts" setup>
import {useStore} from 'vuex'
import Header from "../../components/Header/index.vue"
import Player from "../../components/Player/index.vue"
import FileList from "../../components/FileList/index.vue"
import { ref,computed,onMounted,onBeforeMount } from 'vue'
import fileService from '../../service/fileService';
import { dateEquals, ElTree } from 'element-plus'
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from '@element-plus/icons-vue'

const store = useStore()

interface Tree {
  label: string
  id:string
  level:Number
  children?: Tree[]
}
const treedata: Tree[] = ref([])


const DataInit = async () => {
  let res1 = await fileService.getCategory()
  let res2 = await fileService.getSecondaryCategory()
  store.dispatch('fileModule/getFiles')

  let res1map = new Map()
  let res2map = new Map()

  // 第一个循环，生产一级目录map，k是treedata索引值
  for (let k in res1.data.data.Categories) {
    treedata.value.push({
      label:res1.data.data.Categories[k].name,
      id:res1.data.data.Categories[k].id,
      level:1,
      children:[],
    })

    res1map.set(res1.data.data.Categories[k].id,k)

  }

  // 遍历二级目录
  for (let k2 in res2.data.data.SecondaryCategories) {
    treedata.value[res1map.get(res2.data.data.SecondaryCategories[k2].fatherid)].children.push({
      label: res2.data.data.SecondaryCategories[k2].name,
      id:res2.data.data.SecondaryCategories[k2].id,
      level:2,
    })
    // res2map.set(res2.data.data.SecondaryCategories[k2].id,k2)
  }

};
DataInit()

const handleNodeClick = (data: Tree) => {

  if (data.level != 2) {
    return
  }

  store.commit('fileModule/SET_selectdSecondaryCategory', {id: data.id})

  // for ( let file of res3.data.data.files) {
  //   if ( file.filebelong == data.id) {
  //     console.log(file)
  //   }
  // }
}



const defaultProps = {
  children: 'children',
  label: 'label',
}

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
  border-top: 1px solid #ebebeb;
}
 

.el-tree--highlight-current .el-tree-node.is-current>.el-tree-node__content {
  background-color: #b3cae4;
  color: #fff;
}

</style>