<template>
  <el-card class="time-trend-chart" shadow="hover">
    <template #header>
      <div class="card-header">
        <span class="card-title">时间趋势</span>
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
import { ref, onMounted, onUnmounted, computed, watch } from 'vue'
import * as echarts from 'echarts/core'
import { LineChart, BarChart } from 'echarts/charts'
import { TitleComponent, TooltipComponent, LegendComponent, GridComponent } from 'echarts/components'
import { CanvasRenderer } from 'echarts/renderers'
import { getStatsByTime } from '@/api/stats'
import type { TimeStats } from '@/types/api'

echarts.use([LineChart, BarChart, TitleComponent, TooltipComponent, LegendComponent, GridComponent, CanvasRenderer])

const props = defineProps<{
  startDate?: string
  endDate?: string
}>()

const chartRef = ref<HTMLElement>()
const loading = ref(false)
const error = ref('')
const timeStats = ref<TimeStats | null>(null)
let chartInstance: echarts.ECharts | null = null

const hasData = computed(() => {
  return timeStats.value && timeStats.value.daily_trend.length > 0
})

const initChart = () => {
  if (!chartRef.value) return

  chartInstance = echarts.init(chartRef.value)
  updateChart()

  window.addEventListener('resize', handleResize)
}

const updateChart = () => {
  if (!chartInstance || !timeStats.value) return

  const dates = timeStats.value.daily_trend.map(item => item.date)
  const createdData = timeStats.value.daily_trend.map(item => item.created)
  const completedData = timeStats.value.daily_trend.map(item => item.completed)

  const option = {
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'cross',
        label: {
          backgroundColor: '#6a7985'
        }
      }
    },
    legend: {
      data: ['创建', '完成']
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: dates
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        name: '创建',
        type: 'line',
        stack: 'Total',
        smooth: true,
        lineStyle: {
          width: 0
        },
        showSymbol: false,
        areaStyle: {
          opacity: 0.8,
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: 'rgb(128, 255, 165)'
            },
            {
              offset: 1,
              color: 'rgb(1, 191, 236)'
            }
          ])
        },
        emphasis: {
          focus: 'series'
        },
        data: createdData
      },
      {
        name: '完成',
        type: 'line',
        stack: 'Total',
        smooth: true,
        lineStyle: {
          width: 0
        },
        showSymbol: false,
        areaStyle: {
          opacity: 0.8,
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: 'rgb(0, 221, 255)'
            },
            {
              offset: 1,
              color: 'rgb(77, 119, 255)'
            }
          ])
        },
        emphasis: {
          focus: 'series'
        },
        data: completedData
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
    const response = await getStatsByTime({
      start_date: props.startDate,
      end_date: props.endDate
    })
    timeStats.value = response.data
    updateChart()
  } catch (err: any) {
    error.value = err.message || '加载时间趋势数据失败'
    console.error('Failed to fetch time stats:', err)
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

// 监听日期变化
watch(() => [props.startDate, props.endDate], () => {
  fetchData()
}, { deep: true })
</script>

<style scoped>
.time-trend-chart {
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
  height: 350px;
}

@media (max-width: 768px) {
  .chart-container {
    height: 300px;
  }
}
</style>
