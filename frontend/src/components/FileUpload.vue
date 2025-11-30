<template>
  <div class="file-upload">
    <el-upload
      ref="uploadRef"
      :http-request="handleUpload"
      :before-upload="beforeUpload"
      :show-file-list="showFileList"
      :limit="limit"
      :accept="accept"
      :disabled="disabled"
      :drag="drag"
    >
      <template v-if="drag">
        <el-icon class="el-icon--upload"><upload-filled /></el-icon>
        <div class="el-upload__text">
          将文件拖到此处，或<em>点击上传</em>
        </div>
      </template>
      <el-button v-else type="primary" :disabled="disabled">
        <el-icon><upload /></el-icon>
        {{ buttonText }}
      </el-button>
      <template #tip>
        <div class="el-upload__tip">
          {{ tip }}
        </div>
      </template>
    </el-upload>

    <!-- 上传进度 -->
    <div v-if="uploading" class="upload-progress">
      <el-progress :percentage="uploadProgress" :status="uploadStatus" />
    </div>

    <!-- 预览区域 -->
    <div v-if="fileUrl" class="file-preview">
      <div class="preview-header">
        <span>已上传文件：</span>
        <el-button type="danger" size="small" text @click="clearFile">清除</el-button>
      </div>
      <div class="preview-content">
        <img v-if="isImage" :src="fileUrl" alt="预览" class="preview-image" />
        <div v-else class="preview-file">
          <el-icon><document /></el-icon>
          <a :href="fileUrl" target="_blank">{{ fileName }}</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Upload, UploadFilled, Document } from '@element-plus/icons-vue'
import { uploadFile, isImageFile, validateFileSize, formatFileSize, type UploadFileResponse } from '@/api/upload'

interface Props {
  modelValue?: string        // v-model 绑定的文件 URL
  category?: string          // 文件分类
  maxSize?: number           // 最大文件大小（字节），默认 10MB
  accept?: string            // 接受的文件类型，如 "image/*"
  limit?: number             // 文件数量限制，默认 1
  disabled?: boolean         // 是否禁用
  drag?: boolean             // 是否启用拖拽上传
  buttonText?: string        // 按钮文字
  tip?: string               // 提示文字
  showFileList?: boolean     // 是否显示文件列表
}

interface Emits {
  (e: 'update:modelValue', url: string): void
  (e: 'success', response: UploadFileResponse): void
  (e: 'error', error: Error): void
}

const props = withDefaults(defineProps<Props>(), {
  maxSize: 10 * 1024 * 1024, // 默认 10MB
  accept: 'image/*',
  limit: 1,
  disabled: false,
  drag: false,
  buttonText: '选择文件',
  tip: '只能上传图片文件，且不超过 10MB',
  showFileList: false
})

const emit = defineEmits<Emits>()

const uploadRef = ref()
const uploading = ref(false)
const uploadProgress = ref(0)
const uploadStatus = ref<'success' | 'exception' | 'warning' | ''>('')
const fileUrl = ref(props.modelValue || '')
const fileName = ref('')

// 判断是否为图片
const isImage = computed(() => {
  if (!fileUrl.value) return false
  return fileUrl.value.match(/\.(jpg|jpeg|png|gif|webp)$/i)
})

// 上传前验证
const beforeUpload = (file: File): boolean => {
  // 验证文件大小
  if (!validateFileSize(file, props.maxSize)) {
    ElMessage.error(`文件大小不能超过 ${formatFileSize(props.maxSize)}`)
    return false
  }

  // 验证文件类型（如果是图片）
  if (props.accept.includes('image') && !isImageFile(file)) {
    ElMessage.error('只能上传图片文件（jpg、png、gif、webp）')
    return false
  }

  return true
}

// 自定义上传方法
const handleUpload = async (options: any) => {
  const file = options.file
  
  try {
    uploading.value = true
    uploadProgress.value = 0
    uploadStatus.value = ''

    // 模拟上传进度
    const progressInterval = setInterval(() => {
      if (uploadProgress.value < 90) {
        uploadProgress.value += 10
      }
    }, 200)

    const response = await uploadFile({
      file,
      category: props.category
    })

    clearInterval(progressInterval)
    uploadProgress.value = 100
    uploadStatus.value = 'success'

    if (response.code === 0) {
      fileUrl.value = response.data.url
      fileName.value = response.data.filename
      emit('update:modelValue', response.data.url)
      emit('success', response.data)
      ElMessage.success('上传成功')
      
      // 调用 onSuccess 回调（Element Plus 需要）
      if (options.onSuccess) {
        options.onSuccess(response, file, {})
      }
    } else {
      throw new Error(response.message || '上传失败')
    }
  } catch (error: any) {
    uploadStatus.value = 'exception'
    uploadProgress.value = 0
    emit('error', error)
    
    // 调用 onError 回调（Element Plus 需要）
    if (options.onError) {
      options.onError(error)
    }
    
    ElMessage.error(error.message || '上传失败')
  } finally {
    uploading.value = false
  }
}


// 清除文件
const clearFile = () => {
  fileUrl.value = ''
  fileName.value = ''
  emit('update:modelValue', '')
  uploadRef.value?.clearFiles()
}

// 监听 modelValue 变化
watch(() => props.modelValue, (newValue) => {
  fileUrl.value = newValue || ''
})

// 暴露方法供父组件调用
defineExpose({
  upload: handleUpload,
  clear: clearFile
})
</script>

<style scoped>
.file-upload {
  width: 100%;
}

.upload-progress {
  margin-top: 10px;
}

.file-preview {
  margin-top: 20px;
  padding: 12px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  background-color: #f5f7fa;
}

.preview-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
  font-size: 14px;
  color: #606266;
}

.preview-content {
  display: flex;
  align-items: center;
  justify-content: center;
}

.preview-image {
  max-width: 100%;
  max-height: 300px;
  border-radius: 4px;
}

.preview-file {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #409eff;
}

.preview-file a {
  color: #409eff;
  text-decoration: none;
}

.preview-file a:hover {
  text-decoration: underline;
}
</style>

