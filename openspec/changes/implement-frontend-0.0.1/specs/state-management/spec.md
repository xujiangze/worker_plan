# Spec: State Management

## ADDED Requirements

### Requirement: Plan Store
前端 MUST 使用 Pinia 创建 plan store,用于管理计划相关的状态。

#### Scenario: 初始化 plan store
**Given** 前端项目已配置 Pinia
**When** 创建 `stores/plan.ts` 文件
**Then** 定义 plan store
**And** 定义 state: plans, currentPlan, loading, error
**And** 定义 actions: fetchPlans, fetchPlan, createPlan, updatePlan, deletePlan
**And** 定义 getters: sortedPlans, filteredPlans

#### Scenario: 获取计划列表
**Given** 用户访问计划列表页面
**When** 调用 `fetchPlans()` action
**Then** 系统调用 API 获取计划列表
**And** 更新 `plans` state
**And** 设置 `loading` 为 false
**And** 如果失败,设置 `error` state

#### Scenario: 获取单个计划
**Given** 用户访问计划详情页面
**When** 调用 `fetchPlan(id)` action
**Then** 系统调用 API 获取计划详情
**And** 更新 `currentPlan` state
**And** 设置 `loading` 为 false
**And** 如果失败,设置 `error` state

#### Scenario: 创建计划
**Given** 用户提交创建计划表单
**When** 调用 `createPlan(data)` action
**Then** 系统调用 API 创建计划
**And** 将新计划添加到 `plans` state
**And** 设置 `loading` 为 false
**And** 如果失败,设置 `error` state

#### Scenario: 更新计划
**Given** 用户提交编辑计划表单
**When** 调用 `updatePlan(id, data)` action
**Then** 系统调用 API 更新计划
**And** 更新 `plans` state 中的对应计划
**And** 如果 `currentPlan` 存在,也更新 `currentPlan` state
**And** 设置 `loading` 为 false
**And** 如果失败,设置 `error` state

#### Scenario: 删除计划
**Given** 用户确认删除计划
**When** 调用 `deletePlan(id)` action
**Then** 系统调用 API 删除计划
**And** 从 `plans` state 中移除对应计划
**And** 如果 `currentPlan` 存在且 ID 匹配,清空 `currentPlan` state
**And** 设置 `loading` 为 false
**And** 如果失败,设置 `error` state

#### Scenario: 获取排序后的计划列表
**Given** 需要显示排序后的计划列表
**When** 访问 `sortedPlans` getter
**Then** 系统返回按创建时间降序排列的计划列表

### Requirement: UI Store
前端 MUST 使用 Pinia 创建 ui store,用于管理全局 UI 状态。

#### Scenario: 初始化 ui store
**Given** 前端项目已配置 Pinia
**When** 创建 `stores/ui.ts` 文件
**Then** 定义 ui store
**And** 定义 state: loading, notification, dialog
**And** 定义 actions: showLoading, hideLoading, showNotification, showDialog, hideDialog

#### Scenario: 显示加载状态
**Given** 需要显示加载动画
**When** 调用 `showLoading()` action
**Then** 设置 `loading` state 为 true
**And** 全局加载组件显示

#### Scenario: 隐藏加载状态
**Given** 加载操作完成
**When** 调用 `hideLoading()` action
**Then** 设置 `loading` state 为 false
**And** 全局加载组件隐藏

#### Scenario: 显示通知消息
**Given** 需要显示成功或错误提示
**When** 调用 `showNotification(message, type)` action
**Then** 设置 `notification` state
**And** 通知组件显示消息
**And** 通知类型为 success 或 error
**And** 3 秒后自动隐藏

#### Scenario: 显示对话框
**Given** 需要显示确认对话框
**When** 调用 `showDialog(config)` action
**Then** 设置 `dialog` state
**And** 对话框组件显示
**And** 对话框包含标题、内容、确认按钮、取消按钮

#### Scenario: 隐藏对话框
**Given** 用户点击取消或确认按钮
**When** 调用 `hideDialog()` action
**Then** 清空 `dialog` state
**And** 对话框组件隐藏
