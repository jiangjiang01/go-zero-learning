<template>
  <div class="dict-container p-4">
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
              placeholder="字典名称/字典值/编码"
              clearable
              style="width: 200px"
              @keyup.enter="handleSearch"
            />
          </el-form-item>
          <el-form-item label="字典类型">
            <el-select
              v-model="searchForm.dict_type"
              placeholder="请选择字典类型"
              clearable
              style="width: 200px"
            >
              <el-option
                v-for="type in dictTypes"
                :key="type"
                :label="type"
                :value="type"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="状态">
            <el-select
              v-model="searchForm.status"
              placeholder="请选择状态"
              clearable
              style="width: 120px"
            >
              <el-option
                label="启用"
                :value="1"
              />
              <el-option
                label="禁用"
                :value="0"
              />
            </el-select>
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
              新增字典
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
          label="字典名称"
          width="150"
        />
        <el-table-column
          prop="type"
          label="字典类型"
          width="150"
        />
        <el-table-column
          prop="code"
          label="字典编码"
          width="120"
        />
        <el-table-column
          prop="value"
          label="字典值"
          width="150"
        />
        <el-table-column
          prop="sort"
          label="排序号"
          width="100"
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
    <dict-dialog
      v-model="dialogVisible"
      :dict-id="currentDictId"
      @success="handleDialogSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getDictList,
  deleteDict,
  updateDictStatus,
  getAllDictTypes,
  type DictInfo
} from '@/api/dict'
import DictDialog from './components/DictDialog.vue'

// 搜索表单
const searchForm = reactive({
  keyword: '',
  dict_type: '',
  status: undefined as number | undefined
})

// 表格数据
const tableData = ref<DictInfo[]>([])
const loading = ref(false)

// 分页信息
const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

// 字典类型列表
const dictTypes = ref<string[]>([])

// 对话框
const dialogVisible = ref(false)
const currentDictId = ref<number | null>(null)

// 获取字典列表
const fetchDictList = async () => {
  loading.value = true
  try {
    const res = await getDictList({
      page: pagination.page,
      pageSize: pagination.pageSize,
      keyword: searchForm.keyword || undefined,
      dict_type: searchForm.dict_type || undefined,
      status: searchForm.status
    })
    if (res.code === 200) {
      tableData.value = res.data.list
      pagination.total = res.data.total
    } else {
      ElMessage.error(res.message || '获取字典列表失败')
    }
  } catch (error: any) {
    console.error('获取字典列表失败:', error)
    ElMessage.error('获取字典列表失败')
  } finally {
    loading.value = false
  }
}

// 获取字典类型列表
const fetchDictTypes = async () => {
  try {
    const res = await getAllDictTypes()
    if (res.code === 200) {
      dictTypes.value = res.data
    }
  } catch (error) {
    console.error('获取字典类型列表失败:', error)
  }
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  fetchDictList()
}

// 重置
const handleReset = () => {
  searchForm.keyword = ''
  searchForm.dict_type = ''
  searchForm.status = undefined
  pagination.page = 1
  fetchDictList()
}

// 新增
const handleAdd = () => {
  currentDictId.value = null
  dialogVisible.value = true
}

// 编辑
const handleEdit = (row: DictInfo) => {
  currentDictId.value = row.id
  dialogVisible.value = true
}

// 删除
const handleDelete = async (row: DictInfo) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除字典"${row.name}"吗？`,
      '提示',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    const res = await deleteDict(row.id)
    if (res.code === 200) {
      ElMessage.success('删除成功')
      fetchDictList()
    } else {
      ElMessage.error(res.message || '删除失败')
    }
  } catch (error: any) {
    if (error !== 'cancel') {
      console.error('删除字典失败:', error)
      ElMessage.error('删除失败')
    }
  }
}

// 状态变更
const handleStatusChange = async (row: DictInfo) => {
  const newStatus = row.status === 1 ? 0 : 1
  const statusText = newStatus === 1 ? '启用' : '禁用'
  try {
    const res = await updateDictStatus(row.id, newStatus)
    if (res.code === 200) {
      ElMessage.success(`${statusText}成功`)
      fetchDictList()
    } else {
      ElMessage.error(res.message || `${statusText}失败`)
    }
  } catch (error: any) {
    console.error(`${statusText}字典失败:`, error)
    ElMessage.error(`${statusText}失败`)
  }
}

// 分页大小变更
const handleSizeChange = (size: number) => {
  pagination.pageSize = size
  pagination.page = 1
  fetchDictList()
}

// 页码变更
const handlePageChange = (page: number) => {
  pagination.page = page
  fetchDictList()
}

// 对话框成功回调
const handleDialogSuccess = () => {
  fetchDictList()
  fetchDictTypes() // 刷新字典类型列表
}

// 组件挂载
onMounted(() => {
  fetchDictList()
  fetchDictTypes()
})
</script>

<style scoped>
.dict-container {
  min-height: calc(100vh - 84px);
}
</style>