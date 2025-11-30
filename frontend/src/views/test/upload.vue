<template>
  <div class="upload-test-container p-4">
    <h1 class="text-2xl font-bold mb-6">文件上传测试</h1>

    <el-row :gutter="20">
      <!-- 使用 FileUpload 组件 -->
      <el-col :xs="24" :md="12">
        <el-card shadow="hover">
          <template #header>
            <div class="card-header">
              <el-icon><Picture /></el-icon>
              <span>方式一：使用 FileUpload 组件</span>
            </div>
          </template>
          <FileUpload
            v-model="imageUrl1"
            category="test"
            :max-size="10 * 1024 * 1024"
            accept="image/*"
            tip="只能上传图片文件，且不超过 10MB"
            @success="handleUploadSuccess1"
            @error="handleUploadError"
          />
          <div v-if="imageUrl1" class="mt-4">
            <p class="text-sm text-gray-600 mb-2">上传成功！文件 URL：</p>
            <el-input :model-value="imageUrl1" readonly>
              <template #append>
                <el-button @click="copyUrl(imageUrl1)">复制</el-button>
              </template>
            </el-input>
          </div>
        </el-card>
      </el-col>

      <!-- 使用 API 直接上传 -->
      <el-col :xs="24" :md="12">
        <el-card shadow="hover">
          <template #header>
            <div class="card-header">
              <el-icon><Upload /></el-icon>
              <span>方式二：使用 API 直接上传</span>
            </div>
          </template>
          <el-upload
            :http-request="handleDirectUpload"
            :before-upload="beforeUpload"
            :show-file-list="false"
            accept="image/*"
            :disabled="uploading"
          >
            <el-button type="primary" :loading="uploading">
              <el-icon><Upload /></el-icon>
              选择文件
            </el-button>
            <template #tip>
              <div class="el-upload__tip">只能上传图片文件，且不超过 10MB</div>
            </template>
          </el-upload>
          <div v-if="uploading" class="mt-4">
            <el-progress :percentage="uploadProgress" />
          </div>
          <div v-if="imageUrl2" class="mt-4">
            <p class="text-sm text-gray-600 mb-2">上传成功！</p>
            <img :src="imageUrl2" alt="预览" class="preview-image" />
            <div class="mt-2">
              <el-input :model-value="imageUrl2" readonly>
                <template #append>
                  <el-button @click="copyUrl(imageUrl2)">复制 URL</el-button>
                </template>
              </el-input>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 上传结果列表 -->
    <el-card shadow="hover" class="mt-6">
      <template #header>
        <div class="card-header">
          <el-icon><List /></el-icon>
          <span>上传历史记录</span>
          <el-button
            v-if="uploadResults.length > 0"
            type="danger"
            size="small"
            text
            @click="clearResults"
            style="margin-left: auto"
          >
            清空记录
          </el-button>
        </div>
      </template>
      <el-empty v-if="uploadResults.length === 0" description="暂无上传记录" />
      <el-table v-else :data="uploadResults" style="width: 100%">
        <el-table-column type="index" label="序号" width="60" />
        <el-table-column prop="filename" label="文件名" />
        <el-table-column prop="url" label="URL" min-width="300">
          <template #default="{ row }">
            <el-link :href="row.url" target="_blank" type="primary">
              {{ row.url }}
            </el-link>
          </template>
        </el-table-column>
        <el-table-column prop="size" label="大小" width="100">
          <template #default="{ row }">
            {{ formatFileSize(row.size) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button type="primary" size="small" @click="copyUrl(row.url)">
              复制 URL
            </el-button>
            <el-button type="success" size="small" @click="openUrl(row.url)">
              打开
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 测试说明 -->
    <el-card shadow="hover" class="mt-6">
      <template #header>
        <div class="card-header">
          <el-icon><InfoFilled /></el-icon>
          <span>测试说明</span>
        </div>
      </template>
      <el-alert
        title="测试步骤"
        type="info"
        :closable="false"
        show-icon
      >
        <ol class="test-steps">
          <li>选择一张图片文件（jpg、png、gif、webp 格式）</li>
          <li>文件大小不能超过 10MB</li>
          <li>上传成功后，会显示文件预览和访问 URL</li>
          <li>可以点击"复制 URL"按钮复制文件地址</li>
          <li>可以点击"打开"按钮在新窗口查看文件</li>
          <li>上传记录会保存在下方的表格中</li>
        </ol>
      </el-alert>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Picture, Upload, List, InfoFilled } from '@element-plus/icons-vue'
import FileUpload from '@/components/FileUpload.vue'
import { uploadFile, formatFileSize, type UploadFileResponse } from '@/api/upload'

const imageUrl1 = ref('')
const imageUrl2 = ref('')
const uploading = ref(false)
const uploadProgress = ref(0)
const uploadResults = ref<UploadFileResponse[]>([])

// FileUpload 组件成功回调
const handleUploadSuccess1 = (response: UploadFileResponse) => {
  uploadResults.value.unshift(response)
  ElMessage.success('上传成功！')
}

// 上传错误回调
const handleUploadError = (error: Error) => {
  console.error('上传失败:', error)
}

// 直接上传前的验证
const beforeUpload = (file: File): boolean => {
  if (file.size > 10 * 1024 * 1024) {
    ElMessage.error('文件大小不能超过 10MB')
    return false
  }
  const imageTypes = ['image/jpeg', 'image/png', 'image/gif', 'image/webp']
  if (!imageTypes.includes(file.type)) {
    ElMessage.error('只能上传图片文件（jpg、png、gif、webp）')
    return false
  }
  return true
}

// 直接上传处理
const handleDirectUpload = async (options: any) => {
  uploading.value = true
  uploadProgress.value = 0

  try {
    // 模拟上传进度
    const progressInterval = setInterval(() => {
      if (uploadProgress.value < 90) {
        uploadProgress.value += 10
      }
    }, 200)

    const response = await uploadFile({
      file: options.file,
      category: 'test'
    })

    clearInterval(progressInterval)
    uploadProgress.value = 100

    if (response.code === 0) {
      imageUrl2.value = response.data.url
      uploadResults.value.unshift(response.data)
      ElMessage.success('上传成功！')
      
      if (options.onSuccess) {
        options.onSuccess(response, options.file, {})
      }
    } else {
      throw new Error(response.message || '上传失败')
    }
  } catch (error: any) {
    uploadProgress.value = 0
    ElMessage.error(error.message || '上传失败')
    if (options.onError) {
      options.onError(error)
    }
  } finally {
    uploading.value = false
  }
}

// 复制 URL
const copyUrl = (url: string) => {
  navigator.clipboard.writeText(url).then(() => {
    ElMessage.success('已复制到剪贴板')
  }).catch(() => {
    ElMessage.error('复制失败')
  })
}

// 打开 URL
const openUrl = (url: string) => {
  window.open(url, '_blank')
}

// 清空记录
const clearResults = () => {
  uploadResults.value = []
  ElMessage.success('已清空记录')
}
</script>

<style scoped>
.upload-test-container {
  min-height: calc(100vh - 84px);
}

.card-header {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: bold;
  font-size: 16px;
}

.preview-image {
  max-width: 100%;
  max-height: 300px;
  border-radius: 4px;
  border: 1px solid #dcdfe6;
}

.test-steps {
  margin: 10px 0 0 20px;
  padding: 0;
  line-height: 2;
}

.test-steps li {
  margin-bottom: 8px;
}
</style>

