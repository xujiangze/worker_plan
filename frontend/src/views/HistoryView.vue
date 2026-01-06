<template>
  <div class="history-view">
    <el-card v-loading="loading">
      <!-- 页面头部 -->
      <template #header>
        <div class="card-header">
          <div class="header-left">
            <el-button @click="goBack" :icon="ArrowLeft">返回</el-button>
            <span class="page-title">历史记录 - {{ plan?.title || '加载中...' }}</span>
          </div>
        </div>
      </template>

      <!-- 计划信息卡片 -->
      <div v-if="plan" class="plan-info-card">
        <div class="info-row">
          <span class="label">状态:</span>
          <StatusBadge :status="plan.status" />
        </div>
        <div class="info-row">
          <span class="label">优先级:</span>
          <PriorityBadge :priority="plan.priority" />
        </div>
        <div class="info-row">
          <span class="label">进度:</span>
          <ProgressBar :percentage="plan.progress" />
        </div>
        <div v-if="plan.due_date" class="info-row">
          <span class="label">截止日期:</span>
          <span>{{ formatDate(plan.due_date) }}</span>
        </div>
      </div>

      <!-- 筛选栏 -->
      <div class="filter-bar">
        <el-select
          v-model="filters.changeType"
          placeholder="变更类型"
          clearable
          style="width: 150px; margin-right: 10px"
          @change="handleFilterChange"
        >
          <el-option label="全部" value="" />
          <el-option label="状态" value="Status" />
          <el-option label="进度" value="Progress" />
          <el-option label="信息" value="Info" />
        </el-select>

        <el-date-picker
          v-model="dateRange"
          type="daterange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          value-format="YYYY-MM-DD"
          style="width: 300px"
          @change="handleDateRangeChange"
        />
      </div>

      <!-- 历史记录列表 -->
      <div v-if="!loading" class="history-list">
        <el-empty v-if="!histories || histories.length === 0" description="暂无历史记录" />

        <div v-else class="history-items">
          <el-card v-for="history in histories" :key="history.id" class="history-item" shadow="hover">
            <div class="history-item-header">
              <el-tag :type="getChangeTypeColor(history.change_type)" size="small">
                {{ getChangeTypeLabel(history.change_type) }}
              </el-tag>
              <span class="history-time">{{ formatDateTime(history.changed_at) }}</span>
            </div>
            <div class="history-item-body">
              <div class="field-name">{{ history.field }}</div>
              <div class="field-change">
                <span class="old-value">{{ history.old_value }}</span>
                <el-icon class="arrow-icon"><ArrowRight /></el-icon>
                <span class="new-value">{{ history.new_value }}</span>
              </div>
            </div>
          </el-card>
        </div>
      </div>

      <!-- 分页 -->
      <el-pagination
        v-if="!loading && pagination.total > 0"
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.pageSize"
        :total="pagination.total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        class="pagination"
        @size-change="handleSizeChange"
        @current-change="handlePageChange"
      />
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ArrowLeft, ArrowRight } from '@element-plus/icons-vue'
import { getPlanHistory, getPlanDetail } from '@/api/history'
import type { Plan, PlanHistory, ChangeType } from '@/types/api'
import StatusBadge from '@/components/StatusBadge.vue'
import PriorityBadge from '@/components/PriorityBadge.vue'
import ProgressBar from '@/components/ProgressBar.vue'

const route = useRoute()
const router = useRouter()

// 组件状态
const planId = ref<number>(Number(route.params.planId))
const plan = ref<Plan | null>(null)
const histories = ref<PlanHistory[]>([])
const loading = ref(false)
const dateRange = ref<[string, string] | null>(null)

// 筛选条件
const filters = reactive({
  changeType: '',
  startDate: '',
  endDate: '',
})

// 分页信息
const pagination = reactive({
  page: 1,
  pageSize: 20,
  total: 0,
})

// 获取计划详情
const fetchPlan = async () => {
  try {
    loading.value = true
    const response = await getPlanDetail(planId.value)
    plan.value = response.data
  } catch (error: any) {
    ElMessage.error(error.message || '获取计划详情失败')
  } finally {
    loading.value = false
  }
}

// 获取历史记录
const fetchHistories = async () => {
  try {
    loading.value = true
    const params = {
      page: pagination.page,
      page_size: pagination.pageSize,
      change_type: filters.changeType || undefined,
      start_date: filters.startDate || undefined,
      end_date: filters.endDate || undefined,
    }
    const response = await getPlanHistory(planId.value, params)
    histories.value = response.data.data
    pagination.total = response.data.total
    pagination.page = response.data.page
    pagination.pageSize = response.data.page_size
  } catch (error: any) {
    ElMessage.error(error.message || '获取历史记录失败')
  } finally {
    loading.value = false
  }
}

// 处理筛选条件变化
const handleFilterChange = () => {
  pagination.page = 1
  fetchHistories()
  updateUrlQuery()
}

// 处理时间范围变化
const handleDateRangeChange = (value: [string, string] | null) => {
  if (value) {
    filters.startDate = value[0]
    filters.endDate = value[1]
  } else {
    filters.startDate = ''
    filters.endDate = ''
  }
  pagination.page = 1
  fetchHistories()
  updateUrlQuery()
}

// 处理分页大小变化
const handleSizeChange = (size: number) => {
  pagination.pageSize = size
  pagination.page = 1
  fetchHistories()
  updateUrlQuery()
}

// 处理页码变化
const handlePageChange = (page: number) => {
  pagination.page = page
  fetchHistories()
  updateUrlQuery()
}

// 更新 URL 查询参数
const updateUrlQuery = () => {
  const query: any = {
    page: pagination.page,
    page_size: pagination.pageSize,
  }
  if (filters.changeType) {
    query.change_type = filters.changeType
  }
  if (filters.startDate) {
    query.start_date = filters.startDate
  }
  if (filters.endDate) {
    query.end_date = filters.endDate
  }
  router.replace({ query })
}

// 从 URL 查询参数恢复状态
const restoreFromQuery = () => {
  const query = route.query
  if (query.page) {
    pagination.page = Number(query.page)
  }
  if (query.page_size) {
    pagination.pageSize = Number(query.page_size)
  }
  if (query.change_type) {
    filters.changeType = String(query.change_type)
  }
  if (query.start_date && query.end_date) {
    filters.startDate = String(query.start_date)
    filters.endDate = String(query.end_date)
    dateRange.value = [filters.startDate, filters.endDate]
  }
}

// 返回计划列表
const goBack = () => {
  router.push('/plans')
}

// 格式化日期
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
  })
}

// 格式化日期时间
const formatDateTime = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
  })
}

// 获取变更类型标签
const getChangeTypeLabel = (type?: ChangeType) => {
  const labels: Record<string, string> = {
    Status: '状态',
    Progress: '进度',
    Info: '信息',
  }
  return labels[type || ''] || '未知'
}

// 获取变更类型颜色
const getChangeTypeColor = (type?: ChangeType) => {
  const colors: Record<string, any> = {
    Status: 'primary',
    Progress: 'success',
    Info: 'info',
  }
  return colors[type || ''] || 'info'
}

// 监听路由参数变化
watch(
  () => route.params.planId,
  (newPlanId) => {
    if (newPlanId) {
      planId.value = Number(newPlanId)
      fetchPlan()
      fetchHistories()
    }
  }
)

onMounted(() => {
  restoreFromQuery()
  fetchPlan()
  fetchHistories()
})
</script>

<style scoped>
.history-view {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 10px;
}

.page-title {
  font-size: 16px;
  font-weight: 500;
  color: #303133;
}

.plan-info-card {
  padding: 20px;
  margin-bottom: 20px;
  background-color: #f5f7fa;
  border-radius: 4px;
}

.info-row {
  display: flex;
  align-items: center;
  margin-bottom: 12px;
}

.info-row:last-child {
  margin-bottom: 0;
}

.info-row .label {
  font-size: 14px;
  color: #606266;
  min-width: 80px;
  font-weight: 500;
}

.filter-bar {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  padding: 15px;
  background-color: #f5f7fa;
  border-radius: 4px;
}

.history-list {
  margin-bottom: 20px;
}

.history-items {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.history-item {
  transition: all 0.3s;
}

.history-item:hover {
  transform: translateY(-2px);
}

.history-item-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.history-time {
  font-size: 13px;
  color: #909399;
}

.history-item-body {
  padding-top: 10px;
}

.field-name {
  font-size: 14px;
  font-weight: 500;
  color: #303133;
  margin-bottom: 8px;
}

.field-change {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 13px;
}

.old-value {
  color: #909399;
  text-decoration: line-through;
}

.new-value {
  color: #67c23a;
  font-weight: 500;
}

.arrow-icon {
  color: #909399;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: center;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .history-view {
    padding: 10px;
  }

  .filter-bar {
    flex-direction: column;
    align-items: stretch;
    gap: 10px;
  }

  .filter-bar .el-select,
  .filter-bar .el-date-picker {
    width: 100% !important;
  }

  .history-item-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }

  .field-change {
    flex-direction: column;
    align-items: flex-start;
    gap: 5px;
  }

  .arrow-icon {
    transform: rotate(90deg);
  }
}

@media (min-width: 769px) and (max-width: 1199px) {
  .plan-info-card {
    padding: 15px;
  }

  .filter-bar {
    padding: 12px;
  }
}
</style>
