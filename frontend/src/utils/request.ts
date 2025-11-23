import axios, { type AxiosInstance, type AxiosRequestConfig, type AxiosResponse } from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserStore } from '@/stores/user'
import router from '@/router'

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
      // 1002: 未授权，需要重新登录
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

      // 其他错误
      ElMessage.error(res.message || '请求失败')
      return Promise.reject(new Error(res.message || '请求失败'))
    } else {
      // 成功：返回 data 字段
      return res
    }
  },
  (error) => {
    console.error('响应错误:', error)
    let message = '请求失败'

    if (error.response) {
      const status = error.response.status
      const data = error.response.data
      
      // 如果后端返回了统一格式的错误响应（包含 code 和 message），使用业务错误信息
      if (data && typeof data === 'object' && 'code' in data && 'message' in data) {
        // 这是统一格式的业务错误响应，使用业务错误消息
        message = data.message || '请求失败'
        // 显示业务错误消息
        ElMessage.error(message)
        // 创建一个新的 Error 对象，使用业务错误消息，这样调用方的 catch 块就不会再显示错误了
        return Promise.reject(new Error(message))
      }
      
      // 如果后端返回了错误信息，使用后端的错误信息
      if (data && data.message) {
        message = data.message
      } else {
        switch (status) {
          case 401:
            message = '未授权，请重新登录'
            const userStore = useUserStore()
            userStore.logout().then(() => {
              router.push('/login')
            })
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

    ElMessage.error(message)
    return Promise.reject(new Error(message))
  }
)

export default service

