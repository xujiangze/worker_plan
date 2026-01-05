<template>
  <el-dialog
    v-model="visible"
    :title="isEdit ? '编辑计划' : '新建计划'"
    width="600px"
    @close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-width="100px"
    >
      <el-form-item label="标题" prop="title">
        <el-input
          v-model="formData.title"
          placeholder="请输入计划标题"
          maxlength="255"
          show-word-limit
        />
      </el-form-item>

      <el-form-item label="描述" prop="description">
        <el-input
          v-model="formData.description"
          type="textarea"
          :rows="4"
          placeholder="请输入计划描述"
          maxlength="1000"
          show-word-limit
        />
      </el-form-item>

      <el-form-item label="优先级" prop="priority">
        <el-select v-model="formData.priority" placeholder="请选择优先级" style="width: 100%">
          <el-option label="高" :value="PlanPriority.High" />
          <el-option label="中" :value="PlanPriority.Medium" />
          <el-option label="低" :value="PlanPriority.Low" />
        </el-select>
      </el-form-item>

      <el-form-item label="截止日期" prop="due_date">
        <el-date-picker
          v-model="formData.due_date"
          type="datetime"
          placeholder="请选择截止日期"
          format="YYYY-MM-DD HH:mm:ss"
          value-format="YYYY-MM-DD HH:mm:ss"
          style="width: 100%"
        />
      </el-form-item>
    </el-form>

    <template #footer>
      <el-button @click="handleClose">取消</el-button>
      <el-button type="primary" :loading="loading" @click="handleSubmit">
        保存
      </el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, watch, computed } from 'vue'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import { PlanPriority } from '@/types/api'
import { usePlanStore } from '@/stores/plan'
import { useUiStore } from '@/stores/ui'
import type { Plan, CreatePlanRequest, UpdatePlanRequest } from '@/types/api'

interface Props {
  modelValue: boolean
  plan?: Plan | null
}

interface Emits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'success'): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const planStore = usePlanStore()
const uiStore = useUiStore()

const formRef = ref<FormInstance>()
const loading = ref(false)

const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value),
})

const isEdit = computed(() => !!props.plan)

const formData = reactive<CreatePlanRequest>({
  title: '',
  description: '',
  priority: PlanPriority.Medium,
  due_date: '',
})

const formRules: FormRules = {
  title: [
    { required: true, message: '请输入计划标题', trigger: 'blur' },
    { min: 1, max: 255, message: '标题长度在 1 到 255 个字符', trigger: 'blur' },
  ],
  priority: [
    { required: true, message: '请选择优先级', trigger: 'change' },
  ],
}

// 监听 plan 变化,预填充表单
watch(
  () => props.plan,
  (plan) => {
    if (plan) {
      formData.title = plan.title
      formData.description = plan.description || ''
      formData.priority = plan.priority
      formData.due_date = plan.due_date || ''
    } else {
      resetForm()
    }
  },
  { immediate: true }
)

const resetForm = () => {
  formData.title = ''
  formData.description = ''
  formData.priority = PlanPriority.Medium
  formData.due_date = ''
  formRef.value?.clearValidate()
}

const handleClose = () => {
  visible.value = false
  resetForm()
}

const handleSubmit = async () => {
  if (!formRef.value) return

  try {
    await formRef.value.validate()
    loading.value = true

    if (isEdit.value && props.plan) {
      // 编辑模式
      const updateData: UpdatePlanRequest = {
        title: formData.title,
        description: formData.description,
        priority: formData.priority,
        due_date: formData.due_date || undefined,
      }
      await planStore.updatePlan(props.plan.id, updateData)
      uiStore.showSuccess('编辑成功')
    } else {
      // 创建模式
      const createData: CreatePlanRequest = {
        title: formData.title,
        description: formData.description,
        priority: formData.priority,
        due_date: formData.due_date || undefined,
      }
      await planStore.createPlan(createData)
      uiStore.showSuccess('创建成功')
    }

    emit('success')
    handleClose()
  } catch (error: any) {
    if (error !== false) {
      uiStore.showError(isEdit.value ? '编辑失败' : '创建失败')
    }
  } finally {
    loading.value = false
  }
}
</script>
