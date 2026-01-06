# Change: 修复前端无法展示计划列表问题

## Why
前端无法展示计划列表,后端 API 响应有数据但前端没有显示。经过分析,发现后端返回的 `due_date` 字段是 `*time.Time` 类型,在 JSON 序列化时可能转换为时间戳格式,而前端期望的是 ISO 8601 字符串格式。此外,响应拦截器的数据处理逻辑可能存在问题。

## What Changes
- 修改后端 `Plan` 模型的 `DueDate` 字段,添加自定义 JSON 序列化器,确保返回 ISO 8601 格式的字符串
- 验证前端响应拦截器的数据处理逻辑是否正确
- 添加日志记录以便调试数据传输问题
- 确保前后端数据格式一致

## Impact
- Affected specs: `plan-display` (新增)
- Affected code:
  - `internal/model/plan.go` - 修改 DueDate 字段的 JSON 序列化
  - `frontend/src/utils/request.ts` - 验证响应拦截器
  - `frontend/src/stores/plan.ts` - 添加调试日志
