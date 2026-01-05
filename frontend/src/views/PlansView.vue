<template>
  <div class="plans-view">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>计划列表</span>
          <el-button type="primary" @click="handleCreate">新建计划</el-button>
        </div>
      </template>

      <Loading :loading="loading" text="加载中..." />

      <el-table v-if="!loading" :data="plans" stripe>
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="title" label="标题" />
        <el-table-column label="优先级" width="100">
          <template #default="{ row }">
            <PriorityBadge :priority="row.priority" />
          </template>
        </el-table-column>
        <el-table-column label="状态" width="100">
          <template #default="{ row }">
            <StatusBadge :status="row.status" />
          </template>
        </el-table-column>
        <el-table-column label="进度" width="150">
          <template #default="{ row }">
            <el-progress :percentage="row.progress" />
          </template>
        </el-table-column>
        <el-table-column prop="due_date" label="截止日期" width="180" />
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-if="!loading && total > 0"
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :total="total"
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handlePageChange"
      />
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { usePlanStore } from '@/stores/plan'
import { useUiStore } from '@/stores/ui'
import StatusBadge from '@/components/StatusBadge.vue'
import PriorityBadge from '@/components/PriorityBadge.vue'
import Loading from '@/components/Loading.vue'
import type { Plan } from '@/types/api'

const planStore = usePlanStore()
const uiStore = useUiStore()

const { plans, loading, total, currentPage, pageSize } = planStore
const { showSuccess, showError } = uiStore

onMounted(() => {
  planStore.fetchPlans()
})

const handleCreate = () => {
  uiStore.showInfo('新建计划功能将在 v0.0.2 中实现')
}

const handleEdit = (_plan: Plan) => {
  uiStore.showInfo('编辑计划功能将在 v0.0.2 中实现')
}

const handleDelete = async (row: Plan) => {
  try {
    await planStore.deletePlan(row.id)
    showSuccess('删除成功')
  } catch (error) {
    showError('删除失败')
  }
}

const handleSizeChange = (size: number) => {
  planStore.setPageSize(size)
  planStore.fetchPlans()
}

const handlePageChange = (page: number) => {
  planStore.setPage(page)
  planStore.fetchPlans()
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

.el-pagination {
  margin-top: 20px;
  display: flex;
  justify-content: center;
}
</style>
