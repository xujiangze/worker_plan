import type { EChartsOption } from 'echarts'

// 图表主题配置
export const chartTheme = {
  colors: ['#409EFF', '#67C23A', '#E6A23C', '#F56C6C', '#909399'],
  backgroundColor: '#ffffff',
  textStyle: {
    fontFamily: 'Arial, sans-serif'
  }
}

// 饼图默认配置
export const defaultPieChartOptions: EChartsOption = {
  tooltip: {
    trigger: 'item',
    formatter: '{a} <br/>{b}: {c} ({d}%)'
  },
  legend: {
    orient: 'vertical',
    left: 'left'
  },
  series: [
    {
      name: '统计',
      type: 'pie',
      radius: '50%',
      data: [],
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      }
    }
  ]
}

// 折线图默认配置
export const defaultLineChartOptions: EChartsOption = {
  tooltip: {
    trigger: 'axis'
  },
  legend: {
    data: []
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: {
    type: 'category',
    boundaryGap: false,
    data: []
  },
  yAxis: {
    type: 'value'
  },
  series: []
}

// 柱状图默认配置
export const defaultBarChartOptions: EChartsOption = {
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow'
    }
  },
  legend: {
    data: []
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: {
    type: 'category',
    data: []
  },
  yAxis: {
    type: 'value'
  },
  series: []
}

// 状态映射
export const statusMap: Record<string, string> = {
  Todo: '待开始',
  InProgress: '进行中',
  Done: '已完成',
  Cancelled: '已取消'
}

// 优先级映射
export const priorityMap: Record<string, string> = {
  High: '高',
  Medium: '中',
  Low: '低'
}
