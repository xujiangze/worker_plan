<template>
  <el-tag :type="statusType" size="small">{{ statusText }}</el-tag>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { PlanStatus } from '@/types/api'

interface Props {
  status: PlanStatus
}

const props = defineProps<Props>()

const statusText = computed(() => {
  const statusMap: Record<PlanStatus, string> = {
    [PlanStatus.Todo]: '待办',
    [PlanStatus.InProgress]: '进行中',
    [PlanStatus.Done]: '已完成',
    [PlanStatus.Cancelled]: '已取消',
  }
  return statusMap[props.status] || props.status
})

const statusType = computed(() => {
  const typeMap: Record<PlanStatus, any> = {
    [PlanStatus.Todo]: 'info',
    [PlanStatus.InProgress]: 'warning',
    [PlanStatus.Done]: 'success',
    [PlanStatus.Cancelled]: 'danger',
  }
  return typeMap[props.status] || 'info'
})
</script>
