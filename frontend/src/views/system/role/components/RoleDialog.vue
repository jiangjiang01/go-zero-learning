<template>
  <el-dialog
    v-model="visible"
    :title="roleId ? '编辑角色' : '新增角色'"
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
        label="角色名称"
        prop="name"
      >
        <el-input
          v-model="form.name"
          placeholder="请输入角色名称"
        />
      </el-form-item>

      <el-form-item
        label="角色编码"
        prop="code"
      >
        <el-input
          v-model="form.code"
          :disabled="!!roleId"
          placeholder="请输入角色编码"
        />
      </el-form-item>

      <el-form-item
        label="描述"
        prop="description"
      >
        <el-input
          v-model="form.description"
          type="textarea"
          :rows="3"
          placeholder="请输入描述"
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
import { ref, reactive, watch } from 'vue'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import { getRole, createRole, updateRole, type CreateRoleRequest, type UpdateRoleRequest } from '@/api/role'

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

const formRef = ref<FormInstance>()
const loading = ref(false)
const visible = ref(false)

const form = reactive<CreateRoleRequest & UpdateRoleRequest>({
  name: '',
  code: '',
  description: '',
  status: 1
})

const rules: FormRules = {
  name: [{ required: true, message: '请输入角色名称', trigger: 'blur' }],
  code: [{ required: true, message: '请输入角色编码', trigger: 'blur' }]
}

// 监听 visible 变化
watch(
  () => props.modelValue,
  (val) => {
    visible.value = val
    if (val) {
      if (props.roleId) {
        fetchRoleInfo()
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

// 获取角色信息
const fetchRoleInfo = async () => {
  if (!props.roleId) return
  
  try {
    const res = await getRole(props.roleId)
    Object.assign(form, {
      name: res.data.name,
      code: res.data.code,
      description: res.data.description || '',
      status: res.data.status
    })
  } catch (error: any) {
    ElMessage.error(error.message || '获取角色信息失败')
  }
}

// 重置表单
const resetForm = () => {
  Object.assign(form, {
    name: '',
    code: '',
    description: '',
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
        if (props.roleId) {
          // 编辑
          const updateData: UpdateRoleRequest = {
            name: form.name,
            description: form.description,
            status: form.status
          }
          await updateRole(props.roleId, updateData)
          ElMessage.success('编辑成功')
        } else {
          // 新增
          const createData: CreateRoleRequest = {
            name: form.name,
            code: form.code,
            description: form.description,
            status: form.status
          }
          await createRole(createData)
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

