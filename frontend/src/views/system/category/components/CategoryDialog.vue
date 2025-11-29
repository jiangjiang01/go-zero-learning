<template>
  <el-dialog
    :model-value="visible"
    :title="categoryId ? '编辑分类' : '新增分类'"
    width="600px"
    @update:model-value="handleClose"
    @close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="100px"
    >
      <el-form-item
        label="分类名称"
        prop="name"
      >
        <el-input
          v-model="form.name"
          placeholder="请输入分类名称"
          maxlength="100"
          show-word-limit
        />
      </el-form-item>

      <el-form-item
        label="父分类"
        prop="parent_id"
      >
        <el-select
          v-model="form.parent_id"
          placeholder="请选择父分类（不选则为顶级分类）"
          clearable
          filterable
          style="width: 100%"
        >
          <el-option
            :label="'顶级分类'"
            :value="0"
          />
          <el-option
            v-for="category in parentCategoryOptions"
            :key="category.id"
            :label="category.name"
            :value="category.id"
            :disabled="categoryId === category.id"
          />
        </el-select>
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
          placeholder="请输入分类描述（可选）"
          maxlength="255"
          show-word-limit
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
import { ref, reactive, watch, onMounted, computed } from 'vue'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import {
  getCategoryDetail,
  createCategory,
  updateCategory,
  getCategoryList,
  type CreateCategoryRequest,
  type UpdateCategoryRequest,
  type CategoryInfo
} from '@/api/category'

interface Props {
  visible: boolean
  categoryId?: number | null
}

const props = withDefaults(defineProps<Props>(), {
  categoryId: null
})

interface Emits {
  (e: 'update:visible', visible: boolean): void
  (e: 'success'): void
}

const emit = defineEmits<Emits>()

const formRef = ref<FormInstance>()
const loading = ref(false)
const parentCategoryOptions = ref<CategoryInfo[]>([])

// 计算 visible
const visible = computed({
  get: () => props.visible,
  set: (val) => emit('update:visible', val)
})

const form = reactive({
  name: '',
  desc: '',
  parent_id: 0,
  sort: 0,
  status: 1
})

const rules: FormRules = {
  name: [
    { required: true, message: '请输入分类名称', trigger: 'blur' },
    { min: 1, max: 100, message: '分类名称长度在 1 到 100 个字符', trigger: 'blur' }
  ]
}

// 获取父分类选项（用于下拉选择）
const fetchParentCategories = async () => {
  try {
    const res = await getCategoryList({ all: true })
    // 过滤掉当前编辑的分类（避免循环引用）
    parentCategoryOptions.value = (res.data.categories || []).filter(
      (cat: CategoryInfo) => cat.id !== props.categoryId
    )
  } catch (error) {
    console.error('获取父分类列表失败:', error)
  }
}

// 获取分类详情
const fetchCategoryDetail = async () => {
  if (!props.categoryId) {
    resetForm()
    return
  }

  loading.value = true
  try {
    const res = await getCategoryDetail(props.categoryId)
    const category = res.data
    form.name = category.name
    form.desc = category.desc || ''
    form.parent_id = category.parent_id
    form.sort = category.sort
    form.status = category.status
  } catch (error: any) {
    console.error('获取分类详情失败:', error)
  } finally {
    loading.value = false
  }
}

// 重置表单
const resetForm = () => {
  form.name = ''
  form.desc = ''
  form.parent_id = 0
  form.sort = 0
  form.status = 1
  formRef.value?.clearValidate()
}

// 关闭对话框
const handleClose = () => {
  emit('update:visible', false)
  resetForm()
}

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return

  try {
    await formRef.value.validate()
    loading.value = true

    if (props.categoryId) {
      // 编辑模式
      const updateData: UpdateCategoryRequest = {
        parent_id: form.parent_id || 0
      }
      if (form.name) updateData.name = form.name
      if (form.desc !== undefined) updateData.desc = form.desc
      if (form.sort !== undefined) updateData.sort = form.sort
      if (form.status !== undefined) updateData.status = form.status

      await updateCategory(props.categoryId, updateData)
      ElMessage.success('更新成功')
    } else {
      // 新增模式
      const createData: CreateCategoryRequest = {
        name: form.name,
        desc: form.desc || undefined,
        parent_id: form.parent_id || undefined,
        sort: form.sort || undefined,
        status: form.status
      }
      await createCategory(createData)
      ElMessage.success('创建成功')
    }

    emit('success')
    handleClose()
  } catch (error: any) {
    // 错误消息已经在 request.ts 的响应拦截器中显示过了
    console.error('提交失败:', error)
  } finally {
    loading.value = false
  }
}

// 监听对话框显示
watch(
  () => props.visible,
  (newVal) => {
    if (newVal) {
      fetchParentCategories()
      if (props.categoryId) {
        fetchCategoryDetail()
      } else {
        resetForm()
      }
    }
  },
  { immediate: true }
)
</script>

<style scoped></style>

