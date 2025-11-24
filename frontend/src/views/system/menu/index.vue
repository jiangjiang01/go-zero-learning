<template>
  <div class="menu-container p-4">
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
              placeholder="菜单名称/代码"
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
              新增菜单
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
          label="菜单名称"
          width="150"
        />
        <el-table-column
          prop="code"
          label="菜单代码"
          width="150"
        />
        <el-table-column
          prop="path"
          label="路径"
          width="150"
        />
        <el-table-column
          prop="icon"
          label="图标"
          width="100"
          align="center"
        />
        <el-table-column
          label="类型"
          width="80"
          align="center"
        >
          <template #default="{ row }">
            <el-tag :type="row.type === 1 ? 'primary' : 'info'">
              {{ row.type === 1 ? '菜单' : '按钮' }}
            </el-tag>
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
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          prop="desc"
          label="描述"
          min-width="200"
          show-overflow-tooltip
        />
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
    <menu-dialog
      v-model="dialogVisible"
      :menu-id="currentMenuId"
      @success="handleDialogSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getMenuList,
  deleteMenu,
  type MenuInfo
} from '@/api/menu'
import { formatDateTime } from '@/utils/format'
import MenuDialog from './components/MenuDialog.vue'

const loading = ref(false)
const tableData = ref<MenuInfo[]>([])
const dialogVisible = ref(false)
const currentMenuId = ref<number | null>(null)

const searchForm = reactive({
  keyword: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

// 获取菜单列表
const fetchMenuList = async () => {
  loading.value = true
  try {
    const res = await getMenuList({
      page: pagination.page,
      page_size: pagination.pageSize,
      keyword: searchForm.keyword || undefined
    })
    tableData.value = res.data.menus || []
    pagination.total = res.data.total || 0
  } catch (error: any) {
    // 错误消息已在响应拦截器中统一处理，这里只记录日志
    console.error('获取菜单列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  fetchMenuList()
}

// 重置
const handleReset = () => {
  searchForm.keyword = ''
  pagination.page = 1
  fetchMenuList()
}

// 新增
const handleAdd = () => {
  currentMenuId.value = null
  dialogVisible.value = true
}

// 编辑
const handleEdit = (row: MenuInfo) => {
  currentMenuId.value = row.id
  dialogVisible.value = true
}

// 删除
const handleDelete = async (row: MenuInfo) => {
  ElMessageBox.confirm(`确定要删除菜单 "${row.name}" 吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async () => {
      try {
        await deleteMenu(row.id)
        ElMessage.success('删除成功')
        fetchMenuList()
      } catch (error: any) {
        // 错误消息已经在 request.ts 的响应拦截器中显示过了，这里不需要再次显示
        console.error('删除菜单失败:', error)
      }
    })
    .catch(() => {})
}

// 分页大小变更
const handleSizeChange = () => {
  pagination.page = 1
  fetchMenuList()
}

// 页码变更
const handlePageChange = () => {
  fetchMenuList()
}

// 对话框成功回调
const handleDialogSuccess = () => {
  fetchMenuList()
}

onMounted(() => {
  fetchMenuList()
})
</script>

<style scoped></style>

