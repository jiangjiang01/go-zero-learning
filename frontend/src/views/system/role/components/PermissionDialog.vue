<template>
  <el-dialog
    v-model="visible"
    title="分配权限"
    width="600px"
    @close="handleClose"
  >
    <el-tree
      ref="treeRef"
      :data="menuTree"
      :props="{ children: 'children', label: 'name' }"
      show-checkbox
      node-key="id"
      :default-checked-keys="checkedKeys"
      :default-expand-all="true"
    />

    <template #footer>
      <el-button @click="handleClose">取消</el-button>
      <el-button
        type="primary"
        :loading="loading"
        @click="handleSubmit"
      >
        确定
      </el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import type { ElTree } from 'element-plus'
import { getRoleMenus, assignMenusToRole } from '@/api/role'
import { getMenuList, type MenuInfo } from '@/api/menu'

interface Props {
  modelValue: boolean
  roleId?: number | null
}

const props = withDefaults(defineProps<Props>(), {
  roleId: null
})

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
  success: []
}>()

const treeRef = ref<InstanceType<typeof ElTree>>()
const loading = ref(false)
const visible = ref(false)
const menuTree = ref<MenuInfo[]>([])
const checkedKeys = ref<number[]>([])

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
  
  try {
    // 获取所有菜单
    const menuRes = await getMenuList()
    menuTree.value = menuRes.data || []
    
    // 获取角色已分配的菜单
    const roleMenuRes = await getRoleMenus(props.roleId)
    checkedKeys.value = (roleMenuRes.data || []).map((item: any) => item.id)
  } catch (error: any) {
    ElMessage.error(error.message || '获取数据失败')
  }
}

// 提交
const handleSubmit = async () => {
  if (!treeRef.value || !props.roleId) return
  
  loading.value = true
  try {
    const checkedNodes = treeRef.value.getCheckedKeys() as number[]
    await assignMenusToRole(props.roleId, checkedNodes)
    ElMessage.success('权限分配成功')
    handleClose()
    emit('success')
  } catch (error: any) {
    ElMessage.error(error.message || '权限分配失败')
  } finally {
    loading.value = false
  }
}

// 关闭
const handleClose = () => {
  visible.value = false
  checkedKeys.value = []
}
</script>

<style scoped></style>

