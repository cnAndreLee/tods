<template>
    <el-row class="tac">
      <el-col :span="4">
        <h5 class="mb-2">TODS管理</h5>
        <el-menu
          default-active="2"
          class="el-menu-vertical-demo"
          @open="handleOpen"
          @close="handleClose"
        >
          <el-sub-menu index="1">
            <template #title>
                 <el-icon><UserFilled /></el-icon>
                <span>用户管理</span>
            </template>
            <el-menu-item index="1-1">用户管理</el-menu-item>
          </el-sub-menu>
          <el-sub-menu index="2">
            <template #title>
                <el-icon><Document /></el-icon>
                <span>资源管理</span>
            </template>
            <el-menu-item index="2-1">一级目录管理</el-menu-item>
            <el-menu-item index="2-2">二级目录管理</el-menu-item>
            <el-menu-item index="2-3">文件管理</el-menu-item>
          </el-sub-menu>
        </el-menu>
      </el-col>
    </el-row>
  </template>
  
  <script lang="ts" setup>
  import {
    Document,
    UserFilled,
  } from '@element-plus/icons-vue'
  const handleOpen = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
  }
  const handleClose = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
  }
  </script>
  