<template>
  <el-dialog
    v-model="visible"
    :title="userId ? '编辑用户' : '新增用户'"
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
        label="用户名"
        prop="username"
      >
        <el-input
          v-model="form.username"
          :disabled="!!userId"
          placeholder="请输入用户名"
        />
      </el-form-item>

      <el-form-item
        v-if="!userId"
        label="密码"
        prop="password"
      >
        <el-input
          v-model="form.password"
          type="password"
          placeholder="请输入密码"
          show-password
        />
      </el-form-item>

      <el-form-item
        label="昵称"
        prop="nickname"
      >
        <el-input
          v-model="form.nickname"
          placeholder="请输入昵称"
        />
      </el-form-item>

      <el-form-item
        label="邮箱"
        prop="email"
      >
        <el-input
          v-model="form.email"
          placeholder="请输入邮箱"
        />
      </el-form-item>

      <el-form-item
        label="手机号"
        prop="phone"
      >
        <el-input
          v-model="form.phone"
          placeholder="请输入手机号"
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
import { getUser, createUser, updateUser, type CreateUserRequest, type UpdateUserRequest } from '@/api/user'
import { validateUsername, validatePassword, validateEmail, validatePhone } from '@/utils/validate'

interface Props {
  modelValue: boolean
  userId?: number | null
}

const props = withDefaults(defineProps<Props>(), {
  userId: null
})

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
  success: []
}>()

const formRef = ref<FormInstance>()
const loading = ref(false)
const visible = ref(false)

const form = reactive<CreateUserRequest & UpdateUserRequest>({
  username: '',
  password: '',
  nickname: '',
  email: '',
  phone: '',
  status: 1
})

const rules: FormRules = {
  username: [{ validator: validateUsername, trigger: 'blur' }],
  password: [{ validator: validatePassword, trigger: 'blur' }],
  nickname: [{ required: true, message: '请输入昵称', trigger: 'blur' }],
  email: [{ validator: validateEmail, trigger: 'blur' }],
  phone: [{ validator: validatePhone, trigger: 'blur' }]
}

// 监听 visible 变化
watch(
  () => props.modelValue,
  (val) => {
    visible.value = val
    if (val) {
      if (props.userId) {
        fetchUserInfo()
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

// 获取用户信息
const fetchUserInfo = async () => {
  if (!props.userId) return
  
  try {
    const res = await getUser(props.userId)
    Object.assign(form, {
      username: res.data.username,
      nickname: res.data.nickname,
      email: res.data.email || '',
      phone: res.data.phone || '',
      status: res.data.status
    })
  } catch (error: any) {
    ElMessage.error(error.message || '获取用户信息失败')
  }
}

// 重置表单
const resetForm = () => {
  Object.assign(form, {
    username: '',
    password: '',
    nickname: '',
    email: '',
    phone: '',
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
        if (props.userId) {
          // 编辑
          const updateData: UpdateUserRequest = {
            nickname: form.nickname,
            email: form.email,
            phone: form.phone,
            status: form.status
          }
          await updateUser(props.userId, updateData)
          ElMessage.success('编辑成功')
        } else {
          // 新增
          const createData: CreateUserRequest = {
            username: form.username,
            password: form.password!,
            nickname: form.nickname,
            email: form.email,
            phone: form.phone,
            status: form.status
          }
          await createUser(createData)
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

