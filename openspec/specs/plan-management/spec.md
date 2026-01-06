# plan-management Specification

## Purpose
TBD - created by archiving change implement-backend. Update Purpose after archive.
## Requirements
### Requirement: 创建工作计划
系统 MUST 允许用户创建新的工作计划,包含标题、描述、优先级和截止日期等信息。

#### Scenario: 创建基本工作计划
**Given** 用户已登录系统
**When** 用户发送 POST 请求到 `/api/plans`,请求体包含:
```json
{
  "title": "完成项目文档",
  "description": "编写项目的技术文档和用户手册",
  "priority": "High",
  "due_date": "2026-01-10T23:59:59Z"
}
```
**Then** 系统返回 201 Created 状态码
**And** 响应体包含创建的计划信息:
```json
{
  "id": 1,
  "title": "完成项目文档",
  "description": "编写项目的技术文档和用户手册",
  "priority": "High",
  "status": "Todo",
  "due_date": "2026-01-10T23:59:59Z",
  "progress": 0,
  "created_at": "2026-01-05T12:00:00Z",
  "updated_at": "2026-01-05T12:00:00Z"
}
```
**And** 计划的初始状态为 "Todo"
**And** 计划的初始进度为 0

#### Scenario: 创建计划时缺少必填字段
**Given** 用户已登录系统
**When** 用户发送 POST 请求到 `/api/plans`,请求体缺少必填字段 "title"
**Then** 系统返回 400 Bad Request 状态码
**And** 响应体包含错误信息:
```json
{
  "code": 400,
  "message": "validation error",
  "errors": [
    {
      "field": "title",
      "message": "title is required"
    }
  ]
}
```

#### Scenario: 创建计划时优先级无效
**Given** 用户已登录系统
**When** 用户发送 POST 请求到 `/api/plans`,请求体包含无效的优先级 "Urgent"
**Then** 系统返回 400 Bad Request 状态码
**And** 响应体包含错误信息,说明优先级必须是 High、Medium 或 Low

### Requirement: 查询工作计划列表
系统 MUST 允许用户查询工作计划列表,支持筛选、排序和分页。

#### Scenario: 查询所有工作计划
**Given** 系统中存在多个工作计划
**When** 用户发送 GET 请求到 `/api/plans`
**Then** 系统返回 200 OK 状态码
**And** 响应体包含分页的计划列表:
```json
{
  "total": 20,
  "page": 1,
  "page_size": 20,
  "data": [
    {
      "id": 1,
      "title": "完成项目文档",
      "description": "编写项目的技术文档和用户手册",
      "priority": "High",
      "status": "InProgress",
      "due_date": "2026-01-10T23:59:59Z",
      "progress": 50,
      "created_at": "2026-01-05T12:00:00Z",
      "updated_at": "2026-01-05T14:30:00Z"
    }
  ]
}
```

#### Scenario: 按状态筛选工作计划
**Given** 系统中存在多个不同状态的工作计划
**When** 用户发送 GET 请求到 `/api/plans?status=InProgress`
**Then** 系统返回 200 OK 状态码
**And** 响应体只包含状态为 "InProgress" 的计划

#### Scenario: 按优先级筛选工作计划
**Given** 系统中存在多个不同优先级的工作计划
**When** 用户发送 GET 请求到 `/api/plans?priority=High`
**Then** 系统返回 200 OK 状态码
**And** 响应体只包含优先级为 "High" 的计划

#### Scenario: 按日期范围筛选工作计划
**Given** 系统中存在多个不同截止日期的工作计划
**When** 用户发送 GET 请求到 `/api/plans?start_date=2026-01-01&end_date=2026-01-31`
**Then** 系统返回 200 OK 状态码
**And** 响应体只包含截止日期在指定范围内的计划

#### Scenario: 按创建时间排序
**Given** 系统中存在多个工作计划
**When** 用户发送 GET 请求到 `/api/plans?sort=created_at&order=desc`
**Then** 系统返回 200 OK 状态码
**And** 响应体中的计划按创建时间降序排列

#### Scenario: 分页查询工作计划
**Given** 系统中存在 100 个工作计划
**When** 用户发送 GET 请求到 `/api/plans?page=2&page_size=20`
**Then** 系统返回 200 OK 状态码
**And** 响应体包含第 2 页的 20 个计划
**And** total 字段显示总数为 100

### Requirement: 查询单个工作计划
系统 MUST 允许用户查询单个工作计划的详细信息。

#### Scenario: 查询存在的工作计划
**Given** 系统中存在 ID 为 1 的工作计划
**When** 用户发送 GET 请求到 `/api/plans/1`
**Then** 系统返回 200 OK 状态码
**And** 响应体包含该计划的完整信息

#### Scenario: 查询不存在的工作计划
**Given** 系统中不存在 ID 为 999 的工作计划
**When** 用户发送 GET 请求到 `/api/plans/999`
**Then** 系统返回 404 Not Found 状态码
**And** 响应体包含错误信息,说明计划不存在

### Requirement: 更新工作计划
系统 MUST 允许用户更新工作计划的信息,并记录变更历史。

#### Scenario: 更新工作计划的基本信息
**Given** 系统中存在 ID 为 1 的工作计划
**When** 用户发送 PUT 请求到 `/api/plans/1`,请求体包含:
```json
{
  "title": "完成项目文档(更新)",
  "description": "编写项目的技术文档、用户手册和API文档",
  "priority": "High",
  "due_date": "2026-01-12T23:59:59Z"
}
```
**Then** 系统返回 200 OK 状态码
**And** 响应体包含更新后的计划信息
**And** 系统记录变更历史,包含修改的字段、旧值和新值

#### Scenario: 更新不存在的工作计划
**Given** 系统中不存在 ID 为 999 的工作计划
**When** 用户发送 PUT 请求到 `/api/plans/999`
**Then** 系统返回 404 Not Found 状态码

#### Scenario: 更新计划时字段验证失败
**Given** 系统中存在 ID 为 1 的工作计划
**When** 用户发送 PUT 请求到 `/api/plans/1`,请求体包含无效的优先级
**Then** 系统返回 400 Bad Request 状态码
**And** 响应体包含验证错误信息

### Requirement: 删除工作计划
系统 MUST 允许用户删除工作计划,使用软删除方式保留数据。

#### Scenario: 删除存在的工作计划
**Given** 系统中存在 ID 为 1 的工作计划
**When** 用户发送 DELETE 请求到 `/api/plans/1`
**Then** 系统返回 204 No Content 状态码
**And** 计划被标记为已删除(deleted_at 字段不为空)
**And** 计划的历史记录仍然保留

#### Scenario: 删除不存在的工作计划
**Given** 系统中不存在 ID 为 999 的工作计划
**When** 用户发送 DELETE 请求到 `/api/plans/999`
**Then** 系统返回 404 Not Found 状态码

#### Scenario: 查询已删除的计划
**Given** ID 为 1 的工作计划已被删除
**When** 用户发送 GET 请求到 `/api/plans/1`
**Then** 系统返回 404 Not Found 状态码
**And** 已删除的计划不会出现在列表查询结果中

