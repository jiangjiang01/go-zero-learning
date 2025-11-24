<template>
  <el-dialog
    v-model="visible"
    title="分配角色"
    width="700px"
    @close="handleClose"
  >
    <div v-loading="loading">
      <!-- 当前用户角色列表 -->
      <div class="mb-4">
        <div class="text-sm font-medium mb-2">当前角色：</div>
        <div v-if="userRoles.length === 0" class="text-gray-400 text-sm">
          暂无角色
        </div>
        <el-tag
          v-for="role in userRoles"
          :key="role.id"
          closable
          class="mr-2 mb-2"
          @close="handleRemoveRole(role)"
        >
          {{ role.name }}
        </el-tag>
      </div>

      <!-- 分配新角色 -->
      <div>
        <div class="text-sm font-medium mb-2">添加角色：</div>
        <el-select
          v-model="selectedRoleId"
          placeholder="请选择角色"
          filterable
          style="width: 100%"
          @change="handleAddRole"
        >
          <el-option
            v-for="role in availableRoles"
            :key="role.id"
            :label="role.name"
            :value="role.id"
          >
            <div class="flex justify-between">
              <span>{{ role.name }}</span>
              <span class="text-gray-400 text-xs">{{ role.code }}</span>
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
import { getUserRoles, assignUserRole, removeUserRole } from '@/api/user'
import { getRoleList, type RoleInfo } from '@/api/role'

interface Props {
  modelValue: boolean
  userId: number | null
}

const props = withDefaults(defineProps<Props>(), {
  userId: null
})

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
  success: []
}>()

const visible = ref(false)
const loading = ref(false)
const userRoles = ref<RoleInfo[]>([])
const allRoles = ref<RoleInfo[]>([])
const selectedRoleId = ref<number | null>(null)

// 计算可用角色（排除已分配的角色）
const availableRoles = computed(() => {
  const assignedRoleIds = userRoles.value.map(r => r.id)
  return allRoles.value.filter(role => !assignedRoleIds.includes(role.id))
})

// 监听 visible 变化
watch(
  () => props.modelValue,
  (val) => {
    visible.value = val
    if (val && props.userId) {
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
  if (!props.userId) return
  
  loading.value = true
  try {
    // 并行获取用户角色和所有角色列表
    const [userRolesRes, allRolesRes] = await Promise.all([
      getUserRoles(props.userId),
      getRoleList({ page: 1, page_size: 1000 }) // 获取所有角色
    ])
    
    userRoles.value = userRolesRes.data.roles || []
    allRoles.value = allRolesRes.data.roles || []
    selectedRoleId.value = null
  } catch (error: any) {
    ElMessage.error(error.message || '获取数据失败')
  } finally {
    loading.value = false
  }
}

// 添加角色
const handleAddRole = async (roleId: number) => {
  if (!props.userId || !roleId) return
  
  try {
    await assignUserRole(props.userId, roleId)
    ElMessage.success('角色分配成功')
    selectedRoleId.value = null
    await fetchData()
    emit('success')
  } catch (error: any) {
    ElMessage.error(error.message || '角色分配失败')
    selectedRoleId.value = null
  }
}

// 移除角色
const handleRemoveRole = async (role: RoleInfo) => {
  if (!props.userId) return
  
  ElMessageBox.confirm(`确定要移除角色 "${role.name}" 吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  })
    .then(async () => {
      try {
        await removeUserRole(props.userId!, role.id)
        ElMessage.success('角色移除成功')
        await fetchData()
        emit('success')
      } catch (error: any) {
        ElMessage.error(error.message || '角色移除失败')
      }
    })
    .catch(() => {})
}

// 关闭
const handleClose = () => {
  visible.value = false
  userRoles.value = []
  allRoles.value = []
  selectedRoleId.value = null
}
</script>

<style scoped></style>

