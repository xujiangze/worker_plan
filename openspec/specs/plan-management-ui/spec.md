# plan-management-ui Specification

## Purpose
定义前端计划管理用户界面的规范，包括计划列表、创建/编辑表单、删除功能、状态管理、进度跟踪和数据显示。

## Requirements
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

### Requirement: 计划列表页面
前端 MUST 提供计划列表页面,显示所有工作计划。

#### Scenario: 显示计划列表
**Given** 用户访问计划列表页面 `/plans`
**When** 页面加载完成
**Then** 系统显示所有工作计划的卡片列表
**And** 每个计划卡片显示:
  - 计划标题
  - 计划描述(截断显示)
  - 状态标签(待开始/进行中/已完成/已取消)
  - 优先级标签(高/中/低)
  - 截止日期
  - 进度条(0-100%)
  - 创建时间
**And** 列表按创建时间降序排列

#### Scenario: 加载状态显示
**Given** 用户访问计划列表页面
**When** 数据正在加载
**Then** 系统显示加载动画或骨架屏
**And** 用户无法进行其他操作

#### Scenario: 空状态显示
**Given** 系统中不存在任何工作计划
**When** 用户访问计划列表页面
**Then** 系统显示空状态提示
**And** 提示信息:"暂无工作计划,点击下方按钮创建"
**And** 显示"创建计划"按钮

### Requirement: 计划数据显示
前端 MUST 正确显示计划数据,包括时间字段和可选截止日期。

#### Scenario: 后端返回计划列表时时间字段格式正确
- **WHEN** 后端 API 返回计划列表
- **THEN** `created_at`、`updated_at` 和 `due_date` 字段 SHALL 使用 ISO 8601 格式字符串(例如: "2024-01-06T10:30:00Z")
- **AND** `due_date` 字段 SHALL 在值为 NULL 时返回 null 或不包含该字段

#### Scenario: 前端接收并显示计划数据
- **WHEN** 前端接收到后端返回的计划数据
- **THEN** 前端 SHALL 能够正确解析所有时间字段
- **AND** 计划列表 SHALL 正确显示在页面上
- **AND** 时间信息 SHALL 按照本地化格式展示给用户

#### Scenario: 可选截止日期字段处理
- **WHEN** 用户创建计划时不提供 `due_date` 字段
- **THEN** 后端 SHALL 接受请求并创建计划
- **AND** 数据库中 `due_date` 字段 SHALL 为 NULL
- **AND** 返回的计划对象中 `due_date` SHALL 为 null 或不包含该字段

#### Scenario: 创建计划时提供截止日期
- **WHEN** 用户创建计划时提供有效的 `due_date` 字符串
- **THEN** 后端 SHALL 解析日期字符串并存储
- **AND** 返回的计划对象中 `due_date` SHALL 为 ISO 8601 格式字符串
- **AND** 前端 SHALL 正确显示截止日期信息

#### Scenario: Store 层记录接收到的数据
- **WHEN** store 的 fetchPlans 方法接收到 API 响应
- **THEN** 方法 SHALL 在控制台输出接收到的数据(开发环境)
- **AND** 日志 SHALL 包含 plans 数组的长度和第一个计划对象的关键字段
- **AND** 日志 SHALL 便于快速定位数据格式问题

### Requirement: 创建计划表单
前端 MUST 提供创建计划表单,允许用户输入计划信息并创建新计划。

#### Scenario: 打开创建计划表单
**Given** 用户在计划列表页面
**When** 用户点击"创建计划"按钮
**Then** 系统打开创建计划表单对话框或跳转到 `/plans/create` 页面
**And** 表单包含以下字段:
  - 标题(必填,文本输入框)
  - 描述(选填,多行文本框)
  - 优先级(必填,下拉选择:高/中/低)
  - 截止日期(选填,日期时间选择器)
**And** 表单显示"保存"和"取消"按钮

#### Scenario: 创建计划成功
**Given** 用户在创建计划表单
**When** 用户填写表单:
  - 标题:"完成项目文档"
  - 描述:"编写项目的技术文档和用户手册"
  - 优先级:"高"
  - 截止日期:"2026-01-10"
**And** 用户点击"保存"按钮
**Then** 系统验证表单数据
**And** 系统发送 POST 请求到 `/api/plans`
**And** 系统显示成功提示:"计划创建成功"
**And** 系统关闭表单或跳转到计划列表页面
**And** 新创建的计划显示在列表中

#### Scenario: 创建计划时缺少必填字段
**Given** 用户在创建计划表单
**When** 用户只填写描述,不填写标题
**And** 用户点击"保存"按钮
**Then** 系统显示表单验证错误
**And** 标题字段显示错误提示:"标题不能为空"
**And** 表单不会提交

#### Scenario: 创建计划时标题过长
**Given** 用户在创建计划表单
**When** 用户输入超过 255 个字符的标题
**And** 用户点击"保存"按钮
**Then** 系统显示表单验证错误
**And** 标题字段显示错误提示:"标题不能超过 255 个字符"

#### Scenario: 取消创建计划
**Given** 用户在创建计划表单
**When** 用户点击"取消"按钮
**Then** 系统关闭表单或跳转到计划列表页面
**And** 表单数据不会保存

### Requirement: 编辑计划表单
前端 MUST 提供编辑计划表单,允许用户修改计划信息。

#### Scenario: 打开编辑计划表单
**Given** 用户在计划列表页面
**When** 用户点击某个计划的"编辑"按钮
**Then** 系统打开编辑计划表单对话框或跳转到 `/plans/:id/edit` 页面
**And** 表单预填充当前计划的信息:
  - 标题:"完成项目文档"
  - 描述:"编写项目的技术文档和用户手册"
  - 优先级:"高"
  - 截止日期:"2026-01-10"
**And** 表单显示"保存"和"取消"按钮

#### Scenario: 编辑计划成功
**Given** 用户在编辑计划表单
**When** 用户修改标题为"完成项目文档(更新)"
**And** 用户点击"保存"按钮
**Then** 系统验证表单数据
**And** 系统发送 PUT 请求到 `/api/plans/:id`
**And** 系统显示成功提示:"计划更新成功"
**And** 系统关闭表单或跳转到计划列表页面
**And** 计划信息已更新

#### Scenario: 编辑计划时字段验证失败
**Given** 用户在编辑计划表单
**When** 用户清空标题字段
**And** 用户点击"保存"按钮
**Then** 系统显示表单验证错误
**And** 标题字段显示错误提示:"标题不能为空"

#### Scenario: 取消编辑计划
**Given** 用户在编辑计划表单
**When** 用户点击"取消"按钮
**Then** 系统关闭表单或跳转到计划列表页面
**And** 计划信息不会更新

### Requirement: 删除计划
前端 MUST 提供删除计划功能,并在删除前显示确认对话框。

#### Scenario: 删除计划
**Given** 用户在计划列表页面
**When** 用户点击某个计划的"删除"按钮
**Then** 系统显示确认对话框
**And** 对话框显示提示信息:"确定要删除此计划吗?此操作不可恢复。"
**And** 对话框显示"确定"和"取消"按钮
**When** 用户点击"确定"按钮
**Then** 系统发送 DELETE 请求到 `/api/plans/:id`
**And** 系统显示成功提示:"计划删除成功"
**And** 计划从列表中移除

#### Scenario: 取消删除计划
**Given** 用户在计划列表页面
**When** 用户点击某个计划的"删除"按钮
**And** 系统显示确认对话框
**And** 用户点击"取消"按钮
**Then** 系统关闭确认对话框
**And** 计划不会被删除

#### Scenario: 删除计划失败
**Given** 用户在计划列表页面
**When** 用户点击某个计划的"删除"按钮
**And** 用户点击"确定"按钮
**And** 删除请求失败(网络错误或服务器错误)
**Then** 系统显示错误提示:"删除失败,请稍后重试"
**And** 计划不会被删除

