<template>
  <el-dialog
    v-model="visible"
    :title="menuId ? '编辑菜单' : '新增菜单'"
    width="600px"
    @close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="100px"
    >
      <el-form-item
        label="父菜单"
        prop="parentId"
      >
        <el-tree-select
          v-model="form.parentId"
          :data="menuTreeOptions"
          :props="{ value: 'id', label: 'name', children: 'children' }"
          placeholder="请选择父菜单"
          check-strictly
          clearable
        />
      </el-form-item>

      <el-form-item
        label="菜单名称"
        prop="name"
      >
        <el-input
          v-model="form.name"
          placeholder="请输入菜单名称"
        />
      </el-form-item>

      <el-form-item
        label="路径"
        prop="path"
      >
        <el-input
          v-model="form.path"
          placeholder="请输入路径"
        />
      </el-form-item>

      <el-form-item
        label="组件"
        prop="component"
      >
        <el-input
          v-model="form.component"
          placeholder="请输入组件路径"
        />
      </el-form-item>

      <el-form-item
        label="图标"
        prop="icon"
      >
        <el-input
          v-model="form.icon"
          placeholder="请输入图标名称"
        />
      </el-form-item>

      <el-form-item
        label="排序"
        prop="sort"
      >
        <el-input-number
          v-model="form.sort"
          :min="0"
          placeholder="请输入排序"
        />
      </el-form-item>

      <el-form-item
        label="权限标识"
        prop="permission"
      >
        <el-input
          v-model="form.permission"
          placeholder="请输入权限标识"
        />
      </el-form-item>

      <el-form-item
        label="状态"
        prop="status"
      >
        <el-radio-group v-model="form.status">
          <el-radio :label="1">启用</el-radio>
          <el-radio :label="0">禁用</el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>

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
import { ref, reactive, watch, computed } from 'vue'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import { getMenu, createMenu, updateMenu, type CreateMenuRequest, type UpdateMenuRequest, type MenuInfo } from '@/api/menu'

interface Props {
  modelValue: boolean
  menuId?: number | null
  menuList: MenuInfo[]
}

const props = withDefaults(defineProps<Props>(), {
  menuId: null,
  menuList: () => []
})

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
  success: []
}>()

const formRef = ref<FormInstance>()
const loading = ref(false)
const visible = ref(false)

const form = reactive<CreateMenuRequest & UpdateMenuRequest>({
  parentId: undefined,
  name: '',
  path: '',
  component: '',
  icon: '',
  sort: 0,
  permission: '',
  status: 1
})

const rules: FormRules = {
  name: [{ required: true, message: '请输入菜单名称', trigger: 'blur' }],
  path: [{ required: true, message: '请输入路径', trigger: 'blur' }]
}

// 菜单树选项（排除当前编辑的菜单及其子菜单）
const menuTreeOptions = computed(() => {
  const buildTree = (menus: MenuInfo[], excludeId?: number): any[] => {
    return menus
      .filter((menu) => menu.id !== excludeId)
      .map((menu) => {
        const item: any = {
          id: menu.id,
          name: menu.name,
          children: menu.children ? buildTree(menu.children, excludeId) : []
        }
        return item
      })
  }
  
  return [
    { id: 0, name: '根菜单', children: buildTree(props.menuList, props.menuId || undefined) }
  ]
})

// 监听 visible 变化
watch(
  () => props.modelValue,
  (val) => {
    visible.value = val
    if (val) {
      if (props.menuId) {
        fetchMenuInfo()
      } else {
        resetForm()
      }
    }
  },
  { immediate: true }
)

watch(visible, (val) => {
  emit('update:modelValue', val)
})

// 获取菜单信息
const fetchMenuInfo = async () => {
  if (!props.menuId) return
  
  try {
    const res = await getMenu(props.menuId)
    Object.assign(form, {
      parentId: res.data.parentId || undefined,
      name: res.data.name,
      path: res.data.path,
      component: res.data.component || '',
      icon: res.data.icon || '',
      sort: res.data.sort,
      permission: res.data.permission || '',
      status: res.data.status
    })
  } catch (error: any) {
    ElMessage.error(error.message || '获取菜单信息失败')
  }
}

// 重置表单
const resetForm = () => {
  Object.assign(form, {
    parentId: undefined,
    name: '',
    path: '',
    component: '',
    icon: '',
    sort: 0,
    permission: '',
    status: 1
  })
  formRef.value?.clearValidate()
}

// 提交
const handleSubmit = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        if (props.menuId) {
          // 编辑
          const updateData: UpdateMenuRequest = {
            parentId: form.parentId === 0 ? undefined : form.parentId,
            name: form.name,
            path: form.path,
            component: form.component,
            icon: form.icon,
            sort: form.sort,
            permission: form.permission,
            status: form.status
          }
          await updateMenu(props.menuId, updateData)
          ElMessage.success('编辑成功')
        } else {
          // 新增
          const createData: CreateMenuRequest = {
            parentId: form.parentId === 0 ? undefined : form.parentId,
            name: form.name,
            path: form.path,
            component: form.component,
            icon: form.icon,
            sort: form.sort,
            permission: form.permission,
            status: form.status
          }
          await createMenu(createData)
          ElMessage.success('新增成功')
        }
        handleClose()
        emit('success')
      } catch (error: any) {
        ElMessage.error(error.message || '操作失败')
      } finally {
        loading.value = false
      }
    }
  })
}

// 关闭
const handleClose = () => {
  visible.value = false
  resetForm()
}
</script>

<style scoped></style>

