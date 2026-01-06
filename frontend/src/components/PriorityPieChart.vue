<template>
  <el-card class="priority-pie-chart" shadow="hover">
    <template #header>
      <div class="card-header">
        <span class="card-title">优先级分布</span>
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
import { getStatsByPriority } from '@/api/stats'
import type { PriorityStats } from '@/types/api'
import { priorityMap } from '@/config/charts'

echarts.use([PieChart, TitleComponent, TooltipComponent, LegendComponent, CanvasRenderer])

const chartRef = ref<HTMLElement>()
const loading = ref(false)
const error = ref('')
const priorityStats = ref<PriorityStats[]>([])
let chartInstance: echarts.ECharts | null = null

const hasData = computed(() => {
  return priorityStats.value.length > 0 && priorityStats.value.some(item => item.count > 0)
})

const initChart = () => {
  if (!chartRef.value) return

  chartInstance = echarts.init(chartRef.value)
  updateChart()

  window.addEventListener('resize', handleResize)
}

const updateChart = () => {
  if (!chartInstance) return

  const data = priorityStats.value
    .filter(item => item.count > 0)
    .map(item => ({
      value: item.count,
      name: priorityMap[item.priority] || item.priority
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
        name: '优先级分布',
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
    const response = await getStatsByPriority()
    priorityStats.value = response.data || []
    updateChart()
  } catch (err: any) {
    error.value = err.message || '加载优先级分布数据失败'
    console.error('Failed to fetch priority stats:', err)
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
.priority-pie-chart {
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
