# Project Context

## Purpose
工作计划管理系统,用于帮助用户创建、跟踪和统计工作计划。系统提供计划管理、进度跟踪、统计分析等功能,帮助用户提高工作效率。

## Tech Stack
- **后端**: Go (Golang)
- **前端**: Vue.js
- **数据库**: PostgreSQL
- **ORM**: GORM
- **开发工具**: Claude Code + OpenSpec

## Project Conventions

### Code Style
- Go 代码遵循 Go 官方代码规范和 `gofmt` 格式化
- Vue 代码遵循 Vue.js 官方风格指南
- 使用有意义的变量和函数命名
- 保持代码简洁,避免过度设计

### Architecture Patterns
- 采用前后端分离架构
- RESTful API 设计
- 分层架构: Controller -> Service -> Repository
- 使用依赖注入模式

### Testing Strategy
- 单元测试覆盖核心业务逻辑
- API 集成测试
- 前端组件测试
- 测试覆盖率目标: >70%

### Git Workflow
- 主分支: `main`
- 功能分支: `feature/功能名称`
- 修复分支: `fix/问题描述`
- 使用 Conventional Commits 规范
- 提交前必须通过代码审查

## Domain Context
- **工作计划**: 用户创建的任务或目标,包含标题、描述、优先级、截止日期等
- **计划状态**: 待开始、进行中、已完成、已取消
- **优先级**: 高、中、低
- **统计维度**: 按时间、按状态、按优先级等维度统计

## Important Constraints
- 数据必须持久化存储
- 支持多用户(未来扩展)
- API 响应时间 < 200ms (P95)
- 支持至少 1000 个并发用户

## External Dependencies
- PostgreSQL 数据库
- GORM ORM 库
- Vue.js 前端框架
- Go 标准库及常用第三方库
