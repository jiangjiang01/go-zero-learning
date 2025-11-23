<template>
  <div class="menu-container p-4">
    <el-card>
      <div class="mb-4">
        <el-button
          type="success"
          @click="handleAdd"
        >
          新增菜单
        </el-button>
      </div>

      <!-- 表格 -->
      <el-table
        v-loading="loading"
        :data="tableData"
        border
        row-key="id"
        :tree-props="{ children: 'children' }"
        default-expand-all
      >
        <el-table-column
          prop="name"
          label="菜单名称"
          width="200"
        />
        <el-table-column
          prop="path"
          label="路径"
          width="200"
        />
        <el-table-column
          prop="component"
          label="组件"
          width="200"
        />
        <el-table-column
          prop="icon"
          label="图标"
          width="100"
          align="center"
        >
          <template #default="{ row }">
            <el-icon v-if="row.icon">
              <component :is="row.icon" />
            </el-icon>
          </template>
        </el-table-column>
        <el-table-column
          prop="sort"
          label="排序"
          width="80"
          align="center"
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
          prop="permission"
          label="权限标识"
          width="150"
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
    </el-card>

    <!-- 新增/编辑对话框 -->
    <menu-dialog
      v-model="dialogVisible"
      :menu-id="currentMenuId"
      :menu-list="tableData"
      @success="handleDialogSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getMenuList,
  deleteMenu,
  updateMenuStatus,
  type MenuInfo
} from '@/api/menu'
import MenuDialog from './components/MenuDialog.vue'

const loading = ref(false)
const tableData = ref<MenuInfo[]>([])
const dialogVisible = ref(false)
const currentMenuId = ref<number | null>(null)

// 获取菜单列表
const fetchMenuList = async () => {
  loading.value = true
  try {
    const res = await getMenuList()
    tableData.value = res.data || []
  } catch (error: any) {
    ElMessage.error(error.message || '获取菜单列表失败')
  } finally {
    loading.value = false
  }
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

// 状态变更
const handleStatusChange = async (row: MenuInfo) => {
  const newStatus = row.status === 1 ? 0 : 1
  const statusText = newStatus === 1 ? '启用' : '禁用'
  
  try {
    await updateMenuStatus(row.id, newStatus)
    ElMessage.success(`${statusText}成功`)
    fetchMenuList()
  } catch (error: any) {
    ElMessage.error(error.message || `${statusText}失败`)
  }
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
        ElMessage.error(error.message || '删除失败')
      }
    })
    .catch(() => {})
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
