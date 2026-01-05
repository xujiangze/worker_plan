<template>
  <div v-if="loading" class="loading-container">
    <el-icon class="is-loading" :size="size">
      <Loading />
    </el-icon>
    <span v-if="text" class="loading-text">{{ text }}</span>
  </div>
</template>

<script setup lang="ts">
import { Loading } from '@element-plus/icons-vue'

interface Props {
  loading: boolean
  text?: string
  size?: number
}

withDefaults(defineProps<Props>(), {
  text: '',
  size: 24,
})
</script>

<style scoped>
.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.loading-text {
  margin-top: 10px;
  color: #909399;
  font-size: 14px;
}
</style>
