<template>
  <el-card class="completion-rate-card" shadow="hover">
    <template #header>
      <div class="card-header">
        <span class="card-title">完成率</span>
        <el-button
          v-if="error"
          type="primary"
          size="small"
          @click="fetchData"
          :loading="loading"
        >
          重试
        </el-button>
      </div>
    </template>

    <div v-if="loading" class="loading-container">
      <el-skeleton :rows="2" animated />
    </div>

    <div v-else-if="error" class="error-container">
      <el-result icon="error" title="加载失败" :sub-title="error">
        <template #extra>
          <el-button type="primary" @click="fetchData">重试</el-button>
        </template>
      </el-result>
    </div>

    <div v-else class="completion-rate-content">
      <div class="rate-display">
        <div class="rate-value">{{ completionRate?.completion_rate.toFixed(1) }}%</div>
        <div class="rate-label">总体完成率</div>
      </div>
      <div class="stats-grid">
        <div class="stat-item">
          <div class="stat-value">{{ completionRate?.total_plans || 0 }}</div>
          <div class="stat-label">总计划数</div>
        </div>
        <div class="stat-item">
          <div class="stat-value">{{ completionRate?.completed_plans || 0 }}</div>
          <div class="stat-label">已完成</div>
        </div>
        <div class="stat-item">
          <div class="stat-value">
            {{ (completionRate?.total_plans || 0) - (completionRate?.completed_plans || 0) }}
          </div>
          <div class="stat-label">进行中</div>
        </div>
      </div>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { getCompletionRate } from '@/api/stats'
import type { CompletionRate } from '@/types/api'

const loading = ref(false)
const error = ref('')
const completionRate = ref<CompletionRate | null>(null)

const fetchData = async () => {
  loading.value = true
  error.value = ''

  try {
    const response = await getCompletionRate()
    completionRate.value = response.data
  } catch (err: any) {
    error.value = err.message || '加载完成率数据失败'
    console.error('Failed to fetch completion rate:', err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchData()
})
</script>

<style scoped>
.completion-rate-card {
  height: 100%;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-title {
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.loading-container,
.error-container {
  padding: 20px 0;
}

.completion-rate-content {
  padding: 10px 0;
}

.rate-display {
  text-align: center;
  margin-bottom: 20px;
}

.rate-value {
  font-size: 48px;
  font-weight: 700;
  color: #409eff;
  line-height: 1;
}

.rate-label {
  font-size: 14px;
  color: #909399;
  margin-top: 8px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
  margin-top: 20px;
}

.stat-item {
  text-align: center;
  padding: 12px;
  background-color: #f5f7fa;
  border-radius: 8px;
}

.stat-value {
  font-size: 24px;
  font-weight: 600;
  color: #303133;
  line-height: 1;
}

.stat-label {
  font-size: 12px;
  color: #909399;
  margin-top: 4px;
}

@media (max-width: 768px) {
  .stats-grid {
    grid-template-columns: repeat(3, 1fr);
    gap: 8px;
  }

  .stat-item {
    padding: 8px;
  }

  .stat-value {
    font-size: 18px;
  }

  .stat-label {
    font-size: 10px;
  }
}
</style>
