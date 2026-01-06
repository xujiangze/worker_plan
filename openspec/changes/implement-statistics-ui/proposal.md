# Change: 实现统计分析功能前端界面

## Why
当前系统已经实现了后端统计分析API,但前端StatisticsView.vue只是一个占位符,用户无法查看统计数据。需要实现完整的前端统计界面,让用户能够直观地查看工作计划的统计分析结果,包括按状态、优先级、时间维度的统计以及完成率展示。

## What Changes
- 完善后端GetStatsByTime方法的实现,支持按时间范围统计
- 创建前端统计API调用模块(stats.ts)
- 添加统计相关的TypeScript类型定义
- 实现StatisticsView.vue组件,包含:
  - 按状态统计的饼图展示
  - 按优先级统计的饼图展示
  - 按时间统计的趋势图展示
  - 完成率卡片展示
  - 日期范围选择器
- 集成图表库(如ECharts或Chart.js)用于数据可视化

## Impact
- Affected specs:
  - `statistics` (MODIFIED - 完善时间统计实现)
  - `statistics-ui` (NEW - 新增前端统计界面规范)
- Affected code:
  - 后端: `internal/service/stats_service.go` (完善GetStatsByTime)
  - 后端: `internal/repository/plan_repository.go` (添加时间范围统计方法)
  - 前端: `frontend/src/api/stats.ts` (新增)
  - 前端: `frontend/src/types/api.ts` (添加统计类型)
  - 前端: `frontend/src/views/StatisticsView.vue` (重写)
- Affected docs: 无

## Dependencies
- 依赖现有的后端统计API (`/api/stats/*`)
- 依赖前端Vue.js框架
- 需要选择并集成图表库(ECharts或Chart.js)
