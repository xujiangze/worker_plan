# statistics Specification Delta

## MODIFIED Requirements

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
**And** daily_trend 按日期升序排列

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
**And** completion_rate 为 0
**And** daily_trend 为空数组

#### Scenario: 查询时间统计时部分计划在时间范围内
**Given** 系统中存在 20 个工作计划,其中 10 个在 2026-01-01 至 2026-01-31 之间创建,5 个在此期间完成
**When** 用户发送 GET 请求到 `/api/stats/by-time?start_date=2026-01-01&end_date=2026-01-31`
**Then** 系统返回 200 OK 状态码
**And** created_count 为 10
**And** completed_count 为 5
**And** completion_rate 为 50
**And** daily_trend 包含该时间段内每天的创建和完成数据

#### Scenario: 查询时间统计时包含已删除的计划
**Given** 系统中存在 20 个工作计划,其中 2 个已被软删除
**When** 用户发送 GET 请求到 `/api/stats/by-time?start_date=2026-01-01&end_date=2026-01-31`
**Then** 系统返回 200 OK 状态码
**And** 统计数据只包含未删除的计划
**And** created_count 和 completed_count 不包含已删除的计划

#### Scenario: 查询时间统计时时间范围跨年
**Given** 系统中存在多个工作计划,创建和完成时间分布在 2025-12-25 至 2026-01-10 之间
**When** 用户发送 GET 请求到 `/api/stats/by-time?start_date=2025-12-25&end_date=2026-01-10`
**Then** 系统返回 200 OK 状态码
**And** 响应体包含跨年时间段内的统计信息
**And** daily_trend 包含从 2025-12-25 到 2026-01-10 的每日数据

#### Scenario: 查询时间统计时时间范围过大
**Given** 系统中存在多个工作计划
**When** 用户发送 GET 请求到 `/api/stats/by-time?start_date=2020-01-01&end_date=2026-01-31`
**Then** 系统返回 200 OK 状态码
**And** 系统在 200ms 内返回响应 (P95)
**And** daily_trend 包含所有日期的数据
**And** 响应时间不超过 500ms (P99)
