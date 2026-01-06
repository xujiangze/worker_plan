# plan-display Specification

## Purpose
TBD - created by archiving change fix-plan-display-issue. Update Purpose after archive.
## Requirements
### Requirement: 时间字段序列化格式
系统 SHALL 确保所有时间字段在 JSON 序列化时使用 ISO 8601 格式字符串,以便前端正确解析和显示。

#### Scenario: 后端返回计划列表时时间字段格式正确
- **WHEN** 后端 API 返回计划列表
- **THEN** `created_at`、`updated_at` 和 `due_date` 字段 SHALL 使用 ISO 8601 格式字符串(例如: "2024-01-06T10:30:00Z")
- **AND** `due_date` 字段 SHALL 在值为 NULL 时返回 null 或不包含该字段

#### Scenario: 前端接收并显示计划数据
- **WHEN** 前端接收到后端返回的计划数据
- **THEN** 前端 SHALL 能够正确解析所有时间字段
- **AND** 计划列表 SHALL 正确显示在页面上
- **AND** 时间信息 SHALL 按照本地化格式展示给用户

### Requirement: API 响应数据结构一致性
系统 SHALL 确保所有 API 响应遵循统一的响应结构,包含 `code`、`message` 和 `data` 字段。

#### Scenario: 获取计划列表 API 响应
- **WHEN** 客户端调用 GET /api/plans 接口
- **THEN** 响应 SHALL 包含 `code: 0` 表示成功
- **AND** 响应 SHALL 包含 `message: "success"`
- **AND** 响应的 `data` 字段 SHALL 包含 `items`、`total`、`page`、`page_size` 和 `total_pages`
- **AND** `items` SHALL 是计划对象的数组

#### Scenario: 响应拦截器正确提取数据
- **WHEN** 前端响应拦截器接收到 API 响应
- **THEN** 拦截器 SHALL 提取 `response.data.data` 并返回给调用方
- **AND** 调用方 SHALL 直接获得 `data` 字段的内容,无需再次解包

### Requirement: 可选截止日期字段处理
系统 SHALL 正确处理可选的截止日期字段,支持 NULL 值和有效日期值。

#### Scenario: 创建计划时不提供截止日期
- **WHEN** 用户创建计划时不提供 `due_date` 字段
- **THEN** 后端 SHALL 接受请求并创建计划
- **AND** 数据库中 `due_date` 字段 SHALL 为 NULL
- **AND** 返回的计划对象中 `due_date` SHALL 为 null 或不包含该字段

#### Scenario: 创建计划时提供截止日期
- **WHEN** 用户创建计划时提供有效的 `due_date` 字符串
- **THEN** 后端 SHALL 解析日期字符串并存储
- **AND** 返回的计划对象中 `due_date` SHALL 为 ISO 8601 格式字符串
- **AND** 前端 SHALL 正确显示截止日期信息

### Requirement: 数据传输调试支持
系统 SHALL 提供适当的日志记录机制,便于调试数据传输问题。

#### Scenario: Store 层记录接收到的数据
- **WHEN** store 的 fetchPlans 方法接收到 API 响应
- **THEN** 方法 SHALL 在控制台输出接收到的数据(开发环境)
- **AND** 日志 SHALL 包含 plans 数组的长度和第一个计划对象的关键字段
- **AND** 日志 SHALL 便于快速定位数据格式问题

#### Scenario: API 层记录响应数据
- **WHEN** API 调用成功或失败
- **THEN** 响应拦截器 SHALL 记录关键信息(开发环境)
- **AND** 错误日志 SHALL 包含错误状态码和错误消息
- **AND** 成功日志 SHALL 包含响应数据的摘要信息

