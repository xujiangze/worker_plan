import { defineStore } from 'pinia'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

export const useUiStore = defineStore('ui', () => {
  // 状态
  const loading = ref(false)
  const message = ref<{
    type: 'success' | 'warning' | 'info' | 'error'
    text: string
  } | null>(null)

  // Actions
  const setLoading = (isLoading: boolean) => {
    loading.value = isLoading
  }

  const showSuccess = (message: string) => {
    ElMessage.success(message)
  }

  const showError = (message: string) => {
    ElMessage.error(message)
  }

  const showWarning = (message: string) => {
    ElMessage.warning(message)
  }

  const showInfo = (message: string) => {
    ElMessage.info(message)
  }

  const clearMessage = () => {
    message.value = null
  }

  return {
    // 状态
    loading,
    message,
    // Actions
    setLoading,
    showSuccess,
    showError,
    showWarning,
    showInfo,
    clearMessage,
  }
})
