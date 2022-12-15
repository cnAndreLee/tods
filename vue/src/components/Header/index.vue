<template>
  <el-dropdown role="navigation">
    <span class="el-dropdown-link">
      {{userInfo}}
      <el-icon class="el-icon--right">
        <arrow-down />
      </el-icon>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item>学校：xx</el-dropdown-item>
        <el-dropdown-item>到期日期：</el-dropdown-item>
        <el-dropdown-item @click="logout">退出登录</el-dropdown-item>
        <el-dropdown-item disabled>Action 4</el-dropdown-item>
        <el-dropdown-item divided>Action 5</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
  <router-link v-if="userInfo === 'mxdgg'" id="registerbtn" to="/register">用户管理</router-link>

<!-- <div class="headerdiv">
    <span>{{userInfo}}</span>
    
    <router-link v-if="!userInfo" id="loginbtn" to="/login">登录</router-link>
    
    <button @click="logout" v-if="userInfo" >退出登录</button>
</div> -->
</template>

<script setup>
import { ElNotification as notify } from 'element-plus'
import { ArrowDown } from '@element-plus/icons-vue'
import { computed,onMounted } from 'vue'
// import storageService from '../../service/storagService'
import { useStore } from 'vuex'
import storageService from '../../service/storageService';
import { useRouter } from "vue-router"

const router = useRouter()

const store = useStore()

// 生命周期钩子：组件挂载后执行
onMounted(() => {
    if (storageService.get(storageService.USER_TOKEN) != null) {
        store.dispatch('userModule/info')
    } else {
        router.push('/login')
    }
});

const userInfo = computed( ()  => {
    // return JSON.parse(storageService.get(storageService.USER_INFO))
    return store.state.userModule.userInfo
})
function logout() {
    store.dispatch('userModule/logout')
    router.push('/login')
}

</script>


<style scoped>
.example-showcase .el-dropdown-link {
  cursor: pointer;
  color: var(--el-color-primary);
  display: flex;
  align-items: center;
}
</style>