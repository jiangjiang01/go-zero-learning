<template>
  <el-dialog
    :model-value="visible"
    :title="dialogTitle"
    width="500px"
    @update:model-value="handleClose"
    @close="handleClose"
  >
    <div v-if="order">
      <el-alert
        :title="alertTitle"
        :description="alertDescription"
        type="warning"
        :closable="false"
        show-icon
        style="margin-bottom: 20px"
      />
      <el-descriptions :column="1" border>
        <el-descriptions-item label="订单编号">
          {{ order.order_no }}
        </el-descriptions-item>
        <el-descriptions-item label="当前状态">
          <el-tag :type="getOrderStatusType(order.status)">
            {{ order.status_text }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="订单金额">
          <span class="price-text">¥{{ formatPrice(order.total_amount) }}</span>
        </el-descriptions-item>
      </el-descriptions>
    </div>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="loading">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { updateOrderStatus, formatPrice, getOrderStatusType, getOrderStatusText, type OrderInfo } from '@/api/order'

// Props
interface Props {
  visible: boolean
  order: OrderInfo | null
  targetStatus: number
}

const props = defineProps<Props>()

// Emits
interface Emits {
  (e: 'update:visible', visible: boolean): void
  (e: 'success'): void
}

const emit = defineEmits<Emits>()

// 响应式数据
const loading = ref(false)

// 计算属性
const dialogTitle = computed(() => {
  if (!props.order || !props.targetStatus) return '更新订单状态'
  const targetText = getOrderStatusText(props.targetStatus)
  return `确认${targetText}订单`
})

const alertTitle = computed(() => {
  if (!props.order || !props.targetStatus) return '确认操作'
  const targetText = getOrderStatusText(props.targetStatus)
  return `确认要将订单状态更新为「${targetText}」吗？`
})

const alertDescription = computed(() => {
  if (!props.order) return ''
  const currentText = props.order.status_text
  const targetText = getOrderStatusText(props.targetStatus)
  return `当前状态：${currentText} → 目标状态：${targetText}`
})

// 提交
const handleSubmit = async () => {
  if (!props.order || !props.targetStatus) return

  try {
    loading.value = true
    const response = await updateOrderStatus(props.order.id, {
      status: props.targetStatus
    })

    if (response.code === 0) {
      ElMessage.success('订单状态更新成功')
      emit('success')
      handleClose()
    } else {
      ElMessage.error(response.message || '订单状态更新失败')
    }
  } catch (error) {
    console.error('订单状态更新失败:', error)
  } finally {
    loading.value = false
  }
}

// 关闭对话框
const handleClose = () => {
  emit('update:visible', false)
}
</script>

<style scoped lang="scss">
.price-text {
  color: #f56c6c;
  font-weight: 500;
}
</style>

