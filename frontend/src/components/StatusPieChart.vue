<template>
  <el-card class="status-pie-chart" shadow="hover">
    <template #header>
      <div class="card-header">
        <span class="card-title">状态分布</span>
      </div>
    </template>

    <div v-if="loading" class="loading-container">
      <el-skeleton :rows="3" animated />
    </div>

    <div v-else-if="error" class="error-container">
      <el-result icon="error" title="加载失败" :sub-title="error">
        <template #extra>
          <el-button type="primary" @click="fetchData">重试</el-button>
        </template>
      </el-result>
    </div>

    <div v-else-if="!hasData" class="empty-container">
      <el-empty description="暂无数据" />
    </div>

    <div v-else ref="chartRef" class="chart-container"></div>
  </el-card>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue'
import * as echarts from 'echarts/core'
import { PieChart } from 'echarts/charts'
import { TitleComponent, TooltipComponent, LegendComponent } from 'echarts/components'
import { CanvasRenderer } from 'echarts/renderers'
import { getStatsByStatus } from '@/api/stats'
import type { StatusStats } from '@/types/api'
import { statusMap } from '@/config/charts'

echarts.use([PieChart, TitleComponent, TooltipComponent, LegendComponent, CanvasRenderer])

const chartRef = ref<HTMLElement>()
const loading = ref(false)
const error = ref('')
const statusStats = ref<StatusStats[]>([])
let chartInstance: echarts.ECharts | null = null

const hasData = computed(() => {
  return statusStats.value.length > 0 && statusStats.value.some(item => item.count > 0)
})

const initChart = () => {
  if (!chartRef.value) return

  chartInstance = echarts.init(chartRef.value)
  updateChart()

  window.addEventListener('resize', handleResize)
}

const updateChart = () => {
  if (!chartInstance) return

  const data = statusStats.value
    .filter(item => item.count > 0)
    .map(item => ({
      value: item.count,
      name: statusMap[item.status] || item.status
    }))

  const option = {
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b}: {c} ({d}%)'
    },
    legend: {
      orient: 'vertical',
      left: 'left',
      data: data.map(item => item.name)
    },
    series: [
      {
        name: '状态分布',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        emphasis: {
          label: {
            show: true,
            fontSize: 20,
            fontWeight: 'bold'
          }
        },
        labelLine: {
          show: false
        },
        data: data
      }
    ]
  }

  chartInstance.setOption(option)
}

const handleResize = () => {
  chartInstance?.resize()
}

const fetchData = async () => {
  loading.value = true
  error.value = ''

  try {
    const response = await getStatsByStatus()
    statusStats.value = response.data || []
    updateChart()
  } catch (err: any) {
    error.value = err.message || '加载状态分布数据失败'
    console.error('Failed to fetch status stats:', err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchData()
  initChart()
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
  chartInstance?.dispose()
})
</script>

<style scoped>
.status-pie-chart {
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
.error-container,
.empty-container {
  padding: 20px 0;
  min-height: 300px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.chart-container {
  width: 100%;
  height: 300px;
}

@media (max-width: 768px) {
  .chart-container {
    height: 250px;
  }
}
</style>
