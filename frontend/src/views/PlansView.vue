<template>
  <div class="plans-view">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>计划列表</span>
          <el-button type="primary" @click="handleCreate">新建计划</el-button>
        </div>
      </template>

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
          v-model="filterPriority"
          placeholder="筛选优先级"
          clearable
          style="width: 150px; margin-right: 10px"
          @change="handleFilterChange"
        >
          <el-option label="高" :value="PlanPriority.High" />
          <el-option label="中" :value="PlanPriority.Medium" />
          <el-option label="低" :value="PlanPriority.Low" />
        </el-select>

        <el-select
          v-model="sortBy"
          placeholder="排序方式"
          style="width: 180px"
          @change="handleSortChange"
        >
          <el-option label="创建时间(最新)" value="created_at_desc" />
          <el-option label="创建时间(最早)" value="created_at_asc" />
          <el-option label="截止日期(最近)" value="due_date_asc" />
          <el-option label="截止日期(最远)" value="due_date_desc" />
        </el-select>
      </div>

      <!-- 加载状态 -->
      <Loading v-if="loading" loading text="加载中..." />

      <!-- 空状态 -->
      <el-empty v-else-if="plans.length === 0" description="暂无计划数据" />

      <!-- 计划卡片列表 -->
      <div v-else class="plan-cards">
        <el-card v-for="plan in sortedPlans" :key="plan.id" class="plan-card" shadow="hover">
          <template #header>
            <div class="plan-card-header">
              <div class="plan-title">
                <span class="title-text">{{ plan.title }}</span>
                <div class="plan-badges">
                  <PriorityBadge :priority="plan.priority" />
                  <StatusBadge :status="plan.status" />
                </div>
              </div>
              <div class="plan-actions">
                <el-button size="small" @click="handleEdit(plan)">编辑</el-button>
                <el-button size="small" type="danger" @click="handleDelete(plan)">
                  删除
                </el-button>
              </div>
            </div>
          </template>

          <div class="plan-card-body">
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
              <div class="meta-item">
                <el-icon><TrendCharts /></el-icon>
                <span>进度: {{ plan.progress }}%</span>
              </div>
            </div>

            <div class="plan-progress">
              <el-progress :percentage="plan.progress" :stroke-width="8" />
            </div>
          </div>
        </el-card>
      </div>

      <!-- 分页 -->
      <el-pagination
        v-if="!loading && total > 0"
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :total="total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        class="pagination"
        @size-change="handleSizeChange"
        @current-change="handlePageChange"
      />
    </el-card>

    <!-- 计划表单对话框 -->
    <PlanForm
      v-model="formVisible"
      :plan="currentEditPlan"
      @success="handleFormSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { ElMessageBox } from 'element-plus'
import { Calendar, Clock, TrendCharts } from '@element-plus/icons-vue'
import { usePlanStore } from '@/stores/plan'
import { useUiStore } from '@/stores/ui'
import { PlanStatus, PlanPriority } from '@/types/api'
import StatusBadge from '@/components/StatusBadge.vue'
import PriorityBadge from '@/components/PriorityBadge.vue'
import Loading from '@/components/Loading.vue'
import PlanForm from './PlanForm.vue'
import type { Plan } from '@/types/api'

const route = useRoute()
const planStore = usePlanStore()
const uiStore = useUiStore()

const { plans, loading, total, currentPage, pageSize } = planStore

// 筛选和排序状态
const filterStatus = ref<PlanStatus | ''>('')
const filterPriority = ref<PlanPriority | ''>('')
const sortBy = ref('created_at_desc')

// 表单对话框状态
const formVisible = ref(false)
const currentEditPlan = ref<Plan | null>(null)

// 筛选后的计划列表
const filteredPlans = computed(() => {
  let result = [...plans]

  if (filterStatus.value) {
    result = result.filter((plan) => plan.status === filterStatus.value)
  }

  if (filterPriority.value) {
    result = result.filter((plan) => plan.priority === filterPriority.value)
  }

  return result
})

// 排序后的计划列表
const sortedPlans = computed(() => {
  const result = [...filteredPlans.value]

  switch (sortBy.value) {
    case 'created_at_desc':
      result.sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime())
      break
    case 'created_at_asc':
      result.sort((a, b) => new Date(a.created_at).getTime() - new Date(b.created_at).getTime())
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
  handleRouteChange()
})

// 监听路由变化
watch(
  () => route.path,
  () => {
    handleRouteChange()
  }
)

// 处理路由变化
const handleRouteChange = async () => {
  if (route.name === 'PlanCreate') {
    currentEditPlan.value = null
    formVisible.value = true
  } else if (route.name === 'PlanEdit') {
    const planId = Number(route.params.id)
    if (planId) {
      try {
        const plan = await planStore.fetchPlan(planId)
        currentEditPlan.value = plan
        formVisible.value = true
      } catch (error) {
        uiStore.showError('获取计划详情失败')
      }
    }
  }
}

const handleCreate = () => {
  currentEditPlan.value = null
  formVisible.value = true
}

const handleEdit = (plan: Plan) => {
  currentEditPlan.value = plan
  formVisible.value = true
}

const handleDelete = async (plan: Plan) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除计划 "${plan.title}" 吗?`,
      '删除确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
    )

    await planStore.deletePlan(plan.id)
    uiStore.showSuccess('删除成功')
  } catch (error: any) {
    if (error !== 'cancel') {
      uiStore.showError('删除失败')
    }
  }
}

const handleFilterChange = () => {
  planStore.setFilters({
    status: filterStatus.value || undefined,
    priority: filterPriority.value || undefined,
  })
  planStore.fetchPlans()
}

const handleSortChange = () => {
  // 排序在前端完成,不需要重新请求
}

const handleSizeChange = (size: number) => {
  planStore.setPageSize(size)
  planStore.fetchPlans()
}

const handlePageChange = (page: number) => {
  planStore.setPage(page)
  planStore.fetchPlans()
}

const handleFormSuccess = () => {
  planStore.fetchPlans()
  // 关闭对话框并重置路由
  formVisible.value = false
  currentEditPlan.value = null
  if (route.name === 'PlanCreate' || route.name === 'PlanEdit') {
    // 使用 replace 而不是 push,避免在历史记录中留下创建/编辑路由
    window.history.replaceState({}, '', '/plans')
  }
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
</script>

<style scoped>
.plans-view {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.toolbar {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
  padding: 15px;
  background-color: #f5f7fa;
  border-radius: 4px;
}

.plan-cards {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
  gap: 20px;
  margin-bottom: 20px;
}

.plan-card {
  transition: all 0.3s;
}

.plan-card:hover {
  transform: translateY(-2px);
}

.plan-card-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
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

.plan-card-body {
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
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: center;
}
</style>
