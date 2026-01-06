# 修复创建计划 API 响应处理错误

## 问题概述

前端在提交创建计划给后端接口时出现错误:

```
plan.ts:63 创建计划失败: TypeError: Cannot read properties of undefined (reading 'unshift')
    at Proxy.createPlan (plan.ts:59:19)
    at async handleSubmit (PlanForm.vue:168:7)
```

同时页面展示上也有问题。

## 根本原因分析

经过代码分析,发现问题的根源在于前后端 API 响应数据结构不匹配:

1. **后端响应结构** (`internal/controller/response.go:39-44`):
   ```go
   func CreatedResponse(c *gin.Context, data interface{}) {
       c.JSON(http.StatusCreated, Response{
           Code:    0,
           Message: "success",
           Data:    data,
       })
   }
   ```
   后端返回的格式是: `{code: 0, message: "success", data: {...}}`

2. **前端响应拦截器** (`frontend/src/utils/request.ts:36-39`):
   ```typescript
   request.interceptors.response.use(
     (response: AxiosResponse) => {
       // 统一处理响应数据
       return response.data
     },
   ```
   前端拦截器直接返回 `response.data`,即整个响应对象

3. **前端 Store 期望** (`frontend/src/stores/plan.ts:55-59`):
   ```typescript
   const createPlan = async (data: CreatePlanRequest) => {
     loading.value = true
     try {
       const plan = await planApi.createPlan(data)
       plans.value.unshift(plan)  // 这里期望 plan 是 Plan 对象
   ```
   Store 期望直接得到 Plan 对象,但实际上得到的是包装在 Response 结构中的数据

4. **结果**: 当尝试对 undefined 调用 `unshift` 时,就会报错,因为 `plan` 实际上是 `{code: 0, message: "success", data: {...}}`,而不是 Plan 对象

## 影响范围

- **受影响的功能**: 创建计划功能
- **受影响的文件**:
  - `frontend/src/utils/request.ts` - 响应拦截器
  - `frontend/src/stores/plan.ts` - 计划 Store
  - `frontend/src/api/plan.ts` - 计划 API
- **潜在影响**: 所有使用该 API 的功能都可能受到影响

## 解决方案

### 方案选择

有两种解决方案:

**方案 A**: 修改响应拦截器,自动提取 `data` 字段
- 优点: 一次修改,所有 API 调用都受益,代码更简洁
- 缺点: 需要确保所有 API 响应都遵循统一格式

**方案 B**: 修改 Store 和 API 层,手动提取 `data` 字段
- 优点: 更明确,不依赖响应拦截器的行为
- 缺点: 需要修改多处代码,容易遗漏

**推荐方案**: 方案 A,因为:
1. 后端已经统一使用 Response 结构
2. 前端代码更简洁,减少重复代码
3. 符合 RESTful API 最佳实践

### 实施计划

1. 修改 `frontend/src/utils/request.ts` 响应拦截器,返回 `response.data.data`
2. 验证所有 API 调用是否正常工作
3. 测试创建计划功能
4. 测试其他计划相关功能(列表、详情、更新、删除)

## 验证标准

- [ ] 创建计划功能正常,不再报错
- [ ] 创建成功后,计划列表正确显示新创建的计划
- [ ] 其他计划相关功能(列表、详情、更新、删除)正常工作
- [ ] 页面展示正常,无错误信息
- [ ] 前端控制台无错误日志

## 风险评估

- **风险等级**: 低
- **风险描述**: 修改响应拦截器可能影响所有 API 调用
- **缓解措施**:
  1. 充分测试所有 API 调用
  2. 如果出现问题,可以快速回滚
  3. 添加错误处理,确保向后兼容

## 相关资源

- 错误日志: `plan.ts:63`
- 后端响应结构: `internal/controller/response.go`
- 前端响应拦截器: `frontend/src/utils/request.ts`
- 前端 Store: `frontend/src/stores/plan.ts`
- 前端 API: `frontend/src/api/plan.ts`
