<template>
  <div class="permission-container p-4">
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
              placeholder="权限名称/代码"
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
              新增权限
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
          label="权限名称"
          width="150"
        />
        <el-table-column
          prop="code"
          label="权限代码"
          width="150"
        />
        <el-table-column
          prop="desc"
          label="描述"
          min-width="200"
          show-overflow-tooltip
        />
        <el-table-column
          label="创建时间"
          align="center"
        >
          <template #default="{ row }">
            {{ formatDateTime(row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column
          label="更新时间"
          align="center"
        >
          <template #default="{ row }">
            {{ formatDateTime(row.updated_at) }}
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
    <permission-dialog
      v-model="dialogVisible"
      :permission-id="currentPermissionId"
      @success="handleDialogSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getPermissionList,
  deletePermission,
  type PermissionInfo
} from '@/api/permission'
import { formatDateTime } from '@/utils/format'
import PermissionDialog from './components/PermissionDialog.vue'

const loading = ref(false)
const tableData = ref<PermissionInfo[]>([])
const dialogVisible = ref(false)
const currentPermissionId = ref<number | null>(null)

const searchForm = reactive({
  keyword: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

// 获取权限列表
const fetchPermissionList = async () => {
  loading.value = true
  try {
    const res = await getPermissionList({
      page: pagination.page,
      page_size: pagination.pageSize,
      keyword: searchForm.keyword || undefined
    })
    tableData.value = res.data.permissions || []
    pagination.total = res.data.total || 0
  } catch (error: any) {
    // 错误消息已在响应拦截器中统一处理，这里只记录日志
    console.error('获取权限列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  fetchPermissionList()
}

// 重置
const handleReset = () => {
  searchForm.keyword = ''
  pagination.page = 1
  fetchPermissionList()
}

// 新增
const handleAdd = () => {
  currentPermissionId.value = null
  dialogVisible.value = true
}

// 编辑
const handleEdit = (row: PermissionInfo) => {
  currentPermissionId.value = row.id
  dialogVisible.value = true
}

// 删除
const handleDelete = async (row: PermissionInfo) => {
  ElMessageBox.confirm(`确定要删除权限 "${row.name}" 吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async () => {
      try {
        await deletePermission(row.id)
        ElMessage.success('删除成功')
        fetchPermissionList()
      } catch (error: any) {
        // 错误消息已经在 request.ts 的响应拦截器中显示过了，这里不需要再次显示
        console.error('删除权限失败:', error)
      }
    })
    .catch(() => {})
}

// 分页大小变更
const handleSizeChange = () => {
  pagination.page = 1
  fetchPermissionList()
}

// 页码变更
const handlePageChange = () => {
  fetchPermissionList()
}

// 对话框成功回调
const handleDialogSuccess = () => {
  fetchPermissionList()
}

onMounted(() => {
  fetchPermissionList()
})
</script>

<style scoped></style>

