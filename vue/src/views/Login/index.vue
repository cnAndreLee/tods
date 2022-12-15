<template>
<div>
    用户<input type="text" v-model="data.userInfo.Account" />
    密码<input type="password" v-model="data.userInfo.Key" />
    <button @click="signin">登录</button>
    <p v-if="signinResult != ''">{{signinResult}}</p>
</div>
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
    console.log('signin in')
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
        // console.log("signin failed")
        // console.log(err)
        signinResult.value = "登录失败：" + err.data.msg
    });
};

// return {
//     data,
//     signinResult,
//     signin
// }
</script>

<style scoped>
div {
    margin: auto;
}
</style>