# Design: 统计分析功能前端界面

## Architecture Overview

### Frontend-Backend Interaction
```
StatisticsView.vue
    ↓ (API calls)
stats.ts (API module)
    ↓ (HTTP requests)
Backend API (/api/stats/*)
    ↓ (data queries)
Database (PostgreSQL)
```

### Component Structure
```
StatisticsView.vue
├── StatisticsHeader (标题和日期选择器)
├── CompletionRateCard (完成率卡片)
├── StatsChartsContainer
│   ├── StatusPieChart (状态分布饼图)
│   ├── PriorityPieChart (优先级分布饼图)
│   └── TimeTrendChart (时间趋势图)
└── StatsSummary (统计摘要)
```

## UI/UX Design

### Page Layout
```
+--------------------------------------------------+
|  统计分析                          [日期范围选择]  |
+--------------------------------------------------+
|  +----------------+  +----------------+          |
|  |  完成率: 60%   |  |  总计划: 20    |          |
|  +----------------+  +----------------+          |
+--------------------------------------------------+
|  +----------------+  +----------------+          |
|  |  状态分布      |  |  优先级分布    |          |
|  |  [饼图]        |  |  [饼图]        |          |
|  +----------------+  +----------------+          |
+--------------------------------------------------+
|  时间趋势图                                    |
|  [折线图/柱状图]                               |
+--------------------------------------------------+
```

### Interaction Design
1. **页面加载**: 自动加载最近30天的统计数据
2. **日期范围选择**: 用户可以选择自定义日期范围,选择后自动刷新所有图表
3. **图表交互**:
   - 鼠标悬停显示详细数据
   - 点击图例可以显示/隐藏对应数据
   - 饼图支持点击扇区查看详情
4. **响应式设计**: 在移动端自动调整为单列布局

## Data Flow

### API Data Flow
1. **初始化**: 组件挂载时调用所有统计API
2. **日期变更**: 用户选择日期范围后,重新调用时间统计API
3. **数据转换**: 后端数据转换为图表库所需格式
4. **渲染更新**: 使用Vue响应式系统更新图表

### Data Transformation Examples

#### Status Stats Transformation
```typescript
// Backend response
{
  "Todo": { "count": 5, "percentage": 25 },
  "InProgress": { "count": 8, "percentage": 40 },
  "Done": { "count": 6, "percentage": 30 },
  "Cancelled": { "count": 1, "percentage": 5 }
}

// ECharts format
{
  series: [{
    data: [
      { value: 5, name: '待开始' },
      { value: 8, name: '进行中' },
      { value: 6, name: '已完成' },
      { value: 1, name: '已取消' }
    ]
  }]
}
```

#### Time Stats Transformation
```typescript
// Backend response
{
  "created_count": 15,
  "completed_count": 8,
  "completion_rate": 53.33,
  "daily_trend": [
    { "date": "2026-01-01", "created": 2, "completed": 0 },
    { "date": "2026-01-02", "created": 1, "completed": 1 }
  ]
}

// ECharts format
{
  xAxis: { data: ['2026-01-01', '2026-01-02'] },
  series: [
    { name: '创建', data: [2, 1] },
    { name: '完成', data: [0, 1] }
  ]
}
```

## Technology Choices

### Chart Library Selection: ECharts

**Rationale**:
- 功能强大,支持多种图表类型
- 良好的TypeScript支持
- 丰富的配置选项和主题
- 活跃的社区和文档
- 性能优秀,适合大数据量

**Alternatives Considered**:
- Chart.js: 轻量级但功能相对简单
- Recharts: React专用,不适合Vue项目
- D3.js: 学习曲线陡峭,开发成本高

### State Management
- 使用Vue 3 Composition API的ref/reactive
- 不需要额外的状态管理库(如Pinia),因为统计页面数据相对独立

### Date Handling
- 使用原生Date API
- 日期格式: YYYY-MM-DD (与后端保持一致)

## Performance Considerations

### Backend Optimization
1. **数据库索引**: 在status、priority、created_at、updated_at字段上添加索引
2. **查询优化**: 使用COUNT聚合函数,避免全表扫描
3. **缓存策略**: 考虑对统计数据进行缓存(可选)

### Frontend Optimization
1. **懒加载**: 图表组件按需加载
2. **防抖**: 日期选择器输入防抖,避免频繁API调用
3. **图表优化**: 使用ECharts的渐进式渲染
4. **数据分页**: 时间趋势数据量大时考虑分页或采样

### Performance Targets
- 页面首次加载时间 < 1s
- API响应时间 < 200ms (P95)
- 图表渲染时间 < 100ms
- 日期切换响应时间 < 500ms

## Error Handling

### API Error Handling
1. **网络错误**: 显示友好的错误提示,提供重试按钮
2. **数据格式错误**: 记录错误日志,显示默认空图表
3. **权限错误**: 重定向到登录页面(未来扩展)

### Validation
1. **日期范围验证**: 开始日期不能晚于结束日期
2. **日期格式验证**: 确保日期格式为YYYY-MM-DD
3. **数据完整性验证**: 检查API返回数据的完整性

## Accessibility
- 为图表添加ARIA标签
- 支持键盘导航
- 提供文本替代方案
- 确保颜色对比度符合WCAG标准

## Future Enhancements
1. **数据导出**: 支持导出统计报告(PDF/Excel)
2. **自定义统计**: 允许用户自定义统计维度
3. **实时更新**: 使用WebSocket实现统计数据实时更新
4. **多维度分析**: 支持多维度交叉分析(如按状态和优先级)
5. **数据对比**: 支持不同时间段的数据对比
