<template>
  <el-dialog
    v-model="visible"
    :title="dictId ? '编辑字典' : '新增字典'"
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
        label="字典名称"
        prop="name"
      >
        <el-input
          v-model="form.name"
          placeholder="请输入字典名称"
        />
      </el-form-item>

      <el-form-item
        label="字典类型"
        prop="type"
      >
        <el-input
          v-model="form.type"
          placeholder="请输入字典类型（如：user_status）"
        >
          <template #append>
            <el-select
              v-model="form.type"
              placeholder="选择类型"
              style="width: 150px"
              filterable
              @change="handleTypeChange"
            >
              <el-option
                v-for="type in dictTypes"
                :key="type"
                :label="type"
                :value="type"
              />
            </el-select>
          </template>
        </el-input>
        <div class="text-gray-500 text-xs mt-1">
          提示：同一类型下的编码必须唯一
        </div>
      </el-form-item>

      <el-form-item
        label="字典编码"
        prop="code"
      >
        <el-input
          v-model="form.code"
          placeholder="请输入字典编码（如：1, 0）"
        />
      </el-form-item>

      <el-form-item
        label="字典值"
        prop="value"
      >
        <el-input
          v-model="form.value"
          placeholder="请输入字典值（如：启用，禁用）"
        />
      </el-form-item>

      <el-form-item
        label="排序号"
        prop="sort"
      >
        <el-input-number
          v-model="form.sort"
          :min="0"
          :max="9999"
          placeholder="请输入排序号"
          style="width: 100%"
        />
      </el-form-item>

      <el-form-item
        v-if="dictId"
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
import { ref, reactive, watch, onMounted } from 'vue'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import {
  getDict,
  createDict,
  updateDict,
  getAllDictTypes,
  type CreateDictRequest,
  type UpdateDictRequest
} from '@/api/dict'

interface Props {
  modelValue: boolean
  dictId?: number | null
}

const props = withDefaults(defineProps<Props>(), {
  dictId: null
})

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
  success: []
}>()

const formRef = ref<FormInstance>()
const loading = ref(false)
const visible = ref(false)
const dictTypes = ref<string[]>([])

const form = reactive<CreateDictRequest & UpdateDictRequest & { status?: number }>({
  name: '',
  type: '',
  code: '',
  value: '',
  sort: 0,
  status: 1
})

const rules: FormRules = {
  name: [
    { required: true, message: '请输入字典名称', trigger: 'blur' },
    { min: 1, max: 100, message: '长度在 1 到 100 个字符', trigger: 'blur' }
  ],
  type: [
    { required: true, message: '请输入字典类型', trigger: 'blur' },
    { min: 1, max: 50, message: '长度在 1 到 50 个字符', trigger: 'blur' }
  ],
  code: [
    { required: true, message: '请输入字典编码', trigger: 'blur' },
    { min: 1, max: 50, message: '长度在 1 到 50 个字符', trigger: 'blur' }
  ],
  value: [
    { required: true, message: '请输入字典值', trigger: 'blur' },
    { min: 1, max: 255, message: '长度在 1 到 255 个字符', trigger: 'blur' }
  ],
  sort: [
    { type: 'number', min: 0, max: 9999, message: '排序号范围 0-9999', trigger: 'blur' }
  ]
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

// 监听 visible 变化
watch(
  () => props.modelValue,
  (val) => {
    visible.value = val
    if (val) {
      fetchDictTypes() // 每次打开对话框时刷新字典类型列表
      if (props.dictId) {
        fetchDictInfo()
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

// 获取字典信息
const fetchDictInfo = async () => {
  if (!props.dictId) return
  
  try {
    const res = await getDict(props.dictId)
    if (res.code === 200) {
      Object.assign(form, {
        name: res.data.name,
        type: res.data.type,
        code: res.data.code,
        value: res.data.value,
        sort: res.data.sort,
        status: res.data.status
      })
    }
  } catch (error: any) {
    ElMessage.error(error.message || '获取字典信息失败')
  }
}

// 字典类型变更
const handleTypeChange = (type: string) => {
  form.type = type
}

// 重置表单
const resetForm = () => {
  Object.assign(form, {
    name: '',
    type: '',
    code: '',
    value: '',
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
        if (props.dictId) {
          // 编辑
          const updateData: UpdateDictRequest = {
            name: form.name,
            type: form.type,
            code: form.code,
            value: form.value,
            sort: form.sort,
            status: form.status
          }
          const res = await updateDict(props.dictId, updateData)
          if (res.code === 200) {
            ElMessage.success('编辑成功')
            handleClose()
            emit('success')
          } else {
            ElMessage.error(res.message || '编辑失败')
          }
        } else {
          // 新增
          const createData: CreateDictRequest = {
            name: form.name,
            type: form.type,
            code: form.code,
            value: form.value,
            sort: form.sort || 0
          }
          const res = await createDict(createData)
          if (res.code === 200) {
            ElMessage.success('新增成功')
            handleClose()
            emit('success')
          } else {
            ElMessage.error(res.message || '新增失败')
          }
        }
      } catch (error: any) {
        console.error('操作失败:', error)
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

// 组件挂载时获取字典类型列表
onMounted(() => {
  fetchDictTypes()
})
</script>

<style scoped>
.text-gray-500 {
  color: #909399;
}

.text-xs {
  font-size: 12px;
}

.mt-1 {
  margin-top: 4px;
}
</style>