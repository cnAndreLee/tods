<template>
  <el-row align="middle">
    <el-col :span="3"><span>TODS</span></el-col>
    <el-col :span="18" style="height: 60px"></el-col>
    <el-col :span="3">
      <el-dropdown size="large" role="menu" >
        <el-button type="primary">
          <el-icon style="margin-right:5px"><User /></el-icon> {{userInfo==null? "":userInfo.account}}  <el-icon style="margin-left:5px"><arrow-down /></el-icon>
        </el-button>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item><el-icon><House /></el-icon>学校：{{userInfo==null? "":userInfo.schoolname}}</el-dropdown-item>
            <el-dropdown-item><el-icon><Calendar /></el-icon>到期日期：{{userInfo==null? "":userInfo.outtime}}</el-dropdown-item>
            <el-dropdown-item @click="manage"><el-icon><Management /></el-icon>管理</el-dropdown-item>
            <el-dropdown-item  divided @click="logout"><el-icon><SwitchButton /></el-icon>退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
      <!-- <router-link v-if="userInfo==null? false:userInfo.class === 'admin'" id="registerbtn" to="/register">管理</router-link> -->
    </el-col>
  </el-row>

</template>

<script setup>
import { ElNotification as notify } from 'element-plus'
import { ArrowDown,Calendar,User,House,SwitchButton,Management } from '@element-plus/icons-vue'
import { computed,onMounted } from 'vue'
// import storageService from '../../service/storagService'
import { useStore } from 'vuex'
import storageService from '../../service/storageService';
import { useRouter } from "vue-router"

const router = useRouter()

const store = useStore()

// 生命周期钩子：组件挂载后执行
const init = () => {
  if (storageService.get(storageService.USER_TOKEN) != null) {
      store.dispatch('userModule/info')
  } else {
      router.push('/login')
  }
};
init();

const userInfo = computed( ()  => {
    // return JSON.parse(storageService.get(storageService.USER_INFO))
    return store.state.userModule.userInfo
})
function logout() {
    store.dispatch('userModule/logout')
    router.push('/login')
}
function manage() {
    window.open(router.resolve({path:"/manage"}).href)
    //router.push('/manage')
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