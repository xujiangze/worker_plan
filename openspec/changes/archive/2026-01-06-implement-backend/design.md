# Design: 工作计划管理系统后端架构

## Architecture Overview

### Layered Architecture
采用经典的分层架构模式,确保关注点分离和代码可维护性:

```
┌─────────────────────────────────────┐
│         HTTP Layer (Router)        │  ← 路由和中间件
├─────────────────────────────────────┤
│      Controller Layer (Handler)     │  ← 请求处理和响应
├─────────────────────────────────────┤
│       Service Layer (Business)      │  ← 业务逻辑
├─────────────────────────────────────┤
│    Repository Layer (Data Access)  │  ← 数据访问
├─────────────────────────────────────┤
│         Database (PostgreSQL)       │  ← 数据持久化
└─────────────────────────────────────┘
```

### Project Structure
```
worker_plan/
├── cmd/
│   └── server/
│       └── main.go              # 应用入口
├── internal/
│   ├── config/                  # 配置管理
│   │   └── config.go
│   ├── controller/              # HTTP 处理器
│   │   ├── plan_controller.go
│   │   ├── progress_controller.go
│   │   ├── stats_controller.go
│   │   └── history_controller.go
│   ├── service/                 # 业务逻辑层
│   │   ├── plan_service.go
│   │   ├── progress_service.go
│   │   ├── stats_service.go
│   │   └── history_service.go
│   ├── repository/              # 数据访问层
│   │   ├── plan_repository.go
│   │   ├── plan_history_repository.go
│   │   └── repository.go
│   ├── model/                   # 数据模型
│   │   ├── plan.go
│   │   └── plan_history.go
│   ├── middleware/              # 中间件
│   │   ├── logger.go
│   │   ├── cors.go
│   │   └── recovery.go
│   └── router/                  # 路由配置
│       └── router.go
├── pkg/
│   ├── database/                # 数据库连接
│   │   └── database.go
│   └── utils/                   # 工具函数
│       └── utils.go
├── migrations/                  # 数据库迁移
│   └── 001_init_schema.up.sql
├── tests/                       # 测试文件
│   ├── service/
│   └── integration/
├── go.mod
├── go.sum
└── .env.example
```

## Technology Decisions

### 1. Web Framework: Gin
**选择理由**:
- 高性能,比标准库 net/http 更快
- 中间件支持完善
- 路由功能强大,支持参数绑定和验证
- 社区活跃,文档完善
- 与 GORM 集成良好

**替代方案考虑**:
- Echo: 性能相近,但 Gin 生态更成熟
- 标准库 net/http: 过于底层,开发效率低

### 2. ORM: GORM
**选择理由**:
- 功能完善,支持关联、钩子、事务等
- 自动迁移功能,便于开发
- 支持多种数据库
- 社区活跃,文档完善
- 与 Go 生态集成良好

**替代方案考虑**:
- sqlx: 性能更好,但需要手写 SQL
- sqlc: 类型安全,但学习曲线陡峭

### 3. Database: PostgreSQL
**选择理由**:
- 功能强大,支持 JSON、数组等高级类型
- 事务支持完善
- 性能优秀,适合复杂查询
- 开源免费,社区活跃
- 符合需求文档要求

### 4. Configuration: Viper
**选择理由**:
- 支持多种配置格式(JSON, YAML, ENV)
- 支持配置热重载
- 与环境变量集成良好
- 社区标准

### 5. Logging: Zap
**选择理由**:
- 高性能,零内存分配
- 结构化日志
- 支持多种输出格式
- Uber 开源,质量可靠

## Data Model Design

### Plans Table
```go
type Plan struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Title       string    `gorm:"size:255;not null" json:"title"`
    Description string    `gorm:"type:text" json:"description"`
    Priority    string    `gorm:"size:20;not null" json:"priority"` // High, Medium, Low
    Status      string    `gorm:"size:20;not null" json:"status"` // Todo, InProgress, Done, Cancelled
    DueDate     *time.Time `json:"due_date"`
    Progress    int       `gorm:"not null;default:0" json:"progress"` // 0-100
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    DeletedAt   *time.Time `gorm:"index" json:"-"` // 软删除
}
```

### PlanHistory Table
```go
type PlanHistory struct {
    ID         uint      `gorm:"primaryKey" json:"id"`
    PlanID     uint      `gorm:"not null;index" json:"plan_id"`
    FieldName  string    `gorm:"size:50;not null" json:"field_name"`
    OldValue   string    `gorm:"type:text" json:"old_value"`
    NewValue   string    `gorm:"type:text" json:"new_value"`
    ChangeType string    `gorm:"size:20;not null" json:"change_type"` // Status, Progress, Info
    ChangedAt  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"changed_at"`
}
```

## API Design Principles

### RESTful Conventions
- 使用标准 HTTP 方法: GET(查询), POST(创建), PUT(更新), DELETE(删除), PATCH(部分更新)
- 使用资源路径: `/api/plans`, `/api/plans/:id`
- 使用 HTTP 状态码: 200(成功), 201(创建), 204(无内容), 400(错误请求), 404(未找到), 500(服务器错误)
- 使用 JSON 格式请求和响应

### Response Format
统一响应格式:
```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

错误响应:
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

### Pagination
列表查询支持分页:
```
GET /api/plans?page=1&page_size=20
```

响应:
```json
{
  "total": 100,
  "page": 1,
  "page_size": 20,
  "data": []
}
```

## Business Logic Design

### Plan Management
- **创建计划**: 验证必填字段,设置默认状态为 "Todo",进度为 0
- **更新计划**: 验证字段,记录变更历史
- **删除计划**: 使用软删除,保留历史记录
- **查询计划**: 支持筛选、排序、分页

### Progress Tracking
- **更新状态**: 验证状态转换合法性,记录历史
- **更新进度**: 验证进度范围(0-100),进度达到 100% 时自动更新状态为 "Done"
- **历史记录**: 自动记录所有状态和进度变更

### Statistics
- **按状态统计**: 统计各状态计划数量和占比
- **按优先级统计**: 统计各优先级计划数量和占比
- **按时间统计**: 统计指定时间范围内的创建和完成数量
- **完成率计算**: (已完成计划数 / 总计划数) × 100%

### History Tracking
- **记录变更**: 在更新计划时自动记录变更历史
- **查询历史**: 支持按计划 ID 查询所有历史记录
- **变更类型**: Status(状态变更), Progress(进度变更), Info(信息变更)

## Error Handling Strategy

### Error Types
1. **Validation Error**: 请求参数验证失败 (400)
2. **Not Found Error**: 资源不存在 (404)
3. **Conflict Error**: 资源冲突 (409)
4. **Internal Error**: 服务器内部错误 (500)

### Error Handling Flow
```
Controller → Service → Repository
    ↓           ↓           ↓
  捕获错误    业务逻辑错误  数据库错误
    ↓           ↓           ↓
  转换为 HTTP 响应
```

### Logging Strategy
- 使用结构化日志记录关键操作
- 记录请求日志: 方法、路径、参数、响应时间
- 记录错误日志: 错误类型、堆栈信息
- 记录业务日志: 创建计划、更新状态等

## Testing Strategy

### Unit Tests
- Service 层业务逻辑测试
- Repository 层数据访问测试
- 工具函数测试
- 目标覆盖率: >70%

### Integration Tests
- API 端到端测试
- 数据库集成测试
- 使用测试数据库,不污染生产数据

### Test Tools
- `testing`: Go 标准测试框架
- `testify`: 断言库
- `gomock`: Mock 生成工具
- `sqlmock`: 数据库 Mock

## Security Considerations

### Input Validation
- 使用参数绑定和验证
- 验证字段类型、长度、格式
- 防止 SQL 注入(GORM 参数化查询)

### Data Protection
- 敏感数据使用 HTTPS 传输
- 数据库连接使用 SSL
- 软删除保留数据,可恢复

### Rate Limiting
- API 限流,防止滥用
- 使用令牌桶算法

## Performance Optimization

### Database Optimization
- 为常用查询字段创建索引
- 使用连接池管理数据库连接
- 优化复杂查询,避免 N+1 问题

### Caching Strategy
- 考虑使用 Redis 缓存统计数据
- 缓存热点数据,减少数据库查询

### Query Optimization
- 使用分页避免大量数据查询
- 使用预加载(Preload)减少查询次数

## Deployment Considerations

### Environment Variables
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=worker_plan
SERVER_PORT=8080
LOG_LEVEL=info
```

### Database Migration
- 使用 GORM AutoMigrate 或 golang-migrate
- 版本化管理迁移脚本
- 支持回滚

### Health Check
- 提供 `/health` 端点
- 检查数据库连接状态
- 返回服务健康状态

## Future Extensions

### Authentication & Authorization
- 添加用户认证(JWT)
- 添加权限控制(RBAC)
- 用户只能访问自己的计划

### Real-time Updates
- 使用 WebSocket 推送实时更新
- 使用 Server-Sent Events (SSE)

### Advanced Features
- 计划标签和分类
- 计划提醒和通知
- 计划模板
- 计划导出(Excel, PDF)
