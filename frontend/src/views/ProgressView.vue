<template>
  <div class="progress-view">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>进度跟踪</span>
          <el-button type="primary" @click="refreshData">刷新</el-button>
        </div>
      </template>

      <!-- 统计信息 -->
      <div class="stats-container">
        <el-row :gutter="20">
          <el-col :span="6">
            <el-card class="stat-card">
              <div class="stat-content">
                <div class="stat-value">{{ totalPlans }}</div>
                <div class="stat-label">总计划数</div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card class="stat-card">
              <div class="stat-content">
                <div class="stat-value">{{ completedPlans }}</div>
                <div class="stat-label">已完成</div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card class="stat-card">
              <div class="stat-content">
                <div class="stat-value">{{ inProgressPlans }}</div>
                <div class="stat-label">进行中</div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card class="stat-card">
              <div class="stat-content">
                <div class="stat-value">{{ averageProgress }}%</div>
                <div class="stat-label">平均进度</div>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </div>

      <!-- 筛选和排序工具栏 -->
      <div class="toolbar">
        <el-select
          v-model="filterStatus"
          placeholder="筛选状态"
          clearable
          style="width: 150px; margin-right: 10px"
          @change="handleFilterChange"
        >
          <el-option label="待办" :value="PlanStatus.Todo" />
          <el-option label="进行中" :value="PlanStatus.InProgress" />
          <el-option label="已完成" :value="PlanStatus.Done" />
          <el-option label="已取消" :value="PlanStatus.Cancelled" />
        </el-select>

        <el-select
          v-model="sortBy"
          placeholder="排序方式"
          style="width: 180px"
          @change="handleSortChange"
        >
          <el-option label="进度(从高到低)" value="progress_desc" />
          <el-option label="进度(从低到高)" value="progress_asc" />
          <el-option label="截止日期(最近)" value="due_date_asc" />
          <el-option label="截止日期(最远)" value="due_date_desc" />
        </el-select>
      </div>

      <!-- 加载状态 -->
      <Loading v-if="loading" loading text="加载中..." />

      <!-- 空状态 -->
      <el-empty v-else-if="!plans || plans.length === 0" description="暂无计划数据" />

      <!-- 进度卡片列表 -->
      <div v-else class="progress-cards">
        <el-card v-for="plan in sortedPlans" :key="plan.id" class="progress-card" shadow="hover">
          <div class="progress-card-header">
            <div class="plan-title">
              <span class="title-text">{{ plan.title }}</span>
              <div class="plan-badges">
                <PriorityBadge :priority="plan.priority" />
                <StatusBadge :status="plan.status" />
              </div>
            </div>
            <div class="plan-actions">
              <el-button size="small" @click="goToPlan(plan.id)">查看详情</el-button>
            </div>
          </div>

          <div class="progress-card-body">
            <div v-if="plan.description" class="plan-description">
              {{ plan.description }}
            </div>

            <div class="plan-meta">
              <div class="meta-item">
                <el-icon><Calendar /></el-icon>
                <span>创建时间: {{ formatDate(plan.created_at) }}</span>
              </div>
              <div v-if="plan.due_date" class="meta-item">
                <el-icon><Clock /></el-icon>
                <span>截止日期: {{ formatDate(plan.due_date) }}</span>
              </div>
            </div>

            <!-- 进度条 -->
            <div class="plan-progress">
              <ProgressBar :percentage="plan.progress" />
            </div>

            <!-- 到期提醒 -->
            <div v-if="plan.due_date" class="plan-due-alert">
              <el-tag v-if="isOverdue(plan.due_date)" type="danger" size="small">
                已过期
              </el-tag>
              <el-tag v-else-if="isDueSoon(plan.due_date)" type="warning" size="small">
                即将到期
              </el-tag>
            </div>
          </div>
        </el-card>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Calendar, Clock } from '@element-plus/icons-vue'
import { usePlanStore } from '@/stores/plan'
import { PlanStatus, PlanPriority } from '@/types/api'
import StatusBadge from '@/components/StatusBadge.vue'
import PriorityBadge from '@/components/PriorityBadge.vue'
import Loading from '@/components/Loading.vue'
import ProgressBar from '@/components/ProgressBar.vue'
import { storeToRefs } from 'pinia'

const router = useRouter()
const planStore = usePlanStore()

const { plans, loading } = storeToRefs(planStore)

// 筛选和排序状态
const filterStatus = ref<PlanStatus | ''>('')
const sortBy = ref('progress_desc')

// 统计信息
const totalPlans = computed(() => plans.value?.length || 0)
const completedPlans = computed(() => plans.value?.filter(p => p.status === PlanStatus.Done).length || 0)
const inProgressPlans = computed(() => plans.value?.filter(p => p.status === PlanStatus.InProgress).length || 0)
const averageProgress = computed(() => {
  if (!plans.value || plans.value.length === 0) return 0
  const total = plans.value.reduce((sum, plan) => sum + plan.progress, 0)
  return Math.round(total / plans.value.length)
})

// 筛选后的计划列表
const filteredPlans = computed(() => {
  let result = plans.value ? [...plans.value] : []

  if (filterStatus.value) {
    result = result.filter((plan) => plan.status === filterStatus.value)
  }

  return result
})

// 排序后的计划列表
const sortedPlans = computed(() => {
  const result = [...filteredPlans.value]

  switch (sortBy.value) {
    case 'progress_desc':
      result.sort((a, b) => b.progress - a.progress)
      break
    case 'progress_asc':
      result.sort((a, b) => a.progress - b.progress)
      break
    case 'due_date_asc':
      result.sort((a, b) => {
        if (!a.due_date) return 1
        if (!b.due_date) return -1
        return new Date(a.due_date).getTime() - new Date(b.due_date).getTime()
      })
      break
    case 'due_date_desc':
      result.sort((a, b) => {
        if (!a.due_date) return 1
        if (!b.due_date) return -1
        return new Date(b.due_date).getTime() - new Date(a.due_date).getTime()
      })
      break
  }

  return result
})

onMounted(() => {
  planStore.fetchPlans()
})

const refreshData = () => {
  planStore.fetchPlans()
}

const handleFilterChange = () => {
  // 筛选在前端完成,不需要重新请求
}

const handleSortChange = () => {
  // 排序在前端完成,不需要重新请求
}

const goToPlan = (id: number) => {
  router.push(`/plans/${id}/edit`)
}

const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
  })
}

// 判断是否已过期
const isOverdue = (dueDate: string) => {
  const now = new Date()
  const due = new Date(dueDate)
  return due < now
}

// 判断是否即将到期(7天内)
const isDueSoon = (dueDate: string) => {
  const now = new Date()
  const due = new Date(dueDate)
  const diffTime = due.getTime() - now.getTime()
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  return diffDays > 0 && diffDays <= 7
}
</script>

<style scoped>
.progress-view {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.stats-container {
  margin-bottom: 20px;
}

.stat-card {
  text-align: center;
}

.stat-content {
  padding: 10px;
}

.stat-value {
  font-size: 32px;
  font-weight: bold;
  color: #409eff;
  margin-bottom: 8px;
}

.stat-label {
  font-size: 14px;
  color: #909399;
}

.toolbar {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  padding: 15px;
  background-color: #f5f7fa;
  border-radius: 4px;
}

.progress-cards {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
  gap: 20px;
}

.progress-card {
  transition: all 0.3s;
}

.progress-card:hover {
  transform: translateY(-2px);
}

.progress-card-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 15px;
}

.plan-title {
  flex: 1;
  margin-right: 10px;
}

.title-text {
  font-size: 16px;
  font-weight: 500;
  color: #303133;
}

.plan-badges {
  display: flex;
  gap: 8px;
  margin-top: 8px;
}

.plan-actions {
  display: flex;
  gap: 8px;
}

.progress-card-body {
  padding-top: 10px;
}

.plan-description {
  margin-bottom: 15px;
  color: #606266;
  line-height: 1.6;
  max-height: 100px;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
}

.plan-meta {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 15px;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 6px;
  color: #909399;
  font-size: 13px;
}

.plan-progress {
  margin-top: 10px;
  margin-bottom: 10px;
}

.plan-due-alert {
  margin-top: 10px;
}
</style>
