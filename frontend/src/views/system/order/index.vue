<template>
  <div class="order-container">
    <!-- 搜索栏 -->
    <div class="search-bar">
      <el-form :model="searchForm" inline>
        <el-form-item label="订单编号">
          <el-input
            v-model="searchForm.keyword"
            placeholder="请输入订单编号"
            clearable
            @keyup.enter="handleSearch"
            style="width: 200px"
          />
        </el-form-item>
        <el-form-item label="订单状态">
          <el-select
            v-model="searchForm.status"
            placeholder="请选择订单状态"
            clearable
            style="width: 150px"
          >
            <el-option label="待支付" :value="1" />
            <el-option label="已支付" :value="2" />
            <el-option label="已发货" :value="3" />
            <el-option label="已完成" :value="4" />
            <el-option label="已取消" :value="5" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch" :loading="loading">
            <el-icon><Search /></el-icon>
            搜索
          </el-button>
          <el-button @click="handleReset">
            <el-icon><Refresh /></el-icon>
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 订单表格 -->
    <el-table
      :data="orderList"
      v-loading="loading"
      stripe
      style="width: 100%"
    >
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="order_no" label="订单编号" min-width="180" />
      <el-table-column prop="total_amount" label="订单金额" width="120">
        <template #default="{ row }">
          <span class="price-text">¥{{ formatPrice(row.total_amount) }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="status" label="订单状态" width="120">
        <template #default="{ row }">
          <el-tag :type="getOrderStatusType(row.status)">
            {{ row.status_text }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="remark" label="备注" min-width="150" show-overflow-tooltip />
      <el-table-column prop="created_at" label="创建时间" width="180">
        <template #default="{ row }">
          {{ formatDateTime(row.created_at) }}
        </template>
      </el-table-column>
      <el-table-column label="操作" width="250" fixed="right">
        <template #default="{ row }">
          <el-button type="primary" size="small" @click="handleViewDetail(row)">
            详情
          </el-button>
          <el-button
            v-if="row.status === 1"
            type="success"
            size="small"
            @click="handleUpdateStatus(row, 2)"
          >
            支付
          </el-button>
          <el-button
            v-if="row.status === 1 || row.status === 2"
            type="warning"
            size="small"
            @click="handleUpdateStatus(row, 5)"
          >
            取消
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <div class="pagination">
      <el-pagination
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="pagination.total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>

    <!-- 订单详情对话框 -->
    <OrderDetailDialog
      v-model:visible="detailDialogVisible"
      :order-id="currentOrderId"
      @success="handleDialogSuccess"
    />

    <!-- 订单状态更新对话框 -->
    <OrderStatusDialog
      v-model:visible="statusDialogVisible"
      :order="currentOrder"
      :target-status="targetStatus"
      @success="handleDialogSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Search, Refresh } from '@element-plus/icons-vue'
import { formatDateTime } from '@/utils/format'
import OrderDetailDialog from './components/OrderDetailDialog.vue'
import OrderStatusDialog from './components/OrderStatusDialog.vue'
import {
  getOrderList,
  updateOrderStatus,
  formatPrice,
  getOrderStatusType,
  type OrderInfo
} from '@/api/order'

// 响应式数据
const loading = ref(false)
const orderList = ref<OrderInfo[]>([])
const currentOrderId = ref<number | null>(null)
const currentOrder = ref<OrderInfo | null>(null)
const targetStatus = ref<number>(0)
const detailDialogVisible = ref(false)
const statusDialogVisible = ref(false)

// 搜索表单
const searchForm = reactive({
  keyword: '',
  status: undefined as number | undefined
})

// 分页
const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

// 获取订单列表
const fetchOrderList = async () => {
  try {
    loading.value = true
    const params = {
      page: pagination.page,
      page_size: pagination.pageSize,
      keyword: searchForm.keyword || undefined,
      status: searchForm.status
    }

    const response = await getOrderList(params)

    if (response.code === 0 && response.data) {
      orderList.value = response.data.orders
      pagination.total = response.data.total
    } else {
      ElMessage.error(response.message || '获取订单列表失败')
    }
  } catch (error) {
    console.error('获取订单列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  fetchOrderList()
}

// 重置
const handleReset = () => {
  searchForm.keyword = ''
  searchForm.status = undefined
  pagination.page = 1
  fetchOrderList()
}

// 查看详情
const handleViewDetail = (order: OrderInfo) => {
  currentOrderId.value = order.id
  detailDialogVisible.value = true
}

// 更新订单状态
const handleUpdateStatus = (order: OrderInfo, status: number) => {
  currentOrder.value = order
  targetStatus.value = status
  statusDialogVisible.value = true
}

// 分页大小改变
const handleSizeChange = (size: number) => {
  pagination.pageSize = size
  pagination.page = 1
  fetchOrderList()
}

// 当前页改变
const handleCurrentChange = (page: number) => {
  pagination.page = page
  fetchOrderList()
}

// 对话框成功回调
const handleDialogSuccess = () => {
  fetchOrderList()
}

// 初始化
onMounted(() => {
  fetchOrderList()
})
</script>

<style scoped lang="scss">
.order-container {
  padding: 20px;

  .search-bar {
    margin-bottom: 20px;
    padding: 20px;
    background: #fff;
    border-radius: 4px;
  }

  .pagination {
    margin-top: 20px;
    display: flex;
    justify-content: flex-end;
  }

  .price-text {
    color: #f56c6c;
    font-weight: 500;
  }
}
</style>

