import request from '@/utils/request'
import type { ResponseData } from '@/utils/request'
import type { UserInfo } from './auth'

// 用户列表响应（对接后端接口）
export interface UserListResponse {
  users: UserInfo[]
  total: number
  page: number
  page_size: number
}

// 更新用户请求（对接后端接口）
export interface UpdateUserRequest {
  email?: string
  password?: string
}

/**
 * 获取用户列表
 * @param params - 查询参数
 * @returns Promise
 */
export function getUserList(params?: {
  page?: number
  page_size?: number
  keyword?: string
}): Promise<ResponseData<UserListResponse>> {
  return request({
    url: '/api/users',
    method: 'get',
    params
  })
}

/**
 * 获取用户详情
 * @param id - 用户ID
 * @returns Promise
 */
export function getUser(id: number): Promise<ResponseData<UserInfo>> {
  return request({
    url: `/api/users/${id}`,
    method: 'get'
  })
}

/**
 * 更新当前用户信息
 * @param data - 用户数据
 * @returns Promise
 */
export function updateUser(data: UpdateUserRequest): Promise<ResponseData<UserInfo>> {
  return request({
    url: '/api/users/me',
    method: 'put',
    data
  })
}

/**
 * 删除用户
 * @param id - 用户ID
 * @returns Promise
 */
export function deleteUser(id: number): Promise<ResponseData<{ message: string }>> {
  return request({
    url: `/api/users/${id}`,
    method: 'delete'
  })
}

