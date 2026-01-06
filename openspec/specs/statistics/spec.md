# statistics Specification

## Purpose
TBD - created by archiving change implement-backend. Update Purpose after archive.
## Requirements
### Requirement: 按状态统计
系统 MUST 提供按状态统计工作计划的功能,返回各状态计划的数量和占比。

#### Scenario: 查询所有状态统计
**Given** 系统中存在 20 个工作计划,分布如下:
  - 待开始: 5 个
  - 进行中: 8 个
  - 已完成: 6 个
  - 已取消: 1 个
**When** 用户发送 GET 请求到 `/api/stats/by-status`
**Then** 系统返回 200 OK 状态码
**And** 响应体包含各状态的统计信息:
```json
{
  "Todo": {
    "count": 5,
    "percentage": 25
  },
  "InProgress": {
    "count": 8,
    "percentage": 40
  },
  "Done": {
    "count": 6,
    "percentage": 30
  },
  "Cancelled": {
    "count": 1,
    "percentage": 5
  }
}
```
**And** 百分比总和为 100%

#### Scenario: 查询状态统计时系统为空
**Given** 系统中不存在任何工作计划
**When** 用户发送 GET 请求到 `/api/stats/by-status`
**Then** 系统返回 200 OK 状态码
**And** 响应体包含所有状态,但数量和百分比均为 0:
```json
{
  "Todo": {
    "count": 0,
    "percentage": 0
  },
  "InProgress": {
    "count": 0,
    "percentage": 0
  },
  "Done": {
    "count": 0,
    "percentage": 0
  },
  "Cancelled": {
    "count": 0,
    "percentage": 0
  }
}
```

### Requirement: 按优先级统计
系统 MUST 提供按优先级统计工作计划的功能,返回各优先级计划的数量和占比。

#### Scenario: 查询所有优先级统计
**Given** 系统中存在 20 个工作计划,分布如下:
  - 高优先级: 3 个
  - 中优先级: 10 个
  - 低优先级: 7 个
**When** 用户发送 GET 请求到 `/api/stats/by-priority`
**Then** 系统返回 200 OK 状态码
**And** 响应体包含各优先级的统计信息:
```json
{
  "High": {
    "count": 3,
    "percentage": 15
  },
  "Medium": {
    "count": 10,
    "percentage": 50
  },
  "Low": {
    "count": 7,
    "percentage": 35
  }
}
```
**And** 百分比总和为 100%

#### Scenario: 查询优先级统计时系统为空
**Given** 系统中不存在任何工作计划
**When** 用户发送 GET 请求到 `/api/stats/by-priority`
**Then** 系统返回 200 OK 状态码
**And** 响应体包含所有优先级,但数量和百分比均为 0

### Requirement: 按时间统计
系统 MUST 提供按时间范围统计工作计划的功能,返回指定时间范围内的创建和完成数量,以及每日趋势。

#### Scenario: 查询指定时间范围的统计
**Given** 系统中存在多个工作计划,创建和完成时间分布在 2026-01-01 至 2026-01-31 之间
**When** 用户发送 GET 请求到 `/api/stats/by-time?start_date=2026-01-01&end_date=2026-01-31`
**Then** 系统返回 200 OK 状态码
**And** 响应体包含时间范围内的统计信息:
```json
{
  "created_count": 15,
  "completed_count": 8,
  "completion_rate": 53.33,
  "daily_trend": [
    {
      "date": "2026-01-01",
      "created": 2,
      "completed": 0
    },
    {
      "date": "2026-01-02",
      "created": 1,
      "completed": 1
    },
    {
      "date": "2026-01-03",
      "created": 3,
      "completed": 2
    }
  ]
}
```
**And** completion_rate 计算公式为: (completed_count / created_count) × 100
**And** daily_trend 包含时间范围内每一天的创建和完成数量

#### Scenario: 查询时间统计时缺少日期参数
**Given** 系统中存在多个工作计划
**When** 用户发送 GET 请求到 `/api/stats/by-time` (不包含日期参数)
**Then** 系统返回 400 Bad Request 状态码
**And** 响应体包含错误信息,说明必须提供 start_date 和 end_date 参数

#### Scenario: 查询时间统计时日期格式无效
**Given** 系统中存在多个工作计划
**When** 用户发送 GET 请求到 `/api/stats/by-time?start_date=invalid&end_date=2026-01-31`
**Then** 系统返回 400 Bad Request 状态码
**And** 响应体包含错误信息,说明日期格式无效

#### Scenario: 查询时间统计时开始日期晚于结束日期
**Given** 系统中存在多个工作计划
**When** 用户发送 GET 请求到 `/api/stats/by-time?start_date=2026-01-31&end_date=2026-01-01`
**Then** 系统返回 400 Bad Request 状态码
**And** 响应体包含错误信息,说明开始日期不能晚于结束日期

#### Scenario: 查询时间统计时时间范围内无数据
**Given** 系统中存在多个工作计划,但都在 2026-01-01 之前创建
**When** 用户发送 GET 请求到 `/api/stats/by-time?start_date=2026-02-01&end_date=2026-02-28`
**Then** 系统返回 200 OK 状态码
**And** 响应体显示 created_count 和 completed_count 均为 0
**And** daily_trend 为空数组

### Requirement: 查看完成率
系统 MUST 计算并显示工作计划的总体完成率。

#### Scenario: 查询总体完成率
**Given** 系统中存在 20 个工作计划,其中 12 个已完成
**When** 用户发送 GET 请求到 `/api/stats/completion-rate`
**Then** 系统返回 200 OK 状态码
**And** 响应体包含完成率信息:
```json
{
  "total_plans": 20,
  "completed_plans": 12,
  "completion_rate": 60
}
```
**And** completion_rate 计算公式为: (completed_plans / total_plans) × 100

#### Scenario: 查询完成率时系统为空
**Given** 系统中不存在任何工作计划
**When** 用户发送 GET 请求到 `/api/stats/completion-rate`
**Then** 系统返回 200 OK 状态码
**And** 响应体显示:
```json
{
  "total_plans": 0,
  "completed_plans": 0,
  "completion_rate": 0
}
```

#### Scenario: 查询完成率时所有计划均未完成
**Given** 系统中存在 10 个工作计划,但均未完成
**When** 用户发送 GET 请求到 `/api/stats/completion-rate`
**Then** 系统返回 200 OK 状态码
**And** 响应体显示完成率为 0

### Requirement: 统计数据排除已删除计划
系统 MUST 在计算统计数据时排除已删除的计划。

#### Scenario: 统计数据不包含已删除的计划
**Given** 系统中存在 20 个工作计划,其中 2 个已被软删除
**When** 用户发送 GET 请求到 `/api/stats/completion-rate`
**Then** 系统返回 200 OK 状态码
**And** total_plans 为 18 (排除已删除的 2 个)
**And** 统计数据只包含未删除的计划

### Requirement: 统计数据性能要求
系统 MUST 在合理时间内返回统计数据,确保用户体验。

#### Scenario: 大数据量下的统计性能
**Given** 系统中存在 10,000 个工作计划
**When** 用户发送 GET 请求到 `/api/stats/by-status`
**Then** 系统在 200ms 内返回响应 (P95)
**And** 响应时间不超过 500ms (P99)

