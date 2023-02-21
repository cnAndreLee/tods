<template>
    <el-row><h2>类别管理</h2></el-row>
    <el-row>
        <el-input v-model="inputTitle" placeholder="输入分类名以新增分类" />
        <el-button type="primary" round @click="handleCreateCategory">新增分类</el-button></el-row>
    <el-row>
        <el-table :data="tableData" style="width: 100%">
            <el-table-column prop="title" label="标题" style="width:50%" />
            <el-table-column prop="permissions" label="授权列表" style="width:30%" />
            <el-table-column prop="id" fixed="right" label="操作" style="width:20%">
                <template #default="scope">
                    <el-button type="success" plain @click="handleRename(scope.row)">修改标题</el-button>
                    <el-button type="success" plain @click="handleEdit(scope.row)">编辑权限</el-button>
                    <el-popconfirm title='确认删除？？？'
                    confirm-button-text='Yes'
                    cancel-button-text='No'
                    icon='el-icon-delete'
                    icon-color='red'
                    @confirm='handelDelete(scope.$index)'
                    >
                        <template #reference><el-button type="danger" plain>删除</el-button></template>
                    </el-popconfirm>
                </template>
            </el-table-column>
        </el-table>
    </el-row>

    <!-- 权限dialog -->
    <el-dialog 
        title="请选择允许访问此分类的学校"
        v-model="ChangePermissionDialogVisible"
        width="30%"
        >
        <el-checkbox-group  
        v-model="checkedSchools"
        >
            <el-checkbox v-for="school in schools" :label="school" :key="school"></el-checkbox>
        </el-checkbox-group>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="ChangePermissionDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="changePermission">确认</el-button>
            </span>
        </template>
    </el-dialog>

    <!-- 分类更名对话框 -->
    <el-dialog
        v-model="renameCategoryDialogVisible"
        title="请输入新的合集名称"
        width="30%"
        >
        <el-input v-model="renameInput"></el-input>
        <span>{{ renameInputResult }}</span>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="renameCategoryDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="renameCategory">确认</el-button>
            </span>
        </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ref, computed, h } from 'vue'
import {useStore} from 'vuex'
import { useRouter } from 'vue-router'
import {ElMessage} from 'element-plus'
const store = useStore()
const router = useRouter()

// 表格数据
const tableData = ref([])
// 获取一级分类列表
const getCategoryList = () => {
    tableData.value = []
    store.dispatch('fileModule/getCategories').then((res) => {
        tableData.value = res.data.data.categories
    }).catch((res)=>{
        console.log(res)
    })
}

// 添加分类---------------------------------------------------
const inputTitle = ref("")
const handleCreateCategory = () => {
    // 创建一级分类
    const dtoCreateCategory = {
        "parent_id": "0",
        "title": inputTitle.value,
        "level": 1
    }

    store.dispatch('fileModule/postCategory', dtoCreateCategory).then((res)=>{
        inputTitle.value = ""
        DataInit()
    }).catch((res)=>{
        ElMessage.error(res.data.msg)
    })

}

// 权限编辑相关 -----------------------------------------------------------
const schools = ref([])
const checkedSchools = ref([])
const selectedCategoryID = ref("")
const getUserScholls = () => {
    store.dispatch('userModule/getAllSchools').then((res)=>{
        schools.value = res.data.data.schools
    })
}
getUserScholls()

// 权限dialog显示
const ChangePermissionDialogVisible = ref(false)
const handleEdit = (row) => {
    
    ChangePermissionDialogVisible.value = true
    if (row.permissions != null) {
        checkedSchools.value = row.permissions
    }
    selectedCategoryID.value = row.id
}
const changePermission = () => {
    const payload = {
        "id":selectedCategoryID.value,
        "new_permissions":checkedSchools.value
    }

    store.dispatch('fileModule/changeCategoryPermissions', payload).then((res)=>{
        ChangePermissionDialogVisible.value = false
        ElMessage.success("修改成功")
        getCategoryList()
    }).catch((res)=>{
        ElMessage.error(res.data.msg)
    })
}

// 分类更名相关 -----------------------------------------------------------

const renameCategoryDialogVisible = ref(false)
const renameInput = ref("")
const renameInputResult = ref("")
const handleRename = (row) => {
    
    renameCategoryDialogVisible.value = true
    selectedCategoryID.value = row.id
    renameInput.value = row.title
}

const renameCategory = () => {
    if (renameInput.value == "") {
        renameInputResult.value = "输入错误"
        return
    }

    const dtoRenameCategory = {
        "id": selectedCategoryID.value,
        "new_title": renameInput.value,
    }

    store.dispatch('fileModule/renameCategory', dtoRenameCategory).then((res)=>{
        renameInput.value = ""
        renameCategoryDialogVisible.value = false
        ElMessage.success("修改成功")
        getCategoryList()

    }).catch((res)=>{
        renameInputResult.value =  res.data.msg
    })
}

// 分类删除---------------------------------------------------------------

const handelDelete = (index) => {
    let categoryID = tableData.value[index].id
    store.dispatch('fileModule/deleteCategory',{"id":categoryID}).then((res)=>{
        DataInit()
    }).catch((res)=>{
        ElMessage.error('删除失败: ' + res.data.msg)
    })
}

const DataInit = () => {
    getCategoryList()
}
DataInit()

</script>