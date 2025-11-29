<template>
  <div class="category-container p-4">
    <el-card>
      <!-- 搜索栏 -->
      <div class="mb-4">
        <el-form
          :inline="true"
          :model="searchForm"
          class="demo-form-inline"
        >
          <el-form-item label="关键词">
            <el-input
              v-model="searchForm.keyword"
              placeholder="分类名称"
              clearable
              style="width: 200px"
              @keyup.enter="handleSearch"
            />
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              @click="handleSearch"
            >
              搜索
            </el-button>
            <el-button @click="handleReset">重置</el-button>
            <el-button
              type="success"
              @click="handleAdd"
            >
              新增分类
            </el-button>
          </el-form-item>
        </el-form>
      </div>

      <!-- 表格 -->
      <el-table
        v-loading="loading"
        :data="tableData"
        border
        stripe
      >
        <el-table-column
          type="index"
          label="序号"
          width="60"
          align="center"
        />
        <el-table-column
          prop="name"
          label="分类名称"
          width="150"
        />
        <el-table-column
          prop="desc"
          label="描述"
          min-width="200"
          show-overflow-tooltip
        />
        <el-table-column
          prop="parent_id"
          label="父分类ID"
          width="100"
          align="center"
        >
          <template #default="{ row }">
            {{ row.parent_id === 0 ? '顶级分类' : row.parent_id }}
          </template>
        </el-table-column>
        <el-table-column
          prop="sort"
          label="排序"
          width="80"
          align="center"
        />
        <el-table-column
          label="状态"
          width="80"
          align="center"
        >
          <template #default="{ row }">
            <el-tag :type="getCategoryStatusType(row.status)">
              {{ getCategoryStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          label="创建时间"
          align="center"
          width="180"
        >
          <template #default="{ row }">
            {{ formatDateTime(row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          width="180"
          fixed="right"
        >
          <template #default="{ row }">
            <el-button
              type="primary"
              size="small"
              @click="handleEdit(row)"
            >
              编辑
            </el-button>
            <el-button
              type="danger"
              size="small"
              @click="handleDelete(row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="mt-4 flex justify-end">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.pageSize"
          :total="pagination.total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handlePageChange"
        />
      </div>
    </el-card>

    <!-- 新增/编辑对话框 -->
    <CategoryDialog
      v-model:visible="dialogVisible"
      :category-id="currentCategoryId"
      @success="handleDialogSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getCategoryList,
  deleteCategory,
  getCategoryStatusText,
  getCategoryStatusType,
  type CategoryInfo
} from '@/api/category'
import { formatDateTime } from '@/utils/format'
import CategoryDialog from './components/CategoryDialog.vue'

const loading = ref(false)
const tableData = ref<CategoryInfo[]>([])
const dialogVisible = ref(false)
const currentCategoryId = ref<number | null>(null)

const searchForm = reactive({
  keyword: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

// 获取分类列表
const fetchCategoryList = async () => {
  loading.value = true
  try {
    const res = await getCategoryList({
      page: pagination.page,
      page_size: pagination.pageSize,
      keyword: searchForm.keyword || undefined
    })
    tableData.value = res.data.categories || []
    pagination.total = res.data.total || 0
  } catch (error: any) {
    // 错误消息已在响应拦截器中统一处理，这里只记录日志
    console.error('获取分类列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  fetchCategoryList()
}

// 重置
const handleReset = () => {
  searchForm.keyword = ''
  pagination.page = 1
  fetchCategoryList()
}

// 新增
const handleAdd = () => {
  currentCategoryId.value = null
  dialogVisible.value = true
}

// 编辑
const handleEdit = (row: CategoryInfo) => {
  currentCategoryId.value = row.id
  dialogVisible.value = true
}

// 删除
const handleDelete = async (row: CategoryInfo) => {
  ElMessageBox.confirm(`确定要删除分类 "${row.name}" 吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async () => {
      try {
        await deleteCategory(row.id)
        ElMessage.success('删除成功')
        fetchCategoryList()
      } catch (error: any) {
        // 错误消息已经在 request.ts 的响应拦截器中显示过了，这里不需要再次显示
        console.error('删除分类失败:', error)
      }
    })
    .catch(() => {})
}

// 分页大小变更
const handleSizeChange = () => {
  pagination.page = 1
  fetchCategoryList()
}

// 页码变更
const handlePageChange = () => {
  fetchCategoryList()
}

// 对话框成功回调
const handleDialogSuccess = () => {
  fetchCategoryList()
}

onMounted(() => {
  fetchCategoryList()
})
</script>

<style scoped></style>

