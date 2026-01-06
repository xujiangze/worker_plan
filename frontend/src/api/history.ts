import request from '@/utils/request'
import type { ApiResponse, PlanHistory } from '@/types/api'

// 历史记录查询参数
export interface HistoryQueryParams {
  page?: number
  page_size?: number
  change_type?: string
  start_date?: string
  end_date?: string
}

// 历史记录响应
export interface HistoryResponse {
  total: number
  page: number
  page_size: number
  data: PlanHistory[]
}

// 获取计划历史记录
export const getPlanHistory = (
  planId: number,
  params: HistoryQueryParams
): Promise<ApiResponse<HistoryResponse>> => {
  return request.get(`/plans/${planId}/history`, { params })
}

// 获取计划详情
export const getPlanDetail = (planId: number): Promise<ApiResponse<any>> => {
  return request.get(`/plans/${planId}`)
}
