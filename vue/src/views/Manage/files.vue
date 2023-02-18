<template>
    <el-row><p>资源管理</p></el-row>
    <el-row>
        <el-col :span="8"><p>请先选择需要管理的合集</p></el-col>
        <el-col :span="16"><el-cascader
            v-model="value"
            :options="options"
            :props="props"
            @change="handleChange"
            /></el-col>
    </el-row>

    <el-row>
        <el-col :span="24">
            <el-table :data="tableData" style="width: 100%">
                <el-table-column prop="title" label="标题" style="width:50%" />
                <el-table-column prop="suffix" label="格式" style="width:30%" />
                <el-table-column prop="id" fixed="right" label="操作" style="width:20%">
                    <template #default="scope">
                        <el-popconfirm title='确认删除？？？'
                        confirm-button-text='Yes'
                        cancel-button-text='No'
                        icon='el-icon-delete'
                        icon-color='red'
                        @confirm='handelDelete(scope.$index)'
                        >
                            <template #reference><el-button type="danger" plain>删除</el-button></template>
                        </el-popconfirm>
                    </template></el-table-column>
            </el-table>
        </el-col>
    </el-row>
    <el-row style="height:30px"></el-row>
    <el-row v-if="isSelectedCategory">
        <el-col :span="3"></el-col>
        <el-col :span="18">
            <el-upload
                class="upload-demo"
                drag
                :action="fileUploadUrl"
                :before-upload="beforeUpload"
                :on-success="handleSuccess"
                multiple
                >
                <el-icon class="el-icon--upload"><UploadFilled /></el-icon>
                <div class="el-upload__text">
                    拖拽文件或者<em>点击以上传</em>
                </div>
                <template #tip>
                    <div class="el-upload__tip">
                        视频文件请提前压缩
                    </div>
                </template>
            </el-upload>
        </el-col>
        <el-col :span="3"></el-col>
    </el-row>
</template>

<script lang="ts" setup>
import { ref, computed, h } from 'vue'
import {useStore} from 'vuex'
import { useRouter } from 'vue-router'
import { UploadFilled } from '@element-plus/icons-vue'
import config from '../../config.js'
import { ElMessage, ElMessageBox, ElNotification  } from 'element-plus'
import type { UploadProps } from 'element-plus'

const store = useStore()
const router = useRouter()

// 当前合集id
const ca2id = ref("")
const fileUploadUrl = computed(() => {
    return config.backend + "/api/v1/file" + "?filebelong=" + ca2id.value
})


// 上传前 
const supportFileType = ["videmo/mp4", 
    "application/msword", 
    "application/vnd.openxmlformats-officedocument.wordprocessingml.document", 
    "application/vnd.ms-powerpoint",
    "application/vnd.openxmlformats-officedocument.presentationml.presentation", 
    "application/vnd.ms-excel", 
    "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"]
const beforeUpload : UploadProps['beforeUpload'] = (rawFile) => {
  console.log(rawFile.type)
  if (supportFileType.indexOf(rawFile.type) === -1) {
    // ElMessage.error('不支持的文件类型')
    ElMessage('不支持的文件类型')
    return false
  } 
  return true
}


// 上传成功后
const handleSuccess = () => {
  console.log("upload success")
  getFilesList(ca2id.value)
//   router.go(0)
}

// value[1]为二级分类id
const value = ref([])

const props = {
  expandTrigger: 'hover' as const,
  label: 'title',
  value: 'id',
}

// 是否已经选定分类，用于控制上传控件是否显示
const isSelectedCategory = ref(false)
// 表格数据
const tableData = ref([])

// 获取文件列表
const getFilesList = (id) => {
    console.log("id:"+id)
    store.dispatch('fileModule/getFiles',{"id":id}).then((res)=>{
        tableData.value = res.data.data.files
        isSelectedCategory.value = true
        ca2id.value = id
    }).catch((res)=>{
        console.log(res.data.msg)
        isSelectedCategory.value = false
    })
}
const handleChange = (value) => {
    getFilesList(value[1])
//   store.dispatch('fileModule/getFiles',{"id":value[1]}).then((res)=>{
//     tableData.value = res.data.data.files
//     isSelectedCategory.value = true
//     ca2.value = value[1]
//   }).catch((res)=>{
//     console.log(res.data.msg)
//     isSelectedCategory.value = false
//   })
}

const options = ref([])

const DataInit = () => {
    // 获取分类列表
    store.dispatch('fileModule/getCategories').then((res) => {
      options.value = res.data.data.categories
    }).catch((res)=>{
      console.log(res)
    })
}
DataInit()

// 删除按钮
const handelDelete = (index) => {
    let fileID = tableData.value[index].id
    store.dispatch('fileModule/deleteFile',{"id":fileID}).then((res)=>{
        router.go(0)
    }).catch((res)=>{
        ElMessage.error('删除失败')
    })
}


</script>