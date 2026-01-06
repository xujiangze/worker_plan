<template>
  <el-container class="main-layout">
    <el-header class="header">
      <div class="header-content">
        <h1 class="title">工作计划管理系统</h1>
      </div>
    </el-header>
    <el-container class="content-container">
      <NavigationMenu />
      <el-main class="main">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import NavigationMenu from '@/components/NavigationMenu.vue'
</script>

<style scoped>
.main-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.header {
  background-color: #409eff;
  color: white;
  display: flex;
  align-items: center;
  padding: 0 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  height: 60px;
  flex-shrink: 0;
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

.content-container {
  flex: 1;
  display: flex;
  overflow: hidden;
}

.main {
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
  width: 100%;
  overflow-y: auto;
}

@media (max-width: 768px) {
  .header {
    padding: 0 10px;
  }

  .title {
    font-size: 16px;
  }

  .main {
    padding: 10px;
  }
}
</style>
