<template>
  <el-dialog
    :model-value="visible"
    title="订单详情"
    width="800px"
    @update:model-value="handleClose"
    @close="handleClose"
  >
    <div v-loading="loading" class="order-detail">
      <el-descriptions :column="2" border v-if="order">
        <el-descriptions-item label="订单编号">
          {{ order.order_no }}
        </el-descriptions-item>
        <el-descriptions-item label="订单状态">
          <el-tag :type="getOrderStatusType(order.status)">
            {{ order.status_text }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="订单金额">
          <span class="price-text">¥{{ formatPrice(order.total_amount) }}</span>
        </el-descriptions-item>
        <el-descriptions-item label="创建时间">
          {{ formatDateTime(order.created_at) }}
        </el-descriptions-item>
        <el-descriptions-item label="更新时间">
          {{ formatDateTime(order.updated_at) }}
        </el-descriptions-item>
        <el-descriptions-item label="备注" :span="2">
          {{ order.remark || '无' }}
        </el-descriptions-item>
      </el-descriptions>

      <!-- 订单项列表 -->
      <div class="order-items" v-if="order && order.items && order.items.length > 0">
        <h4>订单商品</h4>
        <el-table :data="order.items" border>
          <el-table-column prop="product_name" label="商品名称" min-width="150" />
          <el-table-column prop="product_desc" label="商品描述" min-width="200" show-overflow-tooltip />
          <el-table-column prop="price" label="单价" width="120">
            <template #default="{ row }">
              <span>¥{{ formatPrice(row.price) }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="quantity" label="数量" width="80" />
          <el-table-column prop="amount" label="小计" width="120">
            <template #default="{ row }">
              <span class="price-text">¥{{ formatPrice(row.amount) }}</span>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">关闭</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { formatDateTime } from '@/utils/format'
import { getOrderDetail, formatPrice, getOrderStatusType, type OrderInfo } from '@/api/order'

// Props
interface Props {
  visible: boolean
  orderId: number | null
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
const order = ref<OrderInfo | null>(null)

// 获取订单详情
const fetchOrderDetail = async () => {
  if (!props.orderId) return

  try {
    loading.value = true
    const response = await getOrderDetail(props.orderId)

    if (response.code === 0 && response.data) {
      order.value = response.data
    } else {
      ElMessage.error(response.message || '获取订单详情失败')
    }
  } catch (error) {
    console.error('获取订单详情失败:', error)
  } finally {
    loading.value = false
  }
}

// 关闭对话框
const handleClose = () => {
  emit('update:visible', false)
  order.value = null
}

// 监听对话框显示和订单ID变化
watch(
  () => [props.visible, props.orderId],
  ([visible, orderId]) => {
    if (visible && orderId) {
      fetchOrderDetail()
    }
  },
  { immediate: true }
)
</script>

<style scoped lang="scss">
.order-detail {
  .order-items {
    margin-top: 20px;

    h4 {
      margin-bottom: 15px;
      font-size: 16px;
      font-weight: 500;
    }
  }

  .price-text {
    color: #f56c6c;
    font-weight: 500;
  }
}
</style>

