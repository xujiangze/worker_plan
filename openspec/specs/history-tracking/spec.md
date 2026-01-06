# history-tracking Specification

## Purpose
TBD - created by archiving change implement-backend. Update Purpose after archive.
## Requirements
### Requirement: 记录计划变更历史
系统 MUST 在计划信息变更时自动记录历史,包括状态变更、进度变更和信息变更。

#### Scenario: 记录状态变更历史
**Given** 系统中存在 ID 为 1 的工作计划,状态为 "Todo"
**When** 用户更新计划状态为 "InProgress"
**Then** 系统在 PlanHistory 表中创建一条记录:
```json
{
  "id": 1,
  "plan_id": 1,
  "field_name": "status",
  "old_value": "Todo",
  "new_value": "InProgress",
  "change_type": "Status",
  "changed_at": "2026-01-05T15:00:00Z"
}
```
**And** changed_at 字段记录变更时间

#### Scenario: 记录进度变更历史
**Given** 系统中存在 ID 为 1 的工作计划,进度为 0
**When** 用户更新计划进度为 50
**Then** 系统在 PlanHistory 表中创建一条记录:
```json
{
  "id": 2,
  "plan_id": 1,
  "field_name": "progress",
  "old_value": "0",
  "new_value": "50",
  "change_type": "Progress",
  "changed_at": "2026-01-05T15:00:00Z"
}
```

#### Scenario: 记录信息变更历史
**Given** 系统中存在 ID 为 1 的工作计划,标题为 "完成项目文档"
**When** 用户更新计划标题为 "完成项目文档(更新)"
**Then** 系统在 PlanHistory 表中创建一条记录:
```json
{
  "id": 3,
  "plan_id": 1,
  "field_name": "title",
  "old_value": "完成项目文档",
  "new_value": "完成项目文档(更新)",
  "change_type": "Info",
  "changed_at": "2026-01-05T15:00:00Z"
}
```

#### Scenario: 记录多个字段变更历史
**Given** 系统中存在 ID 为 1 的工作计划
**When** 用户同时更新计划的标题、描述和优先级
**Then** 系统在 PlanHistory 表中为每个变更字段创建一条记录
**And** 每条记录的 changed_at 时间相同

### Requirement: 查询计划历史
系统 MUST 允许用户查询指定计划的所有历史变更记录。

#### Scenario: 查询计划的所有历史记录
**Given** 系统中存在 ID 为 1 的工作计划,有 5 条历史记录
**When** 用户发送 GET 请求到 `/api/plans/1/history`
**Then** 系统返回 200 OK 状态码
**And** 响应体包含所有历史记录:
```json
{
  "total": 5,
  "data": [
    {
      "id": 1,
      "field_name": "status",
      "old_value": "Todo",
      "new_value": "InProgress",
      "change_type": "Status",
      "changed_at": "2026-01-05T12:30:00Z"
    },
    {
      "id": 2,
      "field_name": "progress",
      "old_value": "0",
      "new_value": "50",
      "change_type": "Progress",
      "changed_at": "2026-01-05T14:00:00Z"
    },
    {
      "id": 3,
      "field_name": "title",
      "old_value": "完成项目文档",
      "new_value": "完成项目文档(更新)",
      "change_type": "Info",
      "changed_at": "2026-01-05T15:00:00Z"
    }
  ]
}
```
**And** 历史记录按 changed_at 降序排列(最新的在前)

#### Scenario: 查询不存在计划的历史记录
**Given** 系统中不存在 ID 为 999 的工作计划
**When** 用户发送 GET 请求到 `/api/plans/999/history`
**Then** 系统返回 404 Not Found 状态码
**And** 响应体包含错误信息,说明计划不存在

#### Scenario: 查询计划历史时无历史记录
**Given** 系统中存在 ID 为 1 的工作计划,但没有任何历史记录
**When** 用户发送 GET 请求到 `/api/plans/1/history`
**Then** 系统返回 200 OK 状态码
**And** 响应体显示:
```json
{
  "total": 0,
  "data": []
}
```

### Requirement: 按变更类型筛选历史记录
系统 MUST 允许用户按变更类型筛选历史记录。

#### Scenario: 筛选状态变更历史
**Given** 系统中存在 ID 为 1 的工作计划,有多种类型的历史记录
**When** 用户发送 GET 请求到 `/api/plans/1/history?change_type=Status`
**Then** 系统返回 200 OK 状态码
**And** 响应体只包含 change_type 为 "Status" 的历史记录

#### Scenario: 筛选进度变更历史
**Given** 系统中存在 ID 为 1 的工作计划,有多种类型的历史记录
**When** 用户发送 GET 请求到 `/api/plans/1/history?change_type=Progress`
**Then** 系统返回 200 OK 状态码
**And** 响应体只包含 change_type 为 "Progress" 的历史记录

#### Scenario: 筛选信息变更历史
**Given** 系统中存在 ID 为 1 的工作计划,有多种类型的历史记录
**When** 用户发送 GET 请求到 `/api/plans/1/history?change_type=Info`
**Then** 系统返回 200 OK 状态码
**And** 响应体只包含 change_type 为 "Info" 的历史记录

### Requirement: 按时间范围筛选历史记录
系统 MUST 允许用户按时间范围筛选历史记录。

#### Scenario: 筛选指定时间范围内的历史记录
**Given** 系统中存在 ID 为 1 的工作计划,历史记录分布在 2026-01-01 至 2026-01-31 之间
**When** 用户发送 GET 请求到 `/api/plans/1/history?start_date=2026-01-01&end_date=2026-01-15`
**Then** 系统返回 200 OK 状态码
**And** 响应体只包含 changed_at 在指定时间范围内的历史记录

#### Scenario: 筛选历史记录时日期格式无效
**Given** 系统中存在 ID 为 1 的工作计划
**When** 用户发送 GET 请求到 `/api/plans/1/history?start_date=invalid&end_date=2026-01-31`
**Then** 系统返回 400 Bad Request 状态码
**And** 响应体包含错误信息,说明日期格式无效

### Requirement: 历史记录分页查询
系统 MUST 支持历史记录的分页查询,避免返回大量数据。

#### Scenario: 分页查询历史记录
**Given** 系统中存在 ID 为 1 的工作计划,有 100 条历史记录
**When** 用户发送 GET 请求到 `/api/plans/1/history?page=1&page_size=20`
**Then** 系统返回 200 OK 状态码
**And** 响应体包含第 1 页的 20 条历史记录:
```json
{
  "total": 100,
  "page": 1,
  "page_size": 20,
  "data": []
}
```

### Requirement: 历史记录保留策略
系统 MUST 保留计划的历史记录,即使计划被删除。

#### Scenario: 删除计划后保留历史记录
**Given** 系统中存在 ID 为 1 的工作计划,有 5 条历史记录
**When** 用户删除该计划
**Then** 计划被标记为已删除(deleted_at 不为空)
**And** 计划的 5 条历史记录仍然保留在数据库中
**And** 历史记录不会被删除

### Requirement: 历史记录数据完整性
系统 MUST 确保历史记录的数据完整性,记录所有变更。

#### Scenario: 记录所有字段的变更
**Given** 系统中存在 ID 为 1 的工作计划
**When** 用户更新计划的标题、描述、优先级、截止日期
**Then** 系统为每个变更字段创建一条历史记录
**And** 每条记录包含正确的 old_value 和 new_value

#### Scenario: 记录空值变更
**Given** 系统中存在 ID 为 1 的工作计划,描述字段为空
**When** 用户更新计划描述为 "新的描述"
**Then** 系统创建一条历史记录:
```json
{
  "field_name": "description",
  "old_value": "",
  "new_value": "新的描述",
  "change_type": "Info"
}
```

#### Scenario: 记录清空字段变更
**Given** 系统中存在 ID 为 1 的工作计划,描述字段为 "旧的描述"
**When** 用户清空计划描述
**Then** 系统创建一条历史记录:
```json
{
  "field_name": "description",
  "old_value": "旧的描述",
  "new_value": "",
  "change_type": "Info"
}
```

### Requirement: 历史记录性能要求
系统 MUST 在合理时间内返回历史记录,确保用户体验。

#### Scenario: 大量历史记录的查询性能
**Given** 系统中存在 ID 为 1 的工作计划,有 1000 条历史记录
**When** 用户发送 GET 请求到 `/api/plans/1/history?page=1&page_size=20`
**Then** 系统在 200ms 内返回响应 (P95)
**And** 响应时间不超过 500ms (P99)

