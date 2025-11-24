<template>
  <div class="user-container p-4">
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
              placeholder="用户名/邮箱"
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
              新增用户
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
          prop="username"
          label="用户名"
        />
        <el-table-column
          prop="email"
          label="邮箱"
        />
        <el-table-column
          label="操作"
          width="240"
          fixed="right"
        >
          <template #default="{ row }">
            <el-button
              type="primary"
              size="small"
              @click="handleEdit(row)"
            >
              查看
            </el-button>
            <el-button
              type="warning"
              size="small"
              @click="handleAssignRole(row)"
            >
              角色
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
    <user-dialog
      v-model="dialogVisible"
      :user-id="currentUserId"
      @success="handleDialogSuccess"
    />

    <!-- 角色分配对话框 -->
    <user-role-dialog
      v-model="roleDialogVisible"
      :user-id="currentUserId"
      @success="handleDialogSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getUserList,
  deleteUser,
  type UserInfo
} from '@/api/user'
import { formatDateTime } from '@/utils/format'
import UserDialog from './components/UserDialog.vue'
import UserRoleDialog from './components/UserRoleDialog.vue'

const loading = ref(false)
const tableData = ref<UserInfo[]>([])
const dialogVisible = ref(false)
const roleDialogVisible = ref(false)
const currentUserId = ref<number | null>(null)

const searchForm = reactive({
  keyword: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

// 获取用户列表
const fetchUserList = async () => {
  loading.value = true
  try {
    const res = await getUserList({
      page: pagination.page,
      page_size: pagination.pageSize,
      keyword: searchForm.keyword || undefined
    })
    tableData.value = res.data.users || []
    pagination.total = res.data.total || 0
  } catch (error: any) {
    // 错误消息已在响应拦截器中统一处理，这里只记录日志
    console.error('获取用户列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  fetchUserList()
}

// 重置
const handleReset = () => {
  searchForm.keyword = ''
  pagination.page = 1
  fetchUserList()
}

// 新增
const handleAdd = () => {
  currentUserId.value = null
  dialogVisible.value = true
}

// 查看/编辑（后端接口只能更新当前用户，这里暂时只支持查看）
const handleEdit = (row: UserInfo) => {
  currentUserId.value = row.id
  dialogVisible.value = true
}

// 分配角色
const handleAssignRole = (row: UserInfo) => {
  currentUserId.value = row.id
  roleDialogVisible.value = true
}

// 删除
const handleDelete = async (row: UserInfo) => {
  ElMessageBox.confirm(`确定要删除用户 "${row.username}" 吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async () => {
      try {
        await deleteUser(row.id)
        ElMessage.success('删除成功')
        fetchUserList()
      } catch (error: any) {
        // 错误消息已经在 request.ts 的响应拦截器中显示过了，这里不需要再次显示
        // 只需要处理业务逻辑（如刷新列表等）
        console.error('删除用户失败:', error)
      }
    })
    .catch(() => {})
}

// 分页大小变更
const handleSizeChange = () => {
  pagination.page = 1
  fetchUserList()
}

// 页码变更
const handlePageChange = () => {
  fetchUserList()
}

// 对话框成功回调
const handleDialogSuccess = () => {
  fetchUserList()
}

onMounted(() => {
  fetchUserList()
})
</script>

<style scoped></style>
