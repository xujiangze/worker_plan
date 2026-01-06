<template>
  <div class="statistics-view">
    <div class="page-header">
      <h1 class="page-title">统计分析</h1>
      <DateRangePicker @change="handleDateChange" @clear="handleDateClear" />
    </div>

    <div class="statistics-content">
      <!-- 完成率卡片 -->
      <el-row :gutter="20" class="mb-20">
        <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
          <CompletionRateCard />
        </el-col>
      </el-row>

      <!-- 饼图区域 -->
      <el-row :gutter="20" class="mb-20">
        <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12">
          <StatusPieChart />
        </el-col>
        <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12">
          <PriorityPieChart />
        </el-col>
      </el-row>

      <!-- 时间趋势图 -->
      <el-row :gutter="20">
        <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
          <TimeTrendChart :start-date="startDate" :end-date="endDate" />
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import CompletionRateCard from '@/components/CompletionRateCard.vue'
import StatusPieChart from '@/components/StatusPieChart.vue'
import PriorityPieChart from '@/components/PriorityPieChart.vue'
import TimeTrendChart from '@/components/TimeTrendChart.vue'
import DateRangePicker from '@/components/DateRangePicker.vue'

const startDate = ref<string>()
const endDate = ref<string>()

const handleDateChange = (start: string, end: string) => {
  startDate.value = start
  endDate.value = end
}

const handleDateClear = () => {
  startDate.value = undefined
  endDate.value = undefined
}
</script>

<style scoped>
.statistics-view {
  padding: 20px;
  background-color: #f5f7fa;
  min-height: 100vh;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  flex-wrap: wrap;
  gap: 16px;
}

.page-title {
  font-size: 24px;
  font-weight: 600;
  color: #303133;
  margin: 0;
}

.statistics-content {
  max-width: 1400px;
  margin: 0 auto;
}

.mb-20 {
  margin-bottom: 20px;
}

@media (max-width: 768px) {
  .statistics-view {
    padding: 10px;
  }

  .page-header {
    flex-direction: column;
    align-items: flex-start;
  }

  .page-title {
    font-size: 20px;
  }

  .mb-20 {
    margin-bottom: 16px;
  }
}
</style>
