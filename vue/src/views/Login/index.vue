<template>
<el-row justify="center" align="middle">
    <el-col :span="3" style="min-height: 800px"></el-col>
    <el-col :span="6" class="logo">
        <img style="width: 100%; height: 480px" src="../../assets/login-box-bg.svg" /></el-col>
    <el-col :span="3"></el-col>
    <el-col :span="3"></el-col>
    <el-col :span="6">
        <el-card class="logincard" shadow="always" style="height:270px">
            <template #header>
                <h2>登录</h2>
            </template>
            <el-input class="inputbox" type="text" v-model="data.userInfo.Account" placeholder="请输入账号" clearable>
                <template #prepend>账号：</template>
            </el-input>

            <el-input class="inputbox" type="password" v-model="data.userInfo.Key" clearable placeholder="请输入密码" show-password>
                <template #prepend>密码：</template>
            </el-input>
            <el-button style="float:right" type="primary" @click="signin">登录</el-button>
            <p v-if="signinResult != ''">{{signinResult}}</p>
        </el-card>
    </el-col>
    <el-col :span="3"></el-col>
</el-row>
</template>

<script setup>
import { reactive, ref } from "vue"
import { useRouter } from "vue-router"
import { useStore } from 'vuex'

const router = useRouter()
const store = useStore()

// data
let data = reactive({
    userInfo :{
        Account: '',
        Key: ''
    }
})
let signinResult = ref(0)

function signin() {

    // 数据校验
    if ( data.userInfo.Account.trim() === '' || data.userInfo.Key.trim() === '' ) {
        signinResult.value = '账号或密码不能为空'
        return ;
    }

    store.dispatch('userModule/login', data.userInfo).then(() => {
        signinResult.value = "登录成功, 正在跳转..."
        // route to home
        setTimeout(() => {router.push('/')}, 500)
    }).catch((err) => {
        console.log(err.data.status)
        if (err.data.status >= 4000) {
            signinResult.value = "登录失败：" + err.data.msg
        } else {
            signinResult.value = "服务器错误，请联系管理员"
        }
        data.userInfo.Key = ""
    });
};

</script>

<style scoped>

/* html {
    background-image: url("../../assets/login-bg.svg");
    background-position: 100%;
    background-repeat: no-repeat;
    background-size: auto 100%;
} */

.logo {
    font-size:200px;
    color: aqua;
}

.inputbox {
    margin-bottom: 8px;
}

</style>