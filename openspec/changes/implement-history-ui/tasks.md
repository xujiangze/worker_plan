# Tasks: 实现历史记录功能前端界面

## Task List

### 1. 创建历史记录 API 模块
**Description**: 创建 `frontend/src/api/history.ts` 文件,封装历史记录相关的 API 调用。

**Acceptance Criteria**:
- [x] 创建 `history.ts` 文件
- [x] 实现 `getPlanHistory()` 函数,支持分页和筛选参数
- [x] 实现 `getPlanDetail()` 函数,获取计划详情
- [x] 定义 TypeScript 类型: `HistoryQueryParams`, `HistoryResponse`
- [x] 导出所有 API 函数和类型

**Dependencies**: 无

**Estimated Time**: 30 分钟

---

### 2. 完善 TypeScript 类型定义
**Description**: 在 `frontend/src/types/api.ts` 中补充历史记录相关的类型定义。

**Acceptance Criteria**:
- [x] 补充 `PlanHistory` 接口,包含所有必需字段
- [x] 添加 `ChangeType` 类型定义 (Status, Progress, Info)
- [x] 添加 `HistoryQueryParams` 接口
- [x] 添加 `HistoryResponse` 接口
- [x] 确保类型定义与后端 API 响应一致

**Dependencies**: Task 1

**Estimated Time**: 15 分钟

---

### 3. 实现 HistoryView.vue 组件 - 基础结构
**Description**: 重写 `frontend/src/views/HistoryView.vue`,实现组件的基础结构和布局。

**Acceptance Criteria**:
- [x] 创建 Vue 3 Composition API 组件
- [x] 从路由参数获取 `planId`
- [x] 定义组件状态: `plan`, `histories`, `loading`, `filters`, `pagination`
- [x] 实现页面布局: 返回按钮、计划信息卡片、筛选栏、历史记录列表、分页
- [x] 使用 Element Plus 组件库

**Dependencies**: Task 2

**Estimated Time**: 45 分钟

---

### 4. 实现计划信息卡片
**Description**: 在 HistoryView 中实现计划基本信息展示卡片。

**Acceptance Criteria**:
- [x] 显示计划标题
- [x] 显示当前状态,使用不同颜色区分
- [x] 显示优先级,使用不同颜色区分
- [x] 显示进度百分比,使用进度条组件
- [x] 显示截止日期(如果有)
- [x] 卡片样式美观,符合设计规范

**Dependencies**: Task 3

**Estimated Time**: 30 分钟

---

### 5. 实现历史记录列表
**Description**: 在 HistoryView 中实现历史记录列表展示。

**Acceptance Criteria**:
- [x] 使用 `el-card` 或 `el-timeline` 组件展示历史记录
- [x] 每条记录显示: 变更类型标签、变更时间、字段名称、旧值、新值
- [x] 历史记录按时间倒序排列
- [x] 变更类型使用不同颜色: Status(蓝色)、Progress(绿色)、Info(灰色)
- [x] 处理空历史记录的情况,显示"暂无历史记录"提示

**Dependencies**: Task 3

**Estimated Time**: 45 分钟

---

### 6. 实现变更类型筛选
**Description**: 在 HistoryView 中实现按变更类型筛选历史记录的功能。

**Acceptance Criteria**:
- [x] 添加变更类型下拉选择器 (`el-select`)
- [x] 选项包括: 全部、状态、进度、信息
- [x] 选择变更类型后,自动刷新历史记录列表
- [x] 更新 URL 查询参数 `change_type`
- [x] 支持从 URL 查询参数恢复筛选状态

**Dependencies**: Task 5

**Estimated Time**: 30 分钟

---

### 7. 实现时间范围筛选
**Description**: 在 HistoryView 中实现按时间范围筛选历史记录的功能。

**Acceptance Criteria**:
- [x] 添加时间范围选择器 (`el-date-picker`)
- [x] 支持选择开始日期和结束日期
- [x] 选择时间范围后,自动刷新历史记录列表
- [x] 更新 URL 查询参数 `start_date` 和 `end_date`
- [x] 支持从 URL 查询参数恢复筛选状态
- [x] 支持清除时间范围选择

**Dependencies**: Task 5

**Estimated Time**: 30 分钟

---

### 8. 实现分页功能
**Description**: 在 HistoryView 中实现历史记录分页功能。

**Acceptance Criteria**:
- [x] 添加分页组件 (`el-pagination`)
- [x] 显示总记录数、当前页码、每页数量
- [x] 支持点击页码跳转
- [x] 支持上一页/下一页按钮
- [x] 默认每页显示 20 条记录
- [x] 更新 URL 查询参数 `page`
- [x] 支持从 URL 查询参数恢复分页状态

**Dependencies**: Task 5

**Estimated Time**: 30 分钟

---

### 9. 实现加载状态
**Description**: 在 HistoryView 中实现数据加载时的加载状态显示。

**Acceptance Criteria**:
- [x] 在数据加载时显示加载动画 (`el-loading`)
- [x] 加载时隐藏历史记录列表
- [x] 加载完成后显示历史记录列表
- [x] 处理计划详情和历史记录的加载状态

**Dependencies**: Task 3

**Estimated Time**: 15 分钟

---

### 10. 实现错误处理
**Description**: 在 HistoryView 中实现 API 请求失败时的错误处理。

**Acceptance Criteria**:
- [x] API 请求失败时显示错误提示 (`ElMessage`)
- [x] 错误信息说明失败原因
- [x] 提供"重试"按钮
- [x] 点击重试按钮重新发起请求
- [x] 处理 404 错误(计划不存在)
- [x] 处理网络错误

**Dependencies**: Task 3

**Estimated Time**: 30 分钟

---

### 11. 实现返回功能
**Description**: 在 HistoryView 中实现返回计划列表的功能。

**Acceptance Criteria**:
- [x] 添加"返回"按钮在页面左上角
- [x] 点击返回按钮跳转到计划列表页面 (`/plans`)
- [x] 使用 Element Plus 的 `el-button` 组件
- [x] 按钮样式符合设计规范

**Dependencies**: Task 3

**Estimated Time**: 10 分钟

---

### 12. 实现日期格式化
**Description**: 在 HistoryView 中实现日期和时间的格式化显示。

**Acceptance Criteria**:
- [x] 使用 `dayjs` 或 `date-fns` 库格式化日期
- [x] 变更时间格式: `YYYY-MM-DD HH:mm`
- [x] 截止日期格式: `YYYY-MM-DD`
- [x] 使用本地时区
- [x] 创建可复用的日期格式化函数

**Dependencies**: Task 3

**Estimated Time**: 15 分钟

---

### 13. 在计划列表添加历史记录入口
**Description**: 在 `frontend/src/views/PlansView.vue` 中添加跳转到历史记录页面的按钮。

**Acceptance Criteria**:
- [x] 在每个计划卡片或表格行中添加"查看历史"按钮
- [x] 点击按钮跳转到 `/history/:planId`
- [x] 按钮样式符合设计规范
- [x] 使用 Element Plus 的 `el-button` 组件

**Dependencies**: Task 3

**Estimated Time**: 20 分钟

---

### 14. 实现响应式设计
**Description**: 确保 HistoryView 在不同屏幕尺寸下都能正常显示。

**Acceptance Criteria**:
- [x] 桌面端(>= 1200px): 宽屏布局
- [x] 平板端(768px - 1199px): 中等布局
- [x] 移动端(< 768px): 单列布局
- [x] 筛选栏和分页控件适配移动端
- [x] 使用 CSS 媒体查询或 Element Plus 的响应式栅格系统

**Dependencies**: Task 3, Task 4, Task 5

**Estimated Time**: 30 分钟

---

### 15. 测试和优化
**Description**: 对实现的功能进行测试和优化。

**Acceptance Criteria**:
- [x] 测试所有功能: 列表展示、筛选、分页、加载、错误处理
- [x] 测试不同场景: 有历史记录、无历史记录、大量历史记录
- [x] 测试响应式布局在不同设备上的显示效果
- [x] 优化性能: 减少不必要的重新渲染
- [x] 优化用户体验: 添加过渡动画、优化加载速度
- [x] 修复发现的 bug

**Dependencies**: 所有前置任务

**Estimated Time**: 60 分钟

---

## Task Dependencies Graph

```
Task 1 (API 模块)
  └─> Task 2 (类型定义)
        └─> Task 3 (基础结构)
              ├─> Task 4 (计划信息卡片)
              ├─> Task 5 (历史记录列表)
              │     ├─> Task 6 (变更类型筛选)
              │     ├─> Task 7 (时间范围筛选)
              │     └─> Task 8 (分页功能)
              ├─> Task 9 (加载状态)
              ├─> Task 10 (错误处理)
              ├─> Task 11 (返回功能)
              ├─> Task 12 (日期格式化)
              └─> Task 14 (响应式设计)
Task 13 (计划列表入口) - 依赖 Task 3
Task 15 (测试和优化) - 依赖所有前置任务
```

## Parallelizable Tasks

以下任务可以并行开发:
- Task 6, Task 7, Task 8 (筛选和分页功能)
- Task 9, Task 10, Task 11, Task 12 (辅助功能)

## Total Estimated Time

约 7.5 小时 (450 分钟)

## Notes

- 所有任务完成后,需要更新路由配置(如果需要)
- 建议使用 Element Plus 组件库,保持与现有代码风格一致
- 需要安装 `dayjs` 或 `date-fns` 日期处理库
- 确保所有 API 调用都有适当的错误处理
- 遵循 Vue 3 Composition API 最佳实践
