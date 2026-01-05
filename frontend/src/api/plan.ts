import request from '@/utils/request'
import type {
  Plan,
  CreatePlanRequest,
  UpdatePlanRequest,
  PaginationParams,
  PaginationResponse,
} from '@/types/api'

// 获取计划列表
export const getPlans = (params: PaginationParams): Promise<PaginationResponse<Plan>> => {
  return request.get('/plans', { params })
}

// 获取单个计划
export const getPlan = (id: number): Promise<Plan> => {
  return request.get(`/plans/${id}`)
}

// 创建计划
export const createPlan = (data: CreatePlanRequest): Promise<Plan> => {
  return request.post('/plans', data)
}

// 更新计划
export const updatePlan = (id: number, data: UpdatePlanRequest): Promise<Plan> => {
  return request.put(`/plans/${id}`, data)
}

// 删除计划
export const deletePlan = (id: number): Promise<void> => {
  return request.delete(`/plans/${id}`)
}

// 更新计划进度
export const updatePlanProgress = (id: number, progress: number): Promise<Plan> => {
  return request.patch(`/plans/${id}/progress`, { progress })
}
