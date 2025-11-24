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
        label="角色代码"
        prop="code"
      >
        <el-input
          v-model="form.code"
          :disabled="!!roleId"
          placeholder="请输入角色代码（如：admin, user）"
        />
        <div class="text-gray-500 text-xs mt-1">
          角色代码用于系统内部标识，创建后不可修改
        </div>
      </el-form-item>

      <el-form-item
        label="描述"
        prop="desc"
      >
        <el-input
          v-model="form.desc"
          type="textarea"
          :rows="3"
          placeholder="请输入角色描述（可选）"
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

const form = reactive<CreateRoleRequest>({
  name: '',
  code: '',
  desc: ''
})

// 验证规则
const validateName = (rule: any, value: string, callback: any) => {
  if (!value) {
    callback(new Error('请输入角色名称'))
  } else if (value.length > 50) {
    callback(new Error('角色名称不能超过50个字符'))
  } else {
    callback()
  }
}

const validateCode = (rule: any, value: string, callback: any) => {
  if (!value) {
    callback(new Error('请输入角色代码'))
  } else if (!/^[a-zA-Z][a-zA-Z0-9_]*$/.test(value)) {
    callback(new Error('角色代码只能包含字母、数字和下划线，且必须以字母开头'))
  } else if (value.length > 50) {
    callback(new Error('角色代码不能超过50个字符'))
  } else {
    callback()
  }
}

const rules: FormRules = {
  name: [{ validator: validateName, trigger: 'blur' }],
  code: [{ validator: validateCode, trigger: 'blur' }],
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
      desc: res.data.desc || ''
    })
  } catch (error: any) {
    // 错误消息已在响应拦截器中统一处理，这里只记录日志
    console.error('获取角色信息失败:', error)
  }
}

// 重置表单
const resetForm = () => {
  Object.assign(form, {
    name: '',
    code: '',
    desc: ''
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
          // 编辑：更新角色
          const updateData: UpdateRoleRequest = {
            name: form.name,
            code: form.code, // 虽然代码不可修改，但需要传递以通过验证
            desc: form.desc || ''
          }
          await updateRole(props.roleId, updateData)
          ElMessage.success('更新成功')
        } else {
          // 新增
          const createData: CreateRoleRequest = {
            name: form.name,
            code: form.code,
            desc: form.desc || ''
          }
          await createRole(createData)
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
</script>

<style scoped></style>

