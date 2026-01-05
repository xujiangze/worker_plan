<template>
  <el-tag :type="priorityType" size="small">{{ priorityText }}</el-tag>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { PlanPriority } from '@/types/api'

interface Props {
  priority: PlanPriority
}

const props = defineProps<Props>()

const priorityText = computed(() => {
  const priorityMap: Record<PlanPriority, string> = {
    [PlanPriority.High]: '高',
    [PlanPriority.Medium]: '中',
    [PlanPriority.Low]: '低',
  }
  return priorityMap[props.priority] || props.priority
})

const priorityType = computed(() => {
  const typeMap: Record<PlanPriority, any> = {
    [PlanPriority.High]: 'danger',
    [PlanPriority.Medium]: 'warning',
    [PlanPriority.Low]: 'info',
  }
  return typeMap[props.priority] || 'info'
})
</script>
