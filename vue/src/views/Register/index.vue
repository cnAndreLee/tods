<template>
    <div>
        用户:<input type="text" v-model="userInfo.Account" />
        密码:<input type="password" v-model="userInfo.Key" />
        <button @click="signup">注册</button>
        <p v-if="signupResult != ''">{{signupResult}}</p>
    </div>
</template>

<script>
import { reactive,ref } from "vue"
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'

export default {
    name: 'Register',
    setup (){
        const router = useRouter()
        const store = useStore()

        // data
        let userInfo = reactive({
            Account: '',
            Key: ''
        })
        let signupResult = ref(0)

        function signup() {
            console.log('signup in')
            if ( userInfo.Account.trim() === '' || userInfo.Key.trim() === '' ) {
                alert('账号或密码不能为空')
                return;
            }

            store.dispatch('userModule/register', userInfo).then(() => {
                signupResult.value = "注册成功, 正在跳转"
                setTimeout(() => {router.push('/')}, 2000)
            }).catch((err) => {
                console.log("signup failed")
                signupResult.value = "注册失败，错误代码：" + err.response.data.status
                console.log('err=>', err.response.data);
            })
            
        }

        return {
            userInfo,
            signupResult,
            signup
        }
    }
}
</script>

<style scoped>
div {
    margin: auto;
}
</style>