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

      <el-form-item label="商品库存" prop="stock">
        <el-input-number
          v-model="form.stock"
          :min="0"
          :max="999999"
          placeholder="请输入商品库存"
          style="width: 200px"
        />
        <span class="price-hint">件</span>
      </el-form-item>

      <el-form-item label="商品状态" prop="status">
        <el-radio-group v-model="form.status">
          <el-radio :label="1">上架</el-radio>
          <el-radio :label="0">下架</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="商品图片">
        <el-upload
          ref="uploadRef"
          :http-request="handleImageUpload"
          :before-upload="beforeImageUpload"
          :file-list="imageList"
          :on-remove="handleImageRemove"
          :limit="10"
          accept="image/*"
          list-type="picture-card"
          :disabled="loading"
        >
          <el-icon><plus /></el-icon>
        </el-upload>
        <div class="image-tip">支持上传多张图片，最多10张，每张不超过10MB</div>
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
import { ElMessage, type FormInstance, type FormRules, type UploadFile, type UploadFiles } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import { createProduct, updateProduct, priceToFen, type ProductInfo } from '@/api/product'
import { uploadFile, isImageFile, validateFileSize, formatFileSize } from '@/api/upload'

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
const uploadRef = ref()
const imageList = ref<UploadFile[]>([])  // 图片列表

// 表单数据
const form = reactive({
  name: '',
  description: '',
  price: 0,
  stock: 0,
  status: 1,
  images: [] as string[]  // 图片URL列表
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
  stock: [
    { required: true, message: '请输入商品库存', trigger: 'blur' },
    { type: 'number', min: 0, max: 999999, message: '库存范围在 0 到 999999 件之间', trigger: 'blur' }
  ],
  status: [
    { required: true, message: '请选择商品状态', trigger: 'change' }
  ]
}

// 重置表单
const resetForm = () => {
  form.name = ''
  form.description = ''
  form.price = 0
  form.stock = 0
  form.status = 1
  form.images = []
  imageList.value = []

  nextTick(() => {
    formRef.value?.clearValidate()
    uploadRef.value?.clearFiles()
  })
}

// 确保图片URL是完整的（如果是相对路径，转换为完整URL）
const ensureFullImageUrl = (url: string): string => {
  if (!url) return url
  // 如果已经是完整URL（以 http:// 或 https:// 开头），直接返回
  if (url.startsWith('http://') || url.startsWith('https://')) {
    return url
  }
  // 如果是相对路径，转换为完整URL（使用当前域名）
  if (url.startsWith('/')) {
    return `${window.location.origin}${url}`
  }
  // 其他情况直接返回
  return url
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
      form.stock = newProduct.stock || 0
      form.status = newProduct.status
      form.images = (newProduct.images || []).map(ensureFullImageUrl)
      
      // 构建图片列表用于显示
      imageList.value = form.images.map((url, index) => ({
        uid: `existing-${index}-${Date.now()}`,
        name: url.split('/').pop() || `image-${index}`,
        url: ensureFullImageUrl(url),
        status: 'success' as const
      }))
    } else {
      // 新增模式，重置表单
      resetForm()
    }
  },
  { immediate: true }
)

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
    const submitData: any = {
      name: form.name.trim(),
      description: form.description.trim(),
      price: priceToFen(form.price), // 元转分
      status: form.status
    }
    
    // 库存字段：新增时如果为0可以不传，编辑时必须传（即使是0）
    if (isEdit.value || form.stock > 0) {
      submitData.stock = form.stock
    }

    // 图片字段：编辑模式下必须传递（即使是空数组），新增模式下如果有图片则传递
    if (isEdit.value) {
      submitData.images = form.images
    } else if (form.images.length > 0) {
      submitData.images = form.images
    }

    let response
    if (isEdit.value && props.product) {
      // 编辑模式
      response = await updateProduct(props.product.id, submitData)
    } else {
      // 新增模式
      response = await createProduct(submitData)
    }

    // 【修复】如果代码执行到这里，说明 response.code === 0（成功）
    // 如果 response.code !== 0，响应拦截器会 reject Promise，代码会进入 catch 块
    // 响应拦截器已经自动显示了错误信息，所以不需要在这里再显示
    if (response.code === 0) {
      ElMessage.success(isEdit.value ? '更新成功' : '创建成功')
      emit('success')
      handleClose()
    }
  } catch (error) {
    // 【修复】表单验证失败时会返回 false，此时不需要处理
    // 如果是 API 错误，响应拦截器已经自动显示了错误信息（如"没有需要更新的字段"）
    // 所以这里不再显示错误提示，避免重复提示
    // 只在控制台输出日志用于调试
    if (error !== false) {
      console.error('提交失败:', error)
      // ❌ 已移除：响应拦截器已经显示了错误信息，这里不再重复显示
      // ElMessage.error(isEdit.value ? '更新失败' : '创建失败')
    }
  } finally {
    loading.value = false
  }
}

// 图片上传前验证
const beforeImageUpload = (file: File): boolean => {
  // 验证文件大小（10MB）
  const maxSize = 10 * 1024 * 1024
  if (!validateFileSize(file, maxSize)) {
    ElMessage.error(`图片大小不能超过 ${formatFileSize(maxSize)}`)
    return false
  }

  // 验证文件类型
  if (!isImageFile(file)) {
    ElMessage.error('只能上传图片文件（jpg、png、gif、webp）')
    return false
  }

  return true
}

// 处理图片上传
const handleImageUpload = async (options: any) => {
  const file = options.file
  
  try {
    const response = await uploadFile({
      file,
      category: 'product'  // 使用 product 分类
    })

    if (response.code === 0) {
      // 添加到图片列表
      const imageUrl = response.data.url
      form.images.push(imageUrl)
      
      // 更新 el-upload 的文件列表
      const uploadFile: UploadFile = {
        uid: Date.now().toString(),
        name: response.data.filename,
        url: imageUrl,
        status: 'success'
      }
      imageList.value.push(uploadFile)
      
      // 调用 onSuccess 回调
      if (options.onSuccess) {
        options.onSuccess(response, file, {})
      }
      
      ElMessage.success('图片上传成功')
    } else {
      throw new Error(response.message || '上传失败')
    }
  } catch (error: any) {
    ElMessage.error(error.message || '图片上传失败')
    // 调用 onError 回调
    if (options.onError) {
      options.onError(error)
    }
  }
}

// 处理图片删除
const handleImageRemove = (file: UploadFile) => {
  // 从图片URL列表中移除
  const index = form.images.findIndex(url => url === file.url)
  if (index > -1) {
    form.images.splice(index, 1)
  }
  
  // 从显示列表中移除
  const listIndex = imageList.value.findIndex(item => item.uid === file.uid)
  if (listIndex > -1) {
    imageList.value.splice(listIndex, 1)
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

.image-tip {
  margin-top: 8px;
  font-size: 12px;
  color: #909399;
}

:deep(.el-upload--picture-card) {
  width: 100px;
  height: 100px;
  line-height: 100px;
}

:deep(.el-upload-list--picture-card .el-upload-list__item) {
  width: 100px;
  height: 100px;
}
</style>
