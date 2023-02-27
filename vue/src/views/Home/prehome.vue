<template>
    <div></div>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router';
import {useStore} from 'vuex'
import storageService from '../../service/storageService';
const store = useStore()
const router = useRouter()
const DataInit = () => {
    // 获取用户信息
    if (storageService.get(storageService.USER_TOKEN) != null) {
      store.dispatch('userModule/info').then(() => {
        router.push('/home')
      }).catch((err)=>{
        store.dispatch('userModule/logout')
        router.push('/login')
      })
    } else {
        store.dispatch('userModule/logout')
        router.push('/login')
    }
    
    
  };
  DataInit()
</script>