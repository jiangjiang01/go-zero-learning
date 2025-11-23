import request from '@/utils/request'
import type { ResponseData } from '@/utils/request'

// 登录请求参数
export interface LoginRequest {
  username: string
  password: string
}

// 登录响应数据
export interface LoginResponse {
  token: string
  user_info: UserInfo
}

// 用户信息（对接后端接口）
export interface UserInfo {
  id: number
  username: string
  email?: string
}

/**
 * 用户登录
 * @param data - 登录数据
 * @returns Promise
 */
export function login(data: LoginRequest): Promise<ResponseData<LoginResponse>> {
  return request({
    url: '/api/users/login',
    method: 'post',
    data
  })
}

/**
 * 用户注册
 * @param data - 注册数据
 * @returns Promise
 */
export function register(data: {
  username: string
  email: string
  password: string
}): Promise<ResponseData<LoginResponse>> {
  return request({
    url: '/api/users',
    method: 'post',
    data
  })
}

/**
 * 获取当前用户信息
 * @returns Promise
 */
export function getUserInfo(): Promise<ResponseData<UserInfo>> {
  return request({
    url: '/api/users/me',
    method: 'get'
  })
}

