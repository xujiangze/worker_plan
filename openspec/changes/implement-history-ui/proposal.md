# Change: 实现历史记录功能前端界面

## Why
当前系统已经实现了后端历史记录API (`/api/plans/{id}/history`),但前端 HistoryView.vue 只是一个占位符,用户无法查看计划的历史变更记录。需要实现完整的前端历史记录界面,让用户能够直观地查看计划的所有变更历史,包括状态变更、进度变更和信息变更。

## What Changes
- 创建前端历史记录 API 调用模块 (history.ts)
- 完善前端 TypeScript 类型定义 (PlanHistory 类型已存在,需补充)
- 实现 HistoryView.vue 组件,包含:
  - 计划基本信息展示
  - 历史记录列表展示
  - 按变更类型筛选功能 (Status, Progress, Info)
  - 按时间范围筛选功能
  - 分页功能
  - 变更详情展示 (字段名、旧值、新值、变更时间)
  - 变更类型标签和颜色区分
- 添加从计划列表跳转到历史记录页面的入口

## Impact
- Affected specs:
  - `history-ui` (NEW - 新增前端历史记录界面规范)
- Affected code:
  - 前端: `frontend/src/api/history.ts` (新增)
  - 前端: `frontend/src/types/api.ts` (补充 PlanHistory 类型)
  - 前端: `frontend/src/views/HistoryView.vue` (重写)
  - 前端: `frontend/src/views/PlansView.vue` (添加跳转按钮)
- Affected docs: 无

## Dependencies
- 依赖现有的后端历史记录 API (`GET /api/plans/{id}/history`)
- 依赖前端 Vue.js 框架和 Element Plus UI 组件库
- 依赖现有的 `history-tracking` 后端规范
