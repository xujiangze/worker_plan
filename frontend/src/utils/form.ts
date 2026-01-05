import { ElMessage } from 'element-plus'

/**
 * 显示表单验证错误
 * @param errors 表单验证错误对象
 */
export const showFormErrors = (errors: Record<string, string[]>) => {
  const messages = Object.values(errors).flat()
  if (messages.length > 0) {
    ElMessage.error(messages[0])
  }
}

/**
 * 显示单个字段验证错误
 * @param field 字段名
 * @param error 错误信息
 */
export const showFieldError = (field: string, error: string) => {
  ElMessage.error(`${field}: ${error}`)
}

/**
 * 清除表单验证错误
 * @param formRef 表单引用
 */
export const clearFormErrors = (formRef: any) => {
  if (formRef && formRef.clearValidate) {
    formRef.clearValidate()
  }
}
