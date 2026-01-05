# Tasks: 实现工作计划管理系统后端

## Phase 1: 项目初始化和基础设施

### 1.1 初始化 Go 项目
- [ ] 创建 Go 模块: `go mod init worker_plan`
- [ ] 创建项目目录结构(cmd, internal, pkg, migrations, tests)
- [ ] 配置 .gitignore 文件
- [ ] 创建 .env.example 配置文件模板

### 1.2 配置依赖管理
- [ ] 添加 Gin 框架依赖
- [ ] 添加 GORM 依赖
- [ ] 添加 PostgreSQL 驱动依赖
- [ ] 添加 Viper 配置管理依赖
- [ ] 添加 Zap 日志依赖
- [ ] 添加测试依赖(testify, gomock, sqlmock)
- [ ] 运行 `go mod tidy` 整理依赖

### 1.3 实现配置管理
- [ ] 创建 `internal/config/config.go` 配置结构体
- [ ] 实现配置加载逻辑(从环境变量和配置文件)
- [ ] 添加配置验证
- [ ] 编写配置加载单元测试

### 1.4 实现数据库连接
- [ ] 创建 `pkg/database/database.go` 数据库连接模块
- [ ] 实现数据库连接池配置
- [ ] 添加数据库连接健康检查
- [ ] 编写数据库连接单元测试

### 1.5 实现日志系统
- [ ] 创建 `internal/middleware/logger.go` 日志中间件
- [ ] 配置 Zap 日志格式和输出
- [ ] 实现请求日志记录
- [ ] 实现错误日志记录

### 1.6 实现中间件
- [ ] 创建 `internal/middleware/cors.go` CORS 中间件
- [ ] 创建 `internal/middleware/recovery.go` 恢复中间件
- [ ] 编写中间件单元测试

## Phase 2: 数据模型和数据库迁移

### 2.1 定义数据模型
- [ ] 创建 `internal/model/plan.go` Plan 模型
- [ ] 创建 `internal/model/plan_history.go` PlanHistory 模型
- [ ] 添加 GORM 标签和 JSON 标签
- [ ] 添加模型验证规则

### 2.2 创建数据库迁移文件
- [ ] 创建 `migrations/001_init_schema.up.sql` 初始化表结构
- [ ] 创建 Plans 表,包含所有字段和索引
- [ ] 创建 PlanHistory 表,包含所有字段和索引
- [ ] 添加外键约束
- [ ] 创建 `migrations/001_init_schema.down.sql` 回滚脚本

### 2.3 实现数据库迁移
- [ ] 集成 golang-migrate 或使用 GORM AutoMigrate
- [ ] 实现迁移命令行工具
- [ ] 测试迁移和回滚功能

## Phase 3: 数据访问层(Repository)

### 3.1 实现 Plan Repository
- [ ] 创建 `internal/repository/plan_repository.go`
- [ ] 实现 Create 方法
- [ ] 实现 FindByID 方法
- [ ] 实现 FindAll 方法(支持筛选、排序、分页)
- [ ] 实现 Update 方法
- [ ] 实现 Delete 方法(软删除)
- [ ] 编写 Plan Repository 单元测试

### 3.2 实现 PlanHistory Repository
- [ ] 创建 `internal/repository/plan_history_repository.go`
- [ ] 实现 Create 方法
- [ ] 实现 FindByPlanID 方法(支持筛选、分页)
- [ ] 编写 PlanHistory Repository 单元测试

### 3.3 实现 Repository 接口
- [ ] 创建 `internal/repository/repository.go` 接口定义
- [ ] 定义 PlanRepository 接口
- [ ] 定义 PlanHistoryRepository 接口
- [ ] 实现依赖注入

## Phase 4: 业务逻辑层(Service)

### 4.1 实现 Plan Service
- [ ] 创建 `internal/service/plan_service.go`
- [ ] 实现 CreatePlan 方法(验证、创建、返回)
- [ ] 实现 GetPlan 方法
- [ ] 实现 GetPlans 方法(筛选、排序、分页)
- [ ] 实现 UpdatePlan 方法(验证、更新、记录历史)
- [ ] 实现 DeletePlan 方法(软删除)
- [ ] 编写 Plan Service 单元测试(使用 Mock)

### 4.2 实现 Progress Service
- [ ] 创建 `internal/service/progress_service.go`
- [ ] 实现 UpdateStatus 方法(验证、更新、记录历史)
- [ ] 实现 UpdateProgress 方法(验证、更新、自动调整状态、记录历史)
- [ ] 实现状态转换验证逻辑
- [ ] 实现进度与状态联动逻辑
- [ ] 编写 Progress Service 单元测试(使用 Mock)

### 4.3 实现 Statistics Service
- [ ] 创建 `internal/service/stats_service.go`
- [ ] 实现 GetStatsByStatus 方法
- [ ] 实现 GetStatsByPriority 方法
- [ ] 实现 GetStatsByTime 方法(时间范围验证、趋势计算)
- [ ] 实现 GetCompletionRate 方法
- [ ] 实现统计数据缓存(可选)
- [ ] 编写 Statistics Service 单元测试(使用 Mock)

### 4.4 实现 History Service
- [ ] 创建 `internal/service/history_service.go`
- [ ] 实现 GetHistoryByPlanID 方法(筛选、分页)
- [ ] 实现 RecordHistory 方法(自动记录变更)
- [ ] 编写 History Service 单元测试(使用 Mock)

## Phase 5: HTTP 处理层(Controller)

### 5.1 实现 Plan Controller
- [ ] 创建 `internal/controller/plan_controller.go`
- [ ] 实现 CreatePlanHandler (POST /api/plans)
- [ ] 实现 GetPlansHandler (GET /api/plans)
- [ ] 实现 GetPlanHandler (GET /api/plans/:id)
- [ ] 实现 UpdatePlanHandler (PUT /api/plans/:id)
- [ ] 实现 DeletePlanHandler (DELETE /api/plans/:id)
- [ ] 添加请求参数绑定和验证
- [ ] 添加错误处理和响应格式化
- [ ] 编写 Plan Controller 单元测试

### 5.2 实现 Progress Controller
- [ ] 创建 `internal/controller/progress_controller.go`
- [ ] 实现 UpdateStatusHandler (PATCH /api/plans/:id/status)
- [ ] 实现 UpdateProgressHandler (PATCH /api/plans/:id/progress)
- [ ] 添加请求参数绑定和验证
- [ ] 添加错误处理和响应格式化
- [ ] 编写 Progress Controller 单元测试

### 5.3 实现 Statistics Controller
- [ ] 创建 `internal/controller/stats_controller.go`
- [ ] 实现 GetStatsByStatusHandler (GET /api/stats/by-status)
- [ ] 实现 GetStatsByPriorityHandler (GET /api/stats/by-priority)
- [ ] 实现 GetStatsByTimeHandler (GET /api/stats/by-time)
- [ ] 实现 GetCompletionRateHandler (GET /api/stats/completion-rate)
- [ ] 添加请求参数绑定和验证
- [ ] 添加错误处理和响应格式化
- [ ] 编写 Statistics Controller 单元测试

### 5.4 实现 History Controller
- [ ] 创建 `internal/controller/history_controller.go`
- [ ] 实现 GetHistoryHandler (GET /api/plans/:id/history)
- [ ] 添加请求参数绑定和验证
- [ ] 添加错误处理和响应格式化
- [ ] 编写 History Controller 单元测试

## Phase 6: 路由和应用启动

### 6.1 实现路由配置
- [ ] 创建 `internal/router/router.go`
- [ ] 配置计划管理路由
- [ ] 配置进度管理路由
- [ ] 配置统计分析路由
- [ ] 配置历史记录路由
- [ ] 添加健康检查路由
- [ ] 添加中间件(日志、CORS、恢复)

### 6.2 实现应用入口
- [ ] 创建 `cmd/server/main.go`
- [ ] 初始化配置
- [ ] 初始化数据库连接
- [ ] 初始化依赖注入
- [ ] 初始化路由
- [ ] 启动 HTTP 服务器
- [ ] 实现优雅关闭

## Phase 7: 集成测试

### 7.1 API 集成测试
- [ ] 创建测试数据库配置
- [ ] 编写计划管理 API 集成测试
- [ ] 编写进度管理 API 集成测试
- [ ] 编写统计分析 API 集成测试
- [ ] 编写历史记录 API 集成测试
- [ ] 测试错误场景和边界情况

### 7.2 性能测试
- [ ] 编写 API 响应时间测试
- [ ] 编写并发请求测试
- [ ] 编写大数据量查询测试
- [ ] 优化性能瓶颈

## Phase 8: 文档和部署准备

### 8.1 编写 API 文档
- [ ] 编写 API 接口文档(使用 Swagger/OpenAPI)
- [ ] 添加请求和响应示例
- [ ] 添加错误码说明

### 8.2 编写部署文档
- [ ] 编写环境配置说明
- [ ] 编写数据库迁移说明
- [ ] 编写部署步骤说明
- [ ] 创建 Dockerfile(可选)

### 8.3 编写开发文档
- [ ] 编写项目结构说明
- [ ] 编写开发环境搭建指南
- [ ] 编写代码规范说明
- [ ] 编写测试指南

## Phase 9: 验收和优化

### 9.1 功能验收
- [ ] 验证所有功能需求已实现
- [ ] 验证所有用户场景可正常使用
- [ ] 验证 API 接口符合设计规范

### 9.2 性能验收
- [ ] 验证 API 响应时间 P95 < 200ms
- [ ] 验证支持 1000 个并发用户
- [ ] 验证数据库查询时间 < 100ms

### 9.3 质量验收
- [ ] 验证单元测试覆盖率 > 70%
- [ ] 修复所有严重 Bug
- [ ] 代码审查和重构

### 9.4 文档验收
- [ ] 验证 API 文档完整
- [ ] 验证部署文档完整
- [ ] 验证开发文档完整

## Dependencies

- Phase 1-2 必须在 Phase 3 之前完成
- Phase 3 必须在 Phase 4 之前完成
- Phase 4 必须在 Phase 5 之前完成
- Phase 5 必须在 Phase 6 之前完成
- Phase 6 必须在 Phase 7 之前完成
- Phase 7 必须在 Phase 8 之前完成
- Phase 8 必须在 Phase 9 之前完成

## Parallelizable Work

以下任务可以并行开发:
- Phase 3.1 和 Phase 3.2 可以并行
- Phase 4.1、Phase 4.2、Phase 4.3、Phase 4.4 可以并行
- Phase 5.1、Phase 5.2、Phase 5.3、Phase 5.4 可以并行
- Phase 7.1 的各个 API 测试可以并行
