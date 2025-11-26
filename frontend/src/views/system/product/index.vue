<template>
  <div class="product-container">
    <!-- 搜索栏 -->
    <div class="search-bar">
      <el-form :model="searchForm" inline>
        <el-form-item label="商品名称">
          <el-input
            v-model="searchForm.keyword"
            placeholder="请输入商品名称"
            clearable
            @keyup.enter="handleSearch"
            style="width: 200px"
          />
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

    <!-- 操作栏 -->
    <div class="toolbar">
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增商品
      </el-button>
    </div>

    <!-- 商品表格 -->
    <el-table
      :data="productList"
      v-loading="loading"
      stripe
      style="width: 100%"
    >
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="name" label="商品名称" min-width="150" />
      <el-table-column prop="description" label="商品描述" min-width="200" show-overflow-tooltip />
      <el-table-column prop="price" label="价格" width="120">
        <template #default="{ row }">
          <span class="price-text">¥{{ formatPrice(row.price) }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="100">
        <template #default="{ row }">
          <el-tag :type="getProductStatusType(row.status)">
            {{ getProductStatusText(row.status) }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="创建时间" width="180">
        <template #default="{ row }">
          {{ formatDateTime(row.created_at) }}
        </template>
      </el-table-column>
      <el-table-column label="操作" width="200" fixed="right">
        <template #default="{ row }">
          <el-button type="primary" size="small" @click="handleEdit(row)">
            编辑
          </el-button>
          <el-button type="danger" size="small" @click="handleDelete(row)">
            删除
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

    <!-- 商品对话框 -->
    <ProductDialog
      v-model:visible="dialogVisible"
      :product="currentProduct"
      @success="handleDialogSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Search, Refresh, Plus } from '@element-plus/icons-vue'
import ProductDialog from './components/ProductDialog.vue'
import { 
  getProductList, 
  deleteProduct, 
  formatPrice, 
  getProductStatusText, 
  getProductStatusType,
  type ProductInfo 
} from '@/api/product'
import { formatDateTime } from '@/utils/format'

// 响应式数据
const loading = ref(false)
const productList = ref<ProductInfo[]>([])
const dialogVisible = ref(false)
const currentProduct = ref<ProductInfo | null>(null)

// 搜索表单
const searchForm = reactive({
  keyword: ''
})

// 分页数据
const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

// 获取商品列表
const fetchProductList = async () => {
  try {
    loading.value = true
    const response = await getProductList({
      page: pagination.page,
      page_size: pagination.pageSize,
      keyword: searchForm.keyword || undefined
    })

    if (response.code === 0) {
      productList.value = response.data.products || []
      pagination.total = response.data.total
      pagination.page = response.data.page
      pagination.pageSize = response.data.page_size
    } else {
      ElMessage.error(response.message || '获取商品列表失败')
    }
  } catch (error) {
    console.error('获取商品列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  fetchProductList()
}

// 重置
const handleReset = () => {
  searchForm.keyword = ''
  pagination.page = 1
  fetchProductList()
}

// 新增商品
const handleAdd = () => {
  currentProduct.value = null
  dialogVisible.value = true
}

// 编辑商品
const handleEdit = (product: ProductInfo) => {
  currentProduct.value = { ...product }
  dialogVisible.value = true
}

// 删除商品
const handleDelete = async (product: ProductInfo) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除商品"${product.name}"吗？`,
      '删除确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )

    const response = await deleteProduct(product.id)
    if (response.code === 0) {
      ElMessage.success('删除成功')
      fetchProductList()
    } else {
      ElMessage.error(response.message || '删除失败')
    }
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除商品失败:', error)
      ElMessage.error('删除失败')
    }
  }
}

// 分页大小改变
const handleSizeChange = (size: number) => {
  pagination.pageSize = size
  pagination.page = 1
  fetchProductList()
}

// 当前页改变
const handleCurrentChange = (page: number) => {
  pagination.page = page
  fetchProductList()
}

// 对话框成功回调
const handleDialogSuccess = () => {
  fetchProductList()
}

// 初始化
onMounted(() => {
  fetchProductList()
})
</script>

<style scoped>
.product-container {
  padding: 20px;
}

.search-bar {
  background: #fff;
  padding: 20px;
  border-radius: 4px;
  margin-bottom: 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.toolbar {
  margin-bottom: 20px;
}

.price-text {
  font-weight: 600;
  color: #e6a23c;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: center;
}
</style>
