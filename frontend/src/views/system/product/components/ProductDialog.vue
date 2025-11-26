<template>
  <el-dialog
    :model-value="visible"
    :title="isEdit ? '编辑商品' : '新增商品'"
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
      <el-form-item label="商品名称" prop="name">
        <el-input
          v-model="form.name"
          placeholder="请输入商品名称"
          maxlength="100"
          show-word-limit
        />
      </el-form-item>

      <el-form-item label="商品描述" prop="description">
        <el-input
          v-model="form.description"
          type="textarea"
          placeholder="请输入商品描述"
          :rows="4"
          maxlength="500"
          show-word-limit
        />
      </el-form-item>

      <el-form-item label="商品价格" prop="price">
        <el-input-number
          v-model="form.price"
          :min="0.01"
          :max="999999"
          :precision="2"
          :step="0.01"
          placeholder="请输入商品价格"
          style="width: 200px"
        />
        <span class="price-hint">元</span>
      </el-form-item>

      <el-form-item label="商品状态" prop="status">
        <el-radio-group v-model="form.status">
          <el-radio :label="1">上架</el-radio>
          <el-radio :label="0">下架</el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="loading">
          {{ isEdit ? '更新' : '创建' }}
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, computed, watch, nextTick } from 'vue'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import { createProduct, updateProduct, priceToFen, type ProductInfo } from '@/api/product'

// Props
interface Props {
  visible: boolean
  product?: ProductInfo | null
}

const props = withDefaults(defineProps<Props>(), {
  product: null
})

// Emits
interface Emits {
  (e: 'update:visible', visible: boolean): void
  (e: 'success'): void
}

const emit = defineEmits<Emits>()

// 响应式数据
const loading = ref(false)
const formRef = ref<FormInstance>()

// 表单数据
const form = reactive({
  name: '',
  description: '',
  price: 0,
  status: 1
})

// 计算属性
const isEdit = computed(() => !!props.product)

// 表单验证规则
const rules: FormRules = {
  name: [
    { required: true, message: '请输入商品名称', trigger: 'blur' },
    { min: 1, max: 100, message: '商品名称长度在 1 到 100 个字符', trigger: 'blur' }
  ],
  price: [
    { required: true, message: '请输入商品价格', trigger: 'blur' },
    { type: 'number', min: 0.01, max: 999999, message: '价格范围在 0.01 到 999999 元之间', trigger: 'blur' }
  ],
  status: [
    { required: true, message: '请选择商品状态', trigger: 'change' }
  ]
}

// 监听商品数据变化
watch(
  () => props.product,
  (newProduct) => {
    if (newProduct) {
      // 编辑模式，填充表单数据
      form.name = newProduct.name
      form.description = newProduct.description
      form.price = newProduct.price / 100 // 分转元
      form.status = newProduct.status
    } else {
      // 新增模式，重置表单
      resetForm()
    }
  },
  { immediate: true }
)

// 重置表单
const resetForm = () => {
  form.name = ''
  form.description = ''
  form.price = 0
  form.status = 1
  
  nextTick(() => {
    formRef.value?.clearValidate()
  })
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
    // 表单验证
    await formRef.value.validate()

    loading.value = true

    // 准备提交数据
    const submitData = {
      name: form.name.trim(),
      description: form.description.trim(),
      price: priceToFen(form.price), // 元转分
      status: form.status
    }

    let response
    if (isEdit.value && props.product) {
      // 编辑模式
      response = await updateProduct(props.product.id, submitData)
    } else {
      // 新增模式
      response = await createProduct(submitData)
    }

    if (response.code === 0) {
      ElMessage.success(isEdit.value ? '更新成功' : '创建成功')
      emit('success')
      handleClose()
    } else {
      ElMessage.error(response.message || (isEdit.value ? '更新失败' : '创建失败'))
    }
  } catch (error) {
    if (error !== false) { // 表单验证失败时会返回 false
      console.error('提交失败:', error)
      ElMessage.error(isEdit.value ? '更新失败' : '创建失败')
    }
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.price-hint {
  margin-left: 10px;
  color: #909399;
  font-size: 14px;
}

.dialog-footer {
  text-align: right;
}

:deep(.el-input-number) {
  width: 100%;
}
</style>
