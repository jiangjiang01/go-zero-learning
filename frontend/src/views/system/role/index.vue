<template>
  <div class="role-container p-4">
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
              placeholder="角色名称/编码"
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
              新增角色
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
          label="角色名称"
          width="150"
        />
        <el-table-column
          prop="code"
          label="角色编码"
          width="150"
        />
        <el-table-column
          prop="description"
          label="描述"
          show-overflow-tooltip
        />
        <el-table-column
          prop="status"
          label="状态"
          width="100"
          align="center"
        >
          <template #default="{ row }">
            <el-tag
              :type="row.status === 1 ? 'success' : 'danger'"
            >
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          prop="created_at"
          label="创建时间"
          width="180"
        />
        <el-table-column
          label="操作"
          width="280"
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
              type="info"
              size="small"
              @click="handlePermission(row)"
            >
              权限
            </el-button>
            <el-button
              :type="row.status === 1 ? 'warning' : 'success'"
              size="small"
              @click="handleStatusChange(row)"
            >
              {{ row.status === 1 ? '禁用' : '启用' }}
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
    <role-dialog
      v-model="dialogVisible"
      :role-id="currentRoleId"
      @success="handleDialogSuccess"
    />

    <!-- 权限分配对话框 -->
    <permission-dialog
      v-model="permissionDialogVisible"
      :role-id="currentRoleId"
      @success="handlePermissionSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getRoleList,
  deleteRole,
  updateRoleStatus,
  type RoleInfo
} from '@/api/role'
import { formatDateTime } from '@/utils/format'
import RoleDialog from './components/RoleDialog.vue'
import PermissionDialog from './components/PermissionDialog.vue'

const loading = ref(false)
const tableData = ref<RoleInfo[]>([])
const dialogVisible = ref(false)
const permissionDialogVisible = ref(false)
const currentRoleId = ref<number | null>(null)

const searchForm = reactive({
  keyword: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

// 获取角色列表
const fetchRoleList = async () => {
  loading.value = true
  try {
    const res = await getRoleList({
      page: pagination.page,
      pageSize: pagination.pageSize,
      keyword: searchForm.keyword || undefined
    })
    tableData.value = res.data.list.map((item: any) => ({
      ...item,
      created_at: formatDateTime(item.created_at)
    }))
    pagination.total = res.data.total
  } catch (error: any) {
    ElMessage.error(error.message || '获取角色列表失败')
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  fetchRoleList()
}

// 重置
const handleReset = () => {
  searchForm.keyword = ''
  pagination.page = 1
  fetchRoleList()
}

// 新增
const handleAdd = () => {
  currentRoleId.value = null
  dialogVisible.value = true
}

// 编辑
const handleEdit = (row: RoleInfo) => {
  currentRoleId.value = row.id
  dialogVisible.value = true
}

// 权限分配
const handlePermission = (row: RoleInfo) => {
  currentRoleId.value = row.id
  permissionDialogVisible.value = true
}

// 状态变更
const handleStatusChange = async (row: RoleInfo) => {
  const newStatus = row.status === 1 ? 0 : 1
  const statusText = newStatus === 1 ? '启用' : '禁用'
  
  try {
    await updateRoleStatus(row.id, newStatus)
    ElMessage.success(`${statusText}成功`)
    fetchRoleList()
  } catch (error: any) {
    ElMessage.error(error.message || `${statusText}失败`)
  }
}

// 删除
const handleDelete = async (row: RoleInfo) => {
  ElMessageBox.confirm(`确定要删除角色 "${row.name}" 吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async () => {
      try {
        await deleteRole(row.id)
        ElMessage.success('删除成功')
        fetchRoleList()
      } catch (error: any) {
        ElMessage.error(error.message || '删除失败')
      }
    })
    .catch(() => {})
}

// 分页大小变更
const handleSizeChange = () => {
  fetchRoleList()
}

// 页码变更
const handlePageChange = () => {
  fetchRoleList()
}

// 对话框成功回调
const handleDialogSuccess = () => {
  fetchRoleList()
}

// 权限分配成功回调
const handlePermissionSuccess = () => {
  ElMessage.success('权限分配成功')
}

onMounted(() => {
  fetchRoleList()
})
</script>

<style scoped></style>
