// 计划状态
export const PlanStatus = {
  Todo: 'Todo',
  InProgress: 'InProgress',
  Done: 'Done',
  Cancelled: 'Cancelled',
} as const

export type PlanStatus = (typeof PlanStatus)[keyof typeof PlanStatus]

// 计划优先级
export const PlanPriority = {
  High: 'High',
  Medium: 'Medium',
  Low: 'Low',
} as const

export type PlanPriority = (typeof PlanPriority)[keyof typeof PlanPriority]

// 计划模型
export interface Plan {
  id: number
  title: string
  description?: string
  priority: PlanPriority
  status: PlanStatus
  due_date?: string
  progress: number
  created_at: string
  updated_at: string
}

// 创建计划请求
export interface CreatePlanRequest {
  title: string
  description?: string
  priority: PlanPriority
  status?: PlanStatus
  due_date?: string
  progress?: number
}

// 更新计划请求
export interface UpdatePlanRequest {
  title?: string
  description?: string
  priority?: PlanPriority
  status?: PlanStatus
  due_date?: string
  progress?: number
}

// 分页参数
export interface PaginationParams {
  page?: number
  page_size?: number
  status?: PlanStatus
  priority?: PlanPriority
  keyword?: string
}

// 分页响应
export interface PaginationResponse<T> {
  items: T[]
  total: number
  page: number
  page_size: number
  total_pages: number
}

// API 响应
export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

// 计划历史记录
export interface PlanHistory {
  id: number
  plan_id: number
  field: string
  old_value: string
  new_value: string
  changed_at: string
}
