<template>
  <div class="dashboard-container p-4">
    <h1 class="text-2xl font-bold mb-6">数据统计 Dashboard</h1>

    <!-- 加载状态 -->
    <div v-if="loading" class="text-center py-20">
      <el-icon class="is-loading" :size="40">
        <Loading />
      </el-icon>
      <p class="mt-4 text-gray-500">加载中...</p>
    </div>

    <!-- 统计数据 -->
    <div v-else>
      <!-- 统计卡片 -->
      <el-row :gutter="20" class="mb-6">
        <!-- 订单统计 -->
        <el-col :xs="24" :sm="12" :md="6">
          <el-card shadow="hover" class="stat-card">
            <div class="stat-content">
              <div class="stat-icon order-icon">
                <Document :size="32" />
              </div>
              <div class="stat-info">
                <div class="stat-value">{{ stats.order_stats?.total_orders || 0 }}</div>
                <div class="stat-label">订单总数</div>
              </div>
            </div>
          </el-card>
        </el-col>

        <el-col :xs="24" :sm="12" :md="6">
          <el-card shadow="hover" class="stat-card">
            <div class="stat-content">
              <div class="stat-icon amount-icon">
                <Money :size="32" />
              </div>
              <div class="stat-info">
                <div class="stat-value">¥{{ formatPrice(stats.order_stats?.total_amount || 0) }}</div>
                <div class="stat-label">总销售额</div>
              </div>
            </div>
          </el-card>
        </el-col>

        <el-col :xs="24" :sm="12" :md="6">
          <el-card shadow="hover" class="stat-card">
            <div class="stat-content">
              <div class="stat-icon product-icon">
                <Box :size="32" />
              </div>
              <div class="stat-info">
                <div class="stat-value">{{ stats.product_stats?.total_products || 0 }}</div>
                <div class="stat-label">商品总数</div>
              </div>
            </div>
          </el-card>
        </el-col>

        <el-col :xs="24" :sm="12" :md="6">
          <el-card shadow="hover" class="stat-card">
            <div class="stat-content">
              <div class="stat-icon user-icon">
                <User :size="32" />
              </div>
              <div class="stat-info">
                <div class="stat-value">{{ stats.user_stats?.total_users || 0 }}</div>
                <div class="stat-label">用户总数</div>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>

      <!-- 详细统计信息 -->
      <el-row :gutter="20" class="mb-6">
        <!-- 订单统计详情 -->
        <el-col :xs="24" :md="12">
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">
                <Document :size="18" />
                <span>订单统计</span>
              </div>
            </template>
            <div class="detail-stats">
              <div class="detail-item">
                <span class="detail-label">今日订单：</span>
                <span class="detail-value">{{ stats.order_stats?.today_orders || 0 }}</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">今日销售额：</span>
                <span class="detail-value">¥{{ formatPrice(stats.order_stats?.today_amount || 0) }}</span>
              </div>
              <div class="detail-item">
                <span class="detail-label">待支付：</span>
                <el-tag type="warning" size="small">{{ stats.order_stats?.status_count?.pending || 0 }}</el-tag>
              </div>
              <div class="detail-item">
                <span class="detail-label">已支付：</span>
                <el-tag type="success" size="small">{{ stats.order_stats?.status_count?.paid || 0 }}</el-tag>
              </div>
              <div class="detail-item">
                <span class="detail-label">已发货：</span>
                <el-tag type="info" size="small">{{ stats.order_stats?.status_count?.shipped || 0 }}</el-tag>
              </div>
              <div class="detail-item">
                <span class="detail-label">已完成：</span>
                <el-tag type="success" size="small">{{ stats.order_stats?.status_count?.completed || 0 }}</el-tag>
              </div>
              <div class="detail-item">
                <span class="detail-label">已取消：</span>
                <el-tag type="danger" size="small">{{ stats.order_stats?.status_count?.canceled || 0 }}</el-tag>
              </div>
            </div>
          </el-card>
        </el-col>

        <!-- 商品统计详情 -->
        <el-col :xs="24" :md="12">
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">
                <Box :size="18" />
                <span>商品统计</span>
              </div>
            </template>
            <div class="detail-stats">
              <div class="detail-item">
                <span class="detail-label">上架商品：</span>
                <el-tag type="success" size="small">{{ stats.product_stats?.onsale_products || 0 }}</el-tag>
              </div>
              <div class="detail-item">
                <span class="detail-label">下架商品：</span>
                <el-tag type="danger" size="small">{{ stats.product_stats?.offsale_products || 0 }}</el-tag>
              </div>
              <div class="detail-item">
                <span class="detail-label">库存预警：</span>
                <el-tag type="warning" size="small">{{ stats.product_stats?.low_stock_products || 0 }}</el-tag>
              </div>
              <div class="detail-item">
                <span class="detail-label">总库存量：</span>
                <span class="detail-value">{{ stats.product_stats?.total_stock || 0 }}</span>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>

      <!-- 用户统计详情 -->
      <el-row :gutter="20" class="mb-6">
        <el-col :xs="24" :md="12">
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">
                <User :size="18" />
                <span>用户统计</span>
              </div>
            </template>
            <div class="detail-stats">
              <div class="detail-item">
                <span class="detail-label">今日新增：</span>
                <el-tag type="success" size="small">{{ stats.user_stats?.today_users || 0 }}</el-tag>
              </div>
              <div class="detail-item">
                <span class="detail-label">活跃用户：</span>
                <el-tag type="info" size="small">{{ stats.user_stats?.active_users || 0 }}</el-tag>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>

      <!-- 图表区域 -->
      <el-row :gutter="20">
        <!-- 订单状态分布饼图 -->
        <el-col :xs="24" :md="12">
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">订单状态分布</div>
            </template>
            <div ref="orderStatusChartRef" style="width: 100%; height: 300px;"></div>
          </el-card>
        </el-col>

        <!-- 商品状态分布饼图 -->
        <el-col :xs="24" :md="12">
          <el-card shadow="hover">
            <template #header>
              <div class="card-header">商品状态分布</div>
            </template>
            <div ref="productStatusChartRef" style="width: 100%; height: 300px;"></div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, nextTick } from 'vue'
import * as echarts from 'echarts'
import { getDashboardStats, formatPrice, type DashboardStatsResponse } from '@/api/dashboard'
import { ElMessage } from 'element-plus'
import {
  User,
  Document,
  Box,
  Money,
  Loading
} from '@element-plus/icons-vue'

const loading = ref(true)
const stats = ref<DashboardStatsResponse>({
  order_stats: {
    total_orders: 0,
    today_orders: 0,
    total_amount: 0,
    today_amount: 0,
    status_count: {
      pending: 0,
      paid: 0,
      shipped: 0,
      completed: 0,
      canceled: 0
    }
  },
  product_stats: {
    total_products: 0,
    onsale_products: 0,
    offsale_products: 0,
    low_stock_products: 0,
    total_stock: 0
  },
  user_stats: {
    total_users: 0,
    today_users: 0,
    active_users: 0
  }
})

const orderStatusChartRef = ref<HTMLDivElement>()
const productStatusChartRef = ref<HTMLDivElement>()

let orderStatusChart: echarts.ECharts | null = null
let productStatusChart: echarts.ECharts | null = null

// 获取统计数据
const fetchStats = async () => {
  try {
    loading.value = true
    const response = await getDashboardStats()
    if (response.code === 0) {
      stats.value = response.data
      // 更新图表
      nextTick(() => {
        updateCharts()
      })
    } else {
      ElMessage.error(response.message || '获取统计数据失败')
    }
  } catch (error) {
    console.error('获取统计数据失败:', error)
  } finally {
    loading.value = false
  }
}

// 更新图表
const updateCharts = () => {
  // 订单状态分布饼图
  if (orderStatusChartRef.value && stats.value.order_stats) {
    if (!orderStatusChart) {
      orderStatusChart = echarts.init(orderStatusChartRef.value)
    }

    const option = {
      tooltip: {
        trigger: 'item',
        formatter: '{a} <br/>{b}: {c} ({d}%)'
      },
      legend: {
        orient: 'vertical',
        left: 'left',
        top: 'middle'
      },
      series: [
        {
          name: '订单状态',
          type: 'pie',
          radius: ['40%', '70%'],
          avoidLabelOverlap: false,
          itemStyle: {
            borderRadius: 10,
            borderColor: '#fff',
            borderWidth: 2
          },
          label: {
            show: false,
            position: 'center'
          },
          emphasis: {
            label: {
              show: true,
              fontSize: '20',
              fontWeight: 'bold'
            }
          },
          labelLine: {
            show: false
          },
          data: [
            { value: stats.value.order_stats.status_count.pending, name: '待支付' },
            { value: stats.value.order_stats.status_count.paid, name: '已支付' },
            { value: stats.value.order_stats.status_count.shipped, name: '已发货' },
            { value: stats.value.order_stats.status_count.completed, name: '已完成' },
            { value: stats.value.order_stats.status_count.canceled, name: '已取消' }
          ]
        }
      ]
    }

    orderStatusChart.setOption(option)
  }

  // 商品状态分布饼图
  if (productStatusChartRef.value && stats.value.product_stats) {
    if (!productStatusChart) {
      productStatusChart = echarts.init(productStatusChartRef.value)
    }

    const option = {
      tooltip: {
        trigger: 'item',
        formatter: '{a} <br/>{b}: {c} ({d}%)'
      },
      legend: {
        orient: 'vertical',
        left: 'left',
        top: 'middle'
      },
      series: [
        {
          name: '商品状态',
          type: 'pie',
          radius: ['40%', '70%'],
          avoidLabelOverlap: false,
          itemStyle: {
            borderRadius: 10,
            borderColor: '#fff',
            borderWidth: 2
          },
          label: {
            show: false,
            position: 'center'
          },
          emphasis: {
            label: {
              show: true,
              fontSize: '20',
              fontWeight: 'bold'
            }
          },
          labelLine: {
            show: false
          },
          data: [
            { value: stats.value.product_stats.onsale_products, name: '上架' },
            { value: stats.value.product_stats.offsale_products, name: '下架' }
          ]
        }
      ]
    }

    productStatusChart.setOption(option)
  }
}

// 窗口大小变化时调整图表
const handleResize = () => {
  orderStatusChart?.resize()
  productStatusChart?.resize()
}

onMounted(() => {
  fetchStats()
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
  orderStatusChart?.dispose()
  productStatusChart?.dispose()
})
</script>

<style scoped>
.dashboard-container {
  min-height: calc(100vh - 84px);
}

.stat-card {
  height: 100%;
}

.stat-content {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 8px 0;
}

.stat-icon {
  width: 64px;
  height: 64px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}

.order-icon {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.amount-icon {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
}

.product-icon {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.user-icon {
  background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
}

.stat-info {
  flex: 1;
}

.stat-value {
  font-size: 28px;
  font-weight: bold;
  color: #303133;
  margin-bottom: 4px;
  line-height: 1.2;
}

.stat-label {
  font-size: 14px;
  color: #909399;
}

.card-header {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: bold;
  font-size: 16px;
}

.detail-stats {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.detail-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8px 0;
  border-bottom: 1px solid #f0f0f0;
}

.detail-item:last-child {
  border-bottom: none;
}

.detail-label {
  color: #606266;
  font-size: 14px;
}

.detail-value {
  color: #303133;
  font-weight: 500;
  font-size: 14px;
}
</style>
