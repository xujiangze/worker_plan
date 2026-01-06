# Design: 历史记录功能前端界面

## Architecture Overview

### Component Structure
```
HistoryView.vue
├── PlanInfoCard (计划基本信息卡片)
├── FilterBar (筛选栏)
│   ├── ChangeTypeFilter (变更类型筛选)
│   └── DateRangeFilter (时间范围筛选)
├── HistoryList (历史记录列表)
│   └── HistoryItem (历史记录项)
└── Pagination (分页组件)
```

### Data Flow
1. 用户从计划列表点击"查看历史"按钮
2. 路由跳转到 `/history/:planId`
3. HistoryView 组件加载时:
   - 从路由参数获取 planId
   - 调用 API 获取计划详情
   - 调用 API 获取历史记录列表
4. 用户可以:
   - 按变更类型筛选历史记录
   - 按时间范围筛选历史记录
   - 翻页查看更多历史记录

## UI/UX Design

### Page Layout
```
┌─────────────────────────────────────────┐
│  返回按钮  历史记录 - [计划标题]         │
├─────────────────────────────────────────┤
│  ┌───────────────────────────────────┐ │
│  │  计划基本信息卡片                  │ │
│  │  标题、状态、优先级、进度等         │ │
│  └───────────────────────────────────┘ │
├─────────────────────────────────────────┤
│  筛选栏: [变更类型▼] [时间范围▼]       │
├─────────────────────────────────────────┤
│  ┌───────────────────────────────────┐ │
│  │  历史记录列表                      │ │
│  │  ┌─────────────────────────────┐ │ │
│  │  │  [状态] 2024-01-01 12:00    │ │ │
│  │  │  状态: Todo → InProgress     │ │ │
│  │  └─────────────────────────────┘ │ │
│  │  ┌─────────────────────────────┐ │ │
│  │  │  [进度] 2024-01-01 12:30    │ │ │
│  │  │  进度: 0% → 50%             │ │ │
│  │  └─────────────────────────────┘ │ │
│  └───────────────────────────────────┘ │
├─────────────────────────────────────────┤
│  分页: < 1 2 3 ... 10 >               │
└─────────────────────────────────────────┘
```

### Visual Design
- **变更类型标签**:
  - Status: 蓝色标签
  - Progress: 绿色标签
  - Info: 灰色标签
- **历史记录项**:
  - 卡片式设计,带阴影
  - 时间显示在右上角
  - 变更内容清晰展示旧值和新值
- **响应式设计**:
  - 移动端: 单列布局
  - 平板: 适当调整间距
  - 桌面: 宽屏布局

## Technical Decisions

### 1. API Integration
- 使用现有的 `/api/plans/{id}/history` API
- 支持分页参数: `page`, `page_size`
- 支持筛选参数: `change_type`, `start_date`, `end_date`

### 2. State Management
- 使用 Vue 3 Composition API
- 使用 `ref` 和 `reactive` 管理组件状态
- 状态包括:
  - `plan`: 计划详情
  - `histories`: 历史记录列表
  - `loading`: 加载状态
  - `filters`: 筛选条件
  - `pagination`: 分页信息

### 3. Date Handling
- 使用 `dayjs` 或 `date-fns` 处理日期格式化
- 显示格式: `YYYY-MM-DD HH:mm`
- 时间范围选择器使用 Element Plus 的 `el-date-picker`

### 4. Error Handling
- API 请求失败时显示错误提示
- 使用 Element Plus 的 `ElMessage` 组件
- 提供重试按钮

### 5. Performance Considerations
- 使用分页避免一次性加载大量数据
- 默认每页显示 20 条记录
- 实现虚拟滚动(如果历史记录非常多)

## Component Specifications

### HistoryView.vue
**Props**: 无
**State**:
- `planId`: number (从路由参数获取)
- `plan`: Plan | null
- `histories`: PlanHistory[]
- `loading`: boolean
- `filters`: { changeType?: string; startDate?: string; endDate?: string }
- `pagination`: { page: number; pageSize: number; total: number }

**Methods**:
- `fetchPlan()`: 获取计划详情
- `fetchHistories()`: 获取历史记录
- `handleFilterChange()`: 处理筛选条件变化
- `handlePageChange()`: 处理分页变化
- `formatDate()`: 格式化日期
- `getChangeTypeLabel()`: 获取变更类型标签
- `getChangeTypeColor()`: 获取变更类型颜色

### history.ts (API Module)
**Functions**:
- `getPlanHistory(planId: number, params: HistoryQueryParams): Promise<ApiResponse<HistoryResponse>>`
- `getPlanDetail(planId: number): Promise<ApiResponse<Plan>>`

**Types**:
```typescript
interface HistoryQueryParams {
  page?: number
  page_size?: number
  change_type?: string
  start_date?: string
  end_date?: string
}

interface HistoryResponse {
  total: number
  page: number
  page_size: number
  data: PlanHistory[]
}
```

## Accessibility
- 使用语义化 HTML 标签
- 为所有交互元素添加适当的 ARIA 标签
- 支持键盘导航
- 提供足够的颜色对比度

## Future Enhancements
- 支持导出历史记录为 CSV/Excel
- 支持历史记录搜索功能
- 支持历史记录对比功能(对比两个时间点的状态)
- 支持历史记录可视化(时间线视图)
