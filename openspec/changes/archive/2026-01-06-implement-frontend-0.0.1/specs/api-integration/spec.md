# Spec: API Integration

## ADDED Requirements

### Requirement: Axios 客户端配置
前端 MUST 配置 Axios 客户端,用于与后端 API 通信。

#### Scenario: 配置基础 Axios 实例
**Given** 前端项目已初始化
**When** 开发者创建 `utils/request.ts` 文件
**Then** 系统创建 Axios 实例
**And** 配置基础 URL(从环境变量读取)
**And** 配置超时时间(默认 10 秒)

#### Scenario: 配置请求拦截器
**Given** Axios 实例已创建
**When** 配置请求拦截器
**Then** 拦截器自动设置 `Content-Type: application/json`
**And** 拦截器自动添加认证 token(如果存在)
**And** 拦截器记录请求日志(开发环境)

#### Scenario: 配置响应拦截器
**Given** Axios 实例已创建
**When** 配置响应拦截器
**Then** 拦截器自动解析响应数据
**And** 拦截器统一处理 HTTP 错误
**And** 拦截器记录响应日志(开发环境)

### Requirement: 计划管理 API
前端 MUST 提供计划管理 API 函数,用于调用后端计划管理接口。

#### Scenario: 获取计划列表
**Given** 用户需要查看所有计划
**When** 调用 `getPlans()` 函数
**Then** 系统发送 GET 请求到 `/api/plans`
**And** 系统返回计划列表数据
**And** 数据类型符合 TypeScript 定义

#### Scenario: 获取单个计划
**Given** 用户需要查看某个计划的详情
**When** 调用 `getPlan(id)` 函数
**Then** 系统发送 GET 请求到 `/api/plans/:id`
**And** 系统返回计划详情数据

#### Scenario: 创建计划
**Given** 用户需要创建新计划
**When** 调用 `createPlan(data)` 函数
**Then** 系统发送 POST 请求到 `/api/plans`
**And** 请求体包含计划数据
**And** 系统返回创建的计划数据

#### Scenario: 更新计划
**Given** 用户需要更新计划信息
**When** 调用 `updatePlan(id, data)` 函数
**Then** 系统发送 PUT 请求到 `/api/plans/:id`
**And** 请求体包含更新数据
**And** 系统返回更新后的计划数据

#### Scenario: 删除计划
**Given** 用户需要删除计划
**When** 调用 `deletePlan(id)` 函数
**Then** 系统发送 DELETE 请求到 `/api/plans/:id`
**And** 系统返回成功响应

### Requirement: 进度管理 API
前端 MUST 提供进度管理 API 函数,用于调用后端进度管理接口。

#### Scenario: 更新计划状态
**Given** 用户需要更新计划状态
**When** 调用 `updatePlanStatus(id, status)` 函数
**Then** 系统发送 PATCH 请求到 `/api/plans/:id/status`
**And** 请求体包含状态数据
**And** 系统返回更新后的计划数据

#### Scenario: 更新计划进度
**Given** 用户需要更新计划进度
**When** 调用 `updatePlanProgress(id, progress)` 函数
**Then** 系统发送 PATCH 请求到 `/api/plans/:id/progress`
**And** 请求体包含进度数据
**And** 系统返回更新后的计划数据

### Requirement: API 类型定义
前端 MUST 使用 TypeScript 定义 API 请求和响应类型。

#### Scenario: 定义计划类型
**Given** 需要定义计划数据结构
**When** 创建 `types/plan.ts` 文件
**Then** 定义 `Plan` 接口
**And** 包含字段: id, title, description, status, priority, progress, deadline, createdAt, updatedAt

#### Scenario: 定义 API 响应类型
**Given** 需要定义 API 响应结构
**When** 创建 `types/api.ts` 文件
**Then** 定义 `ApiResponse<T>` 接口
**And** 包含字段: success, data, message, error

#### Scenario: 定义 API 请求类型
**Given** 需要定义 API 请求结构
**When** 创建 `types/api.ts` 文件
**Then** 定义 `CreatePlanRequest` 接口
**And** 定义 `UpdatePlanRequest` 接口
**And** 定义 `UpdateStatusRequest` 接口
**And** 定义 `UpdateProgressRequest` 接口
