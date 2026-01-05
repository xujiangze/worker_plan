<template>
  <div class="progress-bar">
    <el-progress
      :percentage="percentage"
      :stroke-width="strokeWidth"
      :color="progressColor"
      :show-text="showText"
      :format="format"
    />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  percentage: number
  strokeWidth?: number
  showText?: boolean
  format?: (percentage: number) => string
}

const props = withDefaults(defineProps<Props>(), {
  strokeWidth: 8,
  showText: true,
  format: undefined,
})

// 根据进度百分比计算颜色
const progressColor = computed(() => {
  if (props.percentage < 30) {
    return '#f56c6c' // 红色
  } else if (props.percentage < 70) {
    return '#e6a23c' // 橙色
  } else if (props.percentage < 100) {
    return '#409eff' // 蓝色
  } else {
    return '#67c23a' // 绿色
  }
})
</script>

<style scoped>
.progress-bar {
  width: 100%;
}
</style>
