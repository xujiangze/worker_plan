# common-components Specification

## Purpose
TBD - created by archiving change implement-frontend-0.0.1. Update Purpose after archive.
## Requirements
### Requirement: StatusBadge 组件
前端 MUST 提供状态标签组件,用于显示计划状态。

#### Scenario: 显示待开始状态
**Given** 计划状态为 "pending"
**When** 使用 `<StatusBadge status="pending" />` 组件
**Then** 组件显示"待开始"标签
**And** 标签颜色为灰色
**And** 标签样式为圆角矩形

#### Scenario: 显示进行中状态
**Given** 计划状态为 "in_progress"
**When** 使用 `<StatusBadge status="in_progress" />` 组件
**Then** 组件显示"进行中"标签
**And** 标签颜色为蓝色
**And** 标签样式为圆角矩形

#### Scenario: 显示已完成状态
**Given** 计划状态为 "completed"
**When** 使用 `<StatusBadge status="completed" />` 组件
**Then** 组件显示"已完成"标签
**And** 标签颜色为绿色
**And** 标签样式为圆角矩形

#### Scenario: 显示已取消状态
**Given** 计划状态为 "cancelled"
**When** 使用 `<StatusBadge status="cancelled" />` 组件
**Then** 组件显示"已取消"标签
**And** 标签颜色为红色
**And** 标签样式为圆角矩形

#### Scenario: 显示未知状态
**Given** 计划状态为未知值
**When** 使用 `<StatusBadge status="unknown" />` 组件
**Then** 组件显示"未知"标签
**And** 标签颜色为灰色
**And** 标签样式为圆角矩形

### Requirement: PriorityBadge 组件
前端 MUST 提供优先级标签组件,用于显示计划优先级。

#### Scenario: 显示高优先级
**Given** 计划优先级为 "high"
**When** 使用 `<PriorityBadge priority="high" />` 组件
**Then** 组件显示"高"标签
**And** 标签颜色为红色
**And** 标签样式为圆角矩形

#### Scenario: 显示中优先级
**Given** 计划优先级为 "medium"
**When** 使用 `<PriorityBadge priority="medium" />` 组件
**Then** 组件显示"中"标签
**And** 标签颜色为橙色
**And** 标签样式为圆角矩形

#### Scenario: 显示低优先级
**Given** 计划优先级为 "low"
**When** 使用 `<PriorityBadge priority="low" />` 组件
**Then** 组件显示"低"标签
**And** 标签颜色为绿色
**And** 标签样式为圆角矩形

#### Scenario: 显示未知优先级
**Given** 计划优先级为未知值
**When** 使用 `<PriorityBadge priority="unknown" />` 组件
**Then** 组件显示"未知"标签
**And** 标签颜色为灰色
**And** 标签样式为圆角矩形

### Requirement: Loading 组件
前端 MUST 提供加载动画组件,用于显示加载状态。

#### Scenario: 显示加载动画
**Given** 数据正在加载
**When** 使用 `<Loading />` 组件
**Then** 组件显示旋转的加载图标
**And** 加载图标居中显示
**And** 加载图标颜色为主题色

#### Scenario: 显示带文字的加载动画
**Given** 数据正在加载
**When** 使用 `<Loading text="加载中..." />` 组件
**Then** 组件显示旋转的加载图标
**And** 组件显示"加载中..."文字
**And** 文字位于加载图标下方
**And** 文字颜色为灰色

#### Scenario: 显示全屏加载动画
**Given** 页面正在加载
**When** 使用 `<Loading fullscreen />` 组件
**Then** 组件覆盖整个屏幕
**And** 组件背景为半透明白色
**And** 加载图标居中显示
**And** 用户无法进行其他操作

