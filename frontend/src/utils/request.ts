import axios, { type AxiosInstance, type AxiosResponse } from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserStore } from '@/stores/user'
import router from '@/router'

// 扩展 AxiosRequestConfig 类型，添加 skipErrorHandler 选项
declare module 'axios' {
  export interface AxiosRequestConfig {
    skipErrorHandler?: boolean // 是否跳过响应拦截器的自动错误提示（用于业务代码自定义错误处理）
  }
}

// 响应数据类型
export interface ResponseData<T = any> {
  code: number
  data: T
  message: string
}

// 分页响应
export interface PageResult<T> {
  list: T[]
  total: number
  page: number
  pageSize: number
}

// 创建 axios 实例
const service: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_APP_BASE_URL || '',
  timeout: 10000
})

// 请求拦截器
service.interceptors.request.use(
  (config) => {
    const userStore = useUserStore()
    
    // 设置 Content-Type
    if (!config.headers['Content-Type']) {
      config.headers['Content-Type'] = 'application/json'
    }
    
    // 添加 Authorization 头
    if (userStore.token) {
      config.headers['Authorization'] = `Bearer ${userStore.token}`
    }
    
    return config
  },
  (error) => {
    console.error('请求错误:', error)
    return Promise.reject(error)
  }
)

// 响应拦截器
service.interceptors.response.use(
  (response: AxiosResponse<ResponseData>) => {
    const res = response.data

    // 后端统一响应格式：code: 0 表示成功，非 0 表示失败
    if (res.code !== 0) {
      // 检查是否需要跳过错误处理
      const skipErrorHandler = response.config.skipErrorHandler || false

      // 1002: 未授权，需要重新登录（无论是否跳过都要处理）
      if (res.code === 1002 || res.code === 1006) {
        ElMessageBox.confirm('登录状态已过期，请重新登录', '系统提示', {
          confirmButtonText: '重新登录',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          const userStore = useUserStore()
          userStore.logout().then(() => {
            router.push('/login')
          })
        })
        return Promise.reject(new Error(res.message || '未授权'))
      }

      // 其他错误：只有未设置 skipErrorHandler 时才自动显示
      if (!skipErrorHandler) {
        ElMessage.error(res.message || '请求失败')
      }
      return Promise.reject(new Error(res.message || '请求失败'))
    } else {
      // 成功：返回 data 字段（ResponseData 类型）
      // 注意：拦截器返回 ResponseData，业务代码直接使用 .data 访问
      return res as any
    }
  },
  (error) => {
    console.error('响应错误:', error)
    let message = '请求失败'

    // 检查是否需要跳过错误处理
    const skipErrorHandler = error.config?.skipErrorHandler || false

    if (error.response) {
      const status = error.response.status
      const data = error.response.data
      
      // 如果后端返回了统一格式的错误响应（包含 code 和 message），使用业务错误信息
      if (data && typeof data === 'object' && 'code' in data && 'message' in data) {
        // 这是统一格式的业务错误响应，使用业务错误消息
        message = data.message || '请求失败'
        // 只有未设置 skipErrorHandler 时才显示
        if (!skipErrorHandler) {
          ElMessage.error(message)
        }
        return Promise.reject(new Error(message))
      }
      
      // 如果后端返回了错误信息，使用后端的错误信息
      if (data && data.message) {
        message = data.message
      } else {
        switch (status) {
          case 401:
            message = '未授权，请重新登录'
            // 401 错误需要特殊处理（登出），无论是否跳过都要执行
            if (!skipErrorHandler) {
              const userStore = useUserStore()
              userStore.logout().then(() => {
                router.push('/login')
              })
            }
            break
          case 403:
            message = '拒绝访问'
            break
          case 404:
            message = '请求地址不存在'
            break
          case 405:
            message = '请求方法不允许'
            break
          case 500:
            message = '服务器内部错误'
            break
          default:
            message = `请求失败: ${status}`
        }
      }
    } else if (error.request) {
      message = '网络连接失败'
    }

    // 只有未设置 skipErrorHandler 时才显示
    if (!skipErrorHandler) {
      ElMessage.error(message)
    }
    return Promise.reject(new Error(message))
  }
)

export default service

