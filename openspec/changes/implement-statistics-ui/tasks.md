# Tasks: 实现统计分析功能前端界面

## Backend Tasks

### 1. 完善后端时间统计实现
**Description**: 完善 `GetStatsByTime` 方法,实现按时间范围统计功能。

**Steps**:
- [x] 在 `internal/repository/plan_repository.go` 中添加 `CountByDateRange` 方法
- [x] 在 `internal/repository/plan_repository.go` 中添加 `GetDailyTrend` 方法
- [x] 更新 `internal/service/stats_service.go` 中的 `GetStatsByTime` 方法
- [x] 添加日期格式验证逻辑
- [x] 添加日期范围验证逻辑(开始日期不能晚于结束日期)
- [x] 确保排除已删除的计划

**Validation**:
- 单元测试覆盖所有场景
- API测试验证响应格式
- 性能测试确保响应时间 < 200ms (P95)

**Dependencies**: 无

**Estimated Time**: 2-3 hours

---

### 2. 添加数据库索引
**Description**: 为统计查询添加必要的数据库索引,提升查询性能。

**Steps**:
- [x] 在 `status` 字段上添加索引
- [x] 在 `priority` 字段上添加索引
- [x] 在 `created_at` 字段上添加索引
- [x] 在 `updated_at` 字段上添加索引
- [x] 创建数据库迁移文件
- [x] 测试索引效果

**Validation**:
- 验证索引创建成功
- 性能测试验证查询速度提升
- 确保不影响现有功能

**Dependencies**: Task 1

**Estimated Time**: 1 hour

---

## Frontend Tasks

### 3. 安装和配置图表库
**Description**: 安装 ECharts 图表库并进行基础配置。

**Steps**:
- [x] 在前端项目中安装 ECharts: `npm install echarts`
- [x] 安装 ECharts Vue 组件: `npm install vue-echarts`
- [x] 在 `main.ts` 中注册 ECharts 组件
- [x] 创建图表配置文件 `frontend/src/config/charts.ts`
- [x] 配置图表主题和默认选项

**Validation**:
- 验证 ECharts 安装成功
- 创建测试页面验证图表渲染
- 确保打包构建正常

**Dependencies**: 无

**Estimated Time**: 1 hour

---

### 4. 添加统计类型定义
**Description**: 在 `frontend/src/types/api.ts` 中添加统计相关的 TypeScript 类型定义。

**Steps**:
- [x] 添加 `StatusStats` 接口
- [x] 添加 `PriorityStats` 接口
- [x] 添加 `TimeStats` 接口
- [x] 添加 `DailyTrend` 接口
- [x] 添加 `CompletionRate` 接口
- [x] 添加 `TimeRangeParams` 接口

**Validation**:
- TypeScript 编译无错误
- 类型定义与后端响应格式一致
- 添加类型注释和文档

**Dependencies**: 无

**Estimated Time**: 0.5 hour

---

### 5. 创建统计API模块
**Description**: 创建 `frontend/src/api/stats.ts` 模块,封装所有统计相关的API调用。

**Steps**:
- [x] 创建 `frontend/src/api/stats.ts` 文件
- [x] 实现 `getStatsByStatus` 函数
- [x] 实现 `getStatsByPriority` 函数
- [x] 实现 `getStatsByTime` 函数
- [x] 实现 `getCompletionRate` 函数
- [x] 添加错误处理逻辑
- [x] 添加类型注解

**Validation**:
- API调用返回正确的数据格式
- 错误处理正常工作
- 与后端API集成测试通过

**Dependencies**: Task 4

**Estimated Time**: 1 hour

---

### 6. 实现完成率卡片组件
**Description**: 创建完成率卡片组件,显示总体完成情况。

**Steps**:
- [x] 创建 `frontend/src/components/CompletionRateCard.vue` 组件
- [x] 实现数据获取逻辑
- [x] 实现UI布局和样式
- [x] 添加加载状态
- [x] 添加错误处理
- [x] 添加重试功能
- [x] 添加单元测试

**Validation**:
- 组件正确显示完成率数据
- 加载状态正常显示
- 错误处理正常工作
- 响应式布局正常

**Dependencies**: Task 5

**Estimated Time**: 2 hours

---

### 7. 实现状态分布饼图组件
**Description**: 创建状态分布饼图组件,展示各状态计划的分布。

**Steps**:
- [x] 创建 `frontend/src/components/StatusPieChart.vue` 组件
- [x] 集成 ECharts 饼图
- [x] 实现数据转换逻辑
- [x] 实现鼠标悬停交互
- [x] 实现图例点击交互
- [x] 添加空状态处理
- [x] 添加加载状态
- [x] 添加错误处理
- [x] 添加单元测试

**Validation**:
- 饼图正确显示状态分布
- 交互功能正常工作
- 空状态正常显示
- 响应式布局正常

**Dependencies**: Task 3, Task 5

**Estimated Time**: 3 hours

---

### 8. 实现优先级分布饼图组件
**Description**: 创建优先级分布饼图组件,展示各优先级计划的分布。

**Steps**:
- [x] 创建 `frontend/src/components/PriorityPieChart.vue` 组件
- [x] 集成 ECharts 饼图
- [x] 实现数据转换逻辑
- [x] 实现鼠标悬停交互
- [x] 实现图例点击交互
- [x] 添加空状态处理
- [x] 添加加载状态
- [x] 添加错误处理
- [x] 添加单元测试

**Validation**:
- 饼图正确显示优先级分布
- 交互功能正常工作
- 空状态正常显示
- 响应式布局正常

**Dependencies**: Task 3, Task 5

**Estimated Time**: 3 hours

---

### 9. 实现时间趋势图组件
**Description**: 创建时间趋势图组件,展示指定时间范围内的创建和完成趋势。

**Steps**:
- [x] 创建 `frontend/src/components/TimeTrendChart.vue` 组件
- [x] 集成 ECharts 折线图/柱状图
- [x] 实现数据转换逻辑
- [x] 实现鼠标悬停交互
- [x] 实现缩放和平移功能
- [x] 添加空状态处理
- [x] 添加加载状态
- [x] 添加错误处理
- [x] 添加单元测试

**Validation**:
- 图表正确显示时间趋势
- 交互功能正常工作
- 空状态正常显示
- 响应式布局正常
- 大数据量性能良好

**Dependencies**: Task 3, Task 5

**Estimated Time**: 4 hours

---

### 10. 实现日期范围选择器
**Description**: 创建日期范围选择器组件,允许用户选择自定义时间范围。

**Steps**:
- [x] 创建 `frontend/src/components/DateRangePicker.vue` 组件
- [x] 集成 Element Plus 日期选择器
- [x] 实现日期范围验证
- [x] 实现日期格式化
- [x] 实现默认值设置(最近30天)
- [x] 添加错误提示
- [x] 添加防抖功能
- [x] 添加单元测试

**Validation**:
- 日期选择器正常工作
- 日期验证正常
- 错误提示正常显示
- 防抖功能正常

**Dependencies**: 无

**Estimated Time**: 2 hours

---

### 11. 实现统计页面主组件
**Description**: 重写 `frontend/src/views/StatisticsView.vue` 组件,整合所有统计组件。

**Steps**:
- [x] 重写 `frontend/src/views/StatisticsView.vue` 组件
- [x] 整合完成率卡片组件
- [x] 整合状态分布饼图组件
- [x] 整合优先级分布饼图组件
- [x] 整合时间趋势图组件
- [x] 整合日期范围选择器
- [x] 实现响应式布局
- [x] 实现页面加载状态
- [x] 实现全局错误处理
- [x] 添加页面标题和描述
- [x] 添加单元测试

**Validation**:
- 页面正确显示所有统计组件
- 响应式布局正常工作
- 加载状态正常显示
- 错误处理正常工作
- 页面性能符合要求

**Dependencies**: Task 6, Task 7, Task 8, Task 9, Task 10

**Estimated Time**: 4 hours

---

## Testing Tasks

### 12. 编写后端单元测试
**Description**: 为后端统计功能编写单元测试。

**Steps**:
1. 为 `GetStatsByStatus` 编写测试用例
2. 为 `GetStatsByPriority` 编写测试用例
3. 为 `GetStatsByTime` 编写测试用例
4. 为 `GetCompletionRate` 编写测试用例
5. 测试边界情况(空数据、大数据量)
6. 测试错误处理
7. 确保测试覆盖率 > 80%

**Validation**:
- 所有测试用例通过
- 测试覆盖率达标
- CI/CD 集成测试通过

**Dependencies**: Task 1, Task 2

**Estimated Time**: 3 hours

---

### 13. 编写前端组件测试
**Description**: 为前端统计组件编写单元测试和集成测试。

**Steps**:
1. 为完成率卡片组件编写测试
2. 为状态分布饼图组件编写测试
3. 为优先级分布饼图组件编写测试
4. 为时间趋势图组件编写测试
5. 为日期范围选择器编写测试
6. 为统计页面主组件编写测试
7. 测试用户交互
8. 测试响应式布局
9. 确保测试覆盖率 > 70%

**Validation**:
- 所有测试用例通过
- 测试覆盖率达标
- CI/CD 集成测试通过

**Dependencies**: Task 6, Task 7, Task 8, Task 9, Task 10, Task 11

**Estimated Time**: 4 hours

---

### 14. 编写端到端测试
**Description**: 编写端到端测试,验证整个统计功能的用户流程。

**Steps**:
1. 使用 Cypress 或 Playwright 编写 E2E 测试
2. 测试页面加载
3. 测试日期范围选择
4. 测试图表交互
5. 测试错误处理
6. 测试响应式布局
7. 测试性能指标

**Validation**:
- 所有 E2E 测试通过
- 用户流程验证通过
- 性能指标达标

**Dependencies**: Task 11

**Estimated Time**: 3 hours

---

## Documentation Tasks

### 15. 更新API文档
**Description**: 更新 `docs/API.md` 文档,添加统计API的详细说明。

**Steps**:
1. 添加 `/api/stats/by-status` 端点文档
2. 添加 `/api/stats/by-priority` 端点文档
3. 添加 `/api/stats/by-time` 端点文档
4. 添加 `/api/stats/completion-rate` 端点文档
5. 添加请求和响应示例
6. 添加错误码说明
7. 添加使用示例

**Validation**:
- 文档完整准确
- 示例代码可运行
- 格式符合项目规范

**Dependencies**: Task 1

**Estimated Time**: 1 hour

---

### 16. 更新用户文档
**Description**: 更新用户文档,说明如何使用统计分析功能。

**Steps**:
1. 创建或更新用户手册
2. 添加统计功能说明
3. 添加图表使用说明
4. 添加日期范围选择说明
5. 添加截图和示例
6. 添加常见问题解答

**Validation**:
- 文档清晰易懂
- 截图清晰
- 示例准确

**Dependencies**: Task 11

**Estimated Time**: 2 hours

---

## Parallelizable Work

以下任务可以并行执行:
- **Backend**: Task 1, Task 2 (顺序依赖)
- **Frontend**: Task 3, Task 4, Task 5, Task 10 (可并行)
- **Frontend**: Task 6, Task 7, Task 8, Task 9 (可并行,都依赖 Task 3, Task 5)
- **Testing**: Task 12, Task 13, Task 14 (可并行,依赖各自的功能任务)
- **Documentation**: Task 15, Task 16 (可并行)

## Critical Path

关键路径(最短完成时间):
1. Task 1 (2-3h) → Task 2 (1h) → Task 12 (3h) → Task 15 (1h)
2. Task 3 (1h) + Task 4 (0.5h) + Task 5 (1h) → Task 6 (2h) + Task 7 (3h) + Task 8 (3h) + Task 9 (4h) + Task 10 (2h) → Task 11 (4h) → Task 13 (4h) + Task 14 (3h) → Task 16 (2h)

**Total Estimated Time**: 25-30 hours

## Success Criteria

- [x] 所有后端API正常工作,响应时间 < 200ms (P95)
- [x] 所有前端组件正常渲染,页面加载时间 < 1s
- [x] 所有图表交互功能正常工作
- [x] 响应式布局在桌面、平板、移动端都正常
- [x] 单元测试覆盖率 > 70%
- [x] E2E 测试全部通过
- [x] 文档完整准确
- [x] 性能指标达标
- [x] 无明显的bug或性能问题
