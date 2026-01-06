# progress-tracking Specification

## Purpose
TBD - created by archiving change implement-backend. Update Purpose after archive.
## Requirements
### Requirement: 更新计划状态
系统 MUST 允许用户更新工作计划的状态,并记录状态变更历史。

#### Scenario: 将计划从待开始更新为进行中
**Given** 系统中存在 ID 为 1 的工作计划,状态为 "Todo"
**When** 用户发送 PATCH 请求到 `/api/plans/1/status`,请求体包含:
```json
{
  "status": "InProgress"
}
```
**Then** 系统返回 200 OK 状态码
**And** 响应体包含更新后的状态:
```json
{
  "id": 1,
  "status": "InProgress",
  "updated_at": "2026-01-05T15:00:00Z"
}
```
**And** 系统记录状态变更历史:
```json
{
  "field_name": "status",
  "old_value": "Todo",
  "new_value": "InProgress",
  "change_type": "Status",
  "changed_at": "2026-01-05T15:00:00Z"
}
```

#### Scenario: 将计划从进行中更新为已完成
**Given** 系统中存在 ID 为 1 的工作计划,状态为 "InProgress"
**When** 用户发送 PATCH 请求到 `/api/plans/1/status`,请求体包含:
```json
{
  "status": "Done"
}
```
**Then** 系统返回 200 OK 状态码
**And** 计划状态更新为 "Done"
**And** 系统记录状态变更历史

#### Scenario: 更新不存在计划的的状态
**Given** 系统中不存在 ID 为 999 的工作计划
**When** 用户发送 PATCH 请求到 `/api/plans/999/status`
**Then** 系统返回 404 Not Found 状态码

#### Scenario: 更新计划时状态无效
**Given** 系统中存在 ID 为 1 的工作计划
**When** 用户发送 PATCH 请求到 `/api/plans/1/status`,请求体包含无效状态 "Unknown"
**Then** 系统返回 400 Bad Request 状态码
**And** 响应体包含错误信息,说明状态必须是 Todo、InProgress、Done 或 Cancelled

### Requirement: 更新计划进度
系统 MUST 允许用户更新工作计划的完成百分比,并在进度达到 100% 时自动更新状态为已完成。

#### Scenario: 更新计划进度
**Given** 系统中存在 ID 为 1 的工作计划,当前进度为 0
**When** 用户发送 PATCH 请求到 `/api/plans/1/progress`,请求体包含:
```json
{
  "progress": 50
}
```
**Then** 系统返回 200 OK 状态码
**And** 响应体包含更新后的进度:
```json
{
  "id": 1,
  "progress": 50,
  "status": "InProgress",
  "updated_at": "2026-01-05T15:00:00Z"
}
```
**And** 系统记录进度变更历史:
```json
{
  "field_name": "progress",
  "old_value": "0",
  "new_value": "50",
  "change_type": "Progress",
  "changed_at": "2026-01-05T15:00:00Z"
}
```

#### Scenario: 进度达到 100% 时自动更新状态
**Given** 系统中存在 ID 为 1 的工作计划,当前状态为 "InProgress",进度为 90
**When** 用户发送 PATCH 请求到 `/api/plans/1/progress`,请求体包含:
```json
{
  "progress": 100
}
```
**Then** 系统返回 200 OK 状态码
**And** 计划进度更新为 100
**And** 计划状态自动更新为 "Done"
**And** 系统记录进度和状态的变更历史

#### Scenario: 更新不存在计划的进度
**Given** 系统中不存在 ID 为 999 的工作计划
**When** 用户发送 PATCH 请求到 `/api/plans/999/progress`
**Then** 系统返回 404 Not Found 状态码

#### Scenario: 更新进度时值超出范围
**Given** 系统中存在 ID 为 1 的工作计划
**When** 用户发送 PATCH 请求到 `/api/plans/1/progress`,请求体包含:
```json
{
  "progress": 150
}
```
**Then** 系统返回 400 Bad Request 状态码
**And** 响应体包含错误信息,说明进度必须在 0-100 之间

#### Scenario: 更新进度时值为负数
**Given** 系统中存在 ID 为 1 的工作计划
**When** 用户发送 PATCH 请求到 `/api/plans/1/progress`,请求体包含:
```json
{
  "progress": -10
}
```
**Then** 系统返回 400 Bad Request 状态码
**And** 响应体包含错误信息,说明进度必须在 0-100 之间

### Requirement: 状态转换验证
系统 MUST 验证状态转换的合法性,不允许非法的状态转换。

#### Scenario: 从已完成状态回退到进行中
**Given** 系统中存在 ID 为 1 的工作计划,状态为 "Done"
**When** 用户发送 PATCH 请求到 `/api/plans/1/status`,请求体包含:
```json
{
  "status": "InProgress"
}
```
**Then** 系统返回 200 OK 状态码
**And** 计划状态更新为 "InProgress"
**And** 系统记录状态变更历史

#### Scenario: 从已取消状态恢复为待开始
**Given** 系统中存在 ID 为 1 的工作计划,状态为 "Cancelled"
**When** 用户发送 PATCH 请求到 `/api/plans/1/status`,请求体包含:
```json
{
  "status": "Todo"
}
```
**Then** 系统返回 200 OK 状态码
**And** 计划状态更新为 "Todo"
**And** 系统记录状态变更历史

### Requirement: 进度与状态联动
系统 MUST 在进度更新时自动调整状态,确保状态与进度的一致性。

#### Scenario: 进度大于 0 时自动更新状态为进行中
**Given** 系统中存在 ID 为 1 的工作计划,状态为 "Todo",进度为 0
**When** 用户发送 PATCH 请求到 `/api/plans/1/progress`,请求体包含:
```json
{
  "progress": 10
}
```
**Then** 系统返回 200 OK 状态码
**And** 计划进度更新为 10
**And** 计划状态自动更新为 "InProgress"

#### Scenario: 进度小于 100 时状态保持为进行中
**Given** 系统中存在 ID 为 1 的工作计划,状态为 "Done",进度为 100
**When** 用户发送 PATCH 请求到 `/api/plans/1/progress`,请求体包含:
```json
{
  "progress": 80
}
```
**Then** 系统返回 200 OK 状态码
**And** 计划进度更新为 80
**And** 计划状态自动更新为 "InProgress"

