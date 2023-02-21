<template>
    
    <el-row><p>合集管理</p></el-row>
    <el-row>
        <el-col :span="8"><p>请先选择需要管理的合集</p><el-cascader
            v-model="value"
            :options="options"
            :props="props"
            @change="handleChange"
            /></el-col>
        <el-col :span="8"></el-col>
        <el-col :span="8">
            <el-button 
                v-show="isSelectedCategory" 
                type="success" 
                plain 
                title="添加合集"
                @click="AddcollectionDialogVisible = true">
                <el-icon><Plus /></el-icon></el-button>
            <el-popconfirm title='确认删除此合集？？？'
                confirm-button-text='Yes'
                cancel-button-text='No'
                icon='el-icon-delete'
                icon-color='red'
                @confirm='deleteColletion'
                >
                <template #reference><el-button v-show="isSelectedCollection" type="danger" plain title="删除当前合集">
                    <el-icon><Delete /></el-icon></el-button></template></el-popconfirm>
            
            <el-button 
                v-show="isSelectedCollection" 
                type="warning" 
                plain 
                title="修改合集名"
                @click="renameCollectionDialogVisible = true"
                >
                <el-icon><Switch /></el-icon>></el-button>
        </el-col>
    </el-row>

    <el-row style="height: 30px;"></el-row>
    <el-row>
        <el-col :span="24">
            <el-table :data="tableData" style="width: 100%">
                <el-table-column prop="title" label="标题" style="width:50%" />
                <el-table-column prop="suffix" label="格式" style="width:30%" />
                <el-table-column prop="id" fixed="right" label="操作" style="width:20%">
                    <template #default="scope">
                        <el-popconfirm title='确认删除此文件？？？'
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
    <el-row v-if="isSelectedCollection">
        <el-col :span="3"></el-col>
        <el-col :span="18">
            <el-upload
                class="upload-demo"
                drag
                :action="fileUploadUrl"
                :headers="headers"
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

    <!-- 合集添加对话框 -->
    <el-dialog
        v-model="AddcollectionDialogVisible"
        title="请输入要添加的合集名称"
        width="30%"
        >
        <el-input v-model="addInput"></el-input>
        <span>{{ addInputResult }}</span>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="addInput = '';AddcollectionDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="createCollection">确认</el-button>
            </span>
        </template>
    </el-dialog>

    <!-- 合集更名对话框 -->
    <el-dialog
        v-model="renameCollectionDialogVisible"
        title="请输入新的合集名称"
        width="30%"
        >
        <el-input v-model="renameInput"></el-input>
        <span>{{ renameInputResult }}</span>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="renameInput = '';renameCollectionDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="renameCollection">确认</el-button>
            </span>
        </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ref, computed, h } from 'vue'
import {useStore} from 'vuex'
import { useRouter } from 'vue-router'
import storageService from '../../service/storageService'
import { UploadFilled } from '@element-plus/icons-vue'
import config from '../../config.js'
import { ElMessage, ElMessageBox, ElNotification  } from 'element-plus'
import type { UploadProps } from 'element-plus'

const store = useStore()
const router = useRouter()

let token = storageService.get(storageService.USER_TOKEN)
let headers = { Authorization: `Bearer ${token}`}
const fileUploadUrl = computed(() => {
    return config.backend + "/api/v1/file" + "?filebelong=" + value.value[1]
})


// 是否已经选定分类，用于控制合集操作按钮是否显示
const isSelectedCategory = ref(false)
// 是否已经选定合集，用于控制上传控件是否显示
const isSelectedCollection = ref(false)
// 表格数据
const tableData = ref([])


// 上传前 
const supportFileType = ["video/mp4", 
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


// value[0]为一级分类id,value[1]为二级分类id
const value = ref([])

const props = {
  expandTrigger: 'hover' as const,
  label: 'title',
  value: 'id',
}



// 获取文件列表
const getFilesList = (id) => {
    console.log("id:"+id)
    store.dispatch('fileModule/getFiles',{"id":id}).then((res)=>{
        tableData.value = res.data.data.files
        isSelectedCollection.value = true
    }).catch((res)=>{
        ElMessage.error(res.data.msg)
        resetData()
    })
}

// 处理选择框变化
const handleChange = (value) => {

    isSelectedCategory.value = true
    if (value[1] == undefined ) {
        ElMessage.error("此分类下没有合集, 请添加")
        resetData()
        
        return
    }
    getFilesList(value[1])
}

// 选择框内容
const options = ref([])

const DataInit = () => {
    // 获取分类列表
    store.dispatch('fileModule/getCategories').then((res) => {
        options.value = res.data.data.categories
    }).catch((res)=>{
        console.log(res)
    })
    isSelectedCategory.value = false
    isSelectedCollection.value = false
}
DataInit()

// 上传成功后
const handleSuccess = () => {
  ElMessage.success("上传成功")
  getFilesList(value.value[1])
}

// 删除按钮
const handelDelete = (index) => {
    let fileID = tableData.value[index].id
    store.dispatch('fileModule/deleteFile',{"id":fileID}).then((res)=>{
        getFilesList(value.value[1])
    }).catch((res)=>{
        ElMessage.error(res.data.msg)
    })
}

// 重置数据
const resetData = () => {
    isSelectedCollection.value = false
    tableData.value = []
}

// 添加合集dialog
const AddcollectionDialogVisible = ref(false)
const addInput = ref("")
const addInputResult = ref("")
const createCollection = () => {
    if (addInput.value == "") {
        addInputResult.value = "输入错误"
        return
    }

    // 创建二级分类
    const dtoCreateCategory = {
        "parent_id": value.value[0],
        "title": addInput.value,
        "level": 2
    }

    store.dispatch('fileModule/postCategory', dtoCreateCategory).then((res)=>{
        addInput.value = ""
        AddcollectionDialogVisible.value = false
        ElMessage.success("添加成功")
        DataInit()
    }).catch((res)=>{
        addInputResult.value =  res.data.msg
    })
}


// 删除合集
const deleteColletion = () => {
    store.dispatch('fileModule/deleteCategory',{"id":value.value[1]}).then((res)=>{
        DataInit()

    }).catch((res)=>{
        ElMessage.error('删除失败: ' + res.data.msg)
    })
}

// 合集更名
const renameCollectionDialogVisible = ref(false)
const renameInput = ref("")
const renameInputResult = ref("")
const renameCollection = () => {
    if (renameInput.value == "") {
        renameInputResult.value = "输入错误"
        return
    }

    const dtoRenameCategory = {
        "id": value.value[1],
        "new_title": renameInput.value,
    }

    store.dispatch('fileModule/renameCategory', dtoRenameCategory).then((res)=>{
        renameInput.value = ""
        renameCollectionDialogVisible.value = false
        ElMessage.success("修改成功")
        DataInit()
        isSelectedCategory.value = true
    isSelectedCollection.value = true
    }).catch((res)=>{
        renameInputResult.value =  res.data.msg
    })
}

</script>