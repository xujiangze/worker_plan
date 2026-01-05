# 工作计划管理系统 API 文档

## 基础信息

- Base URL: `http://localhost:8080/api`
- Content-Type: `application/json`

## 统一响应格式

### 成功响应
```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

### 错误响应
```json
{
  "code": 400,
  "message": "error message",
  "errors": []
}
```

### HTTP 状态码
- `200 OK`: 请求成功
- `201 Created`: 创建成功
- `204 No Content`: 删除成功
- `400 Bad Request`: 请求参数错误
- `404 Not Found`: 资源不存在
- `409 Conflict`: 资源冲突
- `500 Internal Server Error`: 服务器内部错误

## API 接口

### 1. 计划管理

#### 1.1 创建计划
- **URL**: `POST /api/plans`
- **描述**: 创建新的工作计划
- **请求体**:
```json
{
  "title": "完成项目文档",
  "description": "编写项目的技术文档和用户手册",
  "priority": "High",
  "due_date": "2024-12-31T23:59:59Z"
}
```
- **响应**: `201 Created`
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "title": "完成项目文档",
    "description": "编写项目的技术文档和用户手册",
    "priority": "High",
    "status": "Todo",
    "due_date": "2024-12-31T23:59:59Z",
    "progress": 0,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 1.2 获取计划列表
- **URL**: `GET /api/plans`
- **描述**: 获取工作计划列表,支持分页和筛选
- **查询参数**:
  - `page`: 页码(默认: 1)
  - `page_size`: 每页数量(默认: 20)
  - `status`: 状态筛选(Todo, InProgress, Done, Cancelled)
  - `priority`: 优先级筛选(High, Medium, Low)
- **响应**: `200 OK`
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 100,
    "page": 1,
    "page_size": 20,
    "data": [
      {
        "id": 1,
        "title": "完成项目文档",
        "description": "编写项目的技术文档和用户手册",
        "priority": "High",
        "status": "Todo",
        "due_date": "2024-12-31T23:59:59Z",
        "progress": 0,
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ]
  }
}
```

#### 1.3 获取单个计划
- **URL**: `GET /api/plans/{id}`
- **描述**: 根据 ID 获取工作计划详情
- **路径参数**:
  - `id`: 计划 ID
- **响应**: `200 OK`
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "title": "完成项目文档",
    "description": "编写项目的技术文档和用户手册",
    "priority": "High",
    "status": "Todo",
    "due_date": "2024-12-31T23:59:59Z",
    "progress": 0,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### 1.4 更新计划
- **URL**: `PUT /api/plans/{id}`
- **描述**: 更新工作计划信息
- **路径参数**:
  - `id`: 计划 ID
- **请求体**:
```json
{
  "title": "完成项目文档(更新)",
  "description": "编写项目的技术文档和用户手册",
  "priority": "Medium",
  "due_date": "2024-12-31T23:59:59Z"
}
```
- **响应**: `200 OK`
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "title": "完成项目文档(更新)",
    "description": "编写项目的技术文档和用户手册",
    "priority": "Medium",
    "status": "Todo",
    "due_date": "2024-12-31T23:59:59Z",
    "progress": 0,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
}
```

#### 1.5 删除计划
- **URL**: `DELETE /api/plans/{id}`
- **描述**: 删除工作计划(软删除)
- **路径参数**:
  - `id`: 计划 ID
- **响应**: `204 No Content`

### 2. 进度管理

#### 2.1 更新状态
- **URL**: `PATCH /api/plans/{id}/status`
- **描述**: 更新工作计划的状态
- **路径参数**:
  - `id`: 计划 ID
- **请求体**:
```json
{
  "status": "InProgress"
}
```
- **响应**: `200 OK`
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "title": "完成项目文档",
    "description": "编写项目的技术文档和用户手册",
    "priority": "High",
    "status": "InProgress",
    "due_date": "2024-12-31T23:59:59Z",
    "progress": 0,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
}
```

#### 2.2 更新进度
- **URL**: `PATCH /api/plans/{id}/progress`
- **描述**: 更新工作计划的进度
- **路径参数**:
  - `id`: 计划 ID
- **请求体**:
```json
{
  "progress": 50
}
```
- **响应**: `200 OK`
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "title": "完成项目文档",
    "description": "编写项目的技术文档和用户手册",
    "priority": "High",
    "status": "InProgress",
    "due_date": "2024-12-31T23:59:59Z",
    "progress": 50,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
}
```

### 3. 统计分析

#### 3.1 按状态统计
- **URL**: `GET /api/stats/by-status`
- **描述**: 统计各状态计划的数量和占比
- **响应**: `200 OK`
```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "status": "Todo",
      "count": 10,
      "percent": 25.0
    },
    {
      "status": "InProgress",
      "count": 15,
      "percent": 37.5
    },
    {
      "status": "Done",
      "count": 10,
      "percent": 25.0
    },
    {
      "status": "Cancelled",
      "count": 5,
      "percent": 12.5
    }
  ]
}
```

#### 3.2 按优先级统计
- **URL**: `GET /api/stats/by-priority`
- **描述**: 统计各优先级计划的数量和占比
- **响应**: `200 OK`
```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "priority": "High",
      "count": 10,
      "percent": 25.0
    },
    {
      "priority": "Medium",
      "count": 20,
      "percent": 50.0
    },
    {
      "priority": "Low",
      "count": 10,
      "percent": 25.0
    }
  ]
}
```

#### 3.3 按时间统计
- **URL**: `GET /api/stats/by-time`
- **描述**: 统计指定时间范围内的计划数量
- **查询参数**:
  - `start_date`: 开始日期
  - `end_date`: 结束日期
- **响应**: `200 OK`
```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "date": "2024-01-01",
      "count": 5
    },
    {
      "date": "2024-01-02",
      "count": 3
    }
  ]
}
```

#### 3.4 获取完成率
- **URL**: `GET /api/stats/completion-rate`
- **描述**: 获取计划完成率统计
- **响应**: `200 OK`
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total_plans": 40,
    "completed_plans": 10,
    "completion_rate": 25.0
  }
}
```

### 4. 历史记录

#### 4.1 获取历史记录
- **URL**: `GET /api/plans/{id}/history`
- **描述**: 获取指定计划的历史变更记录
- **路径参数**:
  - `id`: 计划 ID
- **查询参数**:
  - `page`: 页码(默认: 1)
  - `page_size`: 每页数量(默认: 20)
- **响应**: `200 OK`
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 5,
    "page": 1,
    "page_size": 20,
    "data": [
      {
        "id": 1,
        "plan_id": 1,
        "field_name": "status",
        "old_value": "Todo",
        "new_value": "InProgress",
        "change_type": "Status",
        "changed_at": "2024-01-01T12:00:00Z"
      },
      {
        "id": 2,
        "plan_id": 1,
        "field_name": "progress",
        "old_value": "0",
        "new_value": "50",
        "change_type": "Progress",
        "changed_at": "2024-01-01T12:30:00Z"
      }
    ]
  }
}
```

### 5. 健康检查

#### 5.1 健康检查
- **URL**: `GET /health`
- **描述**: 检查服务健康状态
- **响应**: `200 OK`
```json
{
  "status": "ok"
}
```

## 数据模型

### Plan
| 字段 | 类型 | 说明 | 必填 |
|------|------|------|------|
| id | uint | 计划 ID | - |
| title | string | 标题 | 是 |
| description | string | 描述 | 否 |
| priority | string | 优先级(High, Medium, Low) | 是 |
| status | string | 状态(Todo, InProgress, Done, Cancelled) | 是 |
| due_date | timestamp | 截止日期 | 否 |
| progress | int | 进度(0-100) | 是 |
| created_at | timestamp | 创建时间 | - |
| updated_at | timestamp | 更新时间 | - |

### PlanHistory
| 字段 | 类型 | 说明 | 必填 |
|------|------|------|------|
| id | uint | 历史 ID | - |
| plan_id | uint | 计划 ID | 是 |
| field_name | string | 字段名称 | 是 |
| old_value | string | 旧值 | 否 |
| new_value | string | 新值 | 否 |
| change_type | string | 变更类型(Status, Progress, Info) | 是 |
| changed_at | timestamp | 变更时间 | - |

## 错误码

| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 400 | 请求参数错误 |
| 404 | 资源不存在 |
| 409 | 资源冲突 |
| 500 | 服务器内部错误 |
