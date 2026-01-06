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

// 变更类型
export const ChangeType = {
  Status: 'Status',
  Progress: 'Progress',
  Info: 'Info',
} as const

export type ChangeType = (typeof ChangeType)[keyof typeof ChangeType]

// 计划历史记录
export interface PlanHistory {
  id: number
  plan_id: number
  field: string
  old_value: string
  new_value: string
  changed_at: string
  change_type?: ChangeType
}

// 状态统计
export interface StatusStats {
  status: string
  count: number
  percent: number
}

// 优先级统计
export interface PriorityStats {
  priority: string
  count: number
  percent: number
}

// 每日趋势项
export interface DailyTrend {
  date: string
  created: number
  completed: number
}

// 时间统计
export interface TimeStats {
  created_count: number
  completed_count: number
  completion_rate: number
  daily_trend: DailyTrend[]
}

// 完成率
export interface CompletionRate {
  total_plans: number
  completed_plans: number
  completion_rate: number
}

// 时间范围参数
export interface TimeRangeParams {
  start_date?: string
  end_date?: string
}
