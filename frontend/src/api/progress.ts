import request from '@/utils/request'
import type { PlanHistory, PaginationResponse } from '@/types/api'

// 获取计划历史记录
export const getPlanHistory = (
  planId: number,
  page = 1,
  pageSize = 10
): Promise<PaginationResponse<PlanHistory>> => {
  return request.get(`/plans/${planId}/history`, {
    params: { page, page_size: pageSize },
  })
}

// 获取所有计划的历史记录
export const getAllHistory = (
  page = 1,
  pageSize = 10
): Promise<PaginationResponse<PlanHistory>> => {
  return request.get('/history', {
    params: { page, page_size: pageSize },
  })
}
