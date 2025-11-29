<template>
  <div class="cart-container">
    <!-- 购物车信息卡片 -->
    <el-card class="cart-info-card" v-if="cartInfo">
      <template #header>
        <div class="card-header">
          <span>购物车信息</span>
          <div class="cart-summary">
            <span>商品种类：{{ cartInfo.item_count }} 种</span>
            <span class="total-amount">总金额：¥{{ formatPrice(cartInfo.total_amount) }}</span>
          </div>
        </div>
      </template>
      <div class="cart-actions">
        <el-button type="danger" @click="handleClearCart" :loading="loading">
          <el-icon><Delete /></el-icon>
          清空购物车
        </el-button>
      </div>
    </el-card>

    <!-- 购物车项表格 -->
    <el-table
      :data="cartItems"
      v-loading="loading"
      stripe
      style="width: 100%"
      empty-text="购物车为空"
    >
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="product_name" label="商品名称" min-width="200" />
      <el-table-column prop="price" label="单价" width="120">
        <template #default="{ row }">
          <span class="price-text">¥{{ formatPrice(row.price) }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="quantity" label="数量" width="200">
        <template #default="{ row }">
          <el-input-number
            v-model="row.quantity"
            :min="1"
            :max="999"
            :disabled="updatingItemId === row.id"
            @change="(val) => handleQuantityChange(row, val)"
            style="width: 120px"
          />
        </template>
      </el-table-column>
      <el-table-column prop="amount" label="小计" width="120">
        <template #default="{ row }">
          <span class="price-text">¥{{ formatPrice(row.amount) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="120" fixed="right">
        <template #default="{ row }">
          <el-button
            type="danger"
            size="small"
            @click="handleDeleteItem(row)"
            :loading="deletingItemId === row.id"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 空状态 -->
    <el-empty v-if="!loading && cartItems.length === 0" description="购物车为空" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Delete } from '@element-plus/icons-vue'
import {
  getCart,
  updateCartItem,
  deleteCartItem,
  clearCart,
  formatPrice,
  type CartInfo,
  type CartItemInfo
} from '@/api/cart'

// 响应式数据
const loading = ref(false)
const cartInfo = ref<CartInfo | null>(null)
const cartItems = ref<CartItemInfo[]>([])
const updatingItemId = ref<number | null>(null)
const deletingItemId = ref<number | null>(null)

// 获取购物车数据
const fetchCart = async () => {
  try {
    loading.value = true
    const response = await getCart()

    if (response.code === 0) {
      cartInfo.value = response.data
      cartItems.value = response.data.items || []
    } else {
      ElMessage.error(response.message || '获取购物车失败')
    }
  } catch (error) {
    console.error('获取购物车失败:', error)
  } finally {
    loading.value = false
  }
}

// 处理数量变化
const handleQuantityChange = async (item: CartItemInfo, newQuantity: number | null) => {
  if (newQuantity === null) {
    return
  }

  // 保存原始数量
  const originalQuantity = item.quantity

  // 验证数量
  if (newQuantity < 1) {
    ElMessage.warning('数量不能小于1')
    // 恢复原值
    item.quantity = originalQuantity
    return
  }

  if (newQuantity > 999) {
    ElMessage.warning('数量不能超过999')
    // 恢复原值
    item.quantity = originalQuantity
    return
  }

  // 如果数量没有变化，不需要更新
  if (newQuantity === originalQuantity) {
    return
  }

  try {
    updatingItemId.value = item.id
    const response = await updateCartItem(item.id, { quantity: newQuantity })

    if (response.code === 0) {
      // 更新本地数据
      item.quantity = response.data.quantity
      item.amount = response.data.amount
      // 重新获取购物车以更新总金额
      await fetchCart()
      ElMessage.success('更新成功')
    } else {
      ElMessage.error(response.message || '更新失败')
      // 恢复原值
      item.quantity = originalQuantity
      await fetchCart()
    }
  } catch (error) {
    console.error('更新购物车项失败:', error)
    // 恢复原值
    item.quantity = originalQuantity
    await fetchCart()
  } finally {
    updatingItemId.value = null
  }
}

// 删除购物车项
const handleDeleteItem = async (item: CartItemInfo) => {
  try {
    await ElMessageBox.confirm(`确定要删除商品"${item.product_name}"吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })

    deletingItemId.value = item.id
    const response = await deleteCartItem(item.id)

    if (response.code === 0) {
      ElMessage.success('删除成功')
      // 重新获取购物车
      await fetchCart()
    } else {
      ElMessage.error(response.message || '删除失败')
    }
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除购物车项失败:', error)
    }
  } finally {
    deletingItemId.value = null
  }
}

// 清空购物车
const handleClearCart = async () => {
  try {
    await ElMessageBox.confirm('确定要清空购物车吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })

    loading.value = true
    const response = await clearCart()

    if (response.code === 0) {
      ElMessage.success('清空成功')
      // 重新获取购物车
      await fetchCart()
    } else {
      ElMessage.error(response.message || '清空失败')
    }
  } catch (error) {
    if (error !== 'cancel') {
      console.error('清空购物车失败:', error)
    }
  } finally {
    loading.value = false
  }
}

// 初始化
onMounted(() => {
  fetchCart()
})
</script>

<style scoped lang="scss">
.cart-container {
  padding: 20px;
}

.cart-info-card {
  margin-bottom: 20px;

  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;

    .cart-summary {
      display: flex;
      gap: 20px;
      align-items: center;

      .total-amount {
        font-size: 18px;
        font-weight: bold;
        color: #f56c6c;
      }
    }
  }

  .cart-actions {
    margin-top: 10px;
  }
}

.price-text {
  color: #f56c6c;
  font-weight: 500;
}
</style>

