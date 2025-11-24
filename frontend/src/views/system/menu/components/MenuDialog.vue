<template>
  <el-dialog
    v-model="visible"
    :title="menuId ? '编辑菜单' : '新增菜单'"
    width="700px"
    @close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="100px"
    >
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
        label="菜单代码"
        prop="code"
      >
        <el-input
          v-model="form.code"
          :disabled="!!menuId"
          placeholder="请输入菜单代码（如：system:user）"
        />
        <div class="text-gray-500 text-xs mt-1">
          菜单代码用于权限控制，创建后不可修改
        </div>
      </el-form-item>

      <el-form-item
        label="菜单类型"
        prop="type"
      >
        <el-radio-group v-model="form.type">
          <el-radio :label="1">菜单</el-radio>
          <el-radio :label="2">按钮</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item
        label="父菜单"
        prop="parent_id"
      >
        <el-select
          v-model="form.parent_id"
          placeholder="请选择父菜单（不选则为顶级菜单）"
          clearable
          filterable
          style="width: 100%"
        >
          <el-option
            :label="'顶级菜单'"
            :value="0"
          />
          <el-option
            v-for="menu in parentMenuOptions"
            :key="menu.id"
            :label="menu.name"
            :value="menu.id"
            :disabled="menuId === menu.id"
          />
        </el-select>
      </el-form-item>

      <el-form-item
        label="菜单路径"
        prop="path"
      >
        <el-input
          v-model="form.path"
          placeholder="请输入菜单路径（如：/system/user）"
        />
      </el-form-item>

      <el-form-item
        label="菜单图标"
        prop="icon"
      >
        <el-input
          v-model="form.icon"
          placeholder="请输入菜单图标（如：el-icon-user）"
        />
      </el-form-item>

      <el-form-item
        label="排序"
        prop="sort"
      >
        <el-input-number
          v-model="form.sort"
          :min="0"
          :max="9999"
          placeholder="请输入排序值（数字越小越靠前）"
          style="width: 100%"
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

      <el-form-item
        label="描述"
        prop="desc"
      >
        <el-input
          v-model="form.desc"
          type="textarea"
          :rows="3"
          placeholder="请输入菜单描述（可选）"
        />
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
import { ref, reactive, watch, onMounted } from 'vue'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import {
  getMenu,
  createMenu,
  updateMenu,
  getMenuList,
  type CreateMenuRequest,
  type UpdateMenuRequest,
  type MenuInfo
} from '@/api/menu'

interface Props {
  modelValue: boolean
  menuId?: number | null
}

const props = withDefaults(defineProps<Props>(), {
  menuId: null
})

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
  success: []
}>()

const formRef = ref<FormInstance>()
const loading = ref(false)
const visible = ref(false)
const parentMenuOptions = ref<MenuInfo[]>([])

const form = reactive<CreateMenuRequest & { parent_id: number; sort: number; status: number }>({
  name: '',
  code: '',
  desc: '',
  parent_id: 0,
  path: '',
  icon: '',
  type: 1,
  sort: 0,
  status: 1
})

// 验证规则
const validateName = (rule: any, value: string, callback: any) => {
  if (!value) {
    callback(new Error('请输入菜单名称'))
  } else if (value.length > 50) {
    callback(new Error('菜单名称不能超过50个字符'))
  } else {
    callback()
  }
}

const validateCode = (rule: any, value: string, callback: any) => {
  if (!value) {
    callback(new Error('请输入菜单代码'))
  } else if (!/^[a-zA-Z][a-zA-Z0-9_:]*$/.test(value)) {
    callback(new Error('菜单代码只能包含字母、数字、下划线和冒号，且必须以字母开头'))
  } else if (value.length > 50) {
    callback(new Error('菜单代码不能超过50个字符'))
  } else {
    callback()
  }
}

const rules: FormRules = {
  name: [{ validator: validateName, trigger: 'blur' }],
  code: [{ validator: validateCode, trigger: 'blur' }],
  type: [
    {
      validator: (rule, value, callback) => {
        if (value !== 1 && value !== 2) {
          callback(new Error('菜单类型只能是1（菜单）或2（按钮）'))
        } else {
          callback()
        }
      },
      trigger: 'change'
    }
  ],
  status: [
    {
      validator: (rule, value, callback) => {
        if (value !== 0 && value !== 1) {
          callback(new Error('状态只能是0（禁用）或1（启用）'))
        } else {
          callback()
        }
      },
      trigger: 'change'
    }
  ],
  desc: [
    {
      validator: (rule, value, callback) => {
        if (value && value.length > 255) {
          callback(new Error('描述不能超过255个字符'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}

// 获取父菜单列表（用于下拉选择）
const fetchParentMenus = async () => {
  try {
    const res = await getMenuList({ all: true })
    parentMenuOptions.value = res.data.menus || []
  } catch (error: any) {
    console.error('获取父菜单列表失败:', error)
  }
}

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
      name: res.data.name,
      code: res.data.code,
      desc: res.data.desc || '',
      parent_id: res.data.parent_id || 0,
      path: res.data.path || '',
      icon: res.data.icon || '',
      type: res.data.type,
      sort: res.data.sort || 0,
      status: res.data.status
    })
  } catch (error: any) {
    // 错误消息已在响应拦截器中统一处理，这里只记录日志
    console.error('获取菜单信息失败:', error)
  }
}

// 重置表单
const resetForm = () => {
  Object.assign(form, {
    name: '',
    code: '',
    desc: '',
    parent_id: 0,
    path: '',
    icon: '',
    type: 1,
    sort: 0,
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
          // 编辑：更新菜单（所有字段必填）
          const updateData: UpdateMenuRequest = {
            name: form.name,
            code: form.code,
            desc: form.desc || '',
            parent_id: form.parent_id || 0,
            path: form.path || '',
            icon: form.icon || '',
            type: form.type,
            sort: form.sort || 0,
            status: form.status
          }
          await updateMenu(props.menuId, updateData)
          ElMessage.success('更新成功')
        } else {
          // 新增
          const createData: CreateMenuRequest = {
            name: form.name,
            code: form.code,
            desc: form.desc || '',
            parent_id: form.parent_id || 0,
            path: form.path || '',
            icon: form.icon || '',
            type: form.type,
            sort: form.sort || 0,
            status: form.status || 1
          }
          await createMenu(createData)
          ElMessage.success('新增成功')
        }
        handleClose()
        emit('success')
      } catch (error: any) {
        // 错误消息已经在 request.ts 的响应拦截器中显示过了，这里不需要再次显示
        console.error('操作失败:', error)
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

onMounted(() => {
  fetchParentMenus()
})
</script>

<style scoped></style>

