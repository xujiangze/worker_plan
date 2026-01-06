# api-integration Specification

## Purpose
定义前端与后端 API 集成的规范，包括客户端配置、API 函数、类型定义、响应处理和错误处理。

## Requirements
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
**Then** 拦截器自动提取 `response.data.data` 并返回给调用方
**And** 拦截器统一处理 HTTP 错误
**And** 拦截器记录响应日志(开发环境)

### Requirement: 响应数据提取
前端响应拦截器 MUST 自动提取 `data` 字段从后端的统一响应结构中,确保 API 调用接收实际数据负载而不是包装对象。

#### Scenario: 创建计划 API 返回正确的数据结构
**Given** 后端返回格式为 `{code: 0, message: "success", data: {...}}` 的响应
**When** 前端调用 `planApi.createPlan(data)`
**Then** 响应应该是 `data` 字段中的 Plan 对象,而不是整个响应包装器
**And** store 可以成功调用 `plans.value.unshift(plan)` 而不会出错

#### Scenario: 所有 API 调用接收解包后的数据
**Given** 响应拦截器提取 `response.data.data`
**When** 进行任何 API 调用(getPlans, getPlan, updatePlan, deletePlan 等)
**Then** 响应应该是实际的数据负载
**And** store 可以处理数据而无需手动解包

#### Scenario: 错误响应被正确处理
**Given** 后端返回非零 code 的错误响应
**When** API 调用失败
**Then** 错误应该被错误拦截器捕获并处理
**And** 用户应该看到适当的错误消息

### Requirement: HTTP 错误处理
前端 MUST 在响应拦截器中统一处理 HTTP 错误。

#### Scenario: 处理 400 错误
**Given** API 返回 400 状态码
**When** 响应拦截器捕获错误
**Then** 系统显示错误提示:"请求参数错误"
**And** 错误提示类型为 error
**And** 错误提示持续 3 秒

#### Scenario: 处理 401 错误
**Given** API 返回 401 状态码
**When** 响应拦截器捕获错误
**Then** 系统显示错误提示:"未授权,请重新登录"
**And** 错误提示类型为 error
**And** 系统清除本地存储的 token
**And** 系统跳转到登录页面(如果存在)

#### Scenario: 处理 403 错误
**Given** API 返回 403 状态码
**When** 响应拦截器捕获错误
**Then** 系统显示错误提示:"无权限访问"
**And** 错误提示类型为 error

#### Scenario: 处理 404 错误
**Given** API 返回 404 状态码
**When** 响应拦截器捕获错误
**Then** 系统显示错误提示:"资源不存在"
**And** 错误提示类型为 error

#### Scenario: 处理 500 错误
**Given** API 返回 500 状态码
**When** 响应拦截器捕获错误
**Then** 系统显示错误提示:"服务器错误,请稍后重试"
**And** 错误提示类型为 error

#### Scenario: 处理网络错误
**Given** 网络连接失败
**When** 响应拦截器捕获错误
**Then** 系统显示错误提示:"网络错误,请检查网络连接"
**And** 错误提示类型为 error
**And** 错误提示持续 5 秒

#### Scenario: 处理超时错误
**Given** 请求超时
**When** 响应拦截器捕获错误
**Then** 系统显示错误提示:"请求超时,请稍后重试"
**And** 错误提示类型为 error

### Requirement: 表单验证错误处理
前端 MUST 使用 Element Plus 的表单验证规则处理表单验证错误。

#### Scenario: 显示必填字段错误
**Given** 用户提交表单
**When** 必填字段为空
**Then** 系统显示字段验证错误
**And** 错误提示显示在字段下方
**And** 错误提示文字为红色
**And** 错误提示内容:"此字段为必填项"

#### Scenario: 显示字段长度错误
**Given** 用户提交表单
**When** 字段长度超过限制
**Then** 系统显示字段验证错误
**And** 错误提示显示在字段下方
**And** 错误提示文字为红色
**And** 错误提示内容:"长度不能超过 X 个字符"

#### Scenario: 显示字段格式错误
**Given** 用户提交表单
**When** 字段格式不正确
**Then** 系统显示字段验证错误
**And** 错误提示显示在字段下方
**And** 错误提示文字为红色
**And** 错误提示内容:"格式不正确"

#### Scenario: 清除表单验证错误
**Given** 表单存在验证错误
**When** 用户修改字段内容
**Then** 系统清除该字段的验证错误
**And** 错误提示消失

### Requirement: 成功提示消息
前端 MUST 在操作成功时显示成功提示。

#### Scenario: 显示创建成功提示
**Given** 用户创建计划成功
**When** API 返回成功响应
**Then** 系统显示成功提示:"计划创建成功"
**And** 成功提示类型为 success
**And** 成功提示持续 3 秒

#### Scenario: 显示更新成功提示
**Given** 用户更新计划成功
**When** API 返回成功响应
**Then** 系统显示成功提示:"计划更新成功"
**And** 成功提示类型为 success
**And** 成功提示持续 3 秒

#### Scenario: 显示删除成功提示
**Given** 用户删除计划成功
**When** API 返回成功响应
**Then** 系统显示成功提示:"计划删除成功"
**And** 成功提示类型为 success
**And** 成功提示持续 3 秒

#### Scenario: 显示状态更新成功提示
**Given** 用户更新计划状态成功
**When** API 返回成功响应
**Then** 系统显示成功提示:"状态更新成功"
**And** 成功提示类型为 success
**And** 成功提示持续 3 秒

#### Scenario: 显示进度更新成功提示
**Given** 用户更新计划进度成功
**When** API 返回成功响应
**Then** 系统显示成功提示:"进度更新成功"
**And** 成功提示类型为 success
**And** 成功提示持续 3 秒

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

