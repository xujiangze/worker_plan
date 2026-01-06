import request from '@/utils/request'
import type {
  StatusStats,
  PriorityStats,
  TimeStats,
  CompletionRate,
  TimeRangeParams,
  ApiResponse,
} from '@/types/api'

// 按状态统计
export const getStatsByStatus = (): Promise<ApiResponse<StatusStats[]>> => {
  return request.get('/stats/by-status')
}

// 按优先级统计
export const getStatsByPriority = (): Promise<ApiResponse<PriorityStats[]>> => {
  return request.get('/stats/by-priority')
}

// 按时间统计
export const getStatsByTime = (params: TimeRangeParams): Promise<ApiResponse<TimeStats>> => {
  return request.get('/stats/by-time', { params })
}

// 获取完成率
export const getCompletionRate = (): Promise<ApiResponse<CompletionRate>> => {
  return request.get('/stats/completion-rate')
}
