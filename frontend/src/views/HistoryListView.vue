<template>
  <div class="history-list-view">
    <el-card>
      <template #header>
        <div class="card-header">
          <span class="page-title">历史记录</span>
        </div>
      </template>

      <el-empty description="请先选择一个计划查看历史记录">
        <el-button type="primary" @click="goToPlans">前往计划列表</el-button>
      </el-empty>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const goToPlans = () => {
  router.push('/plans')
}
</script>

<style scoped>
.history-list-view {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.page-title {
  font-size: 18px;
  font-weight: 500;
  color: #303133;
}
</style>
