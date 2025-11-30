import request from '@/utils/request'
import type { ResponseData } from '@/utils/request'

// 文件上传响应
export interface UploadFileResponse {
  url: string      // 文件访问 URL
  filename: string // 文件名
  size: number     // 文件大小（字节）
}

// 文件上传请求参数
export interface UploadFileParams {
  file: File           // 要上传的文件
  category?: string    // 文件分类（可选，如：product、avatar、document 等）
}

/**
 * 上传文件
 * @param params - 上传参数（文件和分类）
 * @returns Promise
 */
export const uploadFile = (params: UploadFileParams): Promise<ResponseData<UploadFileResponse>> => {
  const formData = new FormData()
  formData.append('file', params.file)
  
  // 如果提供了分类，添加到表单数据中
  if (params.category) {
    formData.append('category', params.category)
  }
  
  // 注意：不要手动设置 Content-Type，让浏览器自动设置（包含 boundary）
  // request 拦截器会自动处理 FormData，删除 Content-Type
  return request.post('/api/upload', formData)
}

/**
 * 上传图片（便捷方法）
 * @param file - 图片文件
 * @param category - 文件分类（可选）
 * @returns Promise
 */
export const uploadImage = (file: File, category?: string): Promise<ResponseData<UploadFileResponse>> => {
  return uploadFile({ file, category })
}

/**
 * 格式化文件大小显示
 * @param bytes - 文件大小（字节）
 * @returns 格式化后的文件大小字符串
 */
export const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 B'
  
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  
  return `${(bytes / Math.pow(k, i)).toFixed(2)} ${sizes[i]}`
}

/**
 * 验证文件类型是否为图片
 * @param file - 文件对象
 * @returns 是否为图片
 */
export const isImageFile = (file: File): boolean => {
  const imageTypes = ['image/jpeg', 'image/png', 'image/gif', 'image/webp']
  return imageTypes.includes(file.type)
}

/**
 * 验证文件大小
 * @param file - 文件对象
 * @param maxSize - 最大文件大小（字节），默认 10MB
 * @returns 是否在限制范围内
 */
export const validateFileSize = (file: File, maxSize: number = 10 * 1024 * 1024): boolean => {
  return file.size <= maxSize
}

