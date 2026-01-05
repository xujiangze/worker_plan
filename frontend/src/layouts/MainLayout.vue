<template>
  <el-container class="main-layout">
    <el-header class="header">
      <div class="header-content">
        <h1 class="title">工作计划管理系统</h1>
      </div>
    </el-header>
    <el-main class="main">
      <router-view />
    </el-main>
  </el-container>
</template>

<script setup lang="ts">
// 基础布局组件
</script>

<style scoped>
.main-layout {
  min-height: 100vh;
}

.header {
  background-color: #409eff;
  color: white;
  display: flex;
  align-items: center;
  padding: 0 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.header-content {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
}

.title {
  margin: 0;
  font-size: 20px;
  font-weight: 500;
}

.main {
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
  width: 100%;
}
</style>
