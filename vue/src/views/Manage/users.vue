<template>
    <el-row><h2>用户管理</h2></el-row>
    <el-row>
        <el-col :span="20"></el-col>
        <el-col :span="4">
            <el-button 
                type="success" 
                plain 
                title="添加用户"
                @click="UserAddDialogVisible = true">
                <el-icon><Plus /></el-icon></el-button></el-col>
    </el-row>
    <el-row>
        <el-table :data="tableData" style="width: 100%">
            <el-table-column prop="account" label="用户名" style="width:10%" />
            <el-table-column prop="key" label="密码" style="width:10%" />
            <el-table-column prop="is_admin" label="是否为管理员" style="width:10%" />
            <el-table-column prop="school" label="学校" style="width:16%" />
            <el-table-column prop="out_date" label="到期时间" style="width:16%" />
            <el-table-column prop="remark" label="备注" style="width:20%" />
            
            <el-table-column fixed="right" label="操作" style="width:18%">
                <template #default="scope">
                    <el-button type="success" plain @click="handleChangeKey(scope.row)">修改密码</el-button>
                    <el-popconfirm title='确认删除？？？'
                        confirm-button-text='Yes'
                        cancel-button-text='No'
                        icon='el-icon-delete'
                        icon-color='red'
                        @confirm='handelDelete(scope.$index,scope.row)'
                        >
                            <template #reference><el-button type="danger" plain>删除</el-button></template>
                    </el-popconfirm>
                </template>
            </el-table-column>
        </el-table>
    </el-row>

    

    <!-- 用户添加对话框 -->
    <el-dialog
        v-model="UserAddDialogVisible"
        title="请填写用户信息"
        width="30%"
        >
        <el-form :model="form" label-width="120px">
            <el-form-item label="用户名">
                <el-input v-model="form.account" />
            </el-form-item>
            <el-form-item label="密码">
                <el-input v-model="form.key" />
            </el-form-item>
            <el-form-item label="是否为管理员">
                <el-switch v-model="form.is_admin" />
            </el-form-item>
            <el-form-item label="学校">
                <el-input v-model="form.school" />
            </el-form-item>
            <el-form-item label="到期时间">
                <el-date-picker
                    v-model="form.out_date"
                    type="date"
                    placeholder="请选择日期"
                    value-format="YYYY-MM-DD"
                    :editable="false"
                    :clearable="false"
                    style="width: 100%"
                />
            </el-form-item>
            <el-form-item label="备注">
                <el-input v-model="form.remark" type="textarea" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="UserAddDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="handleAddUser">确认</el-button>
            </span>
        </template>
    </el-dialog>

    <!-- admin改密码对话框 -->
    <el-dialog
        v-model="changeKeyDialogVisible"
        title="请输入新的密码"
        width="30%"
        >
        <el-input v-model="newKey"></el-input>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="changeKeyDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="changeKey">确认</el-button>
            </span>
        </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ref, reactive, h } from 'vue'
import {useStore} from 'vuex'
import { useRouter } from 'vue-router'
import {ElMessage} from 'element-plus'
const store = useStore()
const router = useRouter()

// 表格数据
const tableData = ref([])
// 获取列表
const getUserList = () => {
    tableData.value = []
    store.dispatch('userModule/getAllUsers').then((res) => {
        tableData.value = res.data.data.users
    }).catch((res)=>{
        console.log(res)
    })
}

// 用户删除相关 -----------------------------------------------------------
const handelDelete = (index,row) => {
    if (row.account == "admin") {
        ElMessage.error("admin不可删除")
        return
    }
    console.log("delete user: "+tableData.value[index].account)
    let account = tableData.value[index].account

    store.dispatch('userModule/deleteUser', account).then((res)=>{
        getUserList()
        ElMessage.success("删除成功")
    }).catch((res)=>{
        ElMessage.error(res.data.msg)
    })
}


// 用户添加 -----------------------------------------------------------
// user dialog显示开关
const UserAddDialogVisible = ref(false)
const form = reactive({
    account: '',
    key: '',
    is_admin: false,
    school: '',
    out_date: '',
    remark:''
})
const handleAddUser = () => {

    store.dispatch('userModule/register', form).then((res) => {
        UserAddDialogVisible.value = false
        getUserList()
        ElMessage.success("注册成功")
        
    }).catch((res)=>{
        ElMessage.error(res.data.msg)
    })
}

// admin change key ----------------------------------------------------
const changeKeyDialogVisible = ref(false)
const selectedUserAccount = ref("")
const newKey = ref("")
const handleChangeKey = (row) => {
    changeKeyDialogVisible.value = true
    selectedUserAccount.value = row.account
    newKey.value = row.key
}

const changeKey = () => {
    const keyForm = {
        account:selectedUserAccount.value,
        new_key:newKey.value
    }
    store.dispatch('userModule/adminChangeKey', keyForm).then((res) => {
        changeKeyDialogVisible.value = false
        getUserList()
        ElMessage.success("修改成功")
        
    }).catch((res)=>{
        ElMessage.error(res.data.msg)
    })
}

const DataInit = () => {
    getUserList()
}
DataInit()

</script>