import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Plan, PaginationParams, CreatePlanRequest, UpdatePlanRequest } from '@/types/api'
import * as planApi from '@/api/plan'

export const usePlanStore = defineStore('plan', () => {
  // 状态
  const plans = ref<Plan[]>([])
  const currentPlan = ref<Plan | null>(null)
  const loading = ref(false)
  const total = ref(0)
  const currentPage = ref(1)
  const pageSize = ref(10)
  const filters = ref<PaginationParams>({})

  // 计算属性
  const totalPages = computed(() => Math.ceil(total.value / pageSize.value))

  // Actions
  const fetchPlans = async (params?: PaginationParams) => {
    loading.value = true
    try {
      const response = await planApi.getPlans({
        page: currentPage.value,
        page_size: pageSize.value,
        ...filters.value,
        ...params,
      })
      console.log('fetchPlans - 原始响应:', response)
      console.log('fetchPlans - items:', response.items)
      console.log('fetchPlans - items 类型:', typeof response.items)
      console.log('fetchPlans - items 长度:', response.items?.length)
      plans.value = response.items
      total.value = response.total
      currentPage.value = response.page
      pageSize.value = response.page_size
      console.log('fetchPlans - plans.value:', plans.value)
    } catch (error) {
      console.error('获取计划列表失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const fetchPlan = async (id: number) => {
    loading.value = true
    try {
      const plan = await planApi.getPlan(id)
      currentPlan.value = plan
      return plan
    } catch (error) {
      console.error('获取计划详情失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const createPlan = async (data: CreatePlanRequest) => {
    loading.value = true
    try {
      const plan = await planApi.createPlan(data)
      plans.value.unshift(plan)
      total.value += 1
      return plan
    } catch (error) {
      console.error('创建计划失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const updatePlan = async (id: number, data: UpdatePlanRequest) => {
    loading.value = true
    try {
      const updatedPlan = await planApi.updatePlan(id, data)
      const index = plans.value.findIndex((p) => p.id === id)
      if (index !== -1) {
        plans.value[index] = updatedPlan
      }
      if (currentPlan.value?.id === id) {
        currentPlan.value = updatedPlan
      }
      return updatedPlan
    } catch (error) {
      console.error('更新计划失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const deletePlan = async (id: number) => {
    loading.value = true
    try {
      await planApi.deletePlan(id)
      plans.value = plans.value.filter((p) => p.id !== id)
      total.value -= 1
      if (currentPlan.value?.id === id) {
        currentPlan.value = null
      }
    } catch (error) {
      console.error('删除计划失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const updateProgress = async (id: number, progress: number) => {
    loading.value = true
    try {
      const updatedPlan = await planApi.updatePlanProgress(id, progress)
      const index = plans.value.findIndex((p) => p.id === id)
      if (index !== -1) {
        plans.value[index] = updatedPlan
      }
      if (currentPlan.value?.id === id) {
        currentPlan.value = updatedPlan
      }
      return updatedPlan
    } catch (error) {
      console.error('更新进度失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const updateStatus = async (id: number, status: string) => {
    loading.value = true
    try {
      const updatedPlan = await planApi.updatePlanStatus(id, status)
      const index = plans.value.findIndex((p) => p.id === id)
      if (index !== -1) {
        plans.value[index] = updatedPlan
      }
      if (currentPlan.value?.id === id) {
        currentPlan.value = updatedPlan
      }
      return updatedPlan
    } catch (error) {
      console.error('更新状态失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const setFilters = (newFilters: PaginationParams) => {
    filters.value = { ...newFilters }
    currentPage.value = 1
  }

  const setPage = (page: number) => {
    currentPage.value = page
  }

  const setPageSize = (size: number) => {
    pageSize.value = size
    currentPage.value = 1
  }

  const clearCurrentPlan = () => {
    currentPlan.value = null
  }

  return {
    // 状态
    plans,
    currentPlan,
    loading,
    total,
    currentPage,
    pageSize,
    filters,
    // 计算属性
    totalPages,
    // Actions
    fetchPlans,
    fetchPlan,
    createPlan,
    updatePlan,
    deletePlan,
    updateProgress,
    updateStatus,
    setFilters,
    setPage,
    setPageSize,
    clearCurrentPlan,
  }
})
