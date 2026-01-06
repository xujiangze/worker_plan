<template>
  <div class="date-range-picker">
    <el-date-picker
      v-model="dateRange"
      type="daterange"
      range-separator="至"
      start-placeholder="开始日期"
      end-placeholder="结束日期"
      format="YYYY-MM-DD"
      value-format="YYYY-MM-DD"
      :clearable="true"
      @change="handleDateChange"
      @clear="handleClear"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

const emit = defineEmits<{
  (e: 'change', startDate: string, endDate: string): void
  (e: 'clear'): void
}>()

const dateRange = ref<[string, string] | null>(null)

const handleDateChange = (value: [string, string] | null) => {
  if (value && value.length === 2) {
    const [startDate, endDate] = value
    if (startDate && endDate) {
      emit('change', startDate, endDate)
    }
  }
}

const handleClear = () => {
  dateRange.value = null
  emit('clear')
}

const setDefaultRange = () => {
  // 设置默认为最近30天
  const endDate = new Date()
  const startDate = new Date()
  startDate.setDate(startDate.getDate() - 30)

  const formatDate = (date: Date) => {
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, '0')
    const day = String(date.getDate()).padStart(2, '0')
    return `${year}-${month}-${day}`
  }

  dateRange.value = [formatDate(startDate), formatDate(endDate)]
  emit('change', formatDate(startDate), formatDate(endDate))
}

onMounted(() => {
  setDefaultRange()
})
</script>

<style scoped>
.date-range-picker {
  display: inline-block;
}

:deep(.el-date-editor) {
  width: 280px;
}

@media (max-width: 768px) {
  :deep(.el-date-editor) {
    width: 100%;
  }
}
</style>
