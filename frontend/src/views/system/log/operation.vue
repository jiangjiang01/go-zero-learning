<template>
  <div class="operation-log-container p-4">
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
              placeholder="用户名/路径"
              clearable
              style="width: 200px"
              @keyup.enter="handleSearch"
            />
          </el-form-item>
          <el-form-item label="时间范围">
            <el-date-picker
              v-model="dateRange"
              type="datetimerange"
              range-separator="至"
              start-placeholder="开始时间"
              end-placeholder="结束时间"
              format="YYYY-MM-DD HH:mm:ss"
              value-format="YYYY-MM-DD HH:mm:ss"
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
              type="danger"
              @click="handleClean"
            >
              清理日志
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
          width="120"
        />
        <el-table-column
          prop="method"
          label="请求方法"
          width="100"
          align="center"
        >
          <template #default="{ row }">
            <el-tag
              :type="getMethodType(row.method)"
            >
              {{ row.method }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          prop="path"
          label="请求路径"
          show-overflow-tooltip
        />
        <el-table-column
          prop="ip"
          label="IP地址"
          width="130"
        />
        <el-table-column
          prop="status"
          label="状态"
          width="100"
          align="center"
        >
          <template #default="{ row }">
            <el-tag
              :type="row.status === 200 ? 'success' : 'danger'"
            >
              {{ row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          prop="duration"
          label="耗时(ms)"
          width="100"
          align="center"
        />
        <el-table-column
          prop="created_at"
          label="操作时间"
          width="180"
        />
        <el-table-column
          label="操作"
          width="100"
          fixed="right"
        >
          <template #default="{ row }">
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
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getOperationLogList,
  deleteOperationLog,
  cleanOldOperationLogs,
  type OperationLogInfo
} from '@/api/log'
import { formatDateTime } from '@/utils/format'

const loading = ref(false)
const tableData = ref<OperationLogInfo[]>([])
const dateRange = ref<[string, string] | null>(null)

const searchForm = reactive({
  keyword: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

// 获取日志列表
const fetchLogList = async () => {
  loading.value = true
  try {
    const params: any = {
      page: pagination.page,
      pageSize: pagination.pageSize
    }
    
    if (searchForm.keyword) {
      params.keyword = searchForm.keyword
    }
    
    if (dateRange.value) {
      params.startTime = dateRange.value[0]
      params.endTime = dateRange.value[1]
    }
    
    const res = await getOperationLogList(params)
    tableData.value = res.data.list.map((item: any) => ({
      ...item,
      created_at: formatDateTime(item.created_at)
    }))
    pagination.total = res.data.total
  } catch (error: any) {
    ElMessage.error(error.message || '获取日志列表失败')
  } finally {
    loading.value = false
  }
}

// 获取请求方法类型
const getMethodType = (method: string) => {
  const methodMap: Record<string, string> = {
    GET: 'success',
    POST: 'primary',
    PUT: 'warning',
    DELETE: 'danger'
  }
  return methodMap[method] || 'info'
}

// 搜索
const handleSearch = () => {
  pagination.page = 1
  fetchLogList()
}

// 重置
const handleReset = () => {
  searchForm.keyword = ''
  dateRange.value = null
  pagination.page = 1
  fetchLogList()
}

// 删除
const handleDelete = async (row: OperationLogInfo) => {
  ElMessageBox.confirm('确定要删除这条日志吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async () => {
      try {
        await deleteOperationLog(row.id)
        ElMessage.success('删除成功')
        fetchLogList()
      } catch (error: any) {
        ElMessage.error(error.message || '删除失败')
      }
    })
    .catch(() => {})
}

// 清理日志
const handleClean = () => {
  ElMessageBox.prompt('请输入要保留的天数（例如：30表示保留最近30天的日志）', '清理日志', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    inputPattern: /^\d+$/,
    inputErrorMessage: '请输入有效的天数'
  })
    .then(async ({ value }) => {
      try {
        await cleanOldOperationLogs(Number(value))
        ElMessage.success('清理成功')
        fetchLogList()
      } catch (error: any) {
        ElMessage.error(error.message || '清理失败')
      }
    })
    .catch(() => {})
}

// 分页大小变更
const handleSizeChange = () => {
  fetchLogList()
}

// 页码变更
const handlePageChange = () => {
  fetchLogList()
}

onMounted(() => {
  fetchLogList()
})
</script>

<style scoped></style>
