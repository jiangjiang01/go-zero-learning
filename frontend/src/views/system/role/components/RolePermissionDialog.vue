<template>
  <el-dialog
    v-model="visible"
    title="分配权限"
    width="700px"
    @close="handleClose"
  >
    <div v-loading="loading">
      <!-- 当前角色权限列表 -->
      <div class="mb-4">
        <div class="text-sm font-medium mb-2">当前权限：</div>
        <div v-if="rolePermissions.length === 0" class="text-gray-400 text-sm">
          暂无权限
        </div>
        <el-tag
          v-for="permission in rolePermissions"
          :key="permission.id"
          closable
          class="mr-2 mb-2"
          @close="handleRemovePermission(permission)"
        >
          {{ permission.name }}
        </el-tag>
      </div>

      <!-- 分配新权限 -->
      <div>
        <div class="text-sm font-medium mb-2">添加权限：</div>
        <el-select
          v-model="selectedPermissionId"
          placeholder="请选择权限"
          filterable
          style="width: 100%"
          @change="handleAddPermission"
        >
          <el-option
            v-for="permission in availablePermissions"
            :key="permission.id"
            :label="permission.name"
            :value="permission.id"
          >
            <div class="flex justify-between">
              <span>{{ permission.name }}</span>
              <span class="text-gray-400 text-xs">{{ permission.code }}</span>
            </div>
          </el-option>
        </el-select>
      </div>
    </div>

    <template #footer>
      <el-button @click="handleClose">关闭</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, watch, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getRolePermissions, assignRolePermission, removeRolePermission } from '@/api/role'
import { getPermissionList, type PermissionInfo } from '@/api/permission'

interface Props {
  modelValue: boolean
  roleId: number | null
}

const props = withDefaults(defineProps<Props>(), {
  roleId: null
})

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
  success: []
}>()

const visible = ref(false)
const loading = ref(false)
const rolePermissions = ref<PermissionInfo[]>([])
const allPermissions = ref<PermissionInfo[]>([])
const selectedPermissionId = ref<number | null>(null)

// 计算可用权限（排除已分配的权限）
const availablePermissions = computed(() => {
  const assignedPermissionIds = rolePermissions.value.map(p => p.id)
  return allPermissions.value.filter(permission => !assignedPermissionIds.includes(permission.id))
})

// 监听 visible 变化
watch(
  () => props.modelValue,
  (val) => {
    visible.value = val
    if (val && props.roleId) {
      fetchData()
    }
  },
  { immediate: true }
)

watch(visible, (val) => {
  emit('update:modelValue', val)
})

// 获取数据
const fetchData = async () => {
  if (!props.roleId) return
  
  loading.value = true
  try {
    // 并行获取角色权限和所有权限列表
    const [rolePermissionsRes, allPermissionsRes] = await Promise.all([
      getRolePermissions(props.roleId),
      getPermissionList({ page: 1, page_size: 1000 }) // 获取所有权限
    ])
    
    rolePermissions.value = rolePermissionsRes.data.permissions || []
    allPermissions.value = allPermissionsRes.data.permissions || []
    selectedPermissionId.value = null
  } catch (error: any) {
    // 错误消息已在响应拦截器中统一处理，这里只记录日志
    console.error('获取数据失败:', error)
  } finally {
    loading.value = false
  }
}

// 添加权限
const handleAddPermission = async (permissionId: number) => {
  if (!props.roleId || !permissionId) return
  
  try {
    await assignRolePermission(props.roleId, permissionId)
    ElMessage.success('权限分配成功')
    selectedPermissionId.value = null
    await fetchData()
    emit('success')
  } catch (error: any) {
    // 错误消息已在响应拦截器中统一处理，这里只记录日志
    console.error('权限分配失败:', error)
    selectedPermissionId.value = null
  }
}

// 移除权限
const handleRemovePermission = async (permission: PermissionInfo) => {
  if (!props.roleId) return
  
  ElMessageBox.confirm(`确定要移除权限 "${permission.name}" 吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async () => {
      try {
        await removeRolePermission(props.roleId!, permission.id)
        ElMessage.success('权限移除成功')
        await fetchData()
        emit('success')
      } catch (error: any) {
        // 错误消息已在响应拦截器中统一处理，这里只记录日志
        console.error('权限移除失败:', error)
      }
    })
    .catch(() => {})
}

// 关闭
const handleClose = () => {
  visible.value = false
  rolePermissions.value = []
  allPermissions.value = []
  selectedPermissionId.value = null
}
</script>

<style scoped></style>

